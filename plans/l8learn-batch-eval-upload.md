# L8Learn: Batch Evaluation Upload + Guardian-Approved LLM Processing

## Objective

Redesign the evaluation import flow so a guardian can upload ALL evaluation documents at once, review the cleaned/masked versions for data safety, and only then trigger the LLM to build the student profile from all evaluations together.

## New Flow

```
1. Guardian selects a student
2. Guardian uploads multiple PDF files at once (speech, OT, PT, neuropsych, etc.)
3. System immediately: extracts text, strips headers, masks PII, saves cleaned files
4. Guardian reviews each cleaned file in the browser — verifies no sensitive data
5. Guardian clicks "Process All" — system sends ALL cleaned evaluations to Claude in ONE prompt
6. Claude builds a comprehensive student profile from all evaluations combined
7. Guardian reviews findings, accepts/rejects/edits
8. Guardian clicks "Apply to Profile"
```

## Current State vs Required Changes

| Component | Current | New |
|---|---|---|
| Upload | One file per EvalImport | Multiple files per student, one EvalImport per file |
| Document type | Required before upload | Optional — auto-detected by Claude |
| LLM trigger | Automatic on POST | Manual — guardian clicks "Process All" after review |
| LLM input | Single document text | All documents concatenated |
| Cleaning | Happens + sends to LLM in one step | Cleaning happens on upload, LLM waits for approval |

---

## Phase 1: Proto + Backend Changes

### 1.1 Update `learn-eval.proto`

Make `document_type` optional (it already has a zero value `UNSPECIFIED`, so no proto change needed — just remove the UI requirement).

Add a new processing status value for "cleaned and ready for review":
```protobuf
enum EvalProcessingStatus {
    EVAL_PROCESSING_STATUS_UNSPECIFIED = 0;
    EVAL_PROCESSING_PENDING = 1;
    EVAL_PROCESSING_EXTRACTING = 2;
    EVAL_PROCESSING_COMPLETE = 3;
    EVAL_PROCESSING_FAILED = 4;
    EVAL_PROCESSING_CLEANED = 5;    // NEW: cleaned, waiting for guardian review
    EVAL_PROCESSING_SUBMITTED = 6;  // NEW: submitted to LLM, waiting for response
}
```

Run `make-bindings.sh` after proto change.

### 1.2 Split the callback: clean on POST, LLM on PUT

**POST (upload)**: Extract PDF → strip headers → mask PII → save cleaned file → set status to `CLEANED`. **No LLM call.**

**PUT with special trigger**: "Process All" PUTs only ONE eval with `processingStatus = SUBMITTED`. The After callback on that single PUT loads ALL CLEANED evals for the student and sends them to Claude in one call. The UI does NOT PUT every eval individually — only one trigger PUT.

### 1.3 Update `EvalImportServiceCallback.go`

**BeforeAction (POST)**:
```
1. GenerateID
2. Set status = PENDING
3. Launch goroutine (with 500ms delay to ensure record is persisted first):
   extract PDF → strip → mask → save cleaned → set status = CLEANED
```

The goroutine delay addresses the timing issue: BeforeAction runs before persistence, so the goroutine must wait for the POST to complete before calling saveEvalImport (PUT) to update the status.

**After (PUT)** — when `processingStatus == SUBMITTED`:

Note: After callback signature MUST be `func(interface{}, ifs.Action, ifs.IVNic) error` (3 args, 1 return) — NOT the 4-arg form that caused a panic in the previous plan.

```
1. This fires only ONCE (only one eval is PUT with SUBMITTED)
2. Load ALL EvalImports for this student where status == CLEANED
3. Read each cleaned text file
4. Concatenate all texts with document separators
5. Send combined text to Claude in one prompt
6. Parse findings
7. Store ALL findings on the triggering EvalImport record
8. Set triggering eval status = COMPLETE
9. Set all other CLEANED evals status = COMPLETE (no findings on them)
```

**Findings storage**: All findings go on the single EvalImport that was set to SUBMITTED (the trigger record). Other evals just get status COMPLETE. The guardian reviews findings on that one record.

**Partial cleaning failures**: "Process All" only sends evals with status CLEANED. Failed evals (status FAILED) are skipped. The guardian can see which ones failed and re-upload.

### 1.4 Update `eval_processor.go`

Split into two functions:
- `cleanEvalImport()` — extract, strip, mask, save. No LLM call.
- `processAllEvals()` — load all cleaned evals for student, concatenate, call LLM, parse findings.

### 1.5 Update prompt template

Change `BuildEvalImportPrompt` to accept multiple document texts:

```go
func BuildEvalImportPrompt(documents []EvalDocument, currentProfile string) (string, string)

type EvalDocument struct {
    Text         string
    DocumentType string  // may be "UNSPECIFIED" — Claude auto-detects
    FileName     string
}
```

The user message becomes:
```
DOCUMENT 1 (speech evaluation):
<cleaned text>

DOCUMENT 2 (OT evaluation):
<cleaned text>

DOCUMENT 3 (neuropsych):
<cleaned text>

CURRENT_STUDENT_PROFILE:
<profile JSON>
```

