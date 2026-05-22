# L8Learn: Evaluation Import → LLM → Profile Pipeline

## Objective

Enable a guardian to upload professional evaluation documents (speech/language, OT, PT, psychomotor, neuropsych), have the system automatically extract structured findings via Claude API with full PII protection, present findings for guardian review, and apply approved findings to the StudentProfile.

## Security Model: Two-Phase PII Protection (Option 4)

**Phase 1 (Local):** Extract PDF text, strip headers/footers/demographics, mask all names/dates/IDs using expanded PIIMasker.

**Phase 2 (Claude):** Send only masked clinical content to Claude for structured extraction.

**Phase 3 (Local):** Unmask response, map findings to StudentProfile fields, store as PENDING for guardian review.

**Principle:** No identifiable data leaves the server. Claude sees only masked clinical content like: `[STUDENT_1] demonstrates [DOB_1]-appropriate gross motor skills...`

---

## Analysis

### What EXISTS (ready to use)

| Component | File | Status |
|---|---|---|
| EvalImport service + CRUD | `evalimports/EvalImportService.go` | Working |
| EvalImport proto (findings, contradictions, status) | `proto/learn-eval.proto` | Needs processing status + error fields |
| LLM client factory (Live/Simulate/LogOnly) | `engine/llm_client.go` | Working |
| LiveClient (Anthropic API, daily limits, token tracking) | `engine/llm_live.go` | Working |
| LLMSimulator (deterministic test responses) | `engine/llm_simulator.go` | Working |
| PIIMasker (names, SSN, email, phone, DOB) | `engine/pii_masking.go` | Working |
| BuildEvalImportPrompt template | `engine/prompt_templates.go` | Working |
| PromptBuilder | `engine/prompt_builder.go` | Working |
| PromptLogger | `engine/prompt_logger.go` | Partial (logs to console, not service) |
| ProfileUpdater class | `engine/profile_updater.go` | Working (not called from eval flow) |
| FileStore (encrypted upload/download) | `vendor/.../filestore/` | Working |
| PromptLog service | `adaptive/promptlogs/` | Activated in `activate_adaptive.go` |
| Credential retrieval pattern | `l8agent/go/init.go` line 73 | Reference |
| `AddAnthropicCredentials()` helper | `l8secure/.../defaults.go` line 66 | Reference |
| Service activation framework | `services/activate_all.go` | Working |

### What's STUBBED (needs implementation)

| Component | File | Lines | What's needed |
|---|---|---|---|
| POST callback (upload → LLM → findings) | `EvalImportServiceCallback.go` | 24-32 | Core pipeline |
| PUT callback (approve → update profile) | `EvalImportServiceCallback.go` | 33-40 | Profile application |
| LLM engine initialization | `main.go` / `activate_all.go` | - | Wire API key → LLMClient |
| PromptLogger service posting | `prompt_logger.go` | 58 | POST to PromptLog service |

### What's MISSING (needs creation)

| Component | Description |
|---|---|
| PDF text extractor | Go-native PDF text extraction (`ledongthuc/pdf`) |
| Document header stripper | Regex-based removal of demographic headers/footers |
| Expanded PIIMasker | Add parent names, therapist names, school names, addresses, MRN patterns |
| Finding → Profile field mapper | Maps `profile_section`/`profile_field` to actual StudentProfile nested struct paths |
| EvalImport handler with LLM deps | Handler struct receiving LLM client via SLA args (l8agent pattern) |
| Async eval processor | Background goroutine for PDF extraction + LLM call (callback stays lightweight) |
| Processing status enum + error field | Proto field so UI can show processing/complete/failed state |
| Valid field path list in prompt | Tell Claude exactly which profile_section/profile_field values are valid |
| Profile auto-creation for new students | Handle first eval upload when no StudentProfile exists yet |
| File upload field in EvalImport form | `f.file()` on the form for PDF upload |
| Finding review UI | Detail popup with accept/reject/edit per finding |

---

## Traceability Matrix

