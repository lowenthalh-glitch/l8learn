# L8Learn: Schedule & Content Generation from Student Profile

## Objective

Given a student with a populated profile (from uploaded evaluations), generate personalized daily schedules with AI-created lessons, activities, and worksheets tailored to the student's needs, learning style, attention span, and interests.

## Flow

```
1. Guardian opens a student profile
2. Guardian clicks "Generate Schedule"
3. System reads the student's profile (scores, learning style, attention, program settings)
4. System sends profile to Claude with schedule generation prompt
5. Claude generates a weekly plan: daily session blocks with subjects, activities, durations
6. For each session block, system calls Claude again to generate a full lesson
7. Lessons are saved as GeneratedLesson records with steps, questions, materials
8. Guardian reviews the schedule and lessons
9. Student starts a lesson → status changes to IN_PROGRESS → COMPLETED
10. Feedback loop updates the profile based on performance
```

## What EXISTS (ready to use)

| Component | Status |
|---|---|
| StudentProfile with full program settings | Working — has session length, breaks, learning style, attention span |
| GeneratedLesson service + proto | Working — full structure (steps, questions, options, hints, worksheets) |
| DailySchedule service + proto | Registered — callback STUBBED |
| `BuildGenerateLessonPrompt` template | Working — includes profile, mastery, interests, accommodations |
| LLM client (Claude) | Working — LIVE mode with API key |
| PII masking | Working |
| Prompt logging | Working |
| Lesson feedback loop | Working — updates profile on completion |
| Content library (courses, activities, skills) | Working CRUD |

## What's MISSING

| Component | Description |
|---|---|
| Schedule generation prompt | `BuildSchedulePrompt` — sends profile + constraints → returns weekly block plan |
| Schedule callback LLM wiring | `onScheduleCreate` POST triggers schedule generation via LLM |
| Lesson generation trigger | Schedule generator goroutine generates lessons serially (l8agent tool-loop pattern) |
| Schedule generation UI | "Generate Schedule" button on student profile or schedule page |
| Lesson response parser | Parse Claude's lesson JSON into GeneratedLesson proto |
| Schedule response parser | Parse Claude's schedule JSON into ScheduleBlock protos |
| LLM injection into schedule + genlesson services | Same pattern as eval imports — SetLLMClient |
| Proto: link GeneratedLesson → ScheduleBlock | `schedule_id` + `block_id` on GeneratedLesson (l8erp cross-ref pattern) |
| Proto: progress fields on DailySchedule | `lessons_total` + `lessons_generated` |
| Fix 19 wrong After callback signatures | Systemic l8learn issue — all must be 3-arg `ActionValidateFunc` |
| Profile completeness guard | Verify profile has scores/learning style before generating |
| Token limit check | Warn at 100K chars, fail at 400K (same as eval import) |

---

## Phase 1: Schedule Generation Backend

### 1.0 Proto changes + callback signature fixes

**Proto: Link lessons to schedule blocks (l8erp cross-reference pattern)**

Add to `learn-generated.proto` on GeneratedLesson:
```protobuf
string schedule_id = 30;   // DailySchedule that triggered generation
string block_id = 31;      // ScheduleBlock this lesson is for
```

Add to `learn-homeschool.proto` on DailySchedule:
```protobuf
int32 lessons_total = 11;       // Total academic blocks needing lessons
int32 lessons_generated = 12;   // Count of lessons generated so far
```

Run `make-bindings.sh` after both proto changes.

**Fix all 19 wrong After callback signatures in l8learn**

Systemic issue: 19 callback files have the 4-arg signature `(interface{}, ifs.Action, bool, ifs.IVNic) (interface{}, bool, error)` which causes panics via reflection. Must all be fixed to the correct 3-arg `func(interface{}, ifs.Action, ifs.IVNic) error`.