### 1.6 Token limit safety

Multiple evaluation documents concatenated could exceed Claude's practical context for quality output. The `processAllEvals()` function must:
1. Calculate total character count of all documents + system prompt + profile JSON
2. Log the total: `[EvalProcessor] Total prompt size: %d chars (~%d tokens)`
3. If total exceeds 100,000 chars (~25K tokens): log a warning but proceed (Claude Sonnet handles 200K context)
4. If total exceeds 400,000 chars: fail with error message "Combined evaluations too large for single analysis"

---

## Phase 2: UI Changes

### 2.1 Custom multi-file upload form (replaces standard Add form)

The standard `f.file()` supports only one file. The EvalImport Add modal needs a **custom CRUD handler** (per `special-cases.md` pattern) that:

1. Shows a student reference picker
2. Shows a multi-file dropzone using `<input type="file" multiple accept=".pdf">`
3. On Save:
   - Uploads each file to FileStore separately via `Layer8FileUpload.upload()`
   - For each uploaded file, POSTs a new EvalImport with `{ studentId, filePath: storagePath, documentType: UNSPECIFIED }`
   - Shows progress ("Uploading file 2 of 5...")
4. After all POSTs succeed, refreshes the table

This override is placed in `people-batch-process.js` (or a dedicated `people-eval-upload.js`) — NOT in `students-init.js`.

```javascript
// Custom Add modal for EvalImport — replaces standard single-file form
Students._openAddModal = function(service) {
    if (service.model !== 'EvalImport') { origOpenAdd.call(this, service); return; }
    // Show custom multi-file modal...
};
```

Document type is optional (UNSPECIFIED) — guardian can set it per eval later or Claude auto-detects.

### 2.2 Add "Process All" button to the EvalImport table

When the table shows evaluations for a student with status `CLEANED`:
- Show a "Process All Evaluations" button above the table
- **Single trigger**: Clicking it PUTs only the FIRST CLEANED eval with `processingStatus = SUBMITTED`. NOT all evals — only one PUT triggers the batch LLM call on the backend. This avoids the race condition of multiple After callbacks.
- The backend loads ALL CLEANED evals for that student in one LLM call
- Button is disabled/hidden when no CLEANED evals exist or when processing is in progress

**Student context**: The button uses the `studentId` from the first CLEANED eval in the table to scope the batch. If the table shows evals for multiple students, the button shows a student picker or only appears when filtered to one student.

### 2.3 Keep the "View" link for cleaned files

Already implemented — guardian clicks "View" to inspect each cleaned file before processing.

### 2.4 Add status badges

| Status | Badge | Meaning |
|---|---|---|
| Pending | Yellow | Just uploaded, cleaning in progress |
| Cleaned | Blue | Cleaned and masked, ready for guardian review |
| Submitted | Yellow | Sent to LLM, waiting for response |
| Complete | Green | Findings extracted |
| Failed | Red | Error occurred |

---

## Phase 3: Data Safety Verification

### 3.1 Guardian review step (CRITICAL)

After upload and cleaning, the guardian MUST review each cleaned file before processing:
1. Each eval shows "View" link in the table
2. Guardian clicks View, sees the masked text in a popup
3. If satisfied, proceeds to "Process All"
4. If sensitive data found, guardian can delete the eval and re-upload

### 3.2 What gets sent to Claude

Only text that has been:
1. Extracted from PDF (raw text, no images)
2. Stripped of demographic headers (name, DOB, age, grade, school, address, MRN, insurance)
3. Masked with PII tokens (`[PERSON_1]`, `[DOB_1]`, `[MASKED_SSN]`, `[MASKED_ADDRESS]`, etc.)
4. Reviewed and approved by the guardian

---

## Traceability Matrix

| # | Gap | Phase |
|---|---|---|
| 1 | Proto: add CLEANED and SUBMITTED status values | Phase 1 |
| 2 | Proto change requires make-bindings.sh | Phase 1 |
| 3 | Split callback: clean on POST, LLM on PUT/SUBMITTED | Phase 1 |
| 4 | Split eval_processor into cleanEvalImport + processAllEvals | Phase 1 |
| 5 | Update prompt to accept multiple documents | Phase 1 |
| 6 | Token limit safety check before LLM call | Phase 1 |
| 7 | Goroutine timing: 500ms delay before first saveEvalImport | Phase 1 |
| 8 | Single trigger PUT (not N PUTs) to avoid race condition | Phase 1 |
| 9 | All findings stored on trigger eval; other evals set to COMPLETE | Phase 1 |
| 10 | Partial failures: skip FAILED evals, process only CLEANED | Phase 1 |
| 11 | Remove document type requirement from form | Phase 2 |
| 12 | Multi-file upload creates one EvalImport per file | Phase 2 |
| 13 | "Process All" button — single trigger PUT, student context | Phase 2 |
| 14 | Status badges for new states | Phase 2 |
| 15 | Extract batch-process UI logic into separate file (not students-init.js) | Phase 2 |
| 16 | Mobile: "Process All" button + new status badges | Phase 2 |
| 17 | Mock data: update gen_learn_evals.go with CLEANED/SUBMITTED statuses | Phase 1 |
| 18 | Integration tests in go/tests/ for batch flow | Phase 3 |
| 19 | Guardian review flow documented and enforced | Phase 3 |
| 20 | End-to-end verification | Phase 4 |

