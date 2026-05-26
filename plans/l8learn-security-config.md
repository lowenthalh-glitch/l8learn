# Plan: L8Learn Security Config

## Overview

Create the security config JSON for l8learn at `go/secure/plugin/l8learn/l8learn.json`, defining 7 roles with appropriate allow/deny rules. Update mock data to set `associateIds` on guardian users so that `${associateIds}` deny-scope works.

## Roles Summary

| Role | Scoping | Portal |
|------|---------|--------|
| admin | None (full access) | app.html |
| district-admin | None (full access to educational data) | app.html |
| school-admin | None (full access to educational data) | app.html |
| teacher | None (can view all students) | teacher.html |
| content-author | None (action-level only) | app.html |
| guardian | `${associateIds}` — children's student IDs | guardian.html |
| student | `${userId}` — own studentId | student.html |

Only **guardian** and **student** roles need row-level deny rules. All other roles have unrestricted data access within their allowed actions.

## Entity Inventory by Service Area

### Content (area 10) — 8 entities
Course, Unit, Lesson, Activity, Worksheet, FamilyActivity, RealWorldLesson, Project

### Students (area 20) — 10 entities
Student, Guardian, Teacher, Classroom, School, District, Enrollment, Family, StateCompliance, LearningPod, StudentProfile, EvalImport

### Adaptive (area 30) — 7 entities
LearningPath, SkillMastery, Skill, AdaptationRule, DailySchedule, LLMConfig, LLMPromptLog

### Assessment (area 40) — 4 entities
LearningSession, Score, Benchmark, WorksheetScan

### Analytics (area 50) — 2 entities
ProgressReport, EngagementMetric

### History (area 60) — 5 entities
GrowthRecord, CohortSnapshot, RiskAssessment, StandardMastery, ContentEffect

### Collaboration (area 70) — 4 entities
CollabGroup, CollabMessage, TutorMatch, Challenge

## Role Definitions

### 1. admin
Full access to everything, including security entities (L8User, L8Role, L8Credentials).

```
Allow: all actions on all entity types
No deny rules
```

### 2. district-admin
Full CRUD on all educational entities. View-only on security entities.

```
Allow: all actions on all entity types
Deny: POST/PUT/PATCH/DELETE on L8User, L8Role, L8Credentials (view-only for security)
```

### 3. school-admin
Same as district-admin. Full CRUD on all educational entities, view-only on security.

```
Allow: all actions on all entity types
Deny: POST/PUT/PATCH/DELETE on L8User, L8Role, L8Credentials
```

### 4. teacher
Can view all students (no row scoping). CRUD on assessment-related entities. GET on everything else.

```
Allow: all actions (-999) on Worksheet, WorksheetScan, LearningSession, Score, CollabGroup, CollabMessage, Challenge, TutorMatch, DailySchedule
Allow: GET on all other entity types (content, students, adaptive, analytics, history)
Deny: all actions on L8User, L8Role, L8Credentials (no security access)
Field-level deny: student.accommodationNotes, student.hasIep, student.has504Plan (sensitive — only admin/school-admin/district-admin see these)
```

### 5. content-author
CRUD on content entities. GET on analytics.

```
Allow: all actions on Course, Unit, Lesson, Activity, Skill, FamilyActivity, RealWorldLesson, Project, Benchmark, Worksheet, GeneratedLesson
Allow: GET on ContentEffect, CohortSnapshot
Deny: all actions on L8User, L8Role, L8Credentials
No row-level scoping
```

### 6. guardian
GET on student-related entities scoped to their children. CRUD on family/enrollment. Uses `${associateIds}` with children's student IDs.

```
Allow: GET on Student, StudentProfile, EvalImport, LearningPath, SkillMastery, LearningSession, Score, ProgressReport, EngagementMetric, GrowthRecord, RiskAssessment, StandardMastery, DailySchedule, GeneratedLesson
Allow: GET/PUT on Enrollment (own enrollment management)
Allow: GET/POST/PUT on Family, StateCompliance
Allow: GET on Course, Unit, Lesson, Activity, Skill (public content — read-only)

Deny row-level (studentId not in ${associateIds}):
  Student, StudentProfile, EvalImport, LearningPath, SkillMastery,
  LearningSession, Score, ProgressReport, EngagementMetric,
  GrowthRecord, RiskAssessment, StandardMastery, GeneratedLesson

Deny row-level (guardianId!=${userId}):
  Guardian

Deny row-level (primaryGuardianId!=${userId}):
  Family

Field-level deny:
  student.accommodationNotes, student.hasIep, student.has504Plan
```

### 7. student
GET on own data, POST on learning interactions. Uses `${userId}` (= studentId).

