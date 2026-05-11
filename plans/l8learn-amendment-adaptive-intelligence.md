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
| `EvalImprt` | 20 | `ImportId` | `EvalImport` |
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

### 4.5 WORKSHEET_SCAN Profile Update Prompt

When a scanned worksheet is graded, the AI also analyzes handwriting quality, work patterns, and error patterns to update the student's profile:

```
System: You are analyzing a scanned worksheet for learning insights.
Beyond scoring right/wrong, analyze WHAT THE HANDWRITING AND WORK
PATTERNS reveal about this student's development.

Context:
- Student profile: {masked_profile}
- Worksheet: {subject}, {difficulty}, {question_count} questions
- Extracted answers with confidence scores: {answer_list}
- Handwriting analysis:
  - Overall quality: {quality}
  - Number/letter formation: {formation_notes}
  - Size consistency: {consistency}
  - Erasure count: {erasures}
- Work pattern analysis:
  - Scratch work visible: {yes/no}
  - Strategy detected: {strategy}
  - Completion pattern: {pattern}
  - Problems skipped: {count}
  - Quality degrades over time: {yes/no}

Return JSON:
{
  "fine_motor_updates": { ... },
  "math_updates": { ... },
  "attention_updates": { ... },
  "behavior_updates": { ... },
  "learning_style_updates": { ... },
  "insights": "narrative of what this worksheet reveals",
  "recommendations": "what to adjust in the adaptive engine"
}
```

What scanned worksheets reveal that digital activities cannot:

| Paper Evidence | Insight | Profile Update |
|---------------|---------|----------------|
| Erasure marks | Child self-corrects — positive metacognition | `behavior.successful_supports` |
| Scratch work (tally marks, drawings) | Problem-solving strategy | `math.preferred_tools`, `learning_style.preferred_modes` |
| Handwriting degrades after Q3 | Attention stamina ~8 minutes for written work | `attention.focus_academic_task_minutes` |
| Numbers 6/8 confused | Fine motor issue, not knowledge gap | `fine_motor.name_writing` |
| Neat top, messy bottom, skipped end | Fatigue pattern | `adaptive_settings.maximum_activity_length_minutes` |
| All errors off by same amount | Systematic error (skip-counting), not random | `math.error_patterns` |
| Correct answers but terrible handwriting | Knows content, motor issue | Distinguishes `fine_motor` from `math` |

#### Additional WorksheetScan Proto Fields

Add to the existing `WorksheetScan` message:

```protobuf
// Add to existing WorksheetScan message
HandwritingAnalysis handwriting = 19;
WorkPatternAnalysis work_patterns = 20;
string profile_insights = 21;
string profile_recommendations = 22;
repeated string profile_fields_updated = 23;

// Embedded child of WorksheetScan
message HandwritingAnalysis {
    string overall_quality = 1;            // "pre-writing", "emerging", "developing", "proficient"
    string letter_formation = 2;
    string number_formation = 3;
    string size_consistency = 4;           // "consistent", "degrades_over_time", "inconsistent"
    string line_adherence = 5;
    int32 erasure_count = 6;
    repeated string confused_characters = 7;  // "6/8", "b/d", "p/q"
}

// Embedded child of WorksheetScan
message WorkPatternAnalysis {
    bool scratch_work_present = 1;
    string strategy_detected = 2;          // "tally_marks", "number_line", "drawing", "mental"
    string completion_pattern = 3;         // "complete", "front_loaded", "scattered", "abandoned"
    int32 problems_skipped = 4;
    int32 problems_with_erasures = 5;
    bool quality_degrades = 6;
    int32 estimated_focus_minutes = 7;
}
```

---

### 4.6 EVAL_IMPORT — Professional Evaluation PDF Import

Professionals (speech therapists, occupational therapists, psychologists, reading specialists) produce evaluation PDFs. The AI reads these and maps findings to the student profile automatically, with guardian approval.

#### What Professional PDFs Contain

