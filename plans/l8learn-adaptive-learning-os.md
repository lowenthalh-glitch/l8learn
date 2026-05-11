# L8Learn: Adaptive Learning Operating System

## Overview

L8Learn is an adaptive learning platform built on the Layer 8 ecosystem that personalizes education for each child. It continuously observes how a student learns — not just what they get right or wrong, but how they think, where they hesitate, what strategies they use — and dynamically adjusts the learning path in real-time using AI.

**Project name:** `l8learn`
**Repository:** `github.com/saichler/l8learn`
**PREFIX:** `/learn`
**License:** Apache 2.0

---

## 1. Architecture

### 1.1 Service Areas

| Service Area | Byte ID | Domain | Description |
|-------------|---------|--------|-------------|
| Content | 10 | Curriculum & media | Courses, units, lessons, activities, worksheets, family activities, real-world lessons, projects |
| Students | 20 | People & organizations | Students, guardians, teachers, classrooms, schools, districts, enrollment, families, compliance, pods |
| Adaptive | 30 | Learning engine | Learning paths, skill mastery, adaptation rules, skill graph, daily schedules |
| Assessment | 40 | Evaluation | Interactions, session logs, scores, benchmarks, worksheet scans |
| Analytics | 50 | Reporting | Progress reports, class summaries, district rollups, engagement |
| History | 60 | Historical analytics | Growth records, cohort snapshots, risk assessments, standards mastery, content effectiveness |
| Collaboration | 70 | Social learning | Study groups, AI-moderated chat, peer tutoring, team challenges |

### 1.2 Process Topology

```
l8learn/
├── go/
│   ├── learn/
│   │   ├── common/              # PREFIX, shared constants
│   │   ├── content/             # Content services (area 10)
│   │   │   ├── courses/         # CourseService
│   │   │   ├── units/           # UnitService
│   │   │   ├── lessons/         # LessonService
│   │   │   ├── activities/      # ActivityService (ServiceName: "Activity")
│   │   │   ├── worksheets/      # WorksheetService (ServiceName: "Worksheet")
│   │   │   ├── familyactivities/ # FamilyActivityService (ServiceName: "FamActvty")
│   │   │   ├── realworld/       # RealWorldLessonService (ServiceName: "RealWorld")
│   │   │   └── projects/        # ProjectService (ServiceName: "Project")
│   │   ├── students/            # Student services (area 20)
│   │   │   ├── students/        # StudentService (ServiceName: "Student")
│   │   │   ├── guardians/       # GuardianService (ServiceName: "Guardian")
│   │   │   ├── teachers/        # TeacherService (ServiceName: "Teacher")
│   │   │   ├── classrooms/      # ClassroomService (ServiceName: "Classroom")
│   │   │   ├── schools/         # SchoolService (ServiceName: "School")
│   │   │   ├── districts/       # DistrictService (ServiceName: "District")
│   │   │   ├── enrollments/     # EnrollmentService (ServiceName: "Enroll")
│   │   │   ├── families/        # FamilyService (ServiceName: "Family")
│   │   │   ├── compliance/      # ComplianceService (ServiceName: "Comply")
│   │   │   └── pods/            # LearningPodService (ServiceName: "Pod")
│   │   ├── adaptive/            # Adaptive engine (area 30)
│   │   │   ├── paths/           # PathService (ServiceName: "LearnPath")
│   │   │   ├── mastery/         # MasteryService (ServiceName: "Mastery")
│   │   │   ├── skills/          # SkillService (ServiceName: "Skill")
│   │   │   ├── rules/           # AdaptRuleService (ServiceName: "AdaptRule")
│   │   │   └── schedules/       # ScheduleService (ServiceName: "Schedule")
│   │   ├── assessment/          # Assessment services (area 40)
│   │   │   ├── sessions/        # SessionService (ServiceName: "LearnSess")
│   │   │   ├── scores/          # ScoreService (ServiceName: "Score")
│   │   │   ├── benchmarks/      # BenchmarkService (ServiceName: "Benchmark")
│   │   │   └── worksheetscans/  # WorksheetScanService (ServiceName: "WkshtScan")
│   │   ├── analytics/           # Analytics services (area 50)
│   │   │   ├── progress/        # ProgressService (ServiceName: "Progress")
│   │   │   └── engagement/      # EngagementService (ServiceName: "Engage")
│   │   ├── history/             # Historical analytics (area 60)
│   │   │   ├── growth/          # GrowthService (ServiceName: "Growth")
│   │   │   ├── cohorts/         # CohortService (ServiceName: "Cohort")
│   │   │   ├── risk/            # RiskAssessmentService (ServiceName: "RiskAssmt")
│   │   │   ├── standards/       # StandardMasteryService (ServiceName: "StdMastry")
│   │   │   └── effectiveness/   # ContentEffectService (ServiceName: "CntEffect")
│   │   ├── collab/              # Collaboration Center (area 70)
│   │   │   ├── groups/          # CollabGroupService (ServiceName: "Collab")
│   │   │   ├── messages/        # CollabMessageService (ServiceName: "CollabMsg")
│   │   │   ├── tutoring/        # TutorMatchService (ServiceName: "TutorPair")
│   │   │   └── challenges/      # ChallengeService (ServiceName: "Challenge")
│   │   ├── services/            # Service activation (activate_content.go, etc.)
│   │   ├── ui/
│   │   │   ├── main.go          # UI server + type registration
│   │   │   ├── web/             # Desktop web assets
│   │   │   │   ├── app.html
│   │   │   │   ├── login.html
│   │   │   │   ├── login.json
│   │   │   │   ├── l8ui/        # Shared UI library (submodule)
│   │   │   │   ├── js/          # sections.js, reference registries
│   │   │   │   ├── sections/    # Section HTML per module
│   │   │   │   ├── learn-ui/    # Project-specific UI (nav configs, SVG templates)
│   │   │   │   └── m/           # Mobile web assets
│   │   │   │       ├── app.html
│   │   │   │       └── js/
│   │   │   ├── build.sh
│   │   │   └── Dockerfile
│   │   ├── main/                # Backend main entry point
│   │   │   ├── main.go
│   │   │   ├── build.sh
│   │   │   └── Dockerfile
│   │   └── vnet/                # Virtual network
│   │       ├── main.go
│   │       ├── build.sh
│   │       └── Dockerfile
│   ├── logs/
│   │   ├── agent/
│   │   └── vnet/
│   ├── tests/
│   │   └── mocks/
│   │       ├── cmd/
│   │       ├── data.go
│   │       ├── store.go
│   │       └── gen_*.go
│   ├── types/
│   │   └── learn/               # Generated .pb.go files
│   ├── go.mod
│   ├── run-local.sh
│   └── build-all-images.sh
├── proto/
│   ├── make-bindings.sh
│   ├── learn-content.proto
│   ├── learn-students.proto
│   ├── learn-adaptive.proto
│   ├── learn-assessment.proto
│   ├── learn-analytics.proto
│   ├── learn-history.proto
│   ├── learn-homeschool.proto
│   ├── learn-collab.proto
│   └── api.proto                # Auto-downloaded shared types
└── k8s/
    ├── deploy.sh
    ├── undeploy.sh
    ├── vnet.yaml
    ├── learn.yaml
    ├── web.yaml
    ├── log-vnet.yaml
    └── log-agent.yaml
```

---

## 2. Protobuf Design

### 2.1 learn-content.proto — Curriculum & Content (Service Area 10)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum SubjectType {
    SUBJECT_TYPE_UNSPECIFIED = 0;
    SUBJECT_TYPE_MATH = 1;
    SUBJECT_TYPE_READING = 2;
    SUBJECT_TYPE_SCIENCE = 3;
    SUBJECT_TYPE_WRITING = 4;
    SUBJECT_TYPE_SOCIAL_STUDIES = 5;
}

enum GradeLevel {
    GRADE_LEVEL_UNSPECIFIED = 0;
    GRADE_LEVEL_PRE_K = 1;
    GRADE_LEVEL_K = 2;
    GRADE_LEVEL_1 = 3;
    GRADE_LEVEL_2 = 4;
    GRADE_LEVEL_3 = 5;
    GRADE_LEVEL_4 = 6;
    GRADE_LEVEL_5 = 7;
    GRADE_LEVEL_6 = 8;
    GRADE_LEVEL_7 = 9;
    GRADE_LEVEL_8 = 10;
}

enum DifficultyLevel {
    DIFFICULTY_LEVEL_UNSPECIFIED = 0;
    DIFFICULTY_LEVEL_INTRO = 1;
    DIFFICULTY_LEVEL_EASY = 2;
    DIFFICULTY_LEVEL_MEDIUM = 3;
    DIFFICULTY_LEVEL_HARD = 4;
    DIFFICULTY_LEVEL_CHALLENGE = 5;
}

enum ContentStatus {
    CONTENT_STATUS_UNSPECIFIED = 0;
    CONTENT_STATUS_DRAFT = 1;
    CONTENT_STATUS_REVIEW = 2;
    CONTENT_STATUS_PUBLISHED = 3;
    CONTENT_STATUS_ARCHIVED = 4;
}

enum ActivityType {
    ACTIVITY_TYPE_UNSPECIFIED = 0;
    ACTIVITY_TYPE_INTERACTIVE = 1;      // Drag-drop, manipulatives
    ACTIVITY_TYPE_MULTIPLE_CHOICE = 2;
    ACTIVITY_TYPE_FREE_RESPONSE = 3;
    ACTIVITY_TYPE_MATCHING = 4;
    ACTIVITY_TYPE_ORDERING = 5;         // Sequence ordering
    ACTIVITY_TYPE_FILL_BLANK = 6;
    ACTIVITY_TYPE_READING_PASSAGE = 7;
    ACTIVITY_TYPE_VIDEO = 8;
    ACTIVITY_TYPE_GAME = 9;
}

enum QuestionType {
    QUESTION_TYPE_UNSPECIFIED = 0;
    QUESTION_TYPE_SINGLE_CHOICE = 1;
    QUESTION_TYPE_MULTI_CHOICE = 2;
    QUESTION_TYPE_NUMERIC = 3;
    QUESTION_TYPE_TEXT = 4;
    QUESTION_TYPE_DRAG_DROP = 5;
    QUESTION_TYPE_DRAWING = 6;
}

enum WorksheetStatus {
    WORKSHEET_STATUS_UNSPECIFIED = 0;
    WORKSHEET_STATUS_DRAFT = 1;
    WORKSHEET_STATUS_GENERATED = 2;
    WORKSHEET_STATUS_DISTRIBUTED = 3;
    WORKSHEET_STATUS_SCORED = 4;
}

enum WorksheetLayout {
    WORKSHEET_LAYOUT_UNSPECIFIED = 0;
    WORKSHEET_LAYOUT_STANDARD = 1;           // Questions top-to-bottom
    WORKSHEET_LAYOUT_TWO_COLUMN = 2;         // Side-by-side columns
    WORKSHEET_LAYOUT_GRID = 3;               // Grid of problems (math facts)
    WORKSHEET_LAYOUT_PASSAGE = 4;            // Reading passage + questions below
    WORKSHEET_LAYOUT_MATCHING = 5;           // Left-right matching
    WORKSHEET_LAYOUT_FILL_BLANK = 6;         // Inline blanks in text
}

// ============================================================================
// CONTENT MESSAGES
// ============================================================================

// @PrimeObject
message Course {
    string course_id = 1;
    string name = 2;
    string description = 3;
    SubjectType subject = 4;
    GradeLevel min_grade = 5;
    GradeLevel max_grade = 6;
    ContentStatus status = 7;
    string thumbnail_url = 8;
    repeated string standard_ids = 9;        // Curriculum standard alignment
    repeated string prerequisite_course_ids = 10;
    int32 estimated_hours = 11;
    string author_id = 12;
    map<string, string> custom_fields = 13;
    l8common.AuditInfo audit_info = 14;
}