| # | Gap | Phase |
|---|---|---|
| 1 | EvalImport proto missing processing status enum + error field | Phase 0 |
| 2 | Proto change requires make-bindings.sh | Phase 0 |
| 3 | PDF text extraction library not vendored | Phase 1 |
| 4 | Document header/footer stripping | Phase 2 |
| 5 | PIIMasker missing parent/therapist/school/address/MRN patterns | Phase 2 |
| 6 | PIIMasker race condition — shared instance mutated via SetKnownNames | Phase 2 |
| 7 | LLM engine not initialized at startup | Phase 3 |
| 8 | EvalImport handler not wired with LLM deps via SLA args | Phase 3 |
| 9 | EvalImport POST callback stubbed | Phase 3 |
| 10 | Async processing pattern for heavy LLM work | Phase 3 |
| 11 | Error surfacing from async pipeline (processing status) | Phase 3 |
| 12 | New student has no profile — first eval upload fails on nil profile | Phase 3 |
| 13 | Prompt doesn't tell Claude valid profile_section/profile_field values | Phase 3 |
| 14 | EvalImport PUT callback stubbed | Phase 4 |
| 15 | Finding → Profile field mapper doesn't exist | Phase 4 |
| 16 | File upload field missing from EvalImport form | Phase 5 |
| 17 | Finding review UI doesn't exist | Phase 5 |
| 18 | PromptLogger doesn't post to service | Phase 3 |
| 19 | End-to-end verification | Phase 7 |

---

## Phase 0: Proto Changes + Regeneration

**Goal:** Add processing status and error tracking to EvalImport so the UI can show extraction progress and surface failures from the async pipeline.

### 0.1 Update `proto/learn-eval.proto`

Add processing status enum (zero value = UNSPECIFIED per `proto-enum-zero-value` rule):

```protobuf
enum EvalProcessingStatus {
    EVAL_PROCESSING_STATUS_UNSPECIFIED = 0;
    EVAL_PROCESSING_PENDING = 1;
    EVAL_PROCESSING_EXTRACTING = 2;
    EVAL_PROCESSING_COMPLETE = 3;
    EVAL_PROCESSING_FAILED = 4;
}
```

Add fields to EvalImport (use next available field numbers):

```protobuf
message EvalImport {
    // ... existing fields ...
    EvalProcessingStatus processing_status = 15;
    string error_message = 16;
}
```

### 0.2 Regenerate protobuf bindings

Per `protobuf-generation.md` — **MANDATORY** after any proto change:

```bash
cd proto && ./make-bindings.sh
```

Before running, verify `make-bindings.sh` uses `-i` (not `-it`) on `docker run` commands.

### 0.3 Verify build

```bash
cd go && go build ./...
```

### 0.4 Update UI columns and forms

Add processing status to EvalImport columns in `people-columns.js`:
```javascript
...col.enum('processingStatus', 'Status', null, render.evalProcessingStatus),
```

Add enum + renderer to `people-enums.js` for the processing status values.

**Verification:** `go build ./...` passes. Proto types have `ProcessingStatus` and `ErrorMessage` fields.

---

## Phase 1: PDF Text Extraction

**Goal:** Add Go-native PDF text extraction — no system dependencies.

### 1.1 Add `ledongthuc/pdf` library

Full vendor refresh per ecosystem rules:
```bash
cd go
rm -rf go.sum go.mod vendor
go mod init
GOPROXY=direct GOPRIVATE=github.com go mod tidy
go mod vendor
```

Note: Add `github.com/ledongthuc/pdf` to imports in `pdf_extractor.go` before running `go mod tidy` so it gets pulled.

### 1.2 Create `go/learn/students/evalimports/pdf_extractor.go`

**~60 lines.** Single function:

```go
func ExtractTextFromPDF(fileData []byte) (string, error)
```

- Takes raw (decrypted) PDF bytes
- Uses `ledongthuc/pdf` to open from bytes reader
- Extracts text from all pages, concatenated with page markers
- Returns full text string
- Falls back to empty string with error if PDF has no text layer (scanned image)

### 1.3 Create `go/learn/students/evalimports/file_reader.go`

**~30 lines.** Retrieves file via the FileStore service API (not direct disk I/O):

```go
func ReadUploadedFile(filePath string, vnic ifs.IVNic) ([]byte, error)
```

- Calls the FileStore service through `vnic` service API with an `L8FileDownloadRequest`
- FileStore handles disk I/O and decryption internally
- Returns decrypted raw bytes
- Does NOT duplicate FileStore logic — delegates to the service

**Verification:** Integration test in `go/tests/` with a sample PDF → confirm text extraction works.

---

## Phase 2: Enhanced PII Protection

**Goal:** Expand PIIMasker to strip all identifying information before text leaves the server.

### 2.1 Create `go/learn/students/evalimports/document_sanitizer.go`

**~120 lines.** Two main functions:

#### `StripDocumentHeaders(text string) string`

Professional evaluations have a consistent structure:
- Page 1 header: patient name, DOB, MRN, date of eval, provider info
- Subsequent page headers: name, DOB, page number
- Footer: clinic name, address, phone, fax