| Professional | PDF Contains | Maps To |
|-------------|-------------|---------|
| Speech therapist | Articulation scores, receptive/expressive language levels, therapy goals | `speech.clarity`, `speech.expressive_language`, `speech.current_goals` |
| Occupational therapist | Fine motor scores, grip assessment, visual-motor, sensory processing | `fine_motor.pencil_grip`, `fine_motor.cutting`, `sensory.sensitivities` |
| Psychologist (IEP) | Cognitive scores, processing speed, working memory, attention diagnosis | `attention`, `readiness.academic_readiness`, `adaptive_settings` |
| Developmental pediatrician | Diagnosis, medication, sensory profile, behavior plan | `health.medical_conditions`, `sensory`, `behavior`, `therapy.services` |
| Reading specialist | Reading level, phonemic awareness scores, fluency WPM | `literacy.reading_level`, `literacy.phonemic_awareness`, `literacy.reading_fluency_wpm` |
| School psychologist | Social-emotional assessment, behavior intervention plan | `social_emotional`, `behavior.triggers`, `behavior.redirect_strategies` |

#### LLM Prompt for Evaluation Import

```
System: You are extracting structured findings from a professional
evaluation report for a student. Map every finding to the student
profile schema. Only extract what is explicitly stated — do not infer.
Flag any findings that contradict the current profile.

Input:
- PDF text content: {extracted_text}
- Document type: {detected_type}
- Current student profile: {masked_profile}

Return JSON:
{
  "document_type": "speech_evaluation",
  "professional": "Dr. Sarah Miller, SLP",
  "evaluation_date": "2026-04-15",
  "findings": [
    {
      "profile_section": "speech",
      "profile_field": "expressive_language",
      "current_value": "developing",
      "new_value": "age_appropriate",
      "source_text": "Expressive language skills are now within normal limits",
      "confidence": 0.95
    }
  ],
  "contradictions": [
    {
      "profile_field": "speech.clarity",
      "current_value": "somewhat_clear",
      "document_says": "clear in structured settings, unclear in conversation",
      "recommendation": "Update to context-dependent value"
    }
  ],
  "new_therapy_info": {
    "service_type": "speech",
    "provider_name": "Dr. Sarah Miller",
    "frequency": "twice_weekly",
    "therapy_goals": ["articulation /r/ carryover", "narrative sequencing"],
    "home_practice": ["read aloud 10 min daily", "story retelling after dinner"]
  }
}
```

#### Guardian Approval Flow

The AI extracts findings but the **guardian must approve** before the profile is updated:

1. AI extracts findings → stores as PENDING
2. Guardian receives notification: "New evaluation imported — review needed"
3. Guardian sees each finding with source quote from the PDF
4. For each finding: `[Accept]` `[Reject]` `[Edit]`
5. Contradictions highlighted with AI recommendation
6. On "Accept All" or individual accepts → StudentProfile updated
7. All changes logged to l8events for audit trail
8. Adaptive engine recalibrates with updated profile

#### Protobuf