---

## Phase 4: End-to-End Verification

1. Create a student
2. Upload 3 evaluation PDFs at once (multi-file)
3. Verify all 3 show status "Cleaned" in the table
4. Click "View" on each — verify no sensitive data (names, DOB, addresses masked)
5. Upload 1 more PDF that has no text layer — verify it shows status "Failed" with error message
6. Click "Process All Evaluations" — verify only the 3 CLEANED evals are sent, FAILED is skipped
7. Verify only ONE LLM call is made (check server log for single `[EvalProcessor]` batch)
8. Verify the trigger eval has findings; other 2 CLEANED evals show status "Complete" with 0 findings
9. Verify findings reference source text from ALL 3 documents
10. Review findings, accept some, reject some
11. Apply to profile
12. Verify StudentProfile populated with data from all evaluations
13. Verify token count logged: `Total prompt size: N chars (~M tokens)`

---

## File Summary

### Proto changes
| File | Changes |
|---|---|
| `proto/learn-eval.proto` | Add EVAL_PROCESSING_CLEANED and EVAL_PROCESSING_SUBMITTED |

### Backend changes
| File | Changes |
|---|---|
| `EvalImportServiceCallback.go` | POST: clean only. PUT/SUBMITTED: batch LLM call |
| `eval_processor.go` | Split into `cleanEvalImport()` + `processAllEvals()` |
| `prompt_templates.go` | Accept multiple documents in prompt |

### UI changes (shared core + platform wrappers)
| File | Lines (est.) | Purpose |
|---|---|---|
| `people-batch-process-core.js` (NEW) | ~80 | Shared logic: multi-file upload orchestration, "Process All" iteration (PUT SUBMITTED per eval), progress tracking. Platform-independent — no DOM/popup calls. |
| `people-batch-process.js` (NEW) | ~50 | Desktop wrapper: custom Add modal with multi-file dropzone (Layer8DPopup), "Process All" button (Layer8DNotification). Calls core functions. |
| `m/js/students/people-batch-process-m.js` (NEW) | ~50 | Mobile wrapper: same flow using Layer8MPopup/Layer8MAuth. Calls core functions. |
| `people-forms.js` | — | Remove required document type from EvalImport form |
| `people-enums.js` | — | Add CLEANED and SUBMITTED status labels + badges |

Pattern follows established `people-eval-review-core.js` + desktop/mobile wrapper approach (no behavioral duplication).

### Mock data changes
| File | Changes |
|---|---|
| `gen_learn_evals.go` | Add evals with CLEANED and SUBMITTED statuses to mock distribution |

### Existing files unchanged
- `document_sanitizer.go` — already works
- `pii_masking.go` — already works
- `file_reader.go` — already works
- `pdf_extractor.go` — already works
- `profile_applicator.go` — already works
- `people-eval-review-core.js` — already works

---

## Compliance Checklist

### Project Structure & Architecture
- [x] Follows existing l8learn architecture
- [x] No new services — extends existing EvalImport service
- [x] No deployment changes needed (no new binaries)
- [x] run-local.sh unchanged
- [x] login.json unchanged

### Protobuf Design
- [x] Enum zero value UNSPECIFIED (existing)
- [x] New enum values added after existing ones (no renumbering)
- [x] make-bindings.sh after proto change

### Service Design
- [x] POST callback lightweight (goroutine for cleaning with 500ms delay)
- [x] LLM call only on explicit guardian action (single PUT/SUBMITTED trigger)
- [x] No sensitive data sent without guardian approval
- [x] ServiceName `EvalImprt` unchanged (9 chars)
- [x] After callback signature: `func(interface{}, ifs.Action, ifs.IVNic) error` (3 args, 1 return)
- [x] Single trigger PUT avoids N After callbacks / N LLM calls race condition
- [x] All findings stored on trigger eval; other evals set to COMPLETE
- [x] Partial failures: FAILED evals skipped, only CLEANED processed
- [x] Token limit check before LLM call

### UI Design
- [x] Multi-file upload creates separate records
- [x] Guardian review step before LLM processing
- [x] "View" link for inspecting cleaned files
- [x] Clear status progression visible in table
- [x] Desktop and mobile parity — "Process All" button on both platforms
- [x] Batch-process logic in separate file (not embedded in students-init.js)

### Mock Data
- [x] gen_learn_evals.go updated with CLEANED and SUBMITTED status distributions

### Tests
- [x] Integration tests in go/tests/ for batch processing flow

### Security
- [x] PII masking before saving cleaned file
- [x] Guardian must review cleaned files before LLM processing
- [x] No auto-send to LLM on upload
- [x] Original PDF deleted after cleaning
- [x] Credential access via ISecurityProvider (existing, unchanged)