Strip these using section detection:
```go
// Detect and remove header block (everything before first section heading)
// Common section headings: "BACKGROUND", "HISTORY", "ASSESSMENT", "RESULTS",
// "OBSERVATIONS", "RECOMMENDATIONS", "SUMMARY", "GOALS"
```

Remove lines matching patterns:
- `Patient:`, `Name:`, `DOB:`, `Date of Birth:`, `MRN:`, `Medical Record`
- `Date of Evaluation:`, `Date of Report:`
- `Referred by:`, `Insurance:`, `Policy:`
- `Page X of Y`
- Address-like patterns (number + street name + city/state/zip)

#### `BuildMaskingContext(studentId string, vnic ifs.IVNic) []string`

Loads related entities via service API to build the full known-names list:
1. Load Student record → first name, last name, preferred name
2. Load Guardians for this student → parent first/last names
3. Load Teachers for this student's classroom → teacher names
4. Load School → school name
5. Load District → district name

Returns all names as a flat string list for `PIIMasker.SetKnownNames()`.

### 2.2 Expand `pii_masking.go`

Add new regex patterns (**~30 lines** added to existing file):

```go
var (
    // Existing patterns...
    addressRegex   = regexp.MustCompile(`\b\d{1,5}\s+\w+\s+(St|Ave|Blvd|Dr|Rd|Ln|Way|Ct|Pl|Cir|Pkwy)\.?\b`)
    mrnRegex       = regexp.MustCompile(`\b(?:MRN|Medical Record|Record #|Patient ID)[:\s#]*[\w\-]+\b`)
    insuranceRegex = regexp.MustCompile(`\b(?:Policy|Member ID|Group #|Insurance ID)[:\s#]*[\w\-]+\b`)
    dateFullRegex  = regexp.MustCompile(`\b(?:January|February|March|April|May|June|July|August|September|October|November|December)\s+\d{1,2},?\s+\d{4}\b`)
)
```

Add masking for these in `MaskText()`:
- Addresses → `[MASKED_ADDRESS]`
- MRN → `[MASKED_MRN]`
- Insurance → `[MASKED_INSURANCE]`
- Full date strings → `[DATE_N]` (reversible token)

### 2.3 Fix PIIMasker thread safety

The shared `PIIMasker` instance is passed via the handler to all concurrent eval imports. `SetKnownNames()` mutates the `knownNames` slice — a race condition if two imports process simultaneously.

**Fix:** Change `MaskText()` to accept known names as a parameter instead of reading from instance state:

```go
func (p *PIIMasker) MaskTextWithNames(text string, tokenMap *TokenMap, knownNames []string) string
```

The existing `MaskText()` remains for backwards compatibility (uses `p.knownNames`). The async processor calls `MaskTextWithNames()` with per-request names built by `BuildMaskingContext()`. No shared mutable state.

**Verification:** Feed a real evaluation header through the sanitizer + masker → confirm zero PII in output.

---

## Phase 3: LLM Engine Wiring + POST Callback

**Goal:** Wire the LLM client at startup using the SLA/handler dependency injection pattern (matching l8agent), and implement the eval import POST handler with async processing.

### 3.1 Initialize LLM engine in `main.go`

Add after `ActivateAllServices()` (**~15 lines**):

```go
// Load Anthropic credentials (same pattern as l8agent/go/init.go line 73)
_, _, apiKey, _, err := nic.Resources().Security().Credential("Anthropic", "API_KEY", nic.Resources())
masker := engine.NewPIIMasker()
logger := engine.NewPromptLogger(nic)
var llmClient engine.LLMClient
if apiKey != "" {
    llmClient = engine.NewLLMClient(learn.LLMMode_LLM_MODE_LIVE, apiKey, masker, logger)
    nic.Resources().Logger().Info("LLM client initialized in LIVE mode")
} else {
    llmClient = engine.NewLLMClient(learn.LLMMode_LLM_MODE_SIMULATE, "", masker, logger)
    nic.Resources().Logger().Warning("No Anthropic API key found, using LLM simulator")
}
```

Pass `llmClient` and `masker` to EvalImport service activation via the SLA args pattern.

### 3.2 Modify EvalImport service activation to accept LLM dependencies

Update `EvalImportService.go` to accept and forward LLM dependencies:

```go
func Activate(creds, dbname string, vnic ifs.IVNic, llmClient engine.LLMClient, masker *engine.PIIMasker) {
    handler := newEvalImportHandler(llmClient, masker, vnic)
    common.ActivateService(common.ServiceConfig{
        ServiceName: ServiceName,
        ServiceArea: ServiceArea,
        PrimaryKey:  "ImportId",
        Callback:    handler.buildCallback(vnic),
    }, &learn.EvalImport{}, &learn.EvalImportList{}, creds, dbname, vnic)
}
```

### 3.3 Create EvalImport handler struct

In `EvalImportServiceCallback.go`, replace the function-based callback with a handler struct (following l8agent's `chatHandler` pattern):

```go
type evalImportHandler struct {
    llmClient engine.LLMClient
    masker    *engine.PIIMasker
    vnic      ifs.IVNic
}