```protobuf
enum EvalDocumentType {
    EVAL_DOCUMENT_TYPE_UNSPECIFIED = 0;
    EVAL_DOCUMENT_TYPE_SPEECH = 1;
    EVAL_DOCUMENT_TYPE_OCCUPATIONAL_THERAPY = 2;
    EVAL_DOCUMENT_TYPE_PSYCHOLOGICAL = 3;
    EVAL_DOCUMENT_TYPE_IEP = 4;
    EVAL_DOCUMENT_TYPE_DEVELOPMENTAL = 5;
    EVAL_DOCUMENT_TYPE_READING_SPECIALIST = 6;
    EVAL_DOCUMENT_TYPE_BEHAVIORAL = 7;
    EVAL_DOCUMENT_TYPE_MEDICAL = 8;
    EVAL_DOCUMENT_TYPE_OTHER = 9;
}

enum EvalFindingStatus {
    EVAL_FINDING_STATUS_UNSPECIFIED = 0;
    EVAL_FINDING_STATUS_PENDING = 1;
    EVAL_FINDING_STATUS_ACCEPTED = 2;
    EVAL_FINDING_STATUS_REJECTED = 3;
    EVAL_FINDING_STATUS_EDITED = 4;
}

// @PrimeObject
message EvalImport {
    string import_id = 1;
    string student_id = 2;
    string uploaded_by = 3;
    EvalDocumentType document_type = 4;
    string professional_name = 5;
    int64 evaluation_date = 6;
    string file_path = 7;               // Uploaded PDF via Layer8FileUpload

    repeated EvalFinding findings = 8;
    repeated EvalContradiction contradictions = 9;

    bool all_reviewed = 10;
    int32 accepted_count = 11;
    int32 rejected_count = 12;
    bool applied_to_profile = 13;

    l8common.AuditInfo audit_info = 14;
}

message EvalImportList {
    repeated EvalImport list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child
message EvalFinding {
    string finding_id = 1;
    string profile_section = 2;
    string profile_field = 3;
    string current_value = 4;
    string new_value = 5;
    string source_text = 6;
    double confidence = 7;
    EvalFindingStatus status = 8;
    string edited_value = 9;
}

// Embedded child
message EvalContradiction {
    string profile_field = 1;
    string current_value = 2;
    string document_says = 3;
    string ai_recommendation = 4;
    EvalFindingStatus resolution = 5;
}
```

#### Service

| ServiceName | ServiceArea | PrimaryKey | Model |
|-------------|:-----------:|------------|-------|
| `EvalImprt` | 20 | `ImportId` | `EvalImport` |

#### ServiceCallback

- **After POST** (PDF uploaded): Read PDF text → detect document type → send to LLM (or simulator) with masked profile → store extracted findings as PENDING → notify guardian
- **After PUT** (guardian reviewed): For each ACCEPTED/EDITED finding → update StudentProfile → update therapy services → mark `applied_to_profile = true` → log to l8events → trigger adaptive engine recalibration

---

## 5. Data Completeness — UI Forms, Columns, and Mock Data

### 5.1 UI Forms and Columns for New Services

Each new service MUST have enums, columns, and forms defined in the same phase as the service.

**StudentProfile** — added to Students module (`learn-ui/students/people/`):
- Columns: profileId, studentId, overallDescription (truncated), readiness.academicReadiness, readiness.readingReadiness, readiness.mathReadiness
- Form: tabbed layout — Basic (summary, strengths, challenges, goals), Learning Style, Attention, Motivation, Literacy, Math, Speech, Motor, Sensory, Social-Emotional, Behavior, Technology, Health, Therapy, Adaptive Settings, AI Tutor Settings, Goals
- Detail popup uses Related Resources tab to link back to Student record

**LLMPromptLog** — added to AI Monitor section (`learn-ui/aimonitor/`):
- Columns: logId, type (enum renderer), studentId, containsPii (boolean), systemPromptTokens, userMessageTokens, responseTimeMs, timestamp (date)
- Form: read-only — shows full system prompt, user message, response, PII fields found
- **Immutable**: service rejects PUT and DELETE — UI has no edit/delete buttons

**LLMConfig** — added to AI Monitor section:
- Columns: configId, mode (enum), apiProvider, modelName, piiMaskingEnabled (boolean), promptLoggingEnabled, maxDailyCalls, callsToday
- Form: editable by admin only — mode selector, masking toggle, cost controls

**EvalImport** — added to Students module:
- Columns: importId, studentId, documentType (enum), professionalName, evaluationDate (date), allReviewed (boolean), acceptedCount, rejectedCount
- Form: custom wizard — PDF upload → AI extraction → findings review (accept/reject/edit per finding) → apply to profile

### 5.2 Mock Data for New Services

```
go/tests/mocks/
├── gen_learn_profiles.go         # 100 StudentProfiles with realistic learning data
├── gen_learn_promptlogs.go       # 50 simulated prompt logs (mixed types)
├── gen_learn_evals.go            # 10 EvalImports with findings (speech, OT, IEP)
```