Affected files (batch fix like GenerateID):
- `genlessons/GenLessonServiceCallback.go`
- `worksheets/WorksheetServiceCallback.go`
- `familyactivities/FamilyActivityServiceCallback.go`
- `realworld/RealWorldLessonServiceCallback.go`
- `projects/ProjectServiceCallback.go`
- `enrollments/EnrollmentServiceCallback.go`
- `families/FamilyServiceCallback.go`
- `mastery/MasteryServiceCallback.go`
- And ~11 more (full list from `grep "notify bool" go/learn/`)

### 1.1 Create `BuildSchedulePrompt` template

Add to `prompt_templates.go`:

```go
func BuildSchedulePrompt(profile, programSettings, constraints string) (string, string)
```

System role: "adaptive learning schedule planner for homeschool families"

Rules:
- Generate a 5-day weekly schedule
- Each day has 2 sessions (morning + afternoon) per program settings
- Each session follows the daily structure: movement warmup → academic task → break → therapy task → creative → cleanup
- Respect attention span (maxSeatedWorkMinutes, breakFrequencyMinutes)
- Alternate subjects across days (don't do math every morning)
- Include movement breaks between seated activities
- Factor in weather and parent energy if provided
- Each block must specify: subject, activity type, duration, parent role

Context: student profile JSON, program settings, constraints (weather, energy, appointments)

Return format: JSON array of daily blocks matching the ScheduleBlock proto.

### 1.2 Create `BuildLessonFromSchedulePrompt` template

A variation of `BuildGenerateLessonPrompt` that takes a schedule block as input:

```go
func BuildLessonFromSchedulePrompt(profile, scheduleBlock, mastery string) (string, string)
```

This generates a complete lesson for a specific schedule block — the subject and duration are already determined by the schedule.

### 1.3 Wire LLM into Schedule service

Same pattern as EvalImport:
- Add `SetLLMClient()` to schedule service
- Call from `main.go` after initialization

**CRITICAL: Callback architecture** (learned from eval import):
- `After()` only fires on PUT/PATCH — NOT POST. The schedule generation must be triggered from **BeforeAction**, not After.
- The existing `onScheduleCreate` has the WRONG signature (4-arg, 3-return). Must be fixed to `func(interface{}, ifs.Action, ifs.IVNic) error` (3-arg, 1-return) or moved to BeforeAction.
- Goroutine launched from BeforeAction needs `time.Sleep(500ms)` before first PUT to ensure POST persistence completes.

```go
BeforeAction(func(elem interface{}, action ifs.Action, vnic ifs.IVNic) error {
    if action == ifs.POST {
        sched := elem.(*learn.DailySchedule)
        common.GenerateID(&sched.ScheduleId)
        go func() {
            time.Sleep(500 * time.Millisecond)
            handler.generateSchedule(sched, vnic)
        }()
    }
    return nil
})
```

### 1.4 Lesson generation — driven by schedule generator (l8agent serial loop pattern)

**Who drives generation:** The schedule generator goroutine generates ALL lessons in a **single serial loop** — same pattern as l8agent's chat orchestrator tool loop (`orchestrate.go` lines 177-207). The GenLesson callback does NOT trigger generation — it only handles logging and feedback on completion.

**Flow:**
```
schedule_generator.go goroutine:
  1. Call LLM → get schedule blocks
  2. Save schedule with blocks
  3. Set lessons_total = count of academic blocks
  4. For each academic block (serial loop):
     a. Call BuildLessonFromSchedulePrompt
     b. Parse response → GeneratedLesson
     c. Set lesson.ScheduleId = schedule.ScheduleId
     d. Set lesson.BlockId = block.BlockId
     e. POST lesson with status = READY
     f. Increment lessons_generated on schedule, PUT schedule
     g. If LLM fails → POST lesson with status = GENERATING, error in ai_observation
  5. All done → schedule is complete
```

**Token limit check:** Before each LLM call (schedule and per-lesson), check total prompt size. Warn at 100K chars, fail at 400K chars. Same pattern as eval import.

**Progress tracking:** The schedule record's `lessons_generated` field increments after each lesson. UI polls and shows: "Generating lesson 3 of 10..."

**Long-running generation handling** (up to 30 blocks = 15-30 minutes):
1. UI shows progress bar: `lessons_generated / lessons_total`
2. Partial viewing — blocks with READY lessons are clickable, GENERATING blocks show spinner
3. Generation continues in background if guardian navigates away
4. If a lesson fails, skip and continue. Failed block shows "Generation failed"
5. UI polls schedule every 5 seconds. Shows toast when `lessons_generated == lessons_total`

### 1.5 Shared LLM response parser utility

**Extract `stripCodeFences` + JSON unmarshal** from `eval_response_parser.go` into a shared utility in the engine package. Three files currently need this pattern (eval, schedule, lesson) — per the Second Instance Rule, it must be shared.

Create `engine/llm_response.go`:
```go
// StripCodeFences removes ```json ... ``` markdown wrappers from LLM responses.
func StripCodeFences(s string) string

// ParseLLMResponse strips code fences and unmarshals JSON into the target struct.
func ParseLLMResponse(jsonResponse string, target interface{}) error
```

Then create type-specific parsers that call the shared utility:

`schedules/schedule_response_parser.go`:
```go
func parseScheduleResponse(jsonResponse string) []*learn.ScheduleBlock
```

`genlessons/lesson_response_parser.go`:
```go
func parseLessonResponse(jsonResponse string) *learn.GeneratedLesson
```

Refactor existing `evalimports/eval_response_parser.go` to use `engine.ParseLLMResponse` instead of its own `stripCodeFences`.

---

### 1.6 Profile completeness guard

Before generating a schedule, verify the student profile has minimum required data. Create `schedules/profile_check.go`:

```go
func CheckProfileReady(profile *learn.StudentProfile) (bool, string)
```

Required fields:
- `scores` — at least `overallAcademicReadiness` is non-zero
- `learningStyle` — `bestSessionLengthMinutes` > 0
- `attention` — `focusAcademicTaskMinutes` > 0

If not ready, return `(false, "Profile incomplete: upload and process evaluations first")`. The UI shows this message instead of the schedule form.

---

## Phase 2: Schedule Generation UI

### 2.1 "Generate Schedule" button

Add to the student profile detail popup or the DailySchedule section:
- Button: "Generate Weekly Schedule" (disabled with tooltip if profile incomplete)
- Opens a form with:
  - Student (pre-selected)
  - Week start date
  - Available hours per day (default from profile: 2 sessions × 30 min)
  - Parent energy level (dropdown: low/medium/high)
  - Weather (dropdown: sunny/rainy/cold)
  - Appointments (text area, optional)
- On submit: POST a DailySchedule with constraints → backend generates via LLM

### 2.2 Schedule view

Display the generated schedule as a visual timeline or table:
- Days as columns (Mon-Fri)
- Time blocks as rows with color coding by activity type
- Each block shows: subject, duration, activity type, parent role
- Clickable blocks → open the generated lesson detail

### 2.3 Lesson player / detail view

When clicking a schedule block or generated lesson:
- Show lesson steps in order
- Show questions with hints
- Show materials needed
- Show parent instructions
- "Start Lesson" button → sets status to IN_PROGRESS
- After completion → "Mark Complete" with score entry

---

## Phase 3: End-to-End Flow

### 3.1 The complete cycle

```
Profile populated (from eval imports)
    ↓
Guardian clicks "Generate Schedule"
    ↓
LLM generates weekly plan (5 days × 6 blocks each)
    ↓
For each academic block, LLM generates a full lesson
    ↓
Schedule + lessons saved
    ↓
Guardian reviews schedule
    ↓
Student does a lesson
    ↓
Lesson completed → feedback loop → profile updated
    ↓
Next schedule uses updated profile
```

### 3.2 Verification

**Desktop:**
1. Create a student with NO profile → click "Generate Schedule" → verify disabled with message
2. Upload evaluations → process → populate profile
3. Click "Generate Schedule" → fill form → submit
4. Verify progress indicator: "Generating lesson 1 of N..."
5. Verify schedule blocks appear with correct subjects, durations, activity types
6. Verify READY lessons are clickable, GENERATING blocks show spinner
7. Open a lesson — verify steps, questions, materials, parent instructions
8. Start and complete a lesson — verify feedback loop runs
9. Verify profile is updated after completion
10. Verify `go build ./tests/mocks/` passes with new mock data

**Mobile:**
11. Repeat steps 3-9 on mobile UI
12. Verify `m/app.html` loads all new script tags
13. Verify mobile schedule view renders correctly

---

## Traceability Matrix

| # | Gap | Phase |
|---|---|---|
| 1 | Proto: add schedule_id + block_id to GeneratedLesson (l8erp cross-ref pattern) | Phase 1 |
| 2 | Proto: add lessons_total + lessons_generated to DailySchedule | Phase 1 |
| 3 | Run make-bindings.sh after proto changes | Phase 1 |
| 4 | Fix ALL 19 wrong After callback signatures in l8learn (batch fix) | Phase 1 |
| 5 | `BuildSchedulePrompt` template | Phase 1 |
| 6 | `BuildLessonFromSchedulePrompt` template | Phase 1 |
| 7 | Wire LLM into Schedule service (SetLLMClient + BeforeAction pattern) | Phase 1 |
| 8 | Goroutine timing: 500ms delay before first PUT | Phase 1 |
| 9 | Schedule generator: serial lesson loop (l8agent tool-loop pattern) | Phase 1 |
| 10 | Token limit check before each LLM call (100K warn, 400K fail) | Phase 1 |
| 11 | Extract shared `engine/llm_response.go` utility | Phase 1 |
| 12 | Schedule response parser (uses shared utility) | Phase 1 |
| 13 | Lesson response parser (uses shared utility) | Phase 1 |
| 14 | Refactor eval_response_parser to use shared utility | Phase 1 |
| 15 | Profile completeness guard before generation | Phase 1 |
| 16 | Mock data generators for schedules + generated lessons | Phase 1 |
| 17 | Shared core: schedule-gen-core.js | Phase 2 |
| 18 | Shared core: schedule-view-core.js (with progress bar + partial viewing) | Phase 2 |
| 19 | Shared core: lesson-player-core.js | Phase 2 |
| 20 | Desktop wrappers (3 files) | Phase 2 |
| 21 | Mobile wrappers (3 files) + m/app.html script tags | Phase 2 |
| 22 | UI polling (5 sec) + toast on completion | Phase 2 |
| 23 | Integration tests in go/tests/ | Phase 3 |
| 24 | End-to-end verification (desktop + mobile separately) | Phase 3 |

---

## File Summary

### Proto changes

| File | Changes |
|---|---|
| `proto/learn-homeschool.proto` | Add `lessons_total` and `lessons_generated` fields to DailySchedule |
| `proto/learn-generated.proto` | Add `schedule_id` and `block_id` fields to GeneratedLesson |

### Backend (new files)

| File | Lines (est.) | Purpose |
|---|---|---|
| `engine/llm_response.go` | ~30 | Shared: StripCodeFences + ParseLLMResponse utility |
| `schedules/schedule_generator.go` | ~120 | Async schedule generation: load profile → call LLM → parse → save → generate lessons |
| `schedules/schedule_response_parser.go` | ~40 | Parse Claude JSON → ScheduleBlock protos (uses shared utility) |
| `schedules/profile_check.go` | ~30 | Profile completeness guard before generation |
| `genlessons/lesson_generator.go` | ~100 | Async lesson generation per schedule block |
| `genlessons/lesson_response_parser.go` | ~50 | Parse Claude JSON → GeneratedLesson proto (uses shared utility) |

### Backend (modify)

| File | Changes |
|---|---|
| `prompt_templates.go` | Add `BuildSchedulePrompt` + `BuildLessonFromSchedulePrompt` |
| `ScheduleServiceCallback.go` | Fix callback signature (4-arg → 3-arg), wire LLM, BeforeAction POST → async generation |
| `GenLessonServiceCallback.go` | Add generation trigger for GENERATING status |
| `main.go` | Inject LLM client into schedule + genlesson services |
| `evalimports/eval_response_parser.go` | Refactor to use shared `engine.ParseLLMResponse` |

### Mock data

| File | Changes |
|---|---|
| `gen_learn_schedules.go` (NEW) | Generate sample DailySchedule records with blocks |
| `gen_learn_genlessons.go` (NEW) | Generate sample GeneratedLesson records with steps/questions |
| `learn_phases.go` | Add schedule + genlesson mock data phase |

### UI (new files — shared cores)

| File | Lines (est.) | Purpose |
|---|---|---|
| `learn-ui/schedule/schedule-gen-core.js` | ~80 | Shared: form building, validation, POST orchestration |
| `learn-ui/schedule/schedule-view-core.js` | ~100 | Shared: schedule HTML rendering, block color-coding, progress bar |
| `learn-ui/schedule/lesson-player-core.js` | ~100 | Shared: lesson step rendering, question display, completion flow |

### UI (new files — desktop wrappers)

| File | Lines (est.) | Purpose |
|---|---|---|
| `learn-ui/schedule/schedule-gen.js` | ~50 | Desktop wrapper: Layer8DPopup form |
| `learn-ui/schedule/schedule-view.js` | ~60 | Desktop wrapper: Layer8DPopup timeline |
| `learn-ui/schedule/lesson-player.js` | ~60 | Desktop wrapper: Layer8DPopup lesson view |

### UI (new files — mobile wrappers)

| File | Lines (est.) | Purpose |
|---|---|---|
| `m/js/schedule/schedule-gen-m.js` | ~50 | Mobile wrapper: Layer8MPopup form |
| `m/js/schedule/schedule-view-m.js` | ~50 | Mobile wrapper: Layer8MPopup timeline |
| `m/js/schedule/lesson-player-m.js` | ~50 | Mobile wrapper: Layer8MPopup lesson view |

### UI (modify)

| File | Changes |
|---|---|
| `students-init.js` | Add "Generate Schedule" button to student profile (disabled if profile incomplete) |
| `app.html` | Add script tags for new desktop + core files |
| `m/app.html` | Add script tags for new mobile + core files |

---

## Compliance Checklist

### Project Structure & Architecture
- [x] No new services — extends existing Schedule + GenLesson services
- [x] Proto change: add `lessons_total` + `lessons_generated` to DailySchedule + make-bindings.sh
- [x] No deployment changes
- [x] run-local.sh unchanged
- [x] login.json unchanged

### Service Design
- [x] LLM injection via SetLLMClient pattern (same as eval imports)
- [x] Async generation in goroutines from BeforeAction (NOT After — After only fires on PUT/PATCH)
- [x] Goroutine timing: 500ms delay before first PUT (persistence race condition)
- [x] Existing `onScheduleCreate` wrong signature fixed (4-arg → 3-arg `ActionValidateFunc`)
- [x] After() only fires on PUT/PATCH — this is framework-by-design (confirmed from `service_callback.go` line 88). POST generation must use BeforeAction.
- [x] Lesson generation serial (not parallel) to avoid rate limits
- [x] Completion detection via `custom_fields["lessonsReady"]`

### UI Design
- [x] Shared core + desktop wrapper + mobile wrapper pattern (3 cores + 3 desktop + 3 mobile)
- [x] "Generate Schedule" button disabled if profile incomplete
- [x] Progress indicator for long-running generation ("Generating lesson 3 of 10...")
- [x] Partial viewing — READY lessons clickable, GENERATING blocks show spinner
- [x] `app.html` AND `m/app.html` script tags updated

### Mock Data
- [x] Schedule mock generators with sample blocks
- [x] Generated lesson mock generators with steps/questions
- [x] Phase added to learn_phases.go

### Tests
- [x] Integration tests in go/tests/ for schedule + lesson generation

### Security
- [x] PII masking before LLM calls (student name/family masked)
- [x] All prompts logged via PromptLogger
- [x] Role-based access — guardians can generate, students can view

### Maintainability
- [x] All files under 500 lines
- [x] Shared `engine/llm_response.go` utility — no duplicate stripCodeFences across 3 parsers
- [x] Existing eval_response_parser refactored to use shared utility
- [x] Prompt templates in single file (no duplication)
