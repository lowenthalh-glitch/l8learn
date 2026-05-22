# Infrastructure Feature Request: Cross-Entity Deny-Scope

## Problem Statement

The current deny-scope mechanism in l8secure supports only **direct field matching** — the deny query compares a field on the target entity directly against `${userId}`. This works for entities that have the user's own ID as a field (e.g., `clientId!=${userId}` on ProgressLog in l8physio).

It does NOT support **cross-entity scoping** — where access to entity A depends on a relationship through entity B. This pattern is needed in every Layer 8 project that has hierarchical access control.

## Current Behavior

### What works (direct matching)
```json
{
  "ruleId": "deny-other-students",
  "elemType": "Student",
  "allowed": false,
  "attributes": {
    "Student": "select * from Student where primaryGuardianId!=${userId}"
  }
}
```
This works because `primaryGuardianId` is a field on `Student`, and `${userId}` is the guardian's ID. The L8Query engine resolves the field, compares the values, and filters correctly.

### What doesn't work (cross-entity)
A guardian needs to see only **profiles** for students they own. The `StudentProfile` entity has `studentId` but not `guardianId`. The guardian's access depends on the **Student** entity's `primaryGuardianId`.

**Attempted approach 1: Add `primaryGuardianId` to StudentProfile proto**
```json
{
  "elemType": "StudentProfile",
  "attributes": {
    "StudentProfile": "select * from StudentProfile where primaryGuardianId!=${userId}"
  }
}
```
Result: The field exists in the proto (field 39), data is populated correctly, L8Query can filter on it in regular GET queries, but ScopeView does not filter. The Student deny-scope (field 12, same field name) works. The difference is unknown — possibly related to field registration order in the introspector.

**Attempted approach 2: Subquery**
```json
{
  "attributes": {
    "StudentProfile": "select * from StudentProfile where studentId not in (select studentId from Student where primaryGuardianId=${userId})"
  }
}
```
Result: L8Query does not support subqueries. The `not in (select ...)` syntax is not implemented in the `NotIn` comparator — it only supports static lists `[a,b,c]`.

## Use Cases Across the Ecosystem

| Project | User Role | Target Entity | Scoped Via | Relationship |
|---------|-----------|---------------|------------|--------------|
| **l8learn** | Guardian | StudentProfile | Student | profile.studentId → student.primaryGuardianId=${userId} |
| **l8learn** | Guardian | EvalImport | Student | eval.studentId → student.primaryGuardianId=${userId} |
| **l8learn** | Guardian | DailySchedule | Student | schedule.customFields.studentId → student.primaryGuardianId=${userId} |
| **l8learn** | Guardian | GeneratedLesson | Student | lesson.studentId → student.primaryGuardianId=${userId} |
| **l8learn** | Teacher | Student | Classroom | student.classroomId → classroom.primaryTeacherId=${userId} |
| **probler** | GPU Operator | EventRecord | NetworkDevice | event.deviceId → device.deviceType="GPU" AND device.ownerId=${userId} |
| **probler** | Site Admin | Alert | NetworkDevice | alert.deviceId → device.siteId=${userId} |
| **l8erp** | Sales Rep | SalesOrder | SalesTerritory | order.territoryId → territory.assignedTo=${userId} |
| **l8erp** | Manager | Employee | Department | employee.departmentId → department.managerId=${userId} |
| **l8physio** | Therapist | ProgressLog | TreatmentPlan | log.planId → plan.therapistId=${userId} |

## Proposed Solution

### Option A: `scopeVia` Rule Extension (Recommended)

Add a new field to the deny rule structure that describes the cross-entity relationship:

```json
{
  "ruleId": "g-deny-other-profiles",
  "elemType": "StudentProfile",
  "allowed": false,
  "scopeVia": {
    "throughType": "Student",
    "localField": "studentId",
    "throughField": "studentId",
    "ownerField": "primaryGuardianId"
  }
}
```

**Interpretation**: "Deny access to StudentProfile records where the `studentId` field on the profile does not appear as `studentId` on any Student record where `primaryGuardianId` equals ${userId}."

**Processing in ScopeView**:
1. Load all records of `throughType` (Student) where `ownerField` (primaryGuardianId) == ${userId}
2. Collect the `throughField` (studentId) values → `[STU-SEC-001, STU-SEC-002]`
3. Filter target entities: keep only records where `localField` (studentId) is in the collected set
4. Return filtered results