**Phase ordering**: Profiles depend on StudentIDs (Phase 2 in original). PromptLogs depend on StudentIDs. EvalImports depend on StudentIDs + ProfileIDs.

**store.go additions**:
```go
ProfileIDs    []string
PromptLogIDs  []string
EvalImportIDs []string
```

**LLMConfig**: Seeded as a single record in Phase 1 mock generator with `mode: SIMULATE`.

### 5.3 Prompt Template Extraction (Duplication Prevention)

Sections 4.1–4.6 define 6 prompt templates with similar structure (system prompt, context sections, return format). Per `plan-duplication-audit.md`, these MUST be extracted into a shared template builder:

```go
// go/learn/adaptive/engine/prompt_builder.go (<200 lines)
type PromptBuilder struct {
    promptType  learn.LLMPromptType
    systemRole  string
    rules       []string
    context     map[string]string
    returnFormat string
}

func NewPromptBuilder(promptType learn.LLMPromptType) *PromptBuilder { ... }
func (b *PromptBuilder) SetRole(role string) *PromptBuilder { ... }
func (b *PromptBuilder) AddRule(rule string) *PromptBuilder { ... }
func (b *PromptBuilder) AddContext(key, value string) *PromptBuilder { ... }
func (b *PromptBuilder) SetReturnFormat(format string) *PromptBuilder { ... }
func (b *PromptBuilder) Build() (systemPrompt, userMessage string) { ... }
```

Each prompt type uses the builder — no copy-paste of prompt structure:
```go
// prompt_templates.go (<150 lines) — one function per type, config only
func BuildPathDecisionPrompt(profile, mastery, interactions string) (string, string) {
    return NewPromptBuilder(learn.LLM_PROMPT_TYPE_PATH_DECISION).
        SetRole("adaptive learning engine for {grade} students").
        AddRule("Never assign activities whose prerequisites are not PROFICIENT").
        AddRule("Respect adaptive_settings").
        AddContext("student_profile", profile).
        AddContext("mastery", mastery).
        AddContext("interactions", interactions).
        SetReturnFormat(`{"nextActivities":[...],"reasoning":"..."}`).
        Build()
}
```

**File budget**: `prompt_builder.go` ~200 lines, `prompt_templates.go` ~150 lines. No file exceeds 500 lines.

---

## 5A. LLM Simulator File Structure (Maintainability)

The LLM Simulator is split across multiple files to stay under 500 lines each:

```
go/learn/adaptive/engine/
├── llm_client.go          # ~80 lines — interface + mode switching (SIMULATE/LIVE/LOG_ONLY)
├── llm_simulator.go       # ~120 lines — deterministic response generation per prompt type
├── prompt_builder.go      # ~200 lines — shared template builder (prevents duplication)
├── prompt_templates.go    # ~150 lines — one function per prompt type (config only)
└── prompt_logger.go       # ~100 lines — logs prompts to PromptLog service + PII scan
```

Total: ~650 lines across 5 files. No single file exceeds 500.

---

## 5B. Immutability

| Entity | Immutable? | Backend Enforcement | UI |
|--------|:----------:|--------------------|----|
| LLMPromptLog | YES | Reject PUT and DELETE in ServiceCallback Before hook | Read-only table, no edit/delete buttons |
| LLMConfig | NO (admin-editable) | Allow PUT for admin role only | Edit form for admin |
| StudentProfile | NO (auto-updated) | Allow PUT/PATCH from engine and guardian | Read-only for teacher, editable sections for guardian |
| EvalImport | PARTIAL | Allow PUT (for guardian review) but reject DELETE | Edit (review flow) but no delete |

---

## 5C. Deployment

No new deployable binaries. All 4 new services are added to the existing `learn_demo` backend process via `services/activate_students.go` (Profile, EvalImport) and `services/activate_adaptive.go` (PromptLog, LLMConfig). No new Dockerfile, build.sh, or K8s YAML required.