message CourseList {
    repeated Course list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message Unit {
    string unit_id = 1;
    string course_id = 2;                    // Parent course
    string name = 3;
    string description = 4;
    int32 sequence_order = 5;
    ContentStatus status = 6;
    repeated string skill_ids = 7;           // Skills taught in this unit
    repeated string prerequisite_unit_ids = 8;
    int32 estimated_minutes = 9;
    map<string, string> custom_fields = 10;
    l8common.AuditInfo audit_info = 11;
}

message UnitList {
    repeated Unit list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message Lesson {
    string lesson_id = 1;
    string unit_id = 2;                      // Parent unit
    string name = 3;
    string description = 4;
    int32 sequence_order = 5;
    DifficultyLevel difficulty = 6;
    ContentStatus status = 7;
    repeated string skill_ids = 8;           // Skills practiced
    string instruction_text = 9;             // Markdown content
    string media_url = 10;                   // Video/animation URL
    int32 estimated_minutes = 11;
    map<string, string> custom_fields = 12;
    l8common.AuditInfo audit_info = 13;
}

message LessonList {
    repeated Lesson list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message Activity {
    string activity_id = 1;
    string lesson_id = 2;                    // Parent lesson
    string name = 3;
    string instructions = 4;
    ActivityType activity_type = 5;
    DifficultyLevel difficulty = 6;
    ContentStatus status = 7;
    repeated string skill_ids = 8;
    int32 points_possible = 9;
    int32 estimated_seconds = 10;
    int32 max_attempts = 11;                 // 0 = unlimited
    bool hints_enabled = 12;
    int32 hint_count = 13;
    string config_json = 14;                 // Activity-type-specific config (manipulative params, game config)
    repeated Question questions = 15;        // Embedded child
    map<string, string> custom_fields = 16;
    l8common.AuditInfo audit_info = 17;
}

message ActivityList {
    repeated Activity list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of Activity — NOT a Prime Object
message Question {
    string question_id = 1;
    int32 sequence_order = 2;
    QuestionType question_type = 3;
    string prompt_text = 4;                  // The question (markdown)
    string prompt_media_url = 5;             // Image/diagram for the question
    DifficultyLevel difficulty = 6;
    repeated AnswerOption options = 7;       // For choice-based questions
    string correct_answer = 8;              // For numeric/text questions
    string explanation = 9;                  // Shown after answering
    int32 points = 10;
    repeated string skill_ids = 11;          // Specific skills this question tests
    repeated string hint_texts = 12;         // Progressive hints
}

// Embedded child of Question
message AnswerOption {
    string option_id = 1;
    string text = 2;
    string media_url = 3;
    bool is_correct = 4;
    string feedback = 5;                     // Shown when this option is selected
}

// ============================================================================
// PRINTABLE WORKSHEETS
// ============================================================================

// @PrimeObject
// Teacher-generated printable worksheet, optionally personalized per student
message Worksheet {
    string worksheet_id = 1;
    string name = 2;
    string teacher_id = 3;
    SubjectType subject = 4;
    GradeLevel grade_level = 5;
    DifficultyLevel difficulty = 6;
    WorksheetStatus status = 7;
    WorksheetLayout layout = 8;

    // Content selection
    repeated string skill_ids = 9;           // Skills to draw questions from
    repeated string activity_ids = 10;       // Or pick specific activities
    int32 question_count = 11;
    bool shuffle_questions = 12;             // Randomize order per student
    bool shuffle_options = 13;               // Randomize answer choices
    bool include_word_bank = 14;
    bool include_number_line = 15;
    bool include_grid_paper = 16;
    string instructions_text = 17;           // Custom header instructions
    bool personalized = 18;                  // Generate per-student variants based on mastery

    // Distribution
    repeated string student_ids = 19;        // Empty = whole classroom
    string classroom_id = 20;
    int64 due_date = 21;

    // Generated output
    string pdf_storage_path = 22;            // Via Layer8FileUpload
    string answer_key_path = 23;

    // Scoring (after hand-grading or scan)
    repeated WorksheetScore scores = 24;     // Embedded child

    map<string, string> custom_fields = 25;
    l8common.AuditInfo audit_info = 26;
}

message WorksheetList {
    repeated Worksheet list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child — one per student who received the worksheet
message WorksheetScore {
    string student_id = 1;
    int32 questions_correct = 2;
    int32 questions_total = 3;
    double score_percent = 4;
    int64 scored_date = 5;
    string scored_by = 6;                    // Teacher who graded it
    bool fed_to_mastery = 7;                 // Has this score been fed back to SkillMastery?
}
```

### 2.2 learn-students.proto — People & Organizations (Service Area 20)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum StudentStatus {
    STUDENT_STATUS_UNSPECIFIED = 0;
    STUDENT_STATUS_ACTIVE = 1;
    STUDENT_STATUS_INACTIVE = 2;
    STUDENT_STATUS_GRADUATED = 3;
    STUDENT_STATUS_TRANSFERRED = 4;
}

enum GuardianRelation {
    GUARDIAN_RELATION_UNSPECIFIED = 0;
    GUARDIAN_RELATION_PARENT = 1;
    GUARDIAN_RELATION_GRANDPARENT = 2;
    GUARDIAN_RELATION_SIBLING = 3;
    GUARDIAN_RELATION_LEGAL_GUARDIAN = 4;
    GUARDIAN_RELATION_OTHER = 5;
}

enum TeacherRole {
    TEACHER_ROLE_UNSPECIFIED = 0;
    TEACHER_ROLE_PRIMARY = 1;
    TEACHER_ROLE_SPECIALIST = 2;
    TEACHER_ROLE_AIDE = 3;
    TEACHER_ROLE_SUBSTITUTE = 4;
    TEACHER_ROLE_ADMIN = 5;
}

enum EnrollmentStatus {
    ENROLLMENT_STATUS_UNSPECIFIED = 0;
    ENROLLMENT_STATUS_DRAFT = 1;             // Admin created, not yet invited
    ENROLLMENT_STATUS_INVITED = 2;           // Guardian invite sent
    ENROLLMENT_STATUS_PENDING_CONSENT = 3;   // Guardian registered, consent not yet given
    ENROLLMENT_STATUS_CONSENTED = 4;         // Guardian signed, student not yet active
    ENROLLMENT_STATUS_ACTIVE = 5;            // Student can log in and learn
    ENROLLMENT_STATUS_SUSPENDED = 6;         // Temporarily disabled
    ENROLLMENT_STATUS_WITHDRAWN = 7;         // Left the school
}

enum ConsentType {
    CONSENT_TYPE_UNSPECIFIED = 0;
    CONSENT_TYPE_COPPA = 1;                  // Children's Online Privacy Protection
    CONSENT_TYPE_DATA_COLLECTION = 2;        // General data collection
    CONSENT_TYPE_AI_PERSONALIZATION = 3;     // AI-driven content adaptation
    CONSENT_TYPE_PROGRESS_SHARING = 4;       // Share progress with teacher/school
    CONSENT_TYPE_PHOTO_AVATAR = 5;           // Allow photo upload for avatar
}

// ============================================================================
// PEOPLE MESSAGES
// ============================================================================

// @PrimeObject
message Student {
    string student_id = 1;
    string first_name = 2;
    string last_name = 3;
    string preferred_name = 4;
    GradeLevel grade_level = 5;
    StudentStatus status = 6;
    string classroom_id = 7;                 // Current classroom
    string school_id = 8;
    string district_id = 9;
    int64 date_of_birth = 10;
    string avatar_url = 11;
    string primary_guardian_id = 12;
    string language_preference = 13;         // ISO 639-1 code
    bool has_iep = 14;                       // Individualized Education Program
    bool has_504_plan = 15;                  // Section 504 accommodation
    string accommodation_notes = 16;
    int32 weekly_goal_minutes = 17;          // Target usage per week
    int32 daily_goal_minutes = 18;           // Target usage per day
    int64 enrollment_date = 19;
    map<string, string> custom_fields = 20;
    l8common.AuditInfo audit_info = 21;
}

message StudentList {
    repeated Student list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message Guardian {
    string guardian_id = 1;
    string first_name = 2;
    string last_name = 3;
    string email = 4;
    string phone = 5;
    GuardianRelation relation = 6;
    repeated string student_ids = 7;         // Children linked to this guardian
    bool receives_reports = 8;               // Opt-in for progress reports
    string report_frequency = 9;             // "daily", "weekly", "monthly"
    string language_preference = 10;
    map<string, string> custom_fields = 11;
    l8common.AuditInfo audit_info = 12;
}

message GuardianList {
    repeated Guardian list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message Teacher {
    string teacher_id = 1;
    string first_name = 2;
    string last_name = 3;
    string email = 4;
    TeacherRole role = 5;
    repeated string classroom_ids = 6;
    string school_id = 7;
    string district_id = 8;
    repeated SubjectType subjects = 9;       // Subjects this teacher manages
    map<string, string> custom_fields = 10;
    l8common.AuditInfo audit_info = 11;
}

message TeacherList {
    repeated Teacher list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message Classroom {
    string classroom_id = 1;
    string name = 2;
    GradeLevel grade_level = 3;
    string primary_teacher_id = 4;
    string school_id = 5;
    string academic_year = 6;                // "2026-2027"
    int32 student_count = 7;
    map<string, string> custom_fields = 8;
    l8common.AuditInfo audit_info = 9;
}

message ClassroomList {
    repeated Classroom list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message School {
    string school_id = 1;
    string name = 2;
    string district_id = 3;
    string address_line1 = 4;
    string address_line2 = 5;
    string city = 6;
    string state_province = 7;
    string postal_code = 8;
    string country_code = 9;
    string principal_name = 10;
    string phone = 11;
    string timezone = 12;
    map<string, string> custom_fields = 13;
    l8common.AuditInfo audit_info = 14;
}

message SchoolList {
    repeated School list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
message District {
    string district_id = 1;
    string name = 2;
    string state_province = 3;
    string country_code = 4;
    string admin_contact = 5;
    string phone = 6;
    string license_tier = 7;                 // "basic", "standard", "premium"
    int64 license_expiry = 8;
    int32 max_students = 9;                  // License seat count
    map<string, string> custom_fields = 10;
    l8common.AuditInfo audit_info = 11;
}

message DistrictList {
    repeated District list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// ENROLLMENT
// ============================================================================

// @PrimeObject
// Tracks the full enrollment lifecycle from admin creation through activation
message Enrollment {
    string enrollment_id = 1;
    string student_id = 2;
    string school_id = 3;
    string classroom_id = 4;
    string district_id = 5;
    EnrollmentStatus status = 6;
    string academic_year = 7;                // "2026-2027"

    // Guardian link
    string guardian_id = 8;
    string invite_token = 9;                 // Unique token for guardian invite link
    int64 invite_sent_date = 10;
    int64 invite_expiry = 11;

    // Consent tracking
    repeated ConsentRecord consents = 12;    // Embedded child

    // Onboarding progress
    bool guardian_registered = 13;
    bool consent_complete = 14;
    bool diagnostic_complete = 15;
    bool path_created = 16;
    string diagnostic_benchmark_id = 17;

    // Dates
    int64 enrollment_date = 18;
    int64 activation_date = 19;
    int64 withdrawal_date = 20;
    string withdrawal_reason = 21;

    // Transfer
    string transfer_from_school_id = 22;     // If transferring in
    string transfer_from_district_id = 23;

    map<string, string> custom_fields = 24;
    l8common.AuditInfo audit_info = 25;
}

message EnrollmentList {
    repeated Enrollment list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of Enrollment
message ConsentRecord {
    ConsentType type = 1;
    bool granted = 2;
    int64 signed_date = 3;
    string signed_by = 4;                    // Guardian ID
    string ip_address = 5;                   // For legal record
    string consent_version = 6;              // Version of consent text shown
}
```

### 2.3 learn-adaptive.proto — Adaptive Engine (Service Area 30)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum MasteryLevel {
    MASTERY_LEVEL_UNSPECIFIED = 0;
    MASTERY_LEVEL_NOT_STARTED = 1;
    MASTERY_LEVEL_EMERGING = 2;              // <40% — beginning to understand
    MASTERY_LEVEL_DEVELOPING = 3;            // 40-60% — partial understanding
    MASTERY_LEVEL_PROFICIENT = 4;            // 60-80% — solid understanding
    MASTERY_LEVEL_MASTERED = 5;              // 80-95% — strong command
    MASTERY_LEVEL_EXEMPLARY = 6;             // >95% — can teach others
}

enum PathStatus {
    PATH_STATUS_UNSPECIFIED = 0;
    PATH_STATUS_ACTIVE = 1;
    PATH_STATUS_PAUSED = 2;
    PATH_STATUS_COMPLETED = 3;
}

enum AdaptRuleStatus {
    ADAPT_RULE_STATUS_UNSPECIFIED = 0;
    ADAPT_RULE_STATUS_DRAFT = 1;
    ADAPT_RULE_STATUS_ACTIVE = 2;
    ADAPT_RULE_STATUS_DISABLED = 3;
}

enum AdaptStrategy {
    ADAPT_STRATEGY_UNSPECIFIED = 0;
    ADAPT_STRATEGY_REPEAT = 1;              // Repeat same skill, easier difficulty
    ADAPT_STRATEGY_SCAFFOLD = 2;            // Break skill into sub-skills
    ADAPT_STRATEGY_ALTERNATE = 3;           // Try different activity type
    ADAPT_STRATEGY_REVIEW = 4;              // Go back to prerequisite skill
    ADAPT_STRATEGY_ADVANCE = 5;             // Move to next skill
    ADAPT_STRATEGY_ENRICH = 6;              // Challenge with harder problems
    ADAPT_STRATEGY_BREAK = 7;               // Switch to engagement activity (game)
}

enum AdaptTrigger {
    ADAPT_TRIGGER_UNSPECIFIED = 0;
    ADAPT_TRIGGER_SCORE_BELOW = 1;          // Score dropped below threshold
    ADAPT_TRIGGER_SCORE_ABOVE = 2;          // Score rose above threshold
    ADAPT_TRIGGER_STREAK_CORRECT = 3;       // N correct in a row
    ADAPT_TRIGGER_STREAK_INCORRECT = 4;     // N incorrect in a row
    ADAPT_TRIGGER_TIME_EXCEEDED = 5;        // Spent too long on activity
    ADAPT_TRIGGER_TIME_TOO_FAST = 6;        // Completed suspiciously fast (guessing?)
    ADAPT_TRIGGER_HINTS_EXHAUSTED = 7;      // Used all available hints
    ADAPT_TRIGGER_ENGAGEMENT_DROP = 8;      // Engagement signals declining
    ADAPT_TRIGGER_SESSION_DURATION = 9;     // Session length threshold
    ADAPT_TRIGGER_MASTERY_ACHIEVED = 10;    // Skill mastery level reached target
}

// ============================================================================
// SKILL GRAPH
// ============================================================================

// @PrimeObject
// Skill represents a node in the knowledge graph
message Skill {
    string skill_id = 1;
    string name = 2;
    string description = 3;
    SubjectType subject = 4;
    GradeLevel grade_level = 5;
    string domain = 6;                       // "Number Sense", "Geometry", "Phonics"
    string subdomain = 7;                    // "Fractions", "Angles", "Blends"
    repeated string prerequisite_skill_ids = 8;
    repeated string standard_ids = 9;        // Common Core / state standard codes
    int32 typical_mastery_minutes = 10;      // Expected time to mastery
    map<string, string> custom_fields = 11;
    l8common.AuditInfo audit_info = 12;
}

message SkillList {
    repeated Skill list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// STUDENT MASTERY
// ============================================================================

// @PrimeObject
// One record per student per skill
message SkillMastery {
    string mastery_id = 1;
    string student_id = 2;
    string skill_id = 3;
    MasteryLevel level = 4;
    double confidence = 5;                   // 0.0-1.0 statistical confidence
    int32 attempts_count = 6;
    int32 correct_count = 7;
    double current_accuracy = 8;             // Rolling accuracy (last 10 attempts)
    int64 first_attempted = 9;
    int64 last_attempted = 10;
    int64 mastered_date = 11;                // 0 if not yet mastered
    int32 total_time_seconds = 12;
    repeated MasterySnapshot history = 13;   // Embedded child: periodic snapshots
    map<string, string> custom_fields = 14;
    l8common.AuditInfo audit_info = 15;
}

message SkillMasteryList {
    repeated SkillMastery list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of SkillMastery
message MasterySnapshot {
    int64 timestamp = 1;
    MasteryLevel level = 2;
    double accuracy = 3;
    int32 attempts_in_period = 4;
}

// ============================================================================
// LEARNING PATHS
// ============================================================================

// @PrimeObject
// The AI-managed learning journey for one student in one subject
message LearningPath {
    string path_id = 1;
    string student_id = 2;
    SubjectType subject = 3;
    GradeLevel target_grade = 4;
    PathStatus status = 5;
    string current_activity_id = 6;          // What the student should do next
    string current_skill_id = 7;
    DifficultyLevel current_difficulty = 8;
    int32 activities_completed = 9;
    int32 skills_mastered = 10;
    int32 total_time_seconds = 11;
    int64 started_date = 12;
    int64 last_active = 13;
    repeated PathStep upcoming_queue = 14;   // Embedded: next N activities planned
    repeated AdaptationLog adaptation_log = 15; // Embedded: why AI chose each step
    map<string, string> custom_fields = 16;
    l8common.AuditInfo audit_info = 17;
}

message LearningPathList {
    repeated LearningPath list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of LearningPath
message PathStep {
    string activity_id = 1;
    string skill_id = 2;
    DifficultyLevel difficulty = 3;
    string reason = 4;                       // Why this step was chosen
    int32 sequence = 5;
}

// Embedded child of LearningPath
message AdaptationLog {
    int64 timestamp = 1;
    AdaptTrigger trigger = 2;
    AdaptStrategy strategy = 3;
    string from_activity_id = 4;
    string to_activity_id = 5;
    string reasoning = 6;                    // AI explanation of the decision
}

// ============================================================================
// ADAPTATION RULES
// ============================================================================

// @PrimeObject
// Configurable rules that guide the adaptive engine
message AdaptationRule {
    string rule_id = 1;
    string name = 2;
    string description = 3;
    AdaptRuleStatus status = 4;
    int32 priority = 5;                      // Lower = higher priority

    // Scope
    SubjectType subject_filter = 6;          // 0 = all subjects
    GradeLevel grade_filter = 7;             // 0 = all grades

    // Trigger conditions
    AdaptTrigger trigger = 8;
    int32 trigger_threshold = 9;             // e.g., score < 40, streak >= 5
    int32 trigger_window = 10;               // Number of recent attempts to evaluate

    // Response
    AdaptStrategy strategy = 11;
    DifficultyLevel target_difficulty = 12;  // 0 = relative adjustment

    // Limits
    int32 max_applications_per_session = 13;
    int32 cooldown_seconds = 14;

    map<string, string> custom_fields = 15;
    l8common.AuditInfo audit_info = 16;
}

message AdaptationRuleList {
    repeated AdaptationRule list = 1;
    l8api.L8MetaData metadata = 2;
}
```

### 2.4 learn-assessment.proto — Evaluation (Service Area 40)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum SessionStatus {
    SESSION_STATUS_UNSPECIFIED = 0;
    SESSION_STATUS_ACTIVE = 1;
    SESSION_STATUS_COMPLETED = 2;
    SESSION_STATUS_ABANDONED = 3;            // Closed without finishing
}

enum InteractionResult {
    INTERACTION_RESULT_UNSPECIFIED = 0;
    INTERACTION_RESULT_CORRECT = 1;
    INTERACTION_RESULT_INCORRECT = 2;
    INTERACTION_RESULT_PARTIAL = 3;
    INTERACTION_RESULT_SKIPPED = 4;
    INTERACTION_RESULT_HINT_USED = 5;
    INTERACTION_RESULT_TIMED_OUT = 6;
}

enum BenchmarkType {
    BENCHMARK_TYPE_UNSPECIFIED = 0;
    BENCHMARK_TYPE_DIAGNOSTIC = 1;           // Initial placement
    BENCHMARK_TYPE_FORMATIVE = 2;            // During instruction
    BENCHMARK_TYPE_SUMMATIVE = 3;            // End of unit/course
    BENCHMARK_TYPE_PROGRESS_MONITOR = 4;     // Periodic check
}

enum ScanStatus {
    SCAN_STATUS_UNSPECIFIED = 0;
    SCAN_STATUS_UPLOADED = 1;                // Image received
    SCAN_STATUS_PROCESSING = 2;              // AI extracting answers
    SCAN_STATUS_REVIEW = 3;                  // Has flagged items needing teacher review
    SCAN_STATUS_COMPLETE = 4;                // All graded, scores saved
    SCAN_STATUS_FAILED = 5;                  // AI could not process image
}

// ============================================================================
// SESSION & INTERACTION TRACKING
// ============================================================================

// @PrimeObject
// One record per student login session
message LearningSession {
    string session_id = 1;
    string student_id = 2;
    string path_id = 3;                      // Active learning path
    SessionStatus status = 4;
    int64 start_time = 5;
    int64 end_time = 6;
    int32 duration_seconds = 7;
    int32 activities_attempted = 8;
    int32 activities_completed = 9;
    int32 questions_answered = 10;
    int32 questions_correct = 11;
    int32 hints_used = 12;
    int32 points_earned = 13;
    string device_type = 14;                 // "tablet", "desktop", "phone"
    repeated Interaction interactions = 15;   // Embedded child
    map<string, string> custom_fields = 16;
    l8common.AuditInfo audit_info = 17;
}

message LearningSessionList {
    repeated LearningSession list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of LearningSession
message Interaction {
    string interaction_id = 1;
    string activity_id = 2;
    string question_id = 3;
    string skill_id = 4;
    int64 timestamp = 5;
    InteractionResult result = 6;
    string student_answer = 7;
    string correct_answer = 8;
    int32 time_spent_seconds = 9;
    int32 attempt_number = 10;
    int32 hints_used = 11;
    int32 points_awarded = 12;
    DifficultyLevel difficulty = 13;
    string strategy_detected = 14;           // AI-detected problem-solving strategy
}

// ============================================================================
// SCORES & BENCHMARKS
// ============================================================================

// @PrimeObject
// Aggregated score per student per skill or activity
message Score {
    string score_id = 1;
    string student_id = 2;
    string skill_id = 3;
    string activity_id = 4;                  // Empty if skill-level aggregate
    double score_value = 5;                  // 0.0-100.0
    double percentile = 6;                   // Relative to peers
    int64 scored_date = 7;
    int32 questions_total = 8;
    int32 questions_correct = 9;
    int32 time_spent_seconds = 10;
    map<string, string> custom_fields = 11;
    l8common.AuditInfo audit_info = 12;
}

message ScoreList {
    repeated Score list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
// Periodic assessment (diagnostic, formative, summative)
message Benchmark {
    string benchmark_id = 1;
    string name = 2;
    BenchmarkType type = 3;
    SubjectType subject = 4;
    GradeLevel grade_level = 5;
    ContentStatus status = 6;
    repeated string skill_ids = 7;           // Skills assessed
    repeated string activity_ids = 8;        // Activities included
    int32 time_limit_minutes = 9;            // 0 = untimed
    int32 passing_score = 10;                // Minimum score to pass (0-100)
    map<string, string> custom_fields = 11;
    l8common.AuditInfo audit_info = 12;
}

message BenchmarkList {
    repeated Benchmark list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// WORKSHEET SCANNING — AI-powered grading of printed worksheets
// ============================================================================

// @PrimeObject
// One scan job per worksheet per student
message WorksheetScan {
    string scan_id = 1;
    string worksheet_id = 2;
    string student_id = 3;                   // Resolved from name on page or teacher assignment
    string teacher_id = 4;                   // Who initiated the scan
    ScanStatus status = 5;

    // Input
    string image_path = 6;                   // Scanned image via Layer8FileUpload
    string image_format = 7;                 // "jpeg", "png", "pdf"

    // AI extraction results
    string detected_student_name = 8;        // Name read from the worksheet header
    double extraction_confidence = 9;        // Overall confidence 0.0-1.0
    repeated ScannedAnswer answers = 10;     // Embedded child

    // Scoring
    int32 auto_graded_count = 11;
    int32 flagged_count = 12;
    int32 correct_count = 13;
    int32 total_questions = 14;
    double score_percent = 15;
    bool fed_to_mastery = 16;                // Scores fed back to SkillMastery

    map<string, string> custom_fields = 17;
    l8common.AuditInfo audit_info = 18;
}

message WorksheetScanList {
    repeated WorksheetScan list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of WorksheetScan
message ScannedAnswer {
    string question_id = 1;
    int32 question_number = 2;
    string extracted_text = 3;               // What the AI read from the scan
    string correct_answer = 4;
    InteractionResult result = 5;            // CORRECT, INCORRECT, or UNSPECIFIED (flagged)
    double confidence = 6;                   // Per-answer extraction confidence
    bool needs_review = 7;                   // Flagged for teacher
    string teacher_override = 8;             // Teacher's manual answer (if reviewed)
    string crop_image_path = 9;              // Cropped image of just this answer area
    string ai_reasoning = 10;                // Why AI graded it this way (e.g., "1 1/4 = 5/4, equivalent")
}
```

### 2.5 learn-analytics.proto — Reporting (Service Area 50)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum ReportPeriod {
    REPORT_PERIOD_UNSPECIFIED = 0;
    REPORT_PERIOD_DAILY = 1;
    REPORT_PERIOD_WEEKLY = 2;
    REPORT_PERIOD_MONTHLY = 3;
    REPORT_PERIOD_QUARTERLY = 4;
    REPORT_PERIOD_YEARLY = 5;
}

enum EngagementLevel {
    ENGAGEMENT_LEVEL_UNSPECIFIED = 0;
    ENGAGEMENT_LEVEL_DISENGAGED = 1;
    ENGAGEMENT_LEVEL_LOW = 2;
    ENGAGEMENT_LEVEL_MODERATE = 3;
    ENGAGEMENT_LEVEL_HIGH = 4;
    ENGAGEMENT_LEVEL_EXCEPTIONAL = 5;
}

// ============================================================================
// PROGRESS REPORTS
// ============================================================================

// @PrimeObject
// Generated report for a student over a time period
message ProgressReport {
    string report_id = 1;
    string student_id = 2;
    SubjectType subject = 3;
    ReportPeriod period = 4;
    int64 period_start = 5;
    int64 period_end = 6;

    // Time & usage
    int32 total_time_minutes = 7;
    int32 sessions_count = 8;
    int32 days_active = 9;

    // Achievement
    int32 activities_completed = 10;
    int32 skills_practiced = 11;
    int32 skills_mastered = 12;
    double average_score = 13;
    double score_trend = 14;                 // Positive = improving

    // Engagement
    EngagementLevel engagement = 15;
    int32 streak_days = 16;                  // Consecutive days active

    // Strengths & gaps
    repeated string strength_skill_ids = 17;
    repeated string gap_skill_ids = 18;
    string ai_summary = 19;                  // AI-generated narrative summary
    string ai_recommendations = 20;          // AI-generated next steps

    map<string, string> custom_fields = 21;
    l8common.AuditInfo audit_info = 22;
}

message ProgressReportList {
    repeated ProgressReport list = 1;
    l8api.L8MetaData metadata = 2;
}

// @PrimeObject
// Engagement metrics per student (continuously updated)
message EngagementMetric {
    string metric_id = 1;
    string student_id = 2;
    EngagementLevel current_level = 3;
    int32 current_streak_days = 4;
    int32 longest_streak_days = 5;
    int32 total_points = 6;
    int32 badges_earned = 7;
    double avg_session_minutes = 8;
    double avg_daily_minutes = 9;
    int32 weekly_goal_percent = 10;          // % of weekly goal met
    int64 last_session_date = 11;
    int32 days_since_last_session = 12;
    map<string, string> custom_fields = 13;
    l8common.AuditInfo audit_info = 14;
}

message EngagementMetricList {
    repeated EngagementMetric list = 1;
    l8api.L8MetaData metadata = 2;
}
```

### 2.6 learn-history.proto — Historical Analytics (Service Area 60)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum AggregationLevel {
    AGGREGATION_LEVEL_UNSPECIFIED = 0;
    AGGREGATION_LEVEL_STUDENT = 1;
    AGGREGATION_LEVEL_CLASSROOM = 2;
    AGGREGATION_LEVEL_SCHOOL = 3;
    AGGREGATION_LEVEL_DISTRICT = 4;
}

enum GrowthRating {
    GROWTH_RATING_UNSPECIFIED = 0;
    GROWTH_RATING_WELL_BELOW = 1;            // <25th percentile growth
    GROWTH_RATING_BELOW = 2;                 // 25-40th
    GROWTH_RATING_TYPICAL = 3;               // 40-60th
    GROWTH_RATING_ABOVE = 4;                 // 60-75th
    GROWTH_RATING_WELL_ABOVE = 5;            // >75th percentile growth
}

enum RiskLevel {
    RISK_LEVEL_UNSPECIFIED = 0;
    RISK_LEVEL_ON_TRACK = 1;
    RISK_LEVEL_WATCH = 2;                    // Early warning signals
    RISK_LEVEL_AT_RISK = 3;                  // Intervention needed
    RISK_LEVEL_CRITICAL = 4;                 // Immediate action required
}

enum SnapshotType {
    SNAPSHOT_TYPE_UNSPECIFIED = 0;
    SNAPSHOT_TYPE_WEEKLY = 1;
    SNAPSHOT_TYPE_MONTHLY = 2;
    SNAPSHOT_TYPE_QUARTERLY = 3;
    SNAPSHOT_TYPE_SEMESTER = 4;
    SNAPSHOT_TYPE_YEAR_END = 5;
}

// ============================================================================
// GROWTH MEASUREMENT
// ============================================================================

// @PrimeObject
// Measures growth over time — one record per student per subject per academic year
message GrowthRecord {
    string growth_id = 1;
    string student_id = 2;
    SubjectType subject = 3;
    string academic_year = 4;                // "2026-2027"
    GradeLevel grade_level = 5;

    // Starting point (beginning of year / enrollment diagnostic)
    double baseline_score = 6;
    MasteryLevel baseline_mastery = 7;
    int64 baseline_date = 8;
    int32 baseline_skills_mastered = 9;

    // Current / end point
    double current_score = 10;
    MasteryLevel current_mastery = 11;
    int64 current_date = 12;
    int32 current_skills_mastered = 13;

    // Growth metrics
    double absolute_growth = 14;             // current - baseline
    double growth_percentile = 15;           // Compared to peers at same starting level
    GrowthRating rating = 16;
    double expected_growth = 17;             // Typical growth for this starting level
    double growth_vs_expected = 18;          // actual / expected (>1.0 = above expected)

    // Time investment
    int32 total_time_minutes = 19;
    int32 total_sessions = 20;
    int32 total_activities = 21;

    map<string, string> custom_fields = 22;
    l8common.AuditInfo audit_info = 23;
}

message GrowthRecordList {
    repeated GrowthRecord list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// COHORT ANALYTICS
// ============================================================================

// @PrimeObject
// Aggregated snapshot for a group at a point in time (weekly, monthly, year-end)
message CohortSnapshot {
    string snapshot_id = 1;
    AggregationLevel level = 2;
    string entity_id = 3;                    // classroom_id, school_id, or district_id
    SubjectType subject = 4;
    GradeLevel grade_level = 5;              // 0 = all grades at this entity
    string academic_year = 6;
    SnapshotType type = 7;
    int64 snapshot_date = 8;

    // Population
    int32 total_students = 9;
    int32 active_students = 10;
    int32 inactive_students = 11;

    // Mastery distribution
    int32 students_exemplary = 12;
    int32 students_mastered = 13;
    int32 students_proficient = 14;
    int32 students_developing = 15;
    int32 students_emerging = 16;
    int32 students_not_started = 17;

    // Aggregate scores
    double mean_score = 18;
    double median_score = 19;
    double std_deviation = 20;
    double score_25th_percentile = 21;
    double score_75th_percentile = 22;

    // Growth
    double mean_growth = 23;
    double mean_growth_vs_expected = 24;     // >1.0 = above expected
    int32 students_above_expected = 25;
    int32 students_below_expected = 26;

    // Engagement
    double mean_weekly_minutes = 27;
    double participation_rate = 28;          // % meeting weekly goal
    int32 students_disengaged = 29;          // No activity in 14+ days

    // Skills
    int32 total_skills_in_scope = 30;
    double mean_skills_mastered = 31;
    repeated SkillGap top_gaps = 32;         // Embedded: top 10 weakest skills

    map<string, string> custom_fields = 33;
    l8common.AuditInfo audit_info = 34;
}

message CohortSnapshotList {
    repeated CohortSnapshot list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of CohortSnapshot
message SkillGap {
    string skill_id = 1;
    string skill_name = 2;
    double mastery_rate = 3;                 // % of students at PROFICIENT or above
    int32 students_struggling = 4;
}

// ============================================================================
// RISK PREDICTION
// ============================================================================

// @PrimeObject
// AI-generated early warning for students at risk of falling behind (recalculated weekly)
message RiskAssessment {
    string assessment_id = 1;
    string student_id = 2;
    SubjectType subject = 3;
    RiskLevel risk_level = 4;
    double risk_score = 5;                   // 0.0-1.0 probability of falling behind
    int64 assessed_date = 6;

    // Contributing factors (AI-identified)
    repeated RiskFactor factors = 7;         // Embedded child

    // Recommended interventions
    string ai_recommendation = 8;
    repeated string recommended_skill_ids = 9;

    // Tracking
    bool teacher_acknowledged = 10;
    int64 acknowledged_date = 11;
    string teacher_notes = 12;
    string intervention_plan = 13;

    map<string, string> custom_fields = 14;
    l8common.AuditInfo audit_info = 15;
}

message RiskAssessmentList {
    repeated RiskAssessment list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of RiskAssessment
message RiskFactor {
    string factor_type = 1;                  // "engagement_drop", "mastery_plateau", "session_decline"
    string description = 2;
    double weight = 3;                       // Contribution to risk score
    string evidence = 4;                     // "multiplication stuck at 52% for 3 weeks"
}

// ============================================================================
// STANDARDS MASTERY
// ============================================================================

// @PrimeObject
// Tracks mastery against curriculum standards (Common Core, state standards)
// One record per student per standard per academic year
message StandardMastery {
    string standard_mastery_id = 1;
    string student_id = 2;
    string standard_id = 3;                  // e.g., "4.NF.A.1"
    string standard_description = 4;
    SubjectType subject = 5;
    GradeLevel grade_level = 6;
    string academic_year = 7;

    MasteryLevel level = 8;
    double score = 9;
    int32 skills_in_standard = 10;
    int32 skills_mastered = 11;
    int64 last_assessed = 12;

    map<string, string> custom_fields = 13;
    l8common.AuditInfo audit_info = 14;
}

message StandardMasteryList {
    repeated StandardMastery list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// CURRICULUM EFFECTIVENESS
// ============================================================================

// @PrimeObject
// How effective is each activity/lesson/course at producing mastery?
// Aggregated from all students who attempted it
message ContentEffect {
    string effect_id = 1;
    string content_id = 2;                   // activity_id, lesson_id, or course_id
    string content_type = 3;                 // "Activity", "Lesson", "Course"
    string content_name = 4;
    SubjectType subject = 5;
    string academic_year = 6;

    // Usage
    int32 total_attempts = 7;
    int32 unique_students = 8;
    double mean_completion_rate = 9;
    double mean_time_seconds = 10;

    // Effectiveness
    double mean_score = 11;
    double mean_mastery_gain = 12;           // Avg mastery improvement after this content
    double mastery_gain_per_minute = 13;     // Efficiency: learning per unit time
    int32 students_gained_mastery = 14;      // Moved from <PROFICIENT to >=PROFICIENT

    // Comparative
    double effectiveness_percentile = 15;    // vs other content targeting same skills
    string ai_analysis = 16;                 // AI narrative on why this content works or doesn't

    map<string, string> custom_fields = 17;
    l8common.AuditInfo audit_info = 18;
}

message ContentEffectList {
    repeated ContentEffect list = 1;
    l8api.L8MetaData metadata = 2;
}
```

### 2.7 learn-homeschool.proto — Homeschool Features (Service Areas 10, 20, 30)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum ComplianceType {
    COMPLIANCE_TYPE_UNSPECIFIED = 0;
    COMPLIANCE_TYPE_NOTIFICATION = 1;        // Must notify district of intent to homeschool
    COMPLIANCE_TYPE_ATTENDANCE_LOG = 2;       // Must track days/hours of instruction
    COMPLIANCE_TYPE_PORTFOLIO_REVIEW = 3;     // Annual portfolio reviewed by evaluator
    COMPLIANCE_TYPE_STANDARDIZED_TEST = 4;    // Must take state standardized test
    COMPLIANCE_TYPE_SUBJECT_COVERAGE = 5;     // Must cover specific subjects
    COMPLIANCE_TYPE_PROGRESS_REPORT = 6;      // Must submit progress reports to district
    COMPLIANCE_TYPE_INSTRUCTOR_QUAL = 7;      // Parent must meet qualification requirements
}

enum RealWorldContext {
    REAL_WORLD_CONTEXT_UNSPECIFIED = 0;
    REAL_WORLD_CONTEXT_GROCERY = 1;
    REAL_WORLD_CONTEXT_COOKING = 2;
    REAL_WORLD_CONTEXT_NATURE_WALK = 3;
    REAL_WORLD_CONTEXT_ROAD_TRIP = 4;
    REAL_WORLD_CONTEXT_MUSEUM = 5;
    REAL_WORLD_CONTEXT_LIBRARY = 6;
    REAL_WORLD_CONTEXT_GARDEN = 7;
    REAL_WORLD_CONTEXT_BUILDING = 8;
    REAL_WORLD_CONTEXT_SHOPPING = 9;
    REAL_WORLD_CONTEXT_SPORTS = 10;
}

// ============================================================================
// FAMILY (Service Area 20)
// ============================================================================

// @PrimeObject
// Family unit linking guardians and students for multi-child management
message Family {
    string family_id = 1;
    string name = 2;                         // "The Martinez Family"
    repeated string guardian_ids = 3;
    repeated string student_ids = 4;
    string primary_guardian_id = 5;
    string state_code = 6;                   // For compliance rules
    string timezone = 7;
    string language_preference = 8;
    map<string, string> custom_fields = 9;
    l8common.AuditInfo audit_info = 10;
}

message FamilyList {
    repeated Family list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// STATE COMPLIANCE (Service Area 20)
// ============================================================================

// @PrimeObject
// Tracks homeschool legal compliance per state per family per year
message StateCompliance {
    string compliance_id = 1;
    string family_id = 2;
    string state_code = 3;                   // "CA", "TX", "NY"
    string academic_year = 4;

    // Requirements for this state
    repeated ComplianceRequirement requirements = 5;  // Embedded child

    // Auto-tracked from session data
    int32 instruction_days_logged = 6;
    int32 instruction_days_required = 7;
    int32 instruction_hours_logged = 8;
    int32 instruction_hours_required = 9;
    repeated SubjectType subjects_covered = 10;
    repeated SubjectType subjects_required = 11;

    // Status
    bool notification_filed = 12;
    int64 notification_date = 13;
    bool portfolio_submitted = 14;
    int64 portfolio_date = 15;
    int64 next_deadline = 16;
    string next_action = 17;                 // "File annual assessment by June 30"

    // Auto-generated portfolio
    string portfolio_path = 18;              // Generated PDF via Layer8FileUpload

    map<string, string> custom_fields = 19;
    l8common.AuditInfo audit_info = 20;
}

message StateComplianceList {
    repeated StateCompliance list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of StateCompliance
message ComplianceRequirement {
    ComplianceType type = 1;
    string description = 2;
    int64 deadline = 3;
    bool satisfied = 4;
    string evidence_path = 5;                // Stored document via Layer8FileUpload
}

// ============================================================================
// LEARNING PODS (Service Area 20)
// ============================================================================

// @PrimeObject
// Homeschool co-op / learning pod (3-5 families learning together)
message LearningPod {
    string pod_id = 1;
    string name = 2;                         // "Elm Street Learning Pod"
    repeated string family_ids = 3;
    string organizer_guardian_id = 4;
    string meeting_schedule = 5;             // "Tuesdays and Thursdays 10am-12pm"
    string location = 6;
    repeated string shared_course_ids = 7;
    repeated SubjectType pod_subjects = 8;
    int32 total_students = 9;
    map<string, string> custom_fields = 10;
    l8common.AuditInfo audit_info = 11;
}

message LearningPodList {
    repeated LearningPod list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// FAMILY ACTIVITIES (Service Area 10)
// ============================================================================

// @PrimeObject
// Sibling collaboration activities requiring multiple ages working together
message FamilyActivity {
    string family_activity_id = 1;
    string name = 2;
    string description = 3;
    ActivityType activity_type = 4;

    // Multi-child involvement
    repeated StudentRole roles = 5;          // Embedded: what each child does

    // Cross-subject integration
    repeated SubjectType subjects_covered = 6;
    repeated string skill_ids = 7;

    // Logistics
    int32 estimated_minutes = 8;
    repeated string materials_needed = 9;
    string location = 10;                    // "kitchen", "backyard", "living room"
    bool requires_parent = 11;

    GradeLevel min_grade = 12;
    GradeLevel max_grade = 13;
    ContentStatus status = 14;
    map<string, string> custom_fields = 15;
    l8common.AuditInfo audit_info = 16;
}

message FamilyActivityList {
    repeated FamilyActivity list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of FamilyActivity
message StudentRole {
    string student_id = 1;
    string role_description = 2;             // "Teacher: explain how to halve a recipe"
    GradeLevel role_level = 3;               // Difficulty matched to this child
    repeated string skill_ids = 4;           // What THIS child practices
}

// ============================================================================
// REAL-WORLD LESSONS (Service Area 10)
// ============================================================================

// @PrimeObject
// Turns daily life (grocery trips, cooking, nature walks) into curriculum
message RealWorldLesson {
    string lesson_id = 1;
    string name = 2;
    RealWorldContext context = 3;
    string description = 4;

    // Pre-activity (instructions for parent)
    string parent_instructions = 5;
    repeated string materials_needed = 6;

    // During activity (prompts/challenges)
    repeated RealWorldChallenge challenges = 7;  // Embedded child

    // Post-activity (logging back into the system)
    repeated string skill_ids = 8;
    repeated SubjectType subjects = 9;
    GradeLevel min_grade = 10;
    GradeLevel max_grade = 11;
    int32 estimated_minutes = 12;
    ContentStatus status = 13;
    map<string, string> custom_fields = 14;
    l8common.AuditInfo audit_info = 15;
}

message RealWorldLessonList {
    repeated RealWorldLesson list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of RealWorldLesson
message RealWorldChallenge {
    string challenge_id = 1;
    string prompt = 2;                       // "Estimate the total before checkout"
    string skill_id = 3;
    DifficultyLevel difficulty = 4;
    string success_criteria = 5;             // "Within $2 of actual total"
}

// ============================================================================
// PROJECT-BASED LEARNING (Service Area 10)
// ============================================================================

// @PrimeObject
// Multi-week projects integrating multiple subjects
message Project {
    string project_id = 1;
    string name = 2;                         // "Design a Garden"
    string description = 3;
    repeated SubjectType subjects = 4;
    GradeLevel min_grade = 5;
    GradeLevel max_grade = 6;
    int32 estimated_weeks = 7;

    repeated ProjectMilestone milestones = 8;  // Embedded child

    // Cross-subject skill mapping
    repeated string skill_ids = 9;
    repeated string materials_needed = 10;
    string budget_estimate = 11;             // "$15-25"
    ContentStatus status = 12;
    map<string, string> custom_fields = 13;
    l8common.AuditInfo audit_info = 14;
}

message ProjectList {
    repeated Project list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of Project
message ProjectMilestone {
    string milestone_id = 1;
    int32 week_number = 2;
    string title = 3;                        // "Week 2: Calculate area for each plant"
    string description = 4;
    SubjectType primary_subject = 5;
    repeated string skill_ids = 6;
    repeated string deliverables = 7;        // "Garden layout drawing", "Area calculations"
    bool requires_photo = 8;                 // Parent uploads photo of real-world work
}

// ============================================================================
// DAILY SCHEDULE (Service Area 30)
// ============================================================================

// @PrimeObject
// AI-generated daily plan for the whole family
message DailySchedule {
    string schedule_id = 1;
    string family_id = 2;
    int64 schedule_date = 3;

    repeated ScheduleBlock blocks = 4;       // Embedded: time-ordered blocks

    // Context the AI used to generate this
    int32 available_hours = 5;               // Parent said "we have 4 hours today"
    string parent_energy = 6;                // "low", "medium", "high"
    repeated string appointments = 7;        // "dentist at 2pm" — schedule around it
    string weather = 8;                      // "sunny" → outdoor; "rainy" → indoor

    map<string, string> custom_fields = 9;
    l8common.AuditInfo audit_info = 10;
}

message DailyScheduleList {
    repeated DailySchedule list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of DailySchedule
message ScheduleBlock {
    string block_id = 1;
    int32 start_minute = 2;                  // Minutes from midnight (540 = 9:00 AM)
    int32 duration_minutes = 3;
    string student_id = 4;                   // Empty = family activity
    string activity_type = 5;                // "digital", "hands_on", "reading", "outdoor", "break", "sibling"
    SubjectType subject = 6;
    string description = 7;
    string parent_role = 8;                  // "supervise", "teach", "none" (independent)
    bool requires_parent = 9;
}
```

### 2.8 learn-collab.proto — Collaboration Center (Service Area 70)

```protobuf
syntax = "proto3";
package learn;
option go_package = "./types/learn";
import "l8common.proto";
import "api.proto";

// ============================================================================
// ENUMS
// ============================================================================

enum GroupType {
    GROUP_TYPE_UNSPECIFIED = 0;
    GROUP_TYPE_STUDY = 1;                    // Working on same subject/skills
    GROUP_TYPE_PROJECT = 2;                  // Shared deliverable
    GROUP_TYPE_CHALLENGE = 3;                // Competitive team goal
    GROUP_TYPE_TUTORING = 4;                 // Peer tutor + learner pair
    GROUP_TYPE_BOOK_CLUB = 5;               // Shared reading + discussion
    GROUP_TYPE_POD = 6;                      // Homeschool pod group
}

enum GroupStatus {
    GROUP_STATUS_UNSPECIFIED = 0;
    GROUP_STATUS_ACTIVE = 1;
    GROUP_STATUS_COMPLETED = 2;
    GROUP_STATUS_PAUSED = 3;
    GROUP_STATUS_ARCHIVED = 4;
}

enum MemberRole {
    MEMBER_ROLE_UNSPECIFIED = 0;
    MEMBER_ROLE_LEADER = 1;
    MEMBER_ROLE_MEMBER = 2;
    MEMBER_ROLE_TUTOR = 3;
    MEMBER_ROLE_LEARNER = 4;
}

enum ChatMessageType {
    CHAT_MESSAGE_TYPE_UNSPECIFIED = 0;
    CHAT_MESSAGE_TYPE_TEXT = 1;
    CHAT_MESSAGE_TYPE_QUESTION = 2;          // "How do I do #7?"
    CHAT_MESSAGE_TYPE_EXPLANATION = 3;       // "Here's how I solved it..."
    CHAT_MESSAGE_TYPE_ENCOURAGEMENT = 4;     // "You got this!"
    CHAT_MESSAGE_TYPE_WORK_SHARE = 5;        // Shared approach (not answer)
    CHAT_MESSAGE_TYPE_AI_COACH = 6;          // AI interjection
    CHAT_MESSAGE_TYPE_SYSTEM = 7;            // Milestone, badge, notification
}

enum ModerationAction {
    MODERATION_ACTION_UNSPECIFIED = 0;
    MODERATION_ACTION_APPROVED = 1;          // Auto-approved, appropriate
    MODERATION_ACTION_COACHED = 2;           // AI nudged toward better phrasing
    MODERATION_ACTION_BLOCKED = 3;           // Inappropriate, not delivered
    MODERATION_ACTION_FLAGGED = 4;           // Sent but flagged for parent/teacher review
}

// ============================================================================
// COLLABORATION GROUPS
// ============================================================================

// @PrimeObject
// A collaboration group (study squad, project team, tutor pair, challenge team)
message CollabGroup {
    string group_id = 1;
    string name = 2;
    GroupType type = 3;
    GroupStatus status = 4;
    string description = 5;

    // Membership
    repeated GroupMember members = 6;        // Embedded child
    int32 max_members = 7;
    string classroom_id = 8;                 // Scoped to classroom (or pod_id for homeschool)
    string created_by = 9;                   // teacher_id or student_id

    // Learning focus
    SubjectType subject = 10;
    repeated string skill_ids = 11;
    repeated string activity_ids = 12;       // Shared activities/assignments

    // Goals
    int32 team_score_goal = 13;
    int64 deadline = 14;
    int32 current_team_score = 15;
    int32 team_streak_days = 16;

    // Project-specific (type=PROJECT)
    repeated ProjectDeliverable deliverables = 17;  // Embedded child

    map<string, string> custom_fields = 18;
    l8common.AuditInfo audit_info = 19;
}

message CollabGroupList {
    repeated CollabGroup list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of CollabGroup
message GroupMember {
    string student_id = 1;
    MemberRole role = 2;
    int32 contribution_score = 3;
    int32 messages_sent = 4;
    int32 helps_given = 5;                   // Times they explained something to a peer
    int64 last_active = 6;
    string assigned_task = 7;                // For projects: "Data Collector"
}

// Embedded child of CollabGroup (for project groups)
message ProjectDeliverable {
    string deliverable_id = 1;
    string title = 2;
    string description = 3;
    string assigned_to = 4;                  // student_id
    int32 week_number = 5;
    bool completed = 6;
    string file_path = 7;                    // Uploaded work via Layer8FileUpload
}

// ============================================================================
// AI-MODERATED STUDY CHAT
// ============================================================================

// @PrimeObject
// Chat message in a collaboration group (AI-moderated before delivery)
message CollabMessage {
    string message_id = 1;
    string group_id = 2;
    string sender_id = 3;                    // student_id or "ai" for AI coach
    ChatMessageType type = 4;
    string content = 5;
    int64 timestamp = 6;

    // Moderation
    ModerationAction moderation = 7;
    string moderation_reason = 8;

    // Learning value tracking
    bool contains_explanation = 9;
    bool helped_peer = 10;                   // Recipient marked as helpful
    string referenced_question_id = 11;
    string referenced_skill_id = 12;

    map<string, string> custom_fields = 13;
    l8common.AuditInfo audit_info = 14;
}

message CollabMessageList {
    repeated CollabMessage list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// PEER TUTORING
// ============================================================================

// @PrimeObject
// AI-matched peer tutoring pair (student who mastered teaches student who's struggling)
message TutorMatch {
    string match_id = 1;
    string tutor_id = 2;                     // Student who has mastery
    string learner_id = 3;                   // Student who needs help
    string skill_id = 4;
    string group_id = 5;                     // Created CollabGroup for this pair

    // Tutor qualifications
    MasteryLevel tutor_mastery = 6;          // Must be MASTERED or EXEMPLARY
    double tutor_confidence = 7;

    // Outcome tracking
    MasteryLevel learner_start = 8;
    MasteryLevel learner_end = 9;
    double learner_improvement = 10;
    int32 sessions_completed = 11;
    bool successful = 12;                    // Did learner reach PROFICIENT?

    // Rewards
    int32 tutor_points_earned = 13;
    int32 learner_points_earned = 14;
    repeated string badges_awarded = 15;

    int64 start_date = 16;
    int64 end_date = 17;
    map<string, string> custom_fields = 18;
    l8common.AuditInfo audit_info = 19;
}

message TutorMatchList {
    repeated TutorMatch list = 1;
    l8api.L8MetaData metadata = 2;
}

// ============================================================================
// TEAM CHALLENGES
// ============================================================================

// @PrimeObject
// Competitive/collaborative team goal with leaderboard
message Challenge {
    string challenge_id = 1;
    string name = 2;
    string description = 3;
    string classroom_id = 4;
    SubjectType subject = 5;
    repeated string skill_ids = 6;

    // Rules
    int32 team_size = 7;
    int64 start_date = 8;
    int64 end_date = 9;
    int32 daily_requirement = 10;            // Activities per member per day
    bool weakest_member_bonus = 11;          // 2x if weakest member improves
    int32 target_team_score = 12;

    // Teams
    repeated ChallengeTeam teams = 13;       // Embedded child

    ContentStatus status = 14;
    map<string, string> custom_fields = 15;
    l8common.AuditInfo audit_info = 16;
}

message ChallengeList {
    repeated Challenge list = 1;
    l8api.L8MetaData metadata = 2;
}

// Embedded child of Challenge
message ChallengeTeam {
    string team_id = 1;
    string team_name = 2;
    string team_emoji = 3;                   // "🦊", "🐼"
    repeated string member_ids = 4;
    int32 total_score = 5;
    int32 streak_days = 6;
    int32 rank = 7;
}
```

---

## 3. Service Design

### 3.0 Common Package Interface

The `go/learn/common/` package re-exports Layer 8 framework utilities (from `l8common`) and provides the project-specific PREFIX constant. This follows the `../l8erp/go/erp/common/` pattern exactly.

```go
// go/learn/common/defaults.go
package common

import (
    l8c "github.com/saichler/l8common/go/common"
    "github.com/saichler/l8types/go/ifs"
)

const PREFIX = "/learn"

func CreateResources(alias string, logVnet bool) ifs.IResources {
    return l8c.CreateResources(alias, logVnet)
}

var WaitForSignal = l8c.WaitForSignal
var OpenDBConection = l8c.OpenDBConection
```

**Available helper functions** (inherited from `l8common/go/common/`):

| Function | Purpose |
|----------|---------|
| `common.ActivateService(config, single, list, creds, dbname, vnic)` | Register a service with ORM + web endpoints |
| `common.NewValidation(entity, vnic)` | Create a validation builder for ServiceCallback |
| `common.GenerateID(idField *string)` | Auto-generate UUID for primary key on POST |
| `common.RegisterType(resources, single, list, primaryKey)` | Register type for UI introspection |
| `common.ServiceHandler(serviceName, serviceArea, vnic)` | Get a service handler by name |
| `common.GetEntity(serviceName, serviceArea, filter, vnic)` | Fetch single entity by filter |
| `common.GetEntities(serviceName, serviceArea, filter, vnic)` | Fetch multiple entities by filter |
| `common.CreateWebServer(name, registerFn)` | Create the UI web server |

**ServiceConfig struct** (passed to `ActivateService`):
```go
common.ServiceConfig{
    ServiceName: "Course",           // Max 10 chars
    ServiceArea: byte(10),           // Consistent within module
    PrimaryKey:  "CourseId",         // Matches protobuf JSON field name
    Callback:    newCourseServiceCallback(vnic),
}
```

### 3.0.1 Duplication Prevention

Per `plan-duplication-audit.md`: ALL behavioral logic lives in the Layer 8 framework (`l8common`, `l8secure`, `l8bus`). Individual service files contain **configuration only** — no behavioral code is duplicated across the 38 services.

**What the framework provides (behavioral — shared, never duplicated):**
- `common.ActivateService()` — service registration, ORM setup, web endpoint creation, replication
- `common.NewValidation()` — validation builder with `.Require()`, `.Enum()`, `.StatusTransition()`, `.After()`, `.Build()`
- `common.GenerateID()` — UUID generation for primary keys on POST
- `common.RegisterType()` — type introspection registration
- Parallel activation with semaphore pattern

**What each service file provides (configuration — unique per service, ~15 lines):**
- Service name, area, primary key (3 values)
- Which fields are required (list of field accessors)
- Which fields are enums (field + enum map)
- Which After hooks to attach (service-specific logic only)

**Per-service file budget:**
- `*Service.go`: ~12 lines (constants + Activate function + optional helper)
- `*ServiceCallback.go`: ~10-20 lines (validation builder chain — config values only)
- Total per service: ~25-35 lines of configuration
- **No behavioral code is repeated** — it all delegates to the framework

**Services with special After hooks** (unique behavioral logic that IS per-service):
- Enrollment (invite flow, consent validation, activation cascade)
- Worksheet (PDF generation, personalization)
- WorksheetScan (AI extraction, math equivalence grading)
- DailySchedule (AI schedule generation)
- CollabMessage (AI moderation)
- RiskAssessment (AI risk prediction)

These 6 services have unique After callbacks (~50-100 lines each). The other 32 services have config-only callbacks with no custom behavioral code.

**Proto List types:** The `*List { repeated X list = 1; l8api.L8MetaData metadata = 2; }` pattern is a protobuf convention required by the Layer 8 framework serializer. It cannot be abstracted — protobuf has no generics or macros. This is an unavoidable 4-line repetition per type (framework constraint, not design choice).

### 3.1 Service Registry

| ServiceName | ServiceArea | PrimaryKey | Model | Notes |
|-------------|:-----------:|------------|-------|-------|
| `Course` | 10 | `CourseId` | `Course` | |
| `Unit` | 10 | `UnitId` | `Unit` | |
| `Lesson` | 10 | `LessonId` | `Lesson` | |
| `Activity` | 10 | `ActivityId` | `Activity` | Questions embedded |
| `Worksheet` | 10 | `WorksheetId` | `Worksheet` | WorksheetScore embedded; generates PDF via Layer8FileUpload |
| `FamActvty` | 10 | `FamilyActivityId` | `FamilyActivity` | StudentRole embedded; sibling collaboration |
| `RealWorld` | 10 | `LessonId` | `RealWorldLesson` | RealWorldChallenge embedded; turns daily life into curriculum |
| `Project` | 10 | `ProjectId` | `Project` | ProjectMilestone embedded; multi-week cross-subject projects |
| `Student` | 20 | `StudentId` | `Student` | |
| `Guardian` | 20 | `GuardianId` | `Guardian` | |
| `Teacher` | 20 | `TeacherId` | `Teacher` | |
| `Classroom` | 20 | `ClassroomId` | `Classroom` | |
| `School` | 20 | `SchoolId` | `School` | |
| `District` | 20 | `DistrictId` | `District` | |
| `Enroll` | 20 | `EnrollmentId` | `Enrollment` | ConsentRecord embedded; drives guardian invite + COPPA consent + diagnostic flow |
| `Family` | 20 | `FamilyId` | `Family` | Links guardians + students for multi-child management |
| `Comply` | 20 | `ComplianceId` | `StateCompliance` | ComplianceRequirement embedded; auto-tracks hours, subjects, deadlines |
| `Pod` | 20 | `PodId` | `LearningPod` | Homeschool co-op coordination |
| `LearnPath` | 30 | `PathId` | `LearningPath` | PathStep, AdaptationLog embedded |
| `Mastery` | 30 | `MasteryId` | `SkillMastery` | MasterySnapshot embedded |
| `Skill` | 30 | `SkillId` | `Skill` | |
| `AdaptRule` | 30 | `RuleId` | `AdaptationRule` | |
| `Schedule` | 30 | `ScheduleId` | `DailySchedule` | ScheduleBlock embedded; AI-generated daily family plan |
| `LearnSess` | 40 | `SessionId` | `LearningSession` | Interaction embedded |
| `Score` | 40 | `ScoreId` | `Score` | |
| `Benchmark` | 40 | `BenchmarkId` | `Benchmark` | |
| `WkshtScan` | 40 | `ScanId` | `WorksheetScan` | ScannedAnswer embedded; AI extracts + grades handwritten answers |
| `Progress` | 50 | `ReportId` | `ProgressReport` | |
| `Engage` | 50 | `MetricId` | `EngagementMetric` | |
| `Growth` | 60 | `GrowthId` | `GrowthRecord` | 1 per student per subject per year; updated after each session, frozen at year-end |
| `Cohort` | 60 | `SnapshotId` | `CohortSnapshot` | SkillGap embedded; computed weekly/monthly/year-end per classroom/school/district |
| `RiskAssmt` | 60 | `AssessmentId` | `RiskAssessment` | RiskFactor embedded; AI-generated weekly batch job |
| `StdMastry` | 60 | `StandardMasteryId` | `StandardMastery` | Common Core / state standards; updated when SkillMastery changes |
| `CntEffect` | 60 | `EffectId` | `ContentEffect` | Curriculum effectiveness; computed quarterly |
| `Collab` | 70 | `GroupId` | `CollabGroup` | GroupMember + ProjectDeliverable embedded; study groups, project teams |
| `CollabMsg` | 70 | `MessageId` | `CollabMessage` | AI-moderated before delivery; tracks explanations + help given |
| `TutorPair` | 70 | `MatchId` | `TutorMatch` | AI-matched peer tutoring; tracks learner improvement |
| `Challenge` | 70 | `ChallengeId` | `Challenge` | ChallengeTeam embedded; weakest-member bonus incentivizes helping |

All ServiceNames are 10 characters or less. All ServiceAreas are consistent within their module.

### 3.2 ServiceCallback Patterns

**Callback factory pattern** (follows `../l8erp/go/erp/crm/leads/CrmLeadServiceCallback.go`):

```go
// Example: go/learn/content/courses/CourseServiceCallback.go
func newCourseServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
    return common.NewValidation(&learn.Course{}, vnic).
        Require(func(v interface{}) string { return v.(*learn.Course).CourseId }, "CourseId").
        Require(func(v interface{}) string { return v.(*learn.Course).Name }, "Name").
        Enum(func(v interface{}) int32 { return int32(v.(*learn.Course).Subject) }, "Subject", learn.SubjectType_name).
        Enum(func(v interface{}) int32 { return int32(v.(*learn.Course).Status) }, "Status", learn.ContentStatus_name).
        Build()
}

// Example: go/learn/students/enrollments/EnrollmentServiceCallback.go
func newEnrollmentServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
    return common.NewValidation(&learn.Enrollment{}, vnic).
        Require(func(v interface{}) string { return v.(*learn.Enrollment).EnrollmentId }, "EnrollmentId").
        Require(func(v interface{}) string { return v.(*learn.Enrollment).StudentId }, "StudentId").
        Require(func(v interface{}) string { return v.(*learn.Enrollment).SchoolId }, "SchoolId").
        StatusTransition(enrollmentTransitions()).
        After(onEnrollmentChange).
        Build()
}

// Example: go/learn/assessment/worksheetscans/WorksheetScanServiceCallback.go
func newWorksheetScanServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
    return common.NewValidation(&learn.WorksheetScan{}, vnic).
        Require(func(v interface{}) string { return v.(*learn.WorksheetScan).ScanId }, "ScanId").
        Require(func(v interface{}) string { return v.(*learn.WorksheetScan).WorksheetId }, "WorksheetId").
        After(onScanUploaded).
        Build()
}
```

**Auto-generate ID on POST** — all ServiceCallbacks include `common.GenerateID()` in the Before hook. The validation builder handles this automatically when `Require()` is specified for the primary key field.

**Parallel activation** — `go/learn/services/activate_all.go` uses semaphore-based parallelism (20 workers) matching `../l8erp/go/erp/services/activate_all.go`:
```go
const parallelWorkers = 20
// ... semaphore + sync.WaitGroup pattern (see l8erp for reference)
```

**Special callbacks:**

- **Activity callback**: Validate that `questions` have unique `question_id` values, all `is_correct` options sum to at least 1 per question
- **LearningPath callback (After POST)**: Trigger initial diagnostic benchmark, populate `upcoming_queue` with first activities
- **LearningSession callback (After PUT/PATCH)**: When session completes, update `SkillMastery` records, trigger adaptive engine to refresh `LearningPath`
- **SkillMastery callback (After PATCH)**: When mastery level changes, notify guardian via `l8notify`, log event via `l8events`
- **AdaptationRule callback**: Validate trigger/strategy combinations make sense (e.g., SCORE_ABOVE should not trigger REVIEW strategy)
- **Enrollment callback (After POST)**: Create Student + Guardian records if not existing, generate invite token, send guardian invite email via `l8notify`, update status DRAFT → INVITED
- **Enrollment callback (After PUT)**: On status change to CONSENTED → validate all required ConsentRecords are signed. On status change to ACTIVE → create student login via `ISecurityProvider`, assign "student" role, create initial LearningPaths, schedule diagnostic benchmark, notify teacher + guardian
- **Worksheet callback (After POST)**: Fetch questions from `skill_ids`/`activity_ids`, apply shuffle/difficulty/count filters, render to PDF using Go PDF library, generate answer key PDF, store both via `Layer8FileUpload`, update status DRAFT → GENERATED. If `personalized=true`, generate per-student variants based on each student's `SkillMastery` levels
- **Worksheet callback (After PUT)**: When WorksheetScores are added (manual grading or from scan), update status → SCORED, feed scores to `SkillMastery`, trigger `LearningPath` recalculation
- **WorksheetScan callback (After POST)**: Send scanned image + original worksheet definition to AI (l8agent LLM), parse extracted answers into `ScannedAnswer` records, auto-grade by comparing extracted vs correct (with math equivalence: 5/4 = 1 1/4 = 1.25), flag low-confidence items for teacher review, update status → REVIEW or COMPLETE
- **WorksheetScan callback (After PUT)**: When teacher reviews flagged items and status changes to COMPLETE, recalculate score with teacher overrides, create `WorksheetScore` on parent `Worksheet`, update `SkillMastery`, trigger `LearningPath` recalculation, notify guardian if configured

**Computed service callbacks (historical analytics — triggered by operational data changes):**

- **GrowthRecord**: Updated by `SkillMastery` callback (After PATCH) — recalculates growth metrics after each mastery change. Year-end snapshots frozen as permanent records by scheduled job
- **CohortSnapshot**: Scheduled batch job (weekly during school year, monthly during summer) — aggregates all student data within each classroom, school, and district. Also triggered on-demand by admin dashboard load
- **RiskAssessment**: Weekly AI batch job — for every active student, analyzes mastery trajectory, engagement signals, session patterns, and historical peer data to predict risk. Uses `l8agent` LLM with structured output for factor identification and intervention recommendations
- **StandardMastery**: Updated by `SkillMastery` callback — when a skill's mastery changes, all standards that map to that skill are recalculated (skills map to standards via `Skill.standard_ids`)
- **ContentEffect**: Quarterly batch aggregation — for every activity/lesson/course, computes mean score, mastery gain, completion rate, and efficiency across all students who attempted it. AI generates effectiveness narrative

**Homeschool service callbacks:**

- **Family (After POST)**: Link existing guardians and students by ID. Set up state compliance record for the family's state.
- **StateCompliance (scheduled daily)**: Auto-increment `instruction_hours_logged` from session data. Auto-track `subjects_covered` from activities completed. Alert parent 30 days before deadlines via `l8notify`. Auto-generate portfolio PDF from work samples (worksheets, scores, project photos) when portfolio review is due.
- **DailySchedule (After POST)**: AI generates schedule blocks considering all children's learning paths, parent energy, appointments, weather, and attention spans by age. Inserts sibling collaboration blocks and breaks automatically.
- **FamilyActivity (After PUT)**: When activity marked complete, create `Interaction` records for each participating student, update their respective `SkillMastery`, credit the parent-logged time to `StateCompliance.instruction_hours_logged`.
- **RealWorldLesson (After PUT)**: Same as FamilyActivity — log challenges completed, update mastery, count toward compliance hours. Store parent-uploaded photos in compliance portfolio.
- **Project (After PUT)**: When milestones completed, update skills for participating students. Store deliverable photos/files for portfolio.

**Collaboration service callbacks:**

- **CollabMessage (Before POST)**: AI moderation — block profanity/personal info/bullying. If student shares a direct answer, coach them to explain the approach instead. Rate limit: 20 messages per session per student.
- **CollabMessage (After POST)**: If type is QUESTION, AI checks if it can provide a hint. If EXPLANATION, verify correctness — if wrong, AI gently interjects. If `contains_explanation=true`, award "helper" points to sender.
- **TutorMatch (After POST)**: Create CollabGroup of type TUTORING for the pair. Generate tutor guide prompts via AI. Notify both students. Schedule check-in after 3 sessions.
- **TutorMatch (After PUT)**: On completion, evaluate learner improvement. If successful (learner reached PROFICIENT), award badges to both. Feed tutoring effectiveness into `ContentEffect`. Update both students' `EngagementMetric`.
- **Challenge (After POST)**: Auto-assign students to balanced teams (mix of mastery levels). Create team CollabGroups. Generate daily leaderboard updates.
- **CollabGroup (scheduled daily)**: If team member inactive 2+ days, send gentle nudge. If whole team inactive 3+ days, notify teacher. Update team streak counter.

### 3.3 Adaptive Engine — AI Integration

The adaptive engine runs as a service using the `l8agent` pattern:

```go
// On activity completion:
// 1. Score the interaction
// 2. Update SkillMastery
// 3. Evaluate AdaptationRules (rule-based layer)
// 4. If rules produce no clear action → invoke AI (l8agent LLM)
// 5. AI considers: mastery levels, learning velocity, engagement, time-of-day,
//    what worked for similar students, accommodation requirements
// 6. AI outputs: next activity + difficulty + reasoning
// 7. Update LearningPath.upcoming_queue and adaptation_log
```

The AI layer uses the same orchestration pattern as `l8agent/go/services/chat/orchestrate.go` — tool loop with max iterations, token tracking, and structured output.

#### AI Input Matrix for Path Decisions

The adaptive engine collects the following data to decide each student's next activity:

| Input | Source Service | Updates | What It Tells the AI |
|-------|---------------|---------|---------------------|
| Diagnostic results | `Benchmark` + `LearnSess` | Once (placement) | Starting skill levels |
| Skill prerequisite graph | `Skill` | Static (curriculum) | What CAN vs CANNOT be taught next |
| Interaction data | `LearnSess` (Interaction) | Every activity | How the child thinks, specific error patterns, strategies used |
| Mastery trajectory | `Mastery` (MasterySnapshot) | After each session | Improving, plateauing, or declining per skill |
| Engagement signals | `Engage` | Continuous | Session length, time-of-day patterns, activity type preferences, abandonment patterns |
| Accommodations | `Student` (IEP/504 flags) | Set by admin | Extended time, text-to-speech, visual supports, bilingual needs |
| Peer patterns | `Score` + `Mastery` (aggregate) | Weekly rollup | What activity types worked for similar students at the same level |
| Worksheet scores | `WkshtScan` / `Worksheet` | When scanned/graded | Offline performance, handwritten error patterns (e.g., skip-counting errors) |
| Teacher overrides | `LearnPath` | Teacher sets | Priority skills, blocked progressions, per-student notes |
| Growth trajectory | `Growth` | After each session | Is this student growing faster/slower than expected? Ahead or behind peers at same starting level? |
| Cohort context | `Cohort` | Weekly snapshot | What's typical for this classroom/school? What are the common gaps? |
| Risk signals | `RiskAssmt` | Weekly AI batch | Is this student flagged? What factors contribute? |
| Standards alignment | `StdMastry` | On mastery change | Which standards are met vs unmet? What's needed for grade-level proficiency? |
| Content effectiveness | `CntEffect` | Quarterly | Which activities produced the best outcomes for students with similar profiles? |

The AI balances: **what to teach** (skill graph + mastery gaps), **how to teach it** (activity types + learning style), **when to push vs pull back** (engagement + error patterns + accommodation requirements), and **what the teacher wants** (overrides). It includes contingency logic: "if student gets 2 wrong in a row on activity #2, swap to a confidence builder before retrying."

### 3.4 Type Registration (ui/)

Type registration lives in per-module files under `go/learn/ui/`:
- `go/learn/ui/shared_content.go` — `registerContentTypes()`
- `go/learn/ui/shared_students.go` — `registerStudentTypes()`
- `go/learn/ui/shared_adaptive.go` — `registerAdaptiveTypes()`
- `go/learn/ui/shared_assessment.go` — `registerAssessmentTypes()`
- `go/learn/ui/shared_analytics.go` — `registerAnalyticsTypes()`
- `go/learn/ui/shared_history.go` — `registerHistoryTypes()`
- `go/learn/ui/shared_collab.go` — `registerCollabTypes()`

Entry point (`go/learn/ui/main/main.go`):
```go
func main() {
    svr := common.CreateWebServer("web", ui.RegisterTypes)
    svr.Start()
}
```

Registration pattern (`go/learn/ui/shared_content.go`):
```go
package ui

import (
    "github.com/saichler/l8learn/go/types/learn"
    "github.com/saichler/l8types/go/ifs"
)

func registerContentTypes(resources ifs.IResources) {
    common.RegisterType(resources, &learn.Course{}, &learn.CourseList{}, "CourseId")
    common.RegisterType(resources, &learn.Unit{}, &learn.UnitList{}, "UnitId")
    common.RegisterType(resources, &learn.Lesson{}, &learn.LessonList{}, "LessonId")
    common.RegisterType(resources, &learn.Activity{}, &learn.ActivityList{}, "ActivityId")
    common.RegisterType(resources, &learn.Worksheet{}, &learn.WorksheetList{}, "WorksheetId")
    common.RegisterType(resources, &learn.FamilyActivity{}, &learn.FamilyActivityList{}, "FamilyActivityId")
    common.RegisterType(resources, &learn.RealWorldLesson{}, &learn.RealWorldLessonList{}, "LessonId")
    common.RegisterType(resources, &learn.Project{}, &learn.ProjectList{}, "ProjectId")
}
```

Aggregator (`go/learn/ui/register.go`):
```go
func RegisterTypes(resources ifs.IResources) {
    registerContentTypes(resources)
    registerStudentTypes(resources)
    registerAdaptiveTypes(resources)
    registerAssessmentTypes(resources)
    registerAnalyticsTypes(resources)
    registerHistoryTypes(resources)
    registerCollabTypes(resources)
}
```

---

## 4. UI Design

### 4.1 Desktop Modules

| Module | Section Key | Sub-Modules | Services |
|--------|------------|-------------|----------|
| Content | `content` | Curriculum, Worksheets, Projects | Course, Unit, Lesson, Activity, Worksheet, FamActvty, RealWorld, Project |
| Students | `students` | People, Organizations, Enrollment, Homeschool | Student, Guardian, Teacher, Classroom, School, District, Enroll, Family, Comply, Pod |
| Learning | `learning` | Paths, Skills, Rules, Scheduling | LearnPath, Mastery, Skill, AdaptRule, Schedule |
| Assessment | `assessment` | Sessions, Results, Scans | LearnSess, Score, Benchmark, WkshtScan |
| Analytics | `analytics` | Reports, Engagement | Progress, Engage |
| History | `history` | Growth, Cohorts, Risk, Standards, Effectiveness | Growth, Cohort, RiskAssmt, StdMastry, CntEffect |
| Collaboration | `collab` | Groups, Chat, Tutoring, Challenges | Collab, CollabMsg, TutorPair, Challenge |

### 4.2 View Types per Service

| Service | Views | Notes |
|---------|-------|-------|
| Course | table, kanban (by status) | |
| Activity | table, tree (by lesson > unit > course) | |
| Worksheet | table, calendar (by due date) | PDF download, bulk scoring UI, scan integration |
| Student | table | |
| Classroom | table | |
| Enroll | table, kanban (by enrollment status) | Wizard for new enrollment, guardian invite tracking |
| LearnPath | table, timeline (adaptation log) | |
| Mastery | table, chart (mastery over time) | |
| Skill | table, tree (prerequisite graph) | |
| AdaptRule | table | |
| LearnSess | table, timeline, calendar | |
| Score | table, chart (score trends) | |
| Benchmark | table | |
| WkshtScan | table | Batch scan upload, flagged-item review UI with cropped answer images |
| Progress | table, chart | |
| Engage | table, chart (engagement over time) | |
| Growth | table, chart (growth vs expected, year-over-year line) | Filter by school, grade, subject |
| Cohort | table, chart (mastery distribution bar, score trends line) | Drill-down: district → school → classroom |
| RiskAssmt | table, kanban (by risk level) | Teacher acknowledge flow, intervention tracking |
| StdMastry | table, chart (standards heatmap) | Filter by standard framework, grade |
| CntEffect | table, chart (effectiveness scatter plot) | Sort by mastery gain per minute |
| FamActvty | table, kanban (by subject) | Multi-child role assignments, materials list |
| RealWorld | table, kanban (by context) | Filter by location/context type |
| Project | table, timeline (milestones), gantt | Multi-week view with deliverable tracking |
| Family | table | Multi-child dashboard view |
| Comply | table, calendar (deadlines) | Auto-generated portfolio download, deadline alerts |
| Pod | table | Shared curriculum, meeting schedule |
| Schedule | calendar (daily blocks) | Drag-drop time blocks, parent energy selector |
| Collab | table, kanban (by group status) | Team scores, member activity |
| CollabMsg | timeline (chat view) | AI moderation badges, explanation highlighting |
| TutorPair | table, kanban (by status) | Tutor/learner pairing with improvement tracking |
| Challenge | table, chart (leaderboard) | Team rankings, streak counters |

### 4.3 Dashboard Widgets

```
┌──────────────┬──────────────┬──────────────┬──────────────┐
│  Active      │  Skills      │  Avg Daily   │  Engagement  │
│  Students    │  Mastered    │  Minutes     │  Rate        │
│  1,247       │  8,432       │  23 min      │  78%         │
│  ▲ 12%       │  ▲ 340/week  │  ▼ from 25   │  ▲ from 71%  │
└──────────────┴──────────────┴──────────────┴──────────────┘
```

**Historical Analytics Dashboard (district/school/teacher views):**
```
┌──────────────┬──────────────┬──────────────┬──────────────┐
│  Mean Growth │  On Grade    │  At Risk     │  Standards   │
│  vs Expected │  Level       │  Students    │  Coverage    │
│  1.2x        │  67%         │  312 (6.5%)  │  78%         │
│  ▲ above exp │  ▲ from 61%  │  ▼ from 8.1% │  ▲ from 71%  │
└──────────────┴──────────────┴──────────────┴──────────────┘

Role-specific views:
- District admin: year-over-year trends, school comparison, ROI metrics
- Principal: teacher comparison, school skill gaps, at-risk student list
- Teacher: student skill heatmap, growth sparklines, risk alerts
- Superintendent: investment per student, cost per mastery gain, proficiency trends
```

### 4.4 Student Activity Player

This is the child-facing UI — a separate simplified interface:

- **URL**: `/student.html` (not the admin `app.html`)
- **No l8ui module system** — this is a custom, gamified, child-friendly interface
- **Renders activities** based on `LearningPath.current_activity_id`
- **Sends interactions** back to `LearnSess` service via POST
- **Shows progress** as a visual map/journey (not tables)
- **Gamification**: points, streaks, badges, avatar customization
- **Accessibility**: large touch targets, text-to-speech, high contrast mode
- **Responsive**: works on tablets (primary school device)

---

## 5. Security & Compliance

### 5.1 FERPA Compliance

- All student data (PII) encrypted at rest and in transit (VNet AES-256)
- Role-based access via `ISecurityProvider`:
  - **District admin**: all data within their district
  - **Teacher**: only students in their classrooms
  - **Guardian**: only their linked children
  - **Student**: only their own data (read-only)
- Audit trail via `l8events` for every data access
- Data deletion support (right to erasure)

### 5.2 COPPA Compliance

- No direct data collection from children under 13 without guardian consent
- Guardian consent flow in enrollment process
- Minimal PII: no email/phone for students, only name + school-assigned ID
- No third-party tracking or advertising

### 5.3 Roles

```go
// Security roles defined in security config JSON
"district-admin"   → AllowAll for own district (scoped by district_id)
"school-admin"     → AllowAll for own school
"teacher"          → Allow GET/PUT for own classroom students; GET for content
"guardian"          → Allow GET for linked student progress/reports
"student"          → Allow GET for own path/progress; POST for interactions only
"content-author"   → Allow CRUD for content; GET for analytics
```

### 5.4 Security Config JSON Design

The security config is consumed by `ISecurityProvider` at startup. Follows the deny-before-allow model from `l8secure`.

**Role definitions with L8Rule allow/deny:**

```json
{
  "roles": [
    {
      "roleId": "district-admin",
      "roleName": "District Administrator",
      "rules": {
        "da-allow-all": {
          "elemType": "*",
          "allowed": true,
          "actions": { "1": true, "2": true, "3": true, "4": true, "5": true },
          "attributes": { "row": "districtId=${user.districtId}" }
        }
      }
    },
    {
      "roleId": "teacher",
      "roleName": "Teacher",
      "rules": {
        "t-allow-content-read": {
          "elemType": "Course,Unit,Lesson,Activity,Skill,Benchmark",
          "allowed": true,
          "actions": { "5": true }
        },
        "t-allow-students-rw": {
          "elemType": "Student,LearningPath,SkillMastery,LearningSession,Score,ProgressReport,EngagementMetric,GrowthRecord,RiskAssessment",
          "allowed": true,
          "actions": { "2": true, "3": true, "5": true },
          "attributes": { "row": "classroomId IN (${user.classroomIds})" }
        },
        "t-allow-worksheet-crud": {
          "elemType": "Worksheet,WorksheetScan",
          "allowed": true,
          "actions": { "1": true, "2": true, "3": true, "4": true, "5": true },
          "attributes": { "row": "teacherId=${user.teacherId}" }
        },
        "t-allow-collab-manage": {
          "elemType": "CollabGroup,Challenge,TutorMatch",
          "allowed": true,
          "actions": { "1": true, "2": true, "3": true, "4": true, "5": true },
          "attributes": { "row": "classroomId IN (${user.classroomIds})" }
        },
        "t-deny-sensitive-fields": {
          "elemType": "Student",
          "allowed": false,
          "actions": { "5": true },
          "attributes": { "field": "accommodationNotes,has504Plan" }
        }
      }
    },
    {
      "roleId": "guardian",
      "roleName": "Guardian",
      "rules": {
        "g-allow-child-read": {
          "elemType": "Student,LearningPath,SkillMastery,ProgressReport,EngagementMetric,GrowthRecord,Score",
          "allowed": true,
          "actions": { "5": true },
          "attributes": { "row": "studentId IN (${user.studentIds})" }
        },
        "g-allow-enrollment": {
          "elemType": "Enrollment",
          "allowed": true,
          "actions": { "2": true, "3": true, "5": true },
          "attributes": { "row": "guardianId=${user.guardianId}" }
        },
        "g-allow-family": {
          "elemType": "Family,StateCompliance,DailySchedule",
          "allowed": true,
          "actions": { "1": true, "2": true, "3": true, "5": true },
          "attributes": { "row": "familyId=${user.familyId}" }
        },
        "g-deny-other-students": {
          "elemType": "Student",
          "allowed": false,
          "actions": { "5": true },
          "attributes": { "row": "studentId NOT IN (${user.studentIds})" }
        }
      }
    },
    {
      "roleId": "student",
      "roleName": "Student",
      "rules": {
        "s-allow-own-read": {
          "elemType": "LearningPath,SkillMastery,EngagementMetric",
          "allowed": true,
          "actions": { "5": true },
          "attributes": { "row": "studentId=${user.studentId}" }
        },
        "s-allow-interactions-post": {
          "elemType": "LearningSession",
          "allowed": true,
          "actions": { "1": true, "2": true },
          "attributes": { "row": "studentId=${user.studentId}" }
        },
        "s-allow-collab-participate": {
          "elemType": "CollabMessage",
          "allowed": true,
          "actions": { "1": true, "5": true },
          "attributes": { "row": "senderId=${user.studentId}" }
        },
        "s-allow-content-read": {
          "elemType": "Activity,Lesson,Course",
          "allowed": true,
          "actions": { "5": true }
        },
        "s-deny-other-students": {
          "elemType": "Student,LearningSession,Score",
          "allowed": false,
          "actions": { "5": true },
          "attributes": { "row": "studentId!=${user.studentId}" }
        }
      }
    },
    {
      "roleId": "content-author",
      "roleName": "Content Author",
      "rules": {
        "ca-allow-content-crud": {
          "elemType": "Course,Unit,Lesson,Activity,Skill,FamilyActivity,RealWorldLesson,Project,Benchmark",
          "allowed": true,
          "actions": { "1": true, "2": true, "3": true, "4": true, "5": true }
        },
        "ca-allow-analytics-read": {
          "elemType": "ContentEffect,CohortSnapshot",
          "allowed": true,
          "actions": { "5": true }
        }
      }
    }
  ]
}
```

**Row-level scoping**: Uses L8Query `where` expressions in `attributes.row` to restrict data access by user context (e.g., teacher sees only their classrooms' students, guardian sees only linked children).

**Field-level denials**: Uses `attributes.field` to hide sensitive fields (e.g., teacher cannot see 504 plan details — only the school admin or district admin can).

**Public data** (no row scoping needed): `Course`, `Unit`, `Lesson`, `Activity`, `Skill` — curriculum content is visible to all authenticated users.

**Sensitive data** (strict scoping): `Student.accommodationNotes`, `Student.has504Plan`, `Student.hasIep` — only visible to school-admin and district-admin roles.

### 5.5 Multi-Portal Architecture

Following l8erp's portal pattern, L8Learn serves 4 separate portal HTML files with distinct navigation and feature sets:

| Portal | File | Role | Features |
|--------|------|------|----------|
| Admin | `app.html` | district-admin, school-admin, content-author | Full l8ui module system: all 7 service areas, dashboards, content management |
| Teacher | `teacher.html` | teacher | Classroom dashboard, skill heatmap, risk alerts, worksheet management, collaboration oversight, grade scanning |
| Guardian | `guardian.html` | guardian | Child progress, engagement metrics, report cards, compliance dashboard, daily schedule, family activities |
| Student | `student.html` | student | Gamified activity player, progress map, avatar, streaks, collaboration chat |

**login.json portal routing:**
```json
{
    "app": {
        "portals": {
            "app.html": { "roles": ["district-admin", "school-admin", "content-author"] },
            "teacher.html": { "roles": ["teacher"] },
            "guardian.html": { "roles": ["guardian"] },
            "student.html": { "roles": ["student"] }
        }
    }
}
```

After login, `CreateWebServer` routes the user to the correct portal HTML based on their role assignment.

### 5.6 AI Agent Prompts & Configuration

Following l8erp's AIA pattern (`go/erp/aia/activate.go`), L8Learn defines role-specific AI prompts:

```go
// go/learn/aia/activate.go
var defaultPrompts = []*agent.L8AgentPrompt{
    {
        PromptId:       "teacher-classroom",
        SystemPrompt:   "You are a teaching assistant. Help the teacher understand student progress, identify struggling students, suggest interventions, and recommend activities for the class.",
        AllowedModules: []string{"students", "adaptive", "assessment", "analytics", "history"},
        Category:       agent.L8_AGENT_PROMPT_CATEGORY_ANALYSIS,
    },
    {
        PromptId:       "guardian-progress",
        SystemPrompt:   "You are a parent assistant. Help the guardian understand their child's learning progress, explain what skills are being worked on, and suggest how to support learning at home.",
        AllowedModules: []string{"students", "adaptive", "analytics"},
        Category:       agent.L8_AGENT_PROMPT_CATEGORY_REPORTING,
    },
    {
        PromptId:       "student-tutor",
        SystemPrompt:   "You are a friendly learning buddy for a child. Explain concepts simply, encourage effort, give hints (never answers), and celebrate progress. Use age-appropriate language.",
        AllowedModules: []string{"content", "adaptive"},
        Category:       agent.L8_AGENT_PROMPT_CATEGORY_WORKFLOW,
    },
    {
        PromptId:       "content-analyst",
        SystemPrompt:   "You are a curriculum effectiveness analyst. Help content authors understand which activities are most effective, identify content gaps, and suggest improvements based on student outcome data.",
        AllowedModules: []string{"content", "history"},
        Category:       agent.L8_AGENT_PROMPT_CATEGORY_ANALYSIS,
    },
}
```

**Two-phase AI activation** (must match l8erp pattern):
```go
// In main.go — chat service activates AFTER all other services
services.ActivateAllServices(dbcred, dbname, nic)    // Phase 1: all CRUD services
services.ActivateChatService(dbcred, dbname, nic)    // Phase 2: AI chat (needs full introspector)
evtservices.ActivateEvents(dbcred, dbname, nic)      // Phase 3: event tracking
```

### 5.7 External Service Activation

The following external Layer 8 services MUST be activated in `main.go` alongside L8Learn's own services:

| Service | Package | Purpose | Activation Order |
|---------|---------|---------|-----------------|
| L8Learn services | `services.ActivateAllServices()` | All 38 CRUD services (parallel) | Phase 1 |
| AI Agent chat | `services.ActivateChatService()` | AI orchestration (needs introspector populated) | Phase 2 (after Phase 1) |
| L8Events | `evtservices.ActivateEvents()` | Audit trail, event logging | Phase 3 |
| L8Notify | `notifyservices.ActivateNotify()` | Email/push notifications (guardian reports, alerts) | Phase 3 |

---

## 6. Mock Data

### 6.1 Phase Ordering

```
Phase 1: Foundation (no dependencies)
  - Districts (5)
  - Schools (20, 4 per district)
  - Skills (200, organized by subject/grade/domain)
  - AdaptationRules (15, default rule set)

Phase 2: People (depends on Phase 1)
  - Teachers (80, 4 per school)
  - Classrooms (100, 5 per school)
  - Guardians (500)
  - Students (1000, 10 per classroom, linked to guardians)
  - Enrollments (1000, 1 per student, all status=ACTIVE with consent records)

Phase 3: Content (depends on Phase 1)
  - Courses (20, by subject/grade)
  - Units (100, 5 per course)
  - Lessons (400, 4 per unit)
  - Activities (1200, 3 per lesson, with embedded questions)
  - Benchmarks (40, 2 per course)

Phase 4: Learning State (depends on Phases 1-3)
  - LearningPaths (1000, 1 per student per primary subject)
  - SkillMastery (5000, ~5 skills per student with progress)

Phase 5: Session Data (depends on Phase 4)
  - LearningSessions (3000, ~3 recent sessions per student)
  - Scores (5000, aggregated from sessions)

Phase 6: Worksheets & Scans (depends on Phases 2-4)
  - Worksheets (50, ~1 per classroom per subject, with WorksheetScores)
  - WorksheetScans (200, ~4 per worksheet batch, mixed statuses: COMPLETE + REVIEW)

Phase 7: Analytics (depends on Phases 5-6)
  - ProgressReports (1000, current period per student)
  - EngagementMetrics (1000, 1 per student)

Phase 8: Homeschool Features (depends on Phases 1-3)
  - Families (100, links guardians to students)
  - StateCompliance (100, 1 per family with ComplianceRequirements for their state)
  - LearningPods (20, 5 families per pod)
  - FamilyActivities (150, sibling collaboration activities with StudentRoles)
  - RealWorldLessons (100, across all contexts: grocery, cooking, nature, etc.)
  - Projects (30, multi-week cross-subject projects with milestones)
  - DailySchedules (500, ~5 recent days per family with ScheduleBlocks)

Phase 9: Collaboration (depends on Phases 2-4)
  - CollabGroups (80, mixed types: study=30, project=20, challenge=15, tutoring=15)
  - CollabMessages (2000, ~25 per active group)
  - TutorMatches (50, AI-matched pairs with outcome tracking)
  - Challenges (10, classroom-wide team competitions with ChallengeTeams)

Phase 10: Historical Analytics (depends on Phases 4-9)
  - GrowthRecords (2000, 1 per student per subject for current year + 1 prior year for trend)
  - CohortSnapshots (500, weekly snapshots for classrooms + monthly for schools + quarterly for districts)
  - RiskAssessments (200, ~20% of students flagged at WATCH or above with RiskFactors)
  - StandardMastery (3000, ~3 standards per student per subject)
  - ContentEffect (400, 1 per activity aggregated across all students)
```

### 6.2 Generator Files

```
go/tests/mocks/
├── data.go                          # Curated name arrays (student names, skill names, etc.)
├── store.go                         # ID slices: DistrictIDs, SchoolIDs, StudentIDs, etc.
├── gen_learn_foundation.go          # Phase 1: Districts, Schools, Skills, Rules
├── gen_learn_people.go              # Phase 2: Teachers, Classrooms, Guardians, Students
├── gen_learn_content.go             # Phase 3: Courses, Units, Lessons
├── gen_learn_activities.go          # Phase 3 continued: Activities + Questions
├── gen_learn_paths.go               # Phase 4: LearningPaths, SkillMastery
├── gen_learn_sessions.go            # Phase 5: Sessions, Scores
├── gen_learn_worksheets.go           # Phase 6: Worksheets, WorksheetScans
├── gen_learn_analytics.go           # Phase 7: ProgressReports, EngagementMetrics
├── gen_learn_enrollments.go         # Phase 2 continued: Enrollments with ConsentRecords
├── gen_learn_families.go             # Phase 8: Families, StateCompliance, Pods
├── gen_learn_homeschool_content.go  # Phase 8: FamilyActivities, RealWorldLessons, Projects
├── gen_learn_schedules.go           # Phase 8: DailySchedules
├── gen_learn_collab_groups.go       # Phase 9: CollabGroups, CollabMessages
├── gen_learn_collab_tutoring.go     # Phase 9: TutorMatches, Challenges
├── gen_learn_growth.go              # Phase 10: GrowthRecords (current + prior year)
├── gen_learn_cohorts.go             # Phase 10: CohortSnapshots at all aggregation levels
├── gen_learn_risk.go                # Phase 10: RiskAssessments with RiskFactors
├── gen_learn_standards.go           # Phase 10: StandardMastery records
├── gen_learn_effectiveness.go       # Phase 10: ContentEffect records
├── learn_phases.go                  # Phase orchestration
└── cmd/main.go                      # CLI entry point
```

---

## 7. Deployment

### 7.1 Docker Images

| Image | Directory | Base | K8s Kind |
|-------|-----------|------|----------|
| `saichler/learn` | `go/learn/main/` | `saichler/learn-postgres` | StatefulSet |
| `saichler/learn-web` | `go/learn/ui/` | `saichler/learn-security` | DaemonSet (hostNetwork: true) |
| `saichler/learn-vnet` | `go/learn/vnet/` | `saichler/learn-security` | DaemonSet (hostNetwork: true) |
| `saichler/learn-log-vnet` | `go/logs/vnet/` | `saichler/learn-security` | DaemonSet (hostNetwork: true) |
| `saichler/learn-log-agent` | `go/logs/agent/` | `saichler/learn-security` | DaemonSet |
| `saichler/learn-maint` | `go/maint/` | `alpine` | DaemonSet | Maintenance/downtime page |

### 7.2 K8s Manifests

All manifests include: namespace labels, resource labels, NODE_IP env var, `hdata` volume convention (per `k8s-yaml-required-entries.md`).

**Deployment ordering** (in `k8s/deploy.sh`):
```bash
kubectl apply -f ./vnet.yaml          # 1. Virtual network (must be first)
kubectl apply -f ./log-vnet.yaml      # 2. Log virtual network
kubectl apply -f ./learn.yaml         # 3. Backend services
kubectl apply -f ./web.yaml           # 4. Web UI (hostNetwork: true)
kubectl apply -f ./log-agent.yaml     # 5. Log collection agent
kubectl apply -f ./maint.yaml         # 6. Maintenance page (optional, for planned downtime)
```

**Web tier specifics:**
- `hostNetwork: true` — serves HTTPS directly on host port 2773
- Serves all 4 portal HTML files from the same image
- Portal routing handled by `CreateWebServer` based on `login.json` portal config

### 7.3 run-local.sh

Copied and adapted from `../l8erp/go/run-local.sh`:

**Build order:**
1. `tests/mocks/cmd/` → `demo/mocks_demo`
2. `logs/agent/` → `demo/log-agent_demo`
3. `logs/vnet/` → `demo/log-vnet_demo`
4. `learn/vnet/` → `demo/vnet_demo`
5. `learn/main/` → `demo/learn_demo`
6. `learn/ui/main/` → `demo/ui_demo` (+ copy `./web` directory)

**Startup order:**
```bash
./log-vnet_demo &
./vnet_demo &
sleep 1
./log-agent_demo &
./learn_demo local &
./ui_demo &
sleep 8
# Mock data upload (after services healthy)
./mocks_demo --address https://${EXTERNAL_IP}:2773 --user admin --password admin --insecure
```

**Generates `kill_demo.sh`** to clean up all background processes.

---

## 8. Configuration

### 8.1 login.json

```json
{
    "login": {
        "appTitle": "L8Learn",
        "appDescription": "Adaptive Learning Platform",
        "authEndpoint": "/auth",
        "redirectUrl": "/app.html",
        "sessionTimeout": 30,
        "tfaEnabled": false,
        "showRememberMe": true,
        "showRegister": true
    },
    "app": {
        "dateFormat": "mm/dd/yyyy",
        "apiPrefix": "/learn",
        "healthPath": "/0/Health",
        "portals": {
            "app.html": { "roles": ["district-admin", "school-admin", "content-author"] },
            "teacher.html": { "roles": ["teacher"] },
            "guardian.html": { "roles": ["guardian"] },
            "student.html": { "roles": ["student"] }
        }
    }
}
```

### 8.2 ModConfig

Not needed initially — all modules enabled by default. Can be added later when districts want to disable specific sections.

---

## 9. Compliance Checklist

### Project Structure & Architecture
- [x] Follows l8erp directory layout
- [x] File naming conventions match l8erp patterns (`*Service.go`, `*ServiceCallback.go`)
- [x] common/defaults.go with `PREFIX = "/learn"` and re-exported l8common helpers
- [x] Per-module activation files: `services/activate_content.go`, `activate_students.go`, etc.
- [x] Parallel service activation with semaphore (20 workers) in `activate_all.go`
- [x] UI type registration in per-module files: `ui/shared_content.go`, `ui/shared_students.go`, etc.
- [x] Mock data generators in `go/tests/mocks/` with `cmd/main.go` entry point

### Protobuf Design
- [x] Enum zero values are UNSPECIFIED
- [x] List types use `repeated X list = 1` + `l8api.L8MetaData metadata = 2`
- [x] No direct struct references between Prime Objects — ID fields only
- [x] Child entities (Question, AnswerOption, Interaction, PathStep, etc.) are embedded `repeated` fields
- [x] @PrimeObject annotation on all independent entities

### Service Design
- [x] All ServiceNames are 10 characters or less
- [x] ServiceArea byte constant consistent within each module (10, 20, 30, 40, 50, 60, 70)
- [x] ServiceCallback uses `common.NewValidation()` builder with `.Require()`, `.Enum()`, `.After()`, `.Build()`
- [x] ServiceCallback auto-generates primary key on POST via `common.GenerateID()`
- [x] Types registered via `common.RegisterType(resources, &Single{}, &List{}, "PrimaryKey")`
- [x] Each service has Activate() function calling `common.ActivateService(config, single, list, creds, dbname, vnic)`

### UI Design
- [x] Desktop module integration planned (config, enums, columns, forms, init per sub-module)
- [x] Mobile parity addressed (same module data, mobile nav config)
- [x] Child entities use inline tables (Questions in Activity form)
- [x] l8ui components used throughout (tables, forms, charts, widgets, views)

### Mock Data
- [x] All 38 services have mock data generators planned
- [x] Phase ordering accounts for dependencies
- [x] Generator files stay under 500 lines each

### Deployment
- [x] build.sh + Dockerfile for each deployable
- [x] K8s manifests with all required entries
- [x] run-local.sh adapted from l8erp
- [x] deploy.sh / undeploy.sh

### Security
- [x] Security Config JSON design included (Section 5.4)
- [x] Roles defined with allow/deny rules per entity type
- [x] Row-level scoping via L8Query expressions (districtId, classroomId, studentId)
- [x] Field-level denials for sensitive data (accommodationNotes, has504Plan)
- [x] Users/roles/credentials provisioned via ISecurityProvider (not custom endpoints)

### Configuration
- [x] login.json adapted (appTitle, apiPrefix)
- [x] ModConfig handling addressed (not needed initially)

---

## 10. Implementation Phases

### Phase 1: Foundation
- Proto definitions + `make-bindings.sh`
- Project skeleton (common/defaults.go, main entry points)
- l8ui submodule setup

### Phase 2: Content Services
- Course, Unit, Lesson, Activity services + callbacks
- Worksheet service + PDF generation callback (Go PDF library: `jung-kurt/gofpdf`)
- Content management UI (desktop + mobile)

### Phase 3: People & Enrollment Services
- Student, Guardian, Teacher, Classroom, School, District services
- Enrollment service with full lifecycle callbacks (invite → consent → activate)
- Guardian onboarding page (`enroll.html` — standalone wizard, not admin UI)
- Bulk enrollment via CSV import (L8DataImport mapping template)
- People management UI with enrollment tracker dashboard

### Phase 4: Skill Graph & Mastery
- Skill, SkillMastery services
- Skill graph visualization (tree view)

### Phase 5: Adaptive Engine
- LearningPath, AdaptationRule services
- AI integration via l8agent pattern
- Rule evaluation engine

### Phase 6: Assessment & Worksheet Scanning
- LearningSession, Score, Benchmark services
- WorksheetScan service with AI extraction callback (l8agent LLM integration)
- Math equivalence evaluator (5/4 = 1 1/4 = 1.25)
- Batch scan processing (parallel AI calls for classroom-sized stacks)
- Scan review UI with cropped answer images for flagged items
- Score feedback loop: scan results → WorksheetScore → SkillMastery → LearningPath

### Phase 7: Analytics & Reporting
- ProgressReport, EngagementMetric services
- Dashboard widgets, charts, timeline views
- AI-generated narrative summaries

### Phase 8: Homeschool Features
- Family, StateCompliance, LearningPod services
- FamilyActivity, RealWorldLesson, Project services (content area 10)
- DailySchedule service + AI schedule generation
- Multi-child family dashboard (single view for all children)
- State compliance automation (auto-track hours, subjects, deadlines)
- Portfolio auto-generation (PDF from work samples, scores, photos)
- Real-world lesson logging (parent logs grocery/cooking/nature activities)
- Project-based learning with milestone tracking and photo uploads
- Parent burnout detection + "light week" mode

### Phase 9: Collaboration Center
- CollabGroup, CollabMessage, TutorMatch, Challenge services
- AI moderation pipeline (Before POST: block/coach/flag)
- Peer tutoring matching engine (weekly AI batch: match MASTERED students with EMERGING students)
- Team challenge system with balanced auto-assignment and leaderboards
- Study chat UI with explanation highlighting and helper badges
- Weakest-member bonus scoring (incentivizes helping, not carrying)
- Safety layer: rate limiting, personal info detection, bullying pattern detection
- Teacher collaboration analytics dashboard
- Parent opt-in consent for collaboration features (COPPA)

### Phase 10: Historical Analytics & Risk Prediction
- GrowthRecord service + computation callback triggered by SkillMastery changes
- CohortSnapshot service + scheduled batch computation (weekly/monthly/year-end)
- RiskAssessment service + weekly AI batch job for early warning prediction
- StandardMastery service + automatic update when skills map to standards
- ContentEffect service + quarterly aggregation batch job
- Historical dashboards: district admin (year-over-year, school comparison, ROI), principal (teacher comparison, skill gaps, at-risk list), teacher (student heatmap, growth sparklines, risk alerts), superintendent (investment metrics, proficiency trends)
- Standards coverage heatmap and gap analysis views
- Curriculum effectiveness scatter plot (mastery gain per minute vs usage)

### Phase 11: Student Activity Player
- Child-facing gamified UI (`student.html`)
- Activity renderers per activity type
- Real-time interaction tracking
- Gamification (points, streaks, badges, avatars)

### Phase 12: Guardian & Teacher Portals
- Guardian view: child progress, reports, notifications
- Teacher view: classroom overview, skill heatmap, intervention alerts
- Risk alert integration: teacher sees at-risk students with AI recommendations

### Phase 13: Mock Data & Local Dev
- All mock data generators (10 phases, 38 services)
- run-local.sh
- End-to-end verification

### Phase 14: Deployment
- Docker images, K8s manifests
- deploy.sh / undeploy.sh
- build-all-images.sh

### Phase 15: End-to-End Verification
- Navigate every section, verify data loads
- Verify row clicks open details
- Verify forms submit correctly
- Verify on both desktop and mobile
- Verify student activity player works
- Verify adaptive engine produces reasonable paths
- Verify enrollment flow end-to-end: admin creates → guardian receives invite → consent → student activates → diagnostic → first activity
- Verify worksheet flow: teacher creates → PDF downloads → scores entered manually → mastery updates
- Verify scan flow: teacher uploads image → AI extracts → teacher reviews flags → scores saved → mastery updates
- Verify personalized worksheets generate different content per student based on mastery levels
- Verify growth records update after SkillMastery changes
- Verify cohort snapshots aggregate correctly at classroom → school → district levels
- Verify risk assessments identify at-risk students with meaningful factors and recommendations
- Verify standards mastery updates when underlying skills change
- Verify content effectiveness reflects actual student outcomes
- Verify historical dashboard drill-down: district → school → classroom → student
- Verify year-over-year comparison charts render with prior year mock data
- Verify multi-child family dashboard shows all children's progress in single view
- Verify daily schedule generation considers all children, parent energy, and appointments
- Verify state compliance auto-tracks hours and subjects from session data
- Verify portfolio auto-generation produces PDF with work samples
- Verify real-world lesson logging credits hours to compliance
- Verify family activity completion updates mastery for all participating children
- Verify collaboration group creation and member management
- Verify AI moderation blocks inappropriate messages before delivery
- Verify peer tutoring match creates group and tracks learner improvement
- Verify team challenge leaderboard updates with weakest-member bonus
- Verify study chat rate limiting (20 messages per session)
- Verify COPPA consent required before collaboration features activate