**Pseudocode**:
```go
func (this *_securityProvider) applyScopeVia(rule ScopeViaRule, userId string, elements []interface{}) []interface{} {
    // Step 1: Load owner's related records
    query := fmt.Sprintf("select %s from %s where %s=%s",
        rule.ThroughField, rule.ThroughType, rule.OwnerField, userId)
    related := executeQuery(query)

    // Step 2: Collect allowed IDs
    allowedIds := set{}
    for _, r := range related {
        allowedIds.add(getField(r, rule.ThroughField))
    }

    // Step 3: Filter target elements
    var kept []interface{}
    for _, elem := range elements {
        localValue := getField(elem, rule.LocalField)
        if allowedIds.contains(localValue) {
            kept = append(kept, elem)
        }
    }
    return kept
}
```

### Option B: Resolved Static List at Auth Time

When a user authenticates, pre-resolve their cross-entity scope into a static list and cache it:

1. User logs in as GRD-SEC-001 (guardian)
2. Security provider queries: `select studentId from Student where primaryGuardianId=GRD-SEC-001`
3. Result: `[STU-SEC-001, STU-SEC-002]`
4. Cache as: `${userScope.studentIds} = [STU-SEC-001,STU-SEC-002]`
5. Deny rule becomes: `select * from StudentProfile where studentId not in ${userScope.studentIds}`

**Advantage**: No per-request cross-entity lookup — resolved once at login
**Disadvantage**: Stale if students are added/removed during the session

### Option C: Denormalize the Owner Field (Current Workaround)

Add the owner field (`primaryGuardianId`) to every entity that needs scoping. This is what we attempted — it works for direct matching but the deny-scope query doesn't resolve the field on StudentProfile (possibly an introspector issue).

**Advantage**: Uses existing deny-scope mechanism
**Disadvantage**: Data duplication, must keep in sync, and currently doesn't work for newly added fields

## Current Workaround Behavior

| Entity | Deny-Scope | Status |
|--------|-----------|--------|
| Student | `primaryGuardianId!=${userId}` (field 12, original) | **WORKS** ✅ |
| StudentProfile | `primaryGuardianId!=${userId}` (field 39, added later) | **DOES NOT WORK** ❌ |
| Guardian | `guardianId!=${userId}` (field 1, original) | **WORKS** ✅ |
| EvalImport | `primaryGuardianId!=${userId}` (field 17, added later) | **UNTESTED** (likely doesn't work) |

## Investigation Notes

### What was verified:
1. The field `primaryGuardianId` EXISTS on StudentProfile proto (field 39) ✅
2. The data IS populated correctly (`PROF-SEC-001` has `primaryGuardianId=GRD-SEC-001`) ✅
3. L8Query CAN filter by this field in regular GET requests (`where primaryGuardianId=GRD-SEC-001` returns 2 results) ✅
4. The deny rule IS in the security config with correct syntax ✅
5. The permissions endpoint returns correct allowed actions for the guardian role ✅
6. Student deny-scope with the SAME field name on a different type WORKS ✅
7. No error logs from ScopeView ✅

### What is different between Student (works) and StudentProfile (doesn't work):
- Student: `primaryGuardianId` is field 12 (part of the original proto design)
- StudentProfile: `primaryGuardianId` is field 39 (added after initial proto design)
- Both fields have the same name, same type (string), same JSON serialization
- The L8Query interpreter resolves both fields correctly for regular queries
- Only the ScopeView deny-scope processing differs

### Possible root cause:
The ScopeView's `interpreter.NewQuery()` at line 68 of `ScopeView.go` may use a different introspector state than the regular query path. The regular query path resolves fields through the ORM's type registry (which is rebuilt on startup). The ScopeView may use a cached introspector snapshot that was built before field 39 was added. This would explain why old fields work (Student field 12) but new fields don't (StudentProfile field 39).

## Recommendation

Implement **Option A (scopeVia)** in l8secure. It's:
- Generic — works for any cross-entity relationship
- Declarative — configured in JSON, no custom code per project
- Secure — enforced at the API level, not the UI
- Scalable — one lookup per request type, results can be cached per session
- Required by multiple projects (l8learn, probler, l8erp)

Additionally, investigate why deny-scope with `primaryGuardianId` on StudentProfile (field 39) doesn't work while the same field name on Student (field 12) does. This may be a separate introspector bug affecting all newly added fields in deny-scope rules.

## Affected Projects

- **l8learn** (immediate): Guardian and student data scoping
- **probler** (future): Device-type based event/alert scoping
- **l8erp** (future): Territory/department based order/employee scoping
- **l8physio** (potential): Therapist-to-client relationship scoping beyond direct clientId matching

## Priority

**High** — This is a security issue. Without cross-entity scoping, guardians can see other children's profiles, evaluations, and medical data through the API. The UI can hide it, but the API returns all records.
