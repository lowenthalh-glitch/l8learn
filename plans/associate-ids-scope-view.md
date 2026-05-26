# Plan: Associate IDs on L8User for Cross-Entity ScopeView

## Problem

The current ScopeView mechanism in l8secure only supports `${userId}` as a placeholder for deny rule resolution. This works for direct field matching (e.g., `clientId!=${userId}`) but not for cross-entity relationships where a user needs to see data belonging to a list of associated entities (e.g., a parent seeing only their children's records, a manager seeing only their departments' data).

Multiple projects need this capability — see Cross-Project Applicability below.

## Design

### Approach

1. Add `repeated string associate_ids = 19` to `L8User` in `secure.proto`
2. Extend ScopeView to support a `${associateIds}` placeholder that expands to `[id1,id2,id3]` bracket notation
3. Deny rules use existing L8Query `not in` syntax: `"select * from Entity where fieldId not in ${associateIds}"`
4. ScopeView resolves `${associateIds}` from the authenticated user's `associate_ids` field before parsing the L8Query
5. Consuming projects populate `associate_ids` on their users via the Security API

### How the deny rule works

Given a user with `associate_ids: ["ID-001", "ID-002"]`:

**Rule template:**
```
select * from Entity where entityId not in ${associateIds}
```

**After resolution:**
```
select * from Entity where entityId not in [ID-001,ID-002]
```

**Effect:** Rows where `entityId` is NOT in the list are denied (removed). Only records matching ID-001 and ID-002 are returned.

### Edge case: empty associate_ids

If a user has no `associate_ids`, the placeholder resolves to `[]` (empty list). The `not in []` comparison returns true for every row (nothing is in the empty list), so ALL rows are denied. This is the correct secure default — a user with no associates sees nothing.

## Affected Files (l8secure only)

| File | Change |
|------|--------|
| `proto/secure.proto` | Add `repeated string associate_ids = 19` to `L8User` |
| `go/types/secure/secure.pb.go` | Regenerated |
| `go/secure/provider/aaa.go` | Add `getAssociateIdsByToken(aaaid) []string` method |
| `go/secure/provider/ScopeView.go` | Add `${associateIds}` placeholder resolution |

## Phase 1: Proto Change

**File:** `proto/secure.proto`

Add field 19 to `L8User`:

```protobuf
message L8User {
  // ... existing fields 1-18 ...
  string portal = 18;
  repeated string associate_ids = 19;    // IDs of associated entities for cross-entity ScopeView filtering
}
```

Then regenerate bindings. Before running, verify `make-bindings.sh` uses `-i` (not `-it`) on all `docker run` commands — `-it` requires a TTY and fails in non-interactive environments:

```bash
cd proto && ./make-bindings.sh
```

Verify the generated field exists:
```bash
grep "AssociateIds" go/types/secure/secure.pb.go
```

## Phase 2: AAA Method

**File:** `go/secure/provider/aaa.go`

Add a method to retrieve associate IDs by token, following the same pattern as `getUserIdByToken` (line 334):

```go
// getAssociateIdsByToken returns the associate_ids for the user associated with the given aaaid token (read-locked)
func (aaa *AAA) getAssociateIdsByToken(aaaid string) []string {
    aaa.mu.RLock()
    defer aaa.mu.RUnlock()
    token, ok := aaa.tokens[aaaid]
    if !ok {
        return nil
    }
    user, ok := aaa.users[token.UserId]
    if !ok {
        return nil
    }
    return user.AssociateIds
}
```

## Phase 3: ScopeView Placeholder Resolution

**File:** `go/secure/provider/ScopeView.go`

Extend the row-level filter resolution block (lines 47-67) to also resolve `${associateIds}`:

Current code (lines 48, 67):
```go
userId := this.aaa.getUserIdByToken(aaaid)
// ...
resolvedValue := strings.ReplaceAll(value, "${userId}", userId)
```

Updated code:
```go
userId := this.aaa.getUserIdByToken(aaaid)
associateIds := this.aaa.getAssociateIdsByToken(aaaid)
associateIdsList := "[" + strings.Join(associateIds, ",") + "]"
// ...
resolvedValue := strings.ReplaceAll(value, "${userId}", userId)
resolvedValue = strings.ReplaceAll(resolvedValue, "${associateIds}", associateIdsList)
```

The `associateIds` resolution converts `[]string{"ID-001", "ID-002"}` to `[ID-001,ID-002]` — the bracket notation that L8Query's `not in` comparator expects (`getInStringList` in `l8ql/go/gsql/interpreter/comparators/In.go:139`).

If `associateIds` is nil/empty, `associateIdsList` becomes `[]`, and `not in []` denies all rows — secure default.

## Phase 4: Verification