**run-local.sh update**: Add mock data generators to the build step:
```bash
# Already built: cd tests/mocks/cmd && go build -o ../../../demo/mocks_demo
# No change needed — new generators are compiled into the same mocks_demo binary
```

---

## 5D. Mobile Parity

All new UI sections MUST have mobile equivalents:

| Desktop Section | Mobile Equivalent |
|----------------|-------------------|
| AI Monitor (prompt log table) | Mobile nav → System → AI Monitor (Layer8MTable) |
| Student Profile (tabbed detail) | Mobile popup with tabbed sections (Layer8MPopup with onTabChange) |
| Eval Import (upload + review) | Mobile file upload + approval card list |
| LLM Config (settings) | Mobile form in System → AI Settings |

**Mobile files to create** (Phase 1):
- `m/js/aimonitor/aimonitor-enums.js`
- `m/js/aimonitor/aimonitor-columns.js`
- `m/js/aimonitor/aimonitor-forms.js`
- `m/js/aimonitor/aimonitor-index.js`
- Add to `m/app.html` script loading
- Add to mobile nav config

---

## 6. Implementation Phases (Updated)

### Amendment Phase 1: Student Profile + LLM Infrastructure + Eval Import
- Add `learn-profile.proto`, `learn-llm.proto`, and `learn-eval.proto`
- Run `make-bindings.sh`
- Create Profile, PromptLog, LLMConfig, EvalImport services + ServiceCallbacks
- Register types in `shared_students.go` (Profile, EvalImport) and `shared_adaptive.go` (PromptLog, LLMConfig)
- LLMPromptLog: immutable — reject PUT/DELETE in Before hook
- Create LLM Simulator: `llm_client.go`, `llm_simulator.go`, `prompt_builder.go`, `prompt_templates.go`, `prompt_logger.go`
- Reuse `l8agent/masking` for PII — do NOT build new scanner
- Create AI Monitor UI section (desktop): sidebar entry, prompt log table (Type, Student, PII, Tokens, Time columns), PII summary dashboard, LLM Config editor
- Create AI Monitor UI section (mobile): nav entry, mobile table, mobile config form
- Create Student Profile UI: tabbed form in Students module (desktop + mobile)
- Create Eval Import UI: PDF upload, findings review, accept/reject/edit flow (desktop + mobile)
- Create mock data generators: `gen_learn_profiles.go`, `gen_learn_promptlogs.go`, `gen_learn_evals.go`
- Update `store.go` with ProfileIDs, PromptLogIDs, EvalImportIDs
- Wire into `learn_phases.go`
- **Verify**: Login → AI Monitor → see logged prompts with PII flags. Click Students → click a student → see Profile tab with readiness scores. Upload a sample evaluation PDF → see extracted findings → approve → profile updated. Check mobile: same sections visible.

### Amendment Phase 2: Diagnostic Flow
- Create diagnostic benchmark engine (adaptive placement algorithm)
- Wire into enrollment activation callback (After status → ACTIVE)
- Create diagnostic UI in student player (avatar setup → "not a test" → adaptive questions → results)
- Algorithm: start at grade level, 3-5 questions per skill, advance if >80%, drop if <40%, stop at ceiling/floor
- PromptLog: log the PATH_DECISION prompt that generates the initial learning path
- **Verify**: New student logs in → completes diagnostic → SkillMastery + Profile populated → LearningPath created → first activity appears. Check AI Monitor: diagnostic prompt logged with PII masked.

### Amendment Phase 3: Profile Auto-Update
- Wire SkillMastery callback to update StudentProfile (readiness, literacy, math sections)
- Wire session completion to update learning_style, attention, motivation (from interaction patterns)
- Wire worksheet scan callback to update fine_motor, attention, behavior (from handwriting/work pattern analysis)
- Schedule weekly PROFILE_UPDATE prompt via l8alarms scheduler pattern
- PromptLog: log all profile update prompts
- **Verify**: After 5 sessions, student profile shows learned preferences (activity type preferences, attention minutes, error patterns). Check AI Monitor: PROFILE_UPDATE prompts logged.