func newEvalImportHandler(llmClient engine.LLMClient, masker *engine.PIIMasker, vnic ifs.IVNic) *evalImportHandler {
    return &evalImportHandler{llmClient: llmClient, masker: masker, vnic: vnic}
}

func (h *evalImportHandler) buildCallback(vnic ifs.IVNic) ifs.IServiceCallback {
    return common.NewValidation(&learn.EvalImport{}, vnic).
        Require(func(v interface{}) string { return v.(*learn.EvalImport).ImportId }, "ImportId").
        Require(func(v interface{}) string { return v.(*learn.EvalImport).StudentId }, "StudentId").
        After(h.onEvalImportChange).
        Build()
}
```

### 3.4 Implement POST callback with async processing

The POST callback stays lightweight — sets processing status and spawns a goroutine for heavy work:

```go
func (h *evalImportHandler) onEvalImportChange(elem interface{}, action ifs.Action, notify bool, vnic ifs.IVNic) (interface{}, bool, error) {
    eval := elem.(*learn.EvalImport)

    if action == ifs.POST {
        // Lightweight: set status to PENDING, return immediately
        eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_PENDING
        // Heavy work runs async in goroutine
        go h.processEvalImport(eval, vnic)
    }

    if action == ifs.PUT && eval.AppliedToProfile {
        // Profile application is fast (no LLM call), can run inline
        err := h.applyToProfile(eval, vnic)
        if err != nil {
            return nil, false, err
        }
    }

    return nil, true, nil
}
```

### 3.5 Create async eval processor

In `go/learn/students/evalimports/eval_processor.go` (**~120 lines**):

Every error path sets `ProcessingStatus = FAILED` with `ErrorMessage` so the UI can surface failures. Uses `MaskTextWithNames()` (not `SetKnownNames()`) to avoid race conditions on the shared masker.

```go
func (h *evalImportHandler) processEvalImport(eval *learn.EvalImport, vnic ifs.IVNic) {
    // Mark as extracting
    eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_EXTRACTING
    saveEvalImport(eval, vnic)

    // 1. Read uploaded file via FileStore service API
    fileData, err := ReadUploadedFile(eval.FilePath, vnic)
    if err != nil {
        h.failEval(eval, "Failed to read uploaded file: "+err.Error(), vnic)
        return
    }

    // 2. Extract text from PDF
    pdfText, err := ExtractTextFromPDF(fileData)
    if err != nil {
        h.failEval(eval, "Failed to extract text from PDF: "+err.Error(), vnic)
        return
    }
    if pdfText == "" {
        h.failEval(eval, "PDF has no extractable text layer (scanned image not supported yet)", vnic)
        return
    }

    // 3. Strip document headers/demographics
    pdfText = StripDocumentHeaders(pdfText)

    // 4. Build masking context (load student + family + school names)
    knownNames := BuildMaskingContext(eval.StudentId, vnic)

    // 5. Mask PII using per-request names (thread-safe, no shared state mutation)
    tokenMap := engine.NewTokenMap()
    pdfText = h.masker.MaskTextWithNames(pdfText, tokenMap, knownNames)

    // 6. Load current profile for comparison — create empty if none exists
    currentProfile := loadOrCreateProfile(eval.StudentId, vnic)
    profileJSON, _ := json.Marshal(currentProfile)

    // 7. Build prompt (includes valid field path list)
    systemPrompt, userMessage := engine.BuildEvalImportPrompt(
        pdfText,
        eval.DocumentType.String(),
        string(profileJSON),
    )

    // 8. Call LLM (may take 10-60 seconds)
    response, err := h.llmClient.Call(
        learn.LLMPromptType_LLM_PROMPT_TYPE_EVAL_IMPORT,
        systemPrompt, userMessage, eval.StudentId,
    )
    if err != nil {
        h.failEval(eval, "LLM call failed: "+err.Error(), vnic)
        return
    }

    // 9. Unmask PII in response (restore original values)
    response = tokenMap.Unmask(response)

    // 10. Parse response into findings
    findings, contradictions := parseEvalResponse(response)

    // 11. Update eval with findings (all PENDING), mark complete
    eval.Findings = findings
    eval.Contradictions = contradictions
    eval.AcceptedCount = 0
    eval.RejectedCount = 0
    eval.AllReviewed = false
    eval.AppliedToProfile = false
    eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_COMPLETE
    eval.ErrorMessage = ""
    saveEvalImport(eval, vnic)
}