1. Build: `go build ./...`
2. Verify proto field: `grep "AssociateIds" go/types/secure/secure.pb.go`
3. Verify the new AAA method compiles and returns correct data for a user with `associate_ids` set
4. Verify `${associateIds}` placeholder resolves correctly in ScopeView:
   - User with `associate_ids: ["A", "B"]` → deny rule resolves to `not in [A,B]`
   - User with empty `associate_ids` → deny rule resolves to `not in []` (denies all)
   - User with no deny rules using `${associateIds}` → no change in behavior (backward compatible)
5. Verify existing `${userId}` placeholder still works unchanged (no regression)

## Phase 5: Global Rule Documentation

**File:** `~/.claude/rules/associate-ids-scope-view.md`

Create a global rule describing the `associate_ids` feature so that other Layer 8 projects can consume it when they have user-to-entity list relationships that require ScopeView filtering.

The rule should cover:
- What `associate_ids` is and when to use it (user needs to see data belonging to a list of associated entities)
- How to populate `associate_ids` on users (via Security API PUT or security config JSON)
- How to write deny rules using `${associateIds}` with `not in` syntax
- The empty list behavior (secure default — denies all)
- The ID character constraint (no `,`, `[`, `]` in values)
- Example deny rule patterns for common relationships
- That `${associateIds}` is resolved alongside `${userId}` in ScopeView — both can be used in the same deny rule if needed

## Traceability Matrix

| # | Gap / Action Item | Phase |
|---|-------------------|-------|
| 1 | Add `associate_ids` field to L8User proto | Phase 1 |
| 2 | Regenerate protobuf bindings (check `-i` flag in make-bindings.sh) | Phase 1 |
| 3 | AAA method to retrieve associate IDs by token | Phase 2 |
| 4 | ScopeView `${associateIds}` placeholder resolution | Phase 3 |
| 5 | Build verification | Phase 4 |
| 6 | Placeholder resolution verification (including `not in []` edge case) | Phase 4 |
| 7 | Backward compatibility verification (`${userId}` unchanged) | Phase 4 |
| 8 | Global rule documenting `${associateIds}` for consuming projects | Phase 5 |

## Implementation Concerns

### 1. Verify `not in []` behavior

The plan assumes L8Query's `NotIN` comparator handles an empty list `[]` by matching all rows (denying everything). Looking at `getInStringList` (`In.go:139`), when the input is `[]`, it produces an empty string — `strings.Split("", ",")` returns `[""]`, and the comparator compares each row's value against `""`. This may not panic, but may not behave as expected.

**Action:** Before relying on empty `associate_ids` as the secure default, write a test that verifies `not in []` denies all rows. If it doesn't, add an explicit guard in ScopeView:

```go
if len(associateIds) == 0 && strings.Contains(resolvedValue, "${associateIds}") {
    // No associates — deny all rows by skipping the query and marking all as denied
    return object.New(nil, []interface{}{})
}
```

### 2. ID values must not contain delimiter characters

`getInStringList` splits on `,` between `[` and `]` brackets. If an ID in `associate_ids` contains `,`, `[`, or `]` (e.g., `"ID-[001]"`), the parsed list will be corrupted.

Layer 8 IDs are typically clean alphanumeric strings (e.g., `STU-SEC-001`), so this is unlikely in practice. However, there is no validation enforcing it.

**Action:** Document that `associate_ids` values must not contain `,`, `[`, or `]`. Optionally, add a validation check in the `SetUser` path or in `getAssociateIdsByToken` that logs a warning if any ID contains these characters.

## Supersedes

This plan supersedes the `scopeVia` approach (Option A) from `infra-cross-entity-deny-scope.md`. The `associate_ids` approach is simpler:
- No cross-entity lookup at query time (IDs are pre-resolved on the user record)
- No new proto fields on L8Rule (uses existing `attributes` map with existing `not in` syntax)
- No query engine changes — uses existing L8Query `not in [list]` comparator
- One new placeholder (`${associateIds}`) — minimal ScopeView change
- Fully backward compatible — existing deny rules using only `${userId}` are unaffected

The tradeoff is that `associate_ids` must be kept in sync by consuming projects when relationships change. This is acceptable because the Security API (PUT user) can update them at any time.

## Cross-Project Applicability

The `${associateIds}` placeholder is generic — any Layer 8 project can use it for any "user sees only data belonging to their associated entities" pattern:

| Project | User Role | associate_ids Contains | Deny Rule Pattern |
|---------|-----------|----------------------|-------------------|
| l8learn | Guardian | Children's student IDs | `studentId not in ${associateIds}` |
| l8erp | Sales Rep | Assigned territory IDs | `territoryId not in ${associateIds}` |
| l8erp | Manager | Department IDs | `departmentId not in ${associateIds}` |
| probler | GPU Operator | Owned device IDs | `deviceId not in ${associateIds}` |

Consuming projects are responsible for:
1. Populating `associate_ids` on their users (via Security API or security config JSON)
2. Writing deny rules that reference `${associateIds}` in their security plugin
3. Re-vendoring l8secure after this change is committed