### Amendment Phase 4: Parent Coaching
- Create daily PARENT_COACHING scheduled job (l8alarms scheduler pattern)
- Trigger: once per active family per day (morning)
- In simulate mode: generate coaching tip from simulated response and store it
- Display tip in guardian portal (desktop + mobile)
- PromptLog: log each coaching prompt
- **Verify**: Guardian logs in → sees today's coaching tip with materials and activity suggestion. Check AI Monitor: PARENT_COACHING prompt logged.

### Amendment Phase 5: Risk + Analytics Computation
- Implement weekly RISK_ASSESSMENT batch job (l8alarms scheduler pattern)
- Implement cohort snapshot computation (weekly classroom, monthly school/district)
- Implement growth record computation (triggered by SkillMastery changes)
- Implement content effectiveness computation (quarterly)
- All batch jobs log their prompts to PromptLog
- **Verify**: After 1 week of data → risk assessments appear in History → Risk table. Cohort snapshots generated. Growth records populated. Check AI Monitor: batch prompts logged.

### Amendment Phase 6: Go Live with Real LLM
- Switch LLMConfig mode from SIMULATE to LIVE
- Review ALL prompt logs for PII leaks (admin uses AI Monitor)
- Verify PII masking works (no real names in prompts)
- Set cost controls (max_daily_calls)
- **Verify**: Same prompts, real AI responses, data stays clean. PromptLog shows `mode: LIVE`. No PII flags.

### Amendment Phase 7: End-to-End Verification
For every section affected by this amendment:
1. Admin navigates to AI Monitor → sees prompt log table → clicks a row → sees full prompt + response + PII report
2. Admin changes LLM mode to SIMULATE → sees mode reflected in config table
3. Click Students → click a student → see Profile tab with readiness scores, learning style, attention profile
4. Guardian uploads evaluation PDF → sees extracted findings → approves → StudentProfile updated
5. Guardian rejects a finding → profile NOT updated for that field
6. New student logs in → completes diagnostic → Profile + Mastery + Path created → first activity loads
7. After 5 sessions → profile auto-updates → learning_style reflects observed preferences
8. Guardian logs in → sees today's coaching tip
9. After 1 week → risk assessments, cohort snapshots, growth records populated
10. All of the above verified on BOTH desktop and mobile

Sections to verify:
- [ ] AI Monitor (desktop + mobile)
- [ ] Student Profile (desktop + mobile)
- [ ] Eval Import (desktop + mobile)
- [ ] Diagnostic flow (student player)
- [ ] Profile auto-update (after sessions)
- [ ] Parent coaching (guardian portal, desktop + mobile)
- [ ] Risk + analytics (History section)
- [ ] PromptLog immutability (no edit/delete buttons, PUT rejected)

---

## 6. Proto Import Map

| File | Imports | Why |
|------|---------|-----|
| learn-profile.proto | l8common.proto, api.proto | AuditInfo, L8MetaData. Does NOT need learn-content.proto — no shared types used |
| learn-llm.proto | l8common.proto, api.proto | AuditInfo, L8MetaData |
| learn-eval.proto | l8common.proto, api.proto | AuditInfo, L8MetaData. References EvalDocumentType and EvalFindingStatus (defined in same file) |

---

## 7. Pre-Implementation Checklist

