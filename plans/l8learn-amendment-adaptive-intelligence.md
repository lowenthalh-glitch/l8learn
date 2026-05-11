# L8Learn Amendment: Adaptive Intelligence Layer

## Purpose

This amendment addresses the critical gap between the PRD title ("Adaptive Learning Operating System") and what was built (a school administration CRUD system). The infrastructure is solid — 38 services, VNet mesh, ORM, web UI with 7 modules. What's missing is the **intelligence** that makes it adaptive.

## Scope

This amendment covers:
1. Redesigned Student Profile (rich learning profile, not flat admin record)
2. LLM Simulation Mode (test prompts and data flow without a real LLM)
3. Diagnostic Placement Flow
4. AI Engine activation with real prompt construction
5. Voice mode foundations
6. Parent coaching intelligence
7. Computed analytics (not just CRUD)

---

## 1. Student Learning Profile

### 1.1 Why the Current Model Fails

The current `Student` protobuf has 21 fields — all administrative (name, school, grade, classroom). The adaptive engine needs to know **how the child learns**, not just where they're enrolled.

### 1.2 New Protobuf: learn-profile.proto

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";
import "learn-content.proto";

// ============================================================================
// STUDENT LEARNING PROFILE
// ============================================================================

// @PrimeObject
// The comprehensive learning profile — continuously updated by the adaptive engine.
// One per student. This is the AI's understanding of the child.
message StudentProfile {
    string profile_id = 1;
    string student_id = 2;
    int64 created_date = 3;
    int64 last_updated = 4;

    // Overall summary (AI-generated narrative, updated weekly)
    string overall_description = 5;
    repeated string main_strengths = 6;
    repeated string main_challenges = 7;
    repeated string primary_goals = 8;

    // Readiness scores (0-100, computed from assessment data)
    ReadinessScores readiness = 9;

    // Learning style preferences (observed from interaction data)
    LearningStyle learning_style = 10;

    // Attention & regulation (observed from session patterns)
    AttentionProfile attention = 11;

    // Motivation & engagement (observed from activity preferences)
    MotivationProfile motivation = 12;

    // Literacy profile (computed from mastery data)
    LiteracyProfile literacy = 13;

    // Math profile (computed from mastery data)
    MathProfile math = 14;

    // Speech & language
    SpeechLanguageProfile speech = 15;

    // Fine motor / OT
    FineMotorProfile fine_motor = 16;

    // Gross motor
    GrossMotorProfile gross_motor = 17;

    // Sensory needs
    SensoryProfile sensory = 18;

    // Social-emotional
    SocialEmotionalProfile social_emotional = 19;

    // Behavior patterns (observed from interaction data)
    BehaviorProfile behavior = 20;

    // Technology usage limits (set by guardian)
    TechnologyLimits technology = 21;

    // Health & safety (set by guardian/admin)
    HealthSafety health = 22;

    // Therapy & external services
    TherapyServices therapy = 23;

    // Adaptive engine settings (tuned over time)
    AdaptiveSettings adaptive_settings = 24;

    // AI tutor personality settings (set by guardian, refined by AI)
    AITutorSettings ai_tutor = 25;

    // Goals
    StudentGoals goals = 26;

    // Progress tracking preferences
    ProgressTracking progress_tracking = 27;

    map<string, string> custom_fields = 28;
    l8common.AuditInfo audit_info = 29;
}