```
Allow: GET on LearningPath, SkillMastery, EngagementMetric, Score, ProgressReport, GeneratedLesson (own data)
Allow: POST/PUT on LearningSession (submit interactions)
Allow: GET on Activity, Lesson, Course, Skill (public content)
Allow: POST/GET on CollabMessage (participation)

Deny row-level (studentId!=${userId}):
  Student, LearningPath, SkillMastery, LearningSession, Score,
  ProgressReport, EngagementMetric, GrowthRecord, GeneratedLesson

Field-level deny:
  student.accommodationNotes, student.hasIep, student.has504Plan,
  student.primaryGuardianId
```

## Pre-defined Users

```json
"users": {
  "admin": { "roles": {"admin": true}, "portal": "app.html", "password": "admin" }
}
```

Guardian and student users are created by the mock data generator via the Security API (already implemented in `gen_security_test_data.go`). The mock data needs updating to include `associateIds` on guardian users.

## Mock Data Changes

### Update `createUser` in `learn_phases.go`

Add an `associateIds` parameter for guardian users:

```go
func createUser(client Client, userId, fullName, email, role, portal string, associateIds []string) {
    userData := map[string]interface{}{
        "userId":        userId,
        "fullName":      fullName,
        "email":         email,
        "portal":        portal,
        "password":      map[string]string{"hash": "12345678"},
        "accountStatus": "ACCOUNT_STATUS_ACTIVE",
        "roles":         map[string]bool{role: true},
    }
    if len(associateIds) > 0 {
        userData["associateIds"] = associateIds
    }
    client.Post("/learn/73/users", userData)
}
```

### Update `gen_security_test_data.go`

Pass children's student IDs as `associateIds` for guardian users:

```go
createUser(client, g1.GuardianId, ..., "guardian", "guardian.html", g1.StudentIds)
createUser(client, g2.GuardianId, ..., "guardian", "guardian.html", g2.StudentIds)
createUser(client, g3.GuardianId, ..., "guardian", "guardian.html", g3.StudentIds)

// Student users — no associateIds
createUser(client, s1.StudentId, ..., "student", "student.html", nil)
```

## sysconfig Section

```json
"sysconfig": {
  "dataStoreConfig": { "name": "admin", "type": "postgres" },
  "rxQueueSize": "100000",
  "txQueueSize": "100000",
  "vnetPort": 10005,
  "webConfig": { "endPointPrefix": "/web/", "webPort": 2774 }
}
```

## Deliverables

| # | Deliverable | File |
|---|-------------|------|
| 1 | Security config JSON | `go/secure/plugin/l8learn/l8learn.json` |
| 2 | Update createUser to support associateIds | `go/tests/mocks/learn_phases.go` |
| 3 | Pass associateIds for guardian users | `go/tests/mocks/gen_security_test_data.go` |

## Phase 1: Create Security Config JSON

Create `go/secure/plugin/l8learn/l8learn.json` with all 7 roles, the admin user, credentials, and sysconfig.

## Phase 2: Update Mock Data

1. Add `associateIds` parameter to `createUser` in `learn_phases.go`
2. Update all `createUser` calls in `gen_security_test_data.go` to pass `associateIds` for guardians and `nil` for students
3. Update portal paths: guardian users → `guardian.html`, student users → `student.html`

## Phase 3: Verification

1. `go build ./...` — verify mock data compiles
2. Verify security config JSON is valid JSON (no syntax errors)
3. Verify every `elemType` in the config matches a real protobuf type name
4. Verify every field referenced in deny rules (`studentId`, `guardianId`, `primaryGuardianId`) exists on the target entity's protobuf struct
5. Verify `associateIds` is passed correctly for each guardian user mapping:
   - GRD-SEC-001 (Maria Garcia) → `["STU-SEC-001", "STU-SEC-002"]`
   - GRD-SEC-002 (James Wilson) → `["STU-SEC-003"]`
   - GRD-SEC-003 (Aisha Khan) → `["STU-SEC-004", "STU-SEC-005"]`

## Traceability Matrix

| # | Requirement | Phase |
|---|-------------|-------|
| 1 | admin role — full access | Phase 1 |
| 2 | district-admin role — full educational, view-only security | Phase 1 |
| 3 | school-admin role — full educational, view-only security | Phase 1 |
| 4 | teacher role — view all students, CRUD assessment, field-level deny on sensitive data | Phase 1 |
| 5 | content-author role — CRUD content, GET analytics | Phase 1 |
| 6 | guardian role — scoped to children via ${associateIds} | Phase 1 |
| 7 | student role — scoped to self via ${userId} | Phase 1 |
| 8 | Sensitive field denials (accommodationNotes, hasIep, has504Plan) for teacher/guardian/student | Phase 1 |
| 9 | Pre-defined admin user | Phase 1 |
| 10 | Mock data: associateIds on guardian users | Phase 2 |
| 11 | Mock data: correct portal paths | Phase 2 |
| 12 | Build verification | Phase 3 |
| 13 | JSON validity and field existence verification | Phase 3 |