| Question | Answer |
|----------|--------|
| Canonical reference project? | `../l8vendingmachine` for service patterns, `../l8erp` for UI module structure |
| How many new proto files? | 2 (learn-profile.proto, learn-llm.proto) |
| Types shared across files? | None — both are independent |
| Read-only services? | PromptLog (read-only in UI — created by engine, viewed by admin) |
| Complex After hooks? | Profile (auto-update from mastery changes), PromptLog (PII scan on create) |
| External L8 services needed? | l8agent (for LLM client pattern), l8events (audit logging) |
| Service activation order? | Profile → LLMConfig → PromptLog (config must exist before logging starts) |
| How many portals affected? | Admin (AI Monitor section), Guardian (coaching tip) |
| What user sees Phase 1? | AI Monitor sidebar → prompt log table with Type, Student, PII, Tokens, Time columns |
| Custom UI needed? | AI Monitor section (table + PII summary dashboard) |
| Server port? | 2773 (with security config) or 4443 (without) |
| CreateWebServer pattern? | Same as current l8learn UI — `common.CreateWebServer("web", ui.RegisterTypes)` |
| VNet port? | 10005 (unchanged) |
| API signatures to verify? | `l8c.RegisterType`, `common.ActivateService`, `NewValidation().Build()` |
| l8ui module-factory-core.js? | Yes — already included in app.html |

## 8. Type Registration (Phase 1 — BLOCKING)

These MUST be added to the UI server type registration in the SAME phase as service creation:

```go
// In go/learn/ui/shared_students.go (add to existing function)
l8c.RegisterType(resources, &learn.StudentProfile{}, &learn.StudentProfileList{}, "ProfileId")
l8c.RegisterType(resources, &learn.EvalImport{}, &learn.EvalImportList{}, "ImportId")

// In go/learn/ui/shared_adaptive.go (add to existing function)
l8c.RegisterType(resources, &learn.LLMPromptLog{}, &learn.LLMPromptLogList{}, "LogId")
l8c.RegisterType(resources, &learn.LLMConfig{}, &learn.LLMConfigList{}, "ConfigId")
```

Without these, the AI Monitor tables will show "Cannot find pb for method GET".

## 9. API Signatures to Verify Before Writing Code

Before implementing ANY Go code in this amendment, verify these signatures against `go/vendor/github.com/saichler/l8common/`:

```bash
# RegisterType signature
grep -A3 "func RegisterType" go/vendor/github.com/saichler/l8common/go/common/*.go

# ActivateService config
grep -A10 "type ServiceConfig" go/vendor/github.com/saichler/l8common/go/common/*.go

# NewValidation builder methods
grep "func.*VB.*func" go/vendor/github.com/saichler/l8common/go/common/validation_builder.go | head -10

# CreateResources signature
grep -A3 "func CreateResources" go/vendor/github.com/saichler/l8common/go/common/*.go
```

Do NOT assume signatures from the PRD examples. Read the actual function. Then call it.

## 10. Security — LLM API Credentials

- Anthropic API key stored in security config JSON under `credentials.Anthropic.creds.API_KEY`
- Follows existing pattern from `../l8secure/go/secure/plugin/learn/learn.json`
- Key retrieved at runtime: `nic.Resources().Security().Credential("Anthropic", "API_KEY", resources)`
- Only `admin` role can change `LLMConfig.mode` from SIMULATE to LIVE (security rules in learn.json)
- All LLM calls logged to PromptLog regardless of mode (audit requirement)

## 11. Traceability Matrix

| # | Gap | Amendment Phase |
|---|-----|----------------|
| 1 | Student profile is flat admin record | Phase 1 |
| 2 | No LLM integration | Phase 1 (simulator) → Phase 6 (live) |
| 3 | No PII safety controls | Phase 1 (reuse l8agent/masking) |
| 4 | No prompt visibility | Phase 1 (AI Monitor UI) |
| 5 | No diagnostic placement | Phase 2 |
| 6 | Profile never auto-updates | Phase 3 |
| 7 | No parent coaching intelligence | Phase 4 |
| 8 | Risk prediction is CRUD only | Phase 5 |
| 9 | Computed analytics don't compute | Phase 5 |
| 10 | No cost controls for LLM | Phase 1 (config) → Phase 6 (enforcement) |
| 11 | Worksheet scans don't update student profile | Phase 3 (handwriting + work pattern analysis) |
| 12 | No professional evaluation import | Phase 1 (EvalImport service, PDF extraction, guardian approval) |
| 13 | Existing readOnly configs broken (svc() misuse) | Fixed (committed) |
| 14 | No forms/columns for new services | Phase 1 (Section 5.1 defines all) |
| 15 | No mock data for new services | Phase 1 (Section 5.2 defines generators) |
| 16 | Mobile parity not addressed | Phase 1 (Section 5D defines mobile equivalents) |
| 17 | Prompt template duplication | Phase 1 (Section 5.3 — shared PromptBuilder) |
| 18 | LLMPromptLog immutability not enforced | Phase 1 (Section 5B — reject PUT/DELETE) |
| 19 | LLM Simulator exceeds 500 lines | Phase 1 (Section 5A — split across 5 files) |
| 20 | No final verification phase | Phase 7 (end-to-end smoke test) |
| 21 | No deployment artifact clarification | Section 5C — explicitly: no new binaries |
| 22 | run-local.sh not updated | Section 5C — no change needed (same binary) |