func (h *evalImportHandler) failEval(eval *learn.EvalImport, msg string, vnic ifs.IVNic) {
    eval.ProcessingStatus = learn.EvalProcessingStatus_EVAL_PROCESSING_FAILED
    eval.ErrorMessage = msg
    saveEvalImport(eval, vnic)
    vnic.Resources().Logger().Error("EvalImport ", eval.ImportId, ": ", msg)
}
```

### 3.6 Create `go/learn/students/evalimports/eval_response_parser.go`

**~80 lines.** Parses Claude's JSON response into proto types:

```go
func parseEvalResponse(jsonResponse string) ([]*learn.EvalFinding, []*learn.EvalContradiction)
```

Expected JSON structure (matches `BuildEvalImportPrompt` return format):
```json
{
    "document_type": "speech_language",
    "professional": "...",
    "findings": [
        { "profile_section": "speech", "profile_field": "clarity", "new_value": "...", "source_text": "...", "confidence": 0.95 }
    ],
    "contradictions": [
        { "profile_field": "...", "current_value": "...", "document_says": "...", "ai_recommendation": "..." }
    ]
}
```

### 3.7 Update `BuildEvalImportPrompt` with valid field paths

The current prompt says "Map every finding to the student profile schema" but doesn't tell Claude what the valid values are. Claude would guess field names that may not match the proto, and `profile_applicator.go` would silently skip unknown fields.

Add a rule to the prompt builder listing all valid `profile_section.profile_field` combinations:

```go
AddRule("Valid profile_section values: scores, strengths, challenges, learningStyle, attention, motivation, literacy, math, speech, speech.articulation, speech.expressiveLanguage, speech.receptiveLanguage, speech.languageHistory, fineMotor, fineMotor.nameWritingStatus, grossMotor, sensory, socialEmotional, behavior, dailyLiving, dailyLiving.toileting, dailyLiving.hygiene, dailyLiving.dressing, dailyLiving.feeding, health, adaptiveRules, programSettings")
```

This ensures Claude returns exact matches that `profile_applicator.go` can map without guessing.

### 3.8 Create `loadOrCreateProfile` helper

In `service_helpers.go`:

```go
func loadOrCreateProfile(studentId string, vnic ifs.IVNic) *learn.StudentProfile
```

- Queries `select * from StudentProfile where studentId=<id>`
- If found, returns it
- If not found (first eval for a new student), creates an empty `StudentProfile` with `profileId` auto-generated and `studentId` set, POSTs it to the Profile service, and returns it

This handles the case where a guardian uploads an evaluation before a profile exists.

### 3.9 Fix PromptLogger to POST to PromptLog service

In `prompt_logger.go`, replace the TODO at line 58. The PromptLog service is confirmed activated in `activate_adaptive.go`. POST the log record via `vnic` service API (**~10 lines**).

### 3.8 Update `activate_students.go`

Pass LLM dependencies through to `evalimports.Activate()`:

```go
func collectStudentActivations(creds, dbname string, nic ifs.IVNic, llmClient engine.LLMClient, masker *engine.PIIMasker) []func() {
    return []func(){
        // ... other services ...
        func() { evalimports.Activate(creds, dbname, nic, llmClient, masker) },
    }
}
```

**Verification:** Upload a test PDF → confirm findings appear as PENDING on the EvalImport record (may take 10-30 seconds for LLM response).

---

## Phase 4: Profile Application (PUT Callback)

**Goal:** When guardian approves findings, apply them to StudentProfile.

### 4.1 Create `go/learn/students/evalimports/profile_applicator.go`

**~150 lines.** Maps findings to StudentProfile fields:

```go
func ApplyFindingsToProfile(
    eval *learn.EvalImport,
    profile *learn.StudentProfile,
) error
```

The mapping uses `finding.ProfileSection` + `finding.ProfileField` to navigate the nested StudentProfile struct. Key mappings:

| profile_section | profile_field | StudentProfile path |
|---|---|---|
| `speech` | `clarity` | `profile.Speech.Clarity` |
| `speech` | `therapyNeed` | `profile.Speech.TherapyNeed` |
| `speech.articulation` | `notedSoundChallenges` | `profile.Speech.Articulation.NotedSoundChallenges` |
| `speech.receptiveLanguage` | `needs` | `profile.Speech.ReceptiveLanguage.Needs` |
| `fineMotor` | `handDominance` | `profile.FineMotor.HandDominance` |
| `fineMotor` | `pencilGrip` | `profile.FineMotor.PencilGrip` |
| `grossMotor` | `overallStatus` | `profile.GrossMotor.OverallStatus` |
| `sensory` | `sensoryPattern` | `profile.Sensory.SensoryPattern` |
| `attention` | `structuredTaskStamina` | `profile.Attention.StructuredTaskStamina` |
| `scores` | `readingReadiness` | `profile.Scores.ReadingReadiness` |
| `socialEmotional` | `confidence` | `profile.SocialEmotional.Confidence` |
| `dailyLiving.dressing` | `buttons` | `profile.DailyLiving.Dressing.Buttons` |
| ... | ... | ... |

Implementation: Use a switch/map to navigate the nested struct. For string fields, set directly. For `[]string` fields, append or replace. For `int32` fields (scores), parse and set.

For each finding:
- If `status == ACCEPTED`: use `finding.NewValue` (or `finding.ExtractedValue` per proto)
- If `status == EDITED`: use `finding.EditedValue`
- If `status == REJECTED` or `PENDING`: skip

### 4.2 Implement `applyToProfile` on the handler

In `EvalImportServiceCallback.go`:

```go
func (h *evalImportHandler) applyToProfile(eval *learn.EvalImport, vnic ifs.IVNic) error {
    // 1. Load current StudentProfile
    profile := loadCurrentProfile(eval.StudentId, vnic)
    if profile == nil {
        return fmt.Errorf("no profile found for student %s", eval.StudentId)
    }

    // 2. Apply accepted/edited findings
    err := ApplyFindingsToProfile(eval, profile)
    if err != nil {
        return err
    }

    // 3. Update profile timestamps
    profile.LastUpdated = time.Now().Unix()

    // 4. Save profile via PUT to Profile service
    return saveProfile(profile, vnic)
}
```

### 4.3 Helper functions in `go/learn/students/evalimports/service_helpers.go`

**~60 lines:**

```go
func loadCurrentProfile(studentId string, vnic ifs.IVNic) *learn.StudentProfile
func saveProfile(profile *learn.StudentProfile, vnic ifs.IVNic) error
func saveEvalImport(eval *learn.EvalImport, vnic ifs.IVNic) error
```

Use L8Query via the service framework: `select * from StudentProfile where studentId=<id>`.

**Verification:** Create an EvalImport with ACCEPTED findings → set `appliedToProfile=true` → confirm StudentProfile fields updated.

---

## Phase 5: UI — File Upload + Finding Review

**Goal:** Add PDF upload to the EvalImport form and a finding review popup.

### 5.1 Add file upload to EvalImport form

In `people-forms.js`, update the EvalImport form:

```javascript
EvalImport: f.form('Evaluation Import', [
    f.section('Upload Evaluation', [
        ...f.reference('studentId', 'Student', 'Student'),
        ...f.select('documentType', 'Document Type', enums.EVAL_DOC_TYPE),
        ...f.text('professionalName', 'Professional Name'),
        ...f.date('evaluationDate', 'Evaluation Date'),
        ...f.file('filePath', 'Evaluation Document (PDF)', true)
    ]),
    f.section('Review Status', [
        ...f.number('acceptedCount', 'Accepted', false, { readOnly: true }),
        ...f.number('rejectedCount', 'Rejected', false, { readOnly: true }),
        ...f.checkbox('allReviewed', 'All Reviewed', false, { readOnly: true }),
        ...f.checkbox('appliedToProfile', 'Apply to Profile')
        // Omitted: uploadedBy — auto-set by backend from auth context
    ])
])
```

### 5.2 Create shared finding review logic

Create `go/learn/ui/web/learn-ui/students/people/people-eval-review-core.js` (**~150 lines**):

Platform-independent behavioral logic for the finding review workflow:

```javascript
window.L8EvalReview = {
    // Build HTML for findings table (returns HTML string, no DOM dependency)
    renderFindingsHtml: function(findings, contradictions) { ... },

    // Build HTML for a single finding row with accept/reject/edit buttons
    renderFindingRow: function(finding, index) { ... },

    // Build HTML for contradictions section
    renderContradictionsHtml: function(contradictions) { ... },

    // Confidence color class: green > 0.8, yellow > 0.5, red < 0.5
    confidenceClass: function(score) { ... },

    // Handle accept/reject/edit button click — updates finding status in data
    onFindingAction: function(evalImport, findingIndex, action, editedValue) { ... },

    // Compute acceptedCount / rejectedCount from findings
    recomputeCounts: function(evalImport) { ... },

    // Build "Apply to Profile" PUT payload
    buildApplyPayload: function(evalImport) { ... }
};
```

This file contains **zero DOM manipulation or popup calls** — only data logic and HTML string generation. Both desktop and mobile wrappers call these functions.

### 5.3 Create desktop finding review wrapper

Create `go/learn/ui/web/learn-ui/students/people/people-eval-review.js` (**~60 lines**):

Thin desktop wrapper that:
1. Opens `Layer8DPopup` with content from `L8EvalReview.renderFindingsHtml()`
2. Attaches click handlers that delegate to `L8EvalReview.onFindingAction()`
3. "Apply to Profile" button PUTs via `Layer8DForms.saveRecord()`

Wire into the students init so that clicking an EvalImport row opens this custom detail view.

### 5.4 Create mobile finding review wrapper

Create `go/learn/ui/web/m/js/students/people-eval-review-m.js` (**~60 lines**):

Thin mobile wrapper that:
1. Opens `Layer8MPopup` with content from `L8EvalReview.renderFindingsHtml()`
2. Attaches click handlers that delegate to `L8EvalReview.onFindingAction()`
3. "Apply to Profile" button PUTs via `Layer8MAuth.put()`

Include in `m/app.html`.

### 5.5 Add script tags to `app.html`

After `people-profile-forms.js`:
```html
<script src="learn-ui/students/people/people-eval-review.js"></script>
```

**Verification:** Upload a PDF → confirm findings appear in review popup → accept some, reject some → apply → confirm profile updated.

---

## Phase 7: End-to-End Verification

For the complete pipeline:

1. **Create a student** via the Students UI
2. **Upload a PDF evaluation** via the EvalImport Add form
3. **Verify PII masking**: check PromptLog service — confirm no student names, no parent names, no addresses in the logged system/user prompts
4. **Verify findings extracted**: open the EvalImport detail → confirm findings populated with PENDING status
5. **Review findings**: accept some, reject some, edit one
6. **Apply to profile**: click "Apply to Profile"
7. **Verify profile updated**: open the StudentProfile → confirm accepted findings reflected in the correct sections
8. **Verify rejected findings NOT applied**: confirm rejected values did not change the profile
9. **Test with LLM_MODE_SIMULATE**: set LLMConfig mode to SIMULATE → upload another eval → confirm simulator findings appear (deterministic test data)
10. **Test with no API key**: remove Anthropic credential → restart → confirm graceful fallback to simulator mode

Sections to verify:
- [ ] Student creation
- [ ] PDF upload via FileStore
- [ ] PDF text extraction
- [ ] Header stripping
- [ ] PII masking (check prompt logs)
- [ ] LLM call + response parsing (async — check after 10-30 seconds)
- [ ] Finding storage as PENDING
- [ ] Finding review popup
- [ ] Accept/reject/edit workflow
- [ ] Profile field application
- [ ] Profile timestamp update
- [ ] Simulator fallback
- [ ] Processing status transitions (PENDING → EXTRACTING → COMPLETE or FAILED)
- [ ] Error handling: bad PDF → status FAILED with error message in UI
- [ ] Error handling: LLM timeout → status FAILED with error message in UI
- [ ] Error handling: empty response → status FAILED with error message in UI
- [ ] New student with no profile → profile auto-created on first eval upload
- [ ] Concurrent eval uploads → no race condition on PIIMasker (per-request names)

---

## File Summary

### New files to create

| File | Lines (est.) | Purpose |
|---|---|---|
| `evalimports/pdf_extractor.go` | ~60 | PDF text extraction |
| `evalimports/file_reader.go` | ~30 | FileStore service API call for file retrieval |
| `evalimports/document_sanitizer.go` | ~120 | Header stripping + masking context |
| `evalimports/eval_processor.go` | ~120 | Async pipeline with status tracking and error handling |
| `evalimports/eval_response_parser.go` | ~80 | Parse LLM JSON → proto findings |
| `evalimports/profile_applicator.go` | ~150 | Map findings → StudentProfile fields |
| `evalimports/service_helpers.go` | ~60 | Load/save profile and eval via service API |
| `learn-ui/.../people-eval-review-core.js` | ~150 | Shared finding review logic (platform-independent) |
| `learn-ui/.../people-eval-review.js` | ~60 | Desktop wrapper (Layer8DPopup) |

### Existing files to modify

| File | Changes |
|---|---|
| `proto/learn-eval.proto` | Add EvalProcessingStatus enum + processing_status + error_message fields |
| `EvalImportService.go` | Accept LLMClient + PIIMasker params, pass to handler |
| `EvalImportServiceCallback.go` | Replace function with handler struct, async POST, inline PUT |
| `pii_masking.go` | Add address/MRN/insurance/full-date regexes + MaskTextWithNames() |
| `prompt_templates.go` | Add valid field path list to BuildEvalImportPrompt rules |
| `main.go` | Load Anthropic credential, create LLMClient, pass to services |
| `activate_students.go` | Forward LLM dependencies to evalimports.Activate() |
| `prompt_logger.go` | POST log records to PromptLog service |
| `people-enums.js` | Add EVAL_PROCESSING_STATUS enum + renderer |
| `people-columns.js` | Add processingStatus column to EvalImport |
| `people-forms.js` | Add `f.file()` to EvalImport form |
| `app.html` | Add script tags for eval review JS |

### Dependencies to add

| Library | Purpose | Vendoring |
|---|---|---|
| `github.com/ledongthuc/pdf` | Pure Go PDF text extraction | Full `go mod` refresh sequence |

### Mock data updates

Update `go/tests/mocks/gen_learn_evals.go` to populate `filePath` on generated EvalImport records with a test PDF path so the file upload column renders in the table.

### Mobile files to create

| File | Lines (est.) | Purpose |
|---|---|---|
| `m/js/students/people-eval-review-m.js` | ~60 | Mobile wrapper (Layer8MPopup) |

### Architecture compliance notes

- **Dependency injection**: LLMClient passed via service activation args (l8agent `sla.SetArgs()` pattern), NOT global singletons
- **File I/O**: FileStore service called via service API, NOT direct disk read + manual decryption
- **Callback weight**: POST callback is lightweight (spawns goroutine), returns immediately. Heavy work (PDF extraction, LLM call, parsing) runs async in `processEvalImport()`
- **Vendoring**: Full refresh sequence per `vendor-third-party-code.md` rule
- **File sizes**: All files under 500 lines per `maintainability.md`
- **No generics**: All code uses concrete types and interfaces
- **Project-specific UI**: Finding review UI in `learn-ui/`, NOT in `l8ui/`

---

## Compliance Checklist

### Project Structure & Architecture
- [x] Follows l8erp/l8agent architecture patterns
- [x] Directory names and file organization match ecosystem conventions
- [x] Dependencies injected via SLA/handler pattern (not singletons)

### Protobuf Design
- [x] Proto change (Phase 0): add EvalProcessingStatus enum with UNSPECIFIED zero value per `proto-enum-zero-value` rule
- [x] Proto change: regenerate bindings via `make-bindings.sh` per `protobuf-generation` rule
- [x] EvalFinding/EvalContradiction correctly embedded as `repeated` (not Prime Objects)
- [x] Student referenced by `student_id` (ID), not by struct pointer

### Service Design
- [x] ServiceName `EvalImprt` is 9 chars (under 10 limit)
- [x] ServiceArea 20 consistent within Students module
- [x] POST callback lightweight — heavy work in async goroutine
- [x] File I/O via FileStore service API, not direct disk access

### UI Design
- [x] Forms use correct protobuf JSON field names (verified against .pb.go)
- [x] `uploadedBy` omitted from form — auto-set by backend from auth context
- [x] `acceptedCount`, `rejectedCount`, `allReviewed` marked `readOnly: true` (system-computed)
- [x] Finding review uses custom popup, not generic form
- [x] File upload uses `f.file()` per Layer8FileUpload pattern
- [x] Desktop and mobile parity: review UI for both platforms

### Mock Data
- [x] Existing mock data in `gen_learn_evals.go` — update to populate `filePath`

### Tests
- [x] All tests in `go/tests/` per test-location-and-approach rule

### Deployment & Configuration
- [x] No new deployable services — no Docker/K8s changes needed
- [x] Anthropic credential accessed via ISecurityProvider (security config)
- [x] Graceful fallback to simulator when no API key present
- [x] `run-local.sh` unchanged — API key set in security config

### Security & Thread Safety
- [x] PII masking covers: student names, parent names, therapist names, school names, SSN, DOB, email, phone, address, MRN, insurance
- [x] Document headers stripped before masking
- [x] Only masked clinical content sent to Claude
- [x] All prompts logged to PromptLog service (masked)
- [x] Credential retrieval via ISecurityProvider, not hardcoded
- [x] PIIMasker thread-safe: `MaskTextWithNames()` uses per-request names, no shared state mutation
- [x] Async processor sets ProcessingStatus on every error path — no silent failures
- [x] LLM prompt includes valid field path list — Claude can't return unmapped field names
- [x] New student handled: `loadOrCreateProfile()` creates empty profile if none exists

### Vendoring
- [x] Full vendor refresh sequence for `ledongthuc/pdf`