message StudentProfileList {
    repeated StudentProfile list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// PROFILE SUB-MESSAGES (all embedded children of StudentProfile)
// ============================================================================

message ReadinessScores {
    int32 academic_readiness = 1;          // 0-100
    int32 reading_readiness = 2;
    int32 math_readiness = 3;
    int32 writing_fine_motor = 4;
    int32 speech_language = 5;
    int32 attention_stamina = 6;
    int32 social_emotional = 7;
    int32 independence = 8;
}

message LearningStyle {
    repeated string preferred_modes = 1;    // "visual", "auditory", "kinesthetic", "reading"
    int32 best_session_length_minutes = 2;
    int32 best_activity_length_minutes = 3;
    int32 max_seated_work_minutes = 4;
    int32 break_frequency_minutes = 5;
    string best_time_of_day = 6;           // "morning", "midday", "afternoon"
    repeated string effective_activity_types = 7;  // Observed: which ActivityTypes produce best scores
}

message AttentionProfile {
    int32 focus_preferred_activity_minutes = 1;
    int32 focus_academic_task_minutes = 2;
    repeated string losing_focus_signs = 3;  // "rushing", "random_answers", "long_pauses"
    repeated string helpful_supports = 4;    // "timer_visible", "chunked_problems", "movement_breaks"
}

message MotivationProfile {
    repeated string interests = 1;           // "dinosaurs", "space", "animals", "sports"
    repeated string favorite_activities = 2; // ActivityType preferences observed from completion rates
    repeated string preferred_rewards = 3;   // "badges", "streaks", "avatar_items", "points"
    repeated string avoided_activities = 4;  // ActivityTypes with high abandonment
}

message LiteracyProfile {
    string reading_level = 1;               // "pre-reader", "beginning", "developing", "fluent"
    string letter_recognition = 2;          // "none", "some", "uppercase", "all"
    string phonemic_awareness = 3;          // "none", "beginning", "developing", "proficient"
    string sight_words = 4;                 // "0-25", "25-50", "50-100", "100+"
    string comprehension = 5;              // "literal_only", "inferential", "critical"
    int32 book_stamina_minutes = 6;
    int32 reading_fluency_wpm = 7;
}

message MathProfile {
    string level = 1;                       // "pre-number", "counting", "operations", "fractions", "algebra"
    string counting = 2;                    // range: "to_10", "to_20", "to_100", "to_1000"
    string number_recognition = 3;
    string addition = 4;                    // "not_started", "single_digit", "double_digit", "triple_digit"
    string subtraction = 5;
    string multiplication = 6;
    string division = 7;
    string fractions = 8;
    repeated string preferred_tools = 9;    // "number_line", "blocks", "arrays", "fingers", "mental"
    repeated string error_patterns = 10;    // AI-detected: "skip_counting_off_by_one", "place_value_confusion"
}

message SpeechLanguageProfile {
    bool receives_speech_therapy = 1;
    repeated string current_goals = 2;
    repeated string speech_sounds = 3;      // Sounds the child struggles with
    string clarity = 4;                     // "unclear", "somewhat_clear", "clear", "very_clear"
    string expressive_language = 5;         // "limited", "developing", "age_appropriate"
    string receptive_language = 6;
    repeated string helpful_prompts = 7;    // "repeat_slowly", "visual_cues", "sentence_starters"
}

message FineMotorProfile {
    bool receives_occupational_therapy = 1;
    string hand_dominance = 2;             // "left", "right", "not_established"
    string pencil_grip = 3;               // "fist", "developing", "mature"
    string tracing = 4;                    // "not_started", "developing", "proficient"
    string cutting = 5;
    string name_writing = 6;
    repeated string helpful_tools = 7;      // "thick_pencil", "slant_board", "adapted_scissors"
}

message GrossMotorProfile {
    string energy_level = 1;               // "low", "moderate", "high", "very_high"
    repeated string favorite_movement = 2;  // "jumping", "running", "dancing", "yoga"
    string coordination = 3;
    string balance = 4;
    repeated string movement_breaks = 5;    // Effective break activities observed
}

message SensoryProfile {
    repeated string sensitivities = 1;      // "loud_sounds", "bright_screens", "textures"
    repeated string sensory_seeking = 2;    // "movement", "fidget", "chewing"
    repeated string sensory_supports = 3;   // "noise_canceling", "dim_screen", "fidget_tool"
}

message SocialEmotionalProfile {
    string confidence = 1;                  // "low", "developing", "moderate", "high"
    string peer_interaction = 2;            // "avoids", "parallel_play", "cooperative", "leader"
    string turn_taking = 3;
    string emotion_naming = 4;             // "limited", "basic", "developing", "proficient"
    repeated string frustration_triggers = 5;  // "too_many_wrong", "time_pressure", "new_topics"
    repeated string calming_strategies = 6;    // "deep_breaths", "movement_break", "switch_topic"
}

message BehaviorProfile {
    repeated string avoidance_behaviors = 1;   // "refuses_writing", "rushes_through", "says_too_hard"
    repeated string triggers = 2;              // "consecutive_errors", "long_sessions", "unfamiliar_format"
    repeated string redirect_strategies = 3;   // "offer_choice", "reduce_difficulty", "game_mode"
    repeated string successful_supports = 4;   // What has worked (AI-learned)
}

message TechnologyLimits {
    int32 max_screen_time_daily_minutes = 1;
    int32 max_session_minutes = 2;
    repeated string allowed_uses = 3;          // "learning_only", "games_ok", "videos_ok"
    repeated string activities_to_avoid = 4;   // "timed_tests", "competitive"
}

message HealthSafety {
    repeated string medical_conditions = 1;
    repeated string allergies = 2;
    string vision_concerns = 3;
    string hearing_concerns = 4;
    repeated string safety_concerns = 5;
}

message TherapyServices {
    repeated TherapyService services = 1;
}

message TherapyService {
    string service_type = 1;               // "speech", "occupational", "behavioral", "physical"
    string provider_name = 2;
    string frequency = 3;                  // "weekly", "twice_weekly", "monthly"
    repeated string therapy_goals = 4;
    repeated string home_practice = 5;     // Activities to reinforce at home
}

message AdaptiveSettings {
    int32 default_session_length_minutes = 1;
    int32 maximum_session_length_minutes = 2;
    int32 maximum_activity_length_minutes = 3;
    int32 break_frequency_minutes = 4;
    string difficulty_adjustment_speed = 5;  // "slow", "normal", "fast"
    repeated string error_responses = 6;     // What to do on errors: "encourage", "simplify", "hint", "skip"
    repeated string success_responses = 7;   // What to do on success: "celebrate", "advance", "challenge"
    int32 max_consecutive_errors = 8;        // Before auto-switching to easier content
    int32 max_consecutive_correct = 9;       // Before auto-advancing difficulty
}

message AITutorSettings {
    repeated string personality = 1;         // "encouraging", "patient", "playful", "calm"
    repeated string should_do = 2;           // "use_simple_words", "give_examples", "celebrate_effort"
    repeated string should_avoid = 3;        // "time_pressure", "comparison", "negative_feedback"
    string sentence_length = 4;             // "very_short", "short", "normal"
    string question_style = 5;              // "one_at_a_time", "multiple_choice_first", "open_ended"
    string hint_style = 6;                  // "visual", "verbal", "example_based", "step_by_step"
    string encouragement_frequency = 7;     // "every_answer", "every_few", "on_struggle"
}

message StudentGoals {
    repeated Goal short_term = 1;           // This week
    repeated Goal medium_term = 2;          // This month
    repeated Goal long_term = 3;            // This semester/year
}

message Goal {
    string goal_id = 1;
    string description = 2;
    string skill_id = 3;
    string target = 4;                      // "mastery_level_proficient", "reading_fluency_90wpm"
    string status = 5;                      // "not_started", "in_progress", "achieved"
    int64 target_date = 6;
    int64 achieved_date = 7;
}

message ProgressTracking {
    repeated string daily_metrics = 1;      // "time_on_task", "activities_completed", "accuracy"
    repeated string weekly_metrics = 2;     // "skills_progressed", "mastery_gained", "engagement"
    repeated string mastery_criteria = 3;   // "80%_accuracy_over_10_attempts", "3_consecutive_correct"
}
```

### 1.3 Service

| ServiceName | ServiceArea | PrimaryKey | Model |
|-------------|:-----------:|------------|-------|
| `Profile` | 20 | `ProfileId` | `StudentProfile` |

### 1.4 How the Profile Gets Built

The profile is NOT filled in manually. It's built automatically over time:

```
Day 1:  Guardian fills basic info (health, therapy, technology limits, AI tutor preferences)
        Diagnostic benchmark runs → populates readiness scores, initial literacy/math levels

Week 1: Adaptive engine observes interactions → populates:
        - learning_style (from completion rates by activity type)
        - attention (from session duration patterns)
        - motivation (from engagement signals)

Week 2+: Continuous refinement:
        - behavior patterns (from avoidance/frustration signals)
        - error_patterns in math (AI-detected from wrong answers)
        - social_emotional (from collaboration interactions)
        - adaptive_settings auto-tuned (session length, break frequency, difficulty speed)

Monthly: AI generates updated:
        - overall_description narrative
        - main_strengths / main_challenges / primary_goals
        - readiness scores recalculated
```

---

## 2. LLM Simulation Mode (CRITICAL — Must Be Built Before Real LLM)

### 2.1 Purpose

Before connecting to a real LLM (Anthropic Claude), we need to:
1. **See every prompt** the system would send to the LLM
2. **Verify no data leaks** — no PII, no sensitive student data in prompts
3. **Understand the prompt structure** — what context the AI sees
4. **Test with deterministic responses** — predictable outputs for debugging
5. **Control costs** — don't burn API credits during development

### 2.2 Architecture

```
Normal mode:
  Adaptive Engine → LLM Client → Anthropic API → Response

Simulation mode:
  Adaptive Engine → LLM Client → LLM Simulator → Response
                                      ↓
                              Prompt Log Service
                              (stores every prompt + response for review)
```

### 2.3 New Protobuf: learn-llm.proto

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// LLM PROMPT LOGGING & SIMULATION
// ============================================================================

enum LLMPromptType {
    LLM_PROMPT_TYPE_UNSPECIFIED = 0;
    LLM_PROMPT_TYPE_PATH_DECISION = 1;       // Adaptive engine: choose next activity
    LLM_PROMPT_TYPE_PROFILE_UPDATE = 2;      // Update student profile narrative
    LLM_PROMPT_TYPE_RISK_ASSESSMENT = 3;     // Weekly risk prediction
    LLM_PROMPT_TYPE_PROGRESS_SUMMARY = 4;    // Generate progress report narrative
    LLM_PROMPT_TYPE_PARENT_COACHING = 5;     // Daily teaching tip for parent
    LLM_PROMPT_TYPE_WORKSHEET_SCAN = 6;      // Extract answers from scanned image
    LLM_PROMPT_TYPE_CONTENT_ANALYSIS = 7;    // Analyze content effectiveness
    LLM_PROMPT_TYPE_SCHEDULE_GENERATION = 8; // Generate daily family schedule
    LLM_PROMPT_TYPE_CHAT = 9;               // Interactive AI agent chat
    LLM_PROMPT_TYPE_MODERATION = 10;         // Collaboration message moderation
}

enum LLMMode {
    LLM_MODE_UNSPECIFIED = 0;
    LLM_MODE_LIVE = 1;                       // Real LLM (Anthropic API)
    LLM_MODE_SIMULATE = 2;                   // Simulated responses
    LLM_MODE_LOG_ONLY = 3;                   // Log prompt, don't call LLM, return empty
}

// @PrimeObject
// Every prompt sent to the LLM is logged here for review
message LLMPromptLog {
    string log_id = 1;
    LLMPromptType type = 2;
    string student_id = 3;                   // Which student this is about (empty for system-wide)
    LLMMode mode = 4;                        // Was this live, simulated, or log-only?

    // The prompt
    string system_prompt = 5;                // Full system prompt sent
    string user_message = 6;                 // Full user message sent
    int32 system_prompt_tokens = 7;          // Token count estimate
    int32 user_message_tokens = 8;

    // The response
    string response = 9;                     // Full LLM response (or simulated response)
    int32 response_tokens = 10;
    int64 response_time_ms = 11;             // Latency

    // Data safety audit
    bool contains_student_name = 12;         // Flag: does the prompt contain real names?
    bool contains_pii = 13;                  // Flag: does it contain PII (DOB, address, etc.)?
    repeated string pii_fields_found = 14;   // Which PII fields were detected
    bool data_masked = 15;                   // Was PII masking applied before sending?

    // Metadata
    int64 timestamp = 16;
    string triggered_by = 17;                // What caused this prompt: "activity_completed", "weekly_batch", etc.

    map<string, string> custom_fields = 18;
    l8common.AuditInfo audit_info = 19;
}

message LLMPromptLogList {
    repeated LLMPromptLog list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
// System-wide LLM configuration
message LLMConfig {
    string config_id = 1;
    LLMMode mode = 2;                        // Current operating mode
    string api_provider = 3;                 // "anthropic", "openai", etc.
    string model_name = 4;                   // "claude-sonnet-4-6", etc.
    int32 max_tokens = 5;
    double temperature = 6;
    bool pii_masking_enabled = 7;            // Mask student names/DOB before sending
    bool prompt_logging_enabled = 8;         // Log all prompts (always true in simulate mode)
    int32 max_daily_calls = 9;               // Cost control
    int32 calls_today = 10;

    map<string, string> custom_fields = 11;
    l8common.AuditInfo audit_info = 12;
}

message LLMConfigList {
    repeated LLMConfig list = 1;
    l8api.L8MetaData metadata = 2;
}
```

### 2.4 Services

| ServiceName | ServiceArea | PrimaryKey | Model |
|-------------|:-----------:|------------|-------|
| `Profile` | 20 | `ProfileId` | `StudentProfile` |
| `PromptLog` | 30 | `LogId` | `LLMPromptLog` |
| `LLMConfig` | 30 | `ConfigId` | `LLMConfig` |

### 2.5 LLM Simulator

The simulator replaces the real LLM client. It:
1. **Logs every prompt** to the PromptLog service
2. **Scans for PII** in the prompt (student names, DOB, addresses)
3. **Returns deterministic responses** based on prompt type
4. **Flags data safety issues** in the log

```go
// go/learn/adaptive/engine/llm_simulator.go
type LLMSimulator struct {
    vnic       ifs.IVNic
    piiScanner *PIIScanner
}

func (s *LLMSimulator) Call(promptType LLMPromptType, systemPrompt, userMessage string, studentId string) (string, error) {
    // 1. Scan for PII
    piiReport := s.piiScanner.Scan(systemPrompt + userMessage)

    // 2. Log the prompt
    log := &learn.LLMPromptLog{
        Type:                 promptType,
        StudentId:            studentId,
        Mode:                 learn.LLMMode_LLM_MODE_SIMULATE,
        SystemPrompt:         systemPrompt,
        UserMessage:          userMessage,
        ContainsStudentName:  piiReport.HasNames,
        ContainsPii:          piiReport.HasPII,
        PiiFieldsFound:       piiReport.Fields,
        DataMasked:           false,
        Timestamp:            time.Now().Unix(),
    }

    // 3. Generate deterministic response based on type
    response := s.generateResponse(promptType, userMessage)
    log.Response = response

    // 4. Store log
    // POST to PromptLog service

    return response, nil
}
```

### 2.6 PII Scanner

```go
// go/learn/adaptive/engine/pii_scanner.go
type PIIReport struct {
    HasNames  bool
    HasPII    bool
    Fields    []string
}

func (s *PIIScanner) Scan(text string) *PIIReport {
    report := &PIIReport{}

    // Check for student names (loaded from Student service)
    for _, name := range s.knownStudentNames {
        if strings.Contains(text, name) {
            report.HasNames = true
            report.Fields = append(report.Fields, "student_name: " + name)
        }
    }

    // Check for date patterns (DOB)
    // Check for email patterns
    // Check for phone patterns
    // Check for address patterns

    report.HasPII = report.HasNames || len(report.Fields) > 0
    return report
}
```

### 2.7 PII Masking

When `pii_masking_enabled = true`, student names are replaced before sending:
```
Before: "Jake Martinez (Grade 4) scored 62% on multiplication"
After:  "Student_A (Grade 4) scored 62% on multiplication"
```

The mapping (`Student_A → Jake Martinez`) is stored locally and used to unmask the response before saving.

### 2.8 Simulated Responses

For each prompt type, the simulator returns a predictable response:

```go
func (s *LLMSimulator) generateResponse(promptType LLMPromptType, userMessage string) string {
    switch promptType {
    case learn.LLM_PROMPT_TYPE_PATH_DECISION:
        return `{"nextActivities":[
            {"activityId":"ACT-0001","skillId":"SKL-001","difficulty":3,"reason":"Simulated: targeting weakest skill"},
            {"activityId":"ACT-0002","skillId":"SKL-002","difficulty":2,"reason":"Simulated: confidence builder"}
        ],"reasoning":"Simulated response — review prompt in PromptLog"}`

    case learn.LLM_PROMPT_TYPE_PROFILE_UPDATE:
        return `{"overallDescription":"Simulated profile update","mainStrengths":["simulated"],"mainChallenges":["simulated"]}`

    case learn.LLM_PROMPT_TYPE_RISK_ASSESSMENT:
        return `{"riskLevel":1,"riskScore":0.2,"factors":[{"factorType":"simulated","description":"Simulated risk assessment"}]}`

    case learn.LLM_PROMPT_TYPE_PARENT_COACHING:
        return `{"tip":"Simulated coaching tip — review prompt in PromptLog to see what context the AI receives"}`

    default:
        return `{"simulated":true,"message":"Review the prompt in PromptLog service"}`
    }
}
```

### 2.9 Admin UI for Prompt Review

New section in the admin UI: **"AI Monitor"**

```
┌─────────────────────────────────────────────────────────────┐
│  AI Monitor                                [Mode: SIMULATE] │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  LLM Mode: ○ Live  ● Simulate  ○ Log Only                  │
│  PII Masking: [✓] Enabled                                   │
│  Prompt Logging: [✓] Enabled (always on in simulate mode)   │
│                                                              │
│  ┌─ Recent Prompts ──────────────────────────────────────┐  │
│  │ Type          │ Student  │ PII? │ Tokens │ Time       │  │
│  ├───────────────┼──────────┼──────┼────────┼────────────│  │
│  │ PATH_DECISION │ STU-0042 │ ⚠ YES│ 1,240  │ 2s ago     │  │
│  │ PARENT_COACH  │ STU-0001 │ ✓ NO │ 890    │ 5m ago     │  │
│  │ RISK_ASSESS   │ STU-0099 │ ✓ NO │ 2,100  │ 1h ago     │  │
│  └───────────────┴──────────┴──────┴────────┴────────────┘  │
│                                                              │
│  Click a row to see full prompt + response + PII report     │
│                                                              │
│  ┌─ PII Safety Summary ──────────────────────────────────┐  │
│  │ Total prompts today: 47                                │  │
│  │ Prompts with PII detected: 3 (6.4%)                    │  │
│  │ PII fields found: student_name (3x)                    │  │
│  │ All PII was masked before sending: YES                 │  │
│  └────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Diagnostic Placement Flow

### 3.1 What Happens on First Login

```
Student logs in for the first time
    │
    ▼
Avatar & preference setup (2 min)
    │
    ▼
"Let's find out what you already know!" (not a test)
    │
    ▼
Adaptive diagnostic:
    1. Start at student's enrolled grade level
    2. Present 3-5 questions per skill area
    3. If >80% correct → advance to harder skill
    4. If <40% correct → drop to easier skill
    5. Stop when ceiling and floor are found per subject
    │
    ▼
Results:
    - SkillMastery records created (20-30 skills)
    - StudentProfile.readiness scores populated
    - StudentProfile.literacy / math profiles initialized
    - LearningPath created with first 10 activities queued
    - Enrollment marked diagnostic_complete = true
    │
    ▼
Student starts first real activity immediately
```

### 3.2 Diagnostic Prompt Log

Every diagnostic decision is logged:
```
PromptLog entry:
  Type: PATH_DECISION
  Student: STU-0042
  System prompt: "You are placing a new student. Based on diagnostic results..."
  User message: "Grade 4 student. Results: addition 95%, subtraction 88%,
                 multiplication 45%, division 12%, fractions 0%..."
  Response: "Start multiplication at EASY difficulty. Skip division/fractions
            (prerequisites not met). Recommend Array Builder activities."
```

---

## 4. Prompt Templates

### 4.1 PATH_DECISION Prompt

```
System: You are an adaptive learning engine for {grade_level} students.

Rules:
- Never assign activities whose prerequisite skills are not PROFICIENT
- Respect adaptive_settings (max_consecutive_errors, break_frequency)
- Respect AI tutor settings (personality, should_avoid)
- If attention profile shows losing_focus_signs, insert break activities
- Use the student's preferred_modes and effective_activity_types
- Consider error_patterns to choose targeted activities
- If frustration_triggers include "consecutive_errors", limit hard content

Return JSON: {"nextActivities":[...], "reasoning":"..."}

Context:
- Student profile: {masked_profile_summary}
- Current mastery: {skill_mastery_data}
- Recent interactions: {last_10_interactions}
- Available activities: {filtered_activity_list}
```

### 4.2 PROFILE_UPDATE Prompt

```
System: You are updating a student's learning profile based on recent data.
Analyze the interaction patterns and update the profile sections.
Flag any concerning patterns (regression, disengagement, possible learning difficulties).

Context:
- Current profile: {current_profile}
- Last 7 days of interactions: {interaction_summary}
- Mastery changes: {mastery_deltas}
- Session patterns: {session_data}

Return JSON with updated profile sections only (don't repeat unchanged sections).
```

### 4.3 PARENT_COACHING Prompt

```
System: You are a teaching coach for a homeschool parent.
Generate ONE actionable tip for today based on what the child is working on.
Use simple language. Be encouraging. Never criticize the parent.

Context:
- Child's current focus skills: {current_skills}
- Today's planned activities: {schedule}
- Child's learning style: {learning_style_summary}
- Child's interests: {interests}

Return JSON: {"tip":"...", "activitySuggestion":"...", "materials":"..."}
```

### 4.4 RISK_ASSESSMENT Prompt

```
System: You are an early warning system for student learning.
Analyze this student's recent data and assess risk of falling behind.
Be specific about contributing factors and recommended interventions.

Context:
- Mastery trajectory (last 4 weeks): {mastery_trend}
- Engagement signals: {engagement_data}
- Session frequency: {session_frequency}
- Score trends: {score_trends}
- Peer comparison (anonymous): {cohort_percentiles}

Return JSON: {"riskLevel":"ON_TRACK|WATCH|AT_RISK|CRITICAL",
              "riskScore":0.0-1.0,
              "factors":[{"type":"...","description":"...","weight":0.0-1.0}],
              "recommendation":"..."}
```

---

## 5. Implementation Phases

### Amendment Phase 1: Student Profile + LLM Infrastructure
- Add `learn-profile.proto` and `learn-llm.proto`
- Run `make-bindings.sh`
- Create Profile, PromptLog, LLMConfig services
- Create LLM Simulator with PII scanner
- Create AI Monitor UI section (sidebar entry, prompt log table, PII summary)
- Wire LLM Simulator into adaptive engine
- **Verify**: Login → AI Monitor → see logged prompts with PII flags

### Amendment Phase 2: Diagnostic Flow
- Create diagnostic benchmark engine
- Wire into enrollment activation callback
- Create diagnostic UI in student player
- **Verify**: New student logs in → completes diagnostic → SkillMastery + Profile populated → first activity appears

### Amendment Phase 3: Profile Auto-Update
- Wire SkillMastery callback to update StudentProfile
- Wire session completion to update learning_style, attention, motivation
- Schedule weekly PROFILE_UPDATE prompt
- **Verify**: After 5 sessions, student profile shows learned preferences

### Amendment Phase 4: Parent Coaching
- Wire daily PARENT_COACHING prompt for each active family
- Display tip in guardian portal
- **Verify**: Guardian logs in → sees today's coaching tip with materials

### Amendment Phase 5: Risk + Analytics Computation
- Implement weekly RISK_ASSESSMENT batch job
- Implement cohort snapshot computation
- Implement growth record computation
- **Verify**: After 1 week of data → risk assessments appear → cohort snapshots generated

### Amendment Phase 6: Go Live with Real LLM
- Switch LLMConfig mode from SIMULATE to LIVE
- Review all prompt logs for PII leaks
- Set cost controls (max_daily_calls)
- **Verify**: Same prompts, real AI responses, data stays clean

---

## 6. Proto Import Map

| File | Imports |
|------|---------|
| learn-profile.proto | l8common.proto, api.proto, learn-content.proto |
| learn-llm.proto | l8common.proto, api.proto |

---

## 7. Traceability Matrix

| # | Gap | Amendment Phase |
|---|-----|----------------|
| 1 | Student profile is flat admin record | Phase 1 |
| 2 | No LLM integration | Phase 1 (simulator) → Phase 6 (live) |
| 3 | No PII safety controls | Phase 1 |
| 4 | No prompt visibility | Phase 1 (AI Monitor UI) |
| 5 | No diagnostic placement | Phase 2 |
| 6 | Profile never auto-updates | Phase 3 |
| 7 | No parent coaching intelligence | Phase 4 |
| 8 | Risk prediction is CRUD only | Phase 5 |
| 9 | Computed analytics don't compute | Phase 5 |
| 10 | No cost controls for LLM | Phase 1 (config) → Phase 6 (enforcement) |

---

## 8. Compliance

### Data Safety
- All prompts logged with PII scan results
- PII masking enabled by default in simulate and live modes
- Admin can review every prompt sent to the LLM via AI Monitor
- No student names, DOB, addresses, or health data sent to LLM unless masking fails (flagged)
- FERPA: prompt logs are audit records subject to same access controls as student data

### COPPA
- AI tutor personality and technology limits set by guardian, not child
- Child never sees or interacts with raw LLM — only the adaptive engine's decisions