---

## 12. Ecosystem Reuse (Don't Rebuild What Exists)

The following components ALREADY exist in the Layer 8 ecosystem and MUST be reused — not rebuilt:

| Component | Source Project | Reuse How |
|-----------|---------------|-----------|
| PII masking | `../l8agent/go/masking/proxy.go` | Import `masking.MaskText()`, `masking.MaskJSON()`, `masking.TokenMap` for round-trip mask/unmask. Handles SSN, email, phone, money patterns. Do NOT build a new PII scanner. |
| LLM client | `../l8agent/go/llm/client.go` | Import the Anthropic Claude client. Supports system prompt + messages + tools. 60s timeout, 4096 max tokens. |
| Related records popup | `l8ui/related/layer8d-related-resources.js` | Use to show Student → StudentProfile link in detail popup. Adds "Related" tab automatically. |
| Scheduled jobs | `../l8alarms/go/alm/escalation/scheduler.go` | Follow the scheduler pattern for weekly risk assessment batch, daily parent coaching, and quarterly content effectiveness computation. |
| Mobile UI framework | `../l8vendingmachine/go/vend/ui/web/m/app.html` | Copy structure for mobile module registry, nav config, forms, and reference pickers. |

### What This Changes in the Amendment

1. **Section 2.6 (PII Scanner)**: Replace custom `PIIScanner` with `l8agent/masking` package. The `masking.MaskText()` function already does regex-based PII detection. The `TokenMap` handles the `Student_A → Jake Martinez` round-trip.

2. **Section 2.5 (LLM Simulator)**: The simulator wraps the existing `l8agent/llm/client.go` — in simulate mode it skips the HTTP call and returns deterministic responses, but uses the same prompt construction and token counting.

3. **Phase 1 scope reduction**: No PII scanner to build from scratch. Import `l8agent/masking` and wire it into the prompt logging pipeline.

### readOnly Services Fix

The existing `history-config.js` and `assessment-config.js` pass `{ readOnly: true }` as the 6th argument to `svc()`, but `svc(key, label, icon, endpoint, model, viewType, ...)` expects a string. This silently fails — the services render without readOnly protection.

**Fix**: Use inline service objects instead of `svc()` for readOnly services:
```javascript
// WRONG — svc() 6th arg is viewType, not options
svc('growth', 'Growth', 'icon', '/60/Growth', 'GrowthRecord', { readOnly: true })

// CORRECT — inline object with readOnly property
{ key: 'growth', label: 'Growth', icon: 'icon', endpoint: '/60/Growth', model: 'GrowthRecord', readOnly: true }
```

---

## 13. Compliance

### Data Safety
- All prompts logged with PII scan results
- PII masking enabled by default in simulate and live modes
- Admin can review every prompt sent to the LLM via AI Monitor
- No student names, DOB, addresses, or health data sent to LLM unless masking fails (flagged)
- FERPA: prompt logs are audit records subject to same access controls as student data

### COPPA
- AI tutor personality and technology limits set by guardian, not child
- Child never sees or interacts with raw LLM — only the adaptive engine's decisions
