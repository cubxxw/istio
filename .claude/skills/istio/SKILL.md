```markdown
# istio Development Patterns

> Auto-generated skill from repository analysis

## Overview
This skill teaches best practices and common workflows for contributing to the [istio](https://github.com/istio/istio) codebase, a large-scale service mesh project written in Go. It covers coding conventions, test maintenance, and the standard approach to handling flaky tests and CI reliability. The patterns here are distilled from real repository history and should help new contributors quickly align with project standards.

## Coding Conventions

### File Naming
- Use `snake_case` for all file names.
  - **Example:** `main_test.go`, `secretcontroller_test.go`

### Import Style
- Use **relative imports** within the module.
  - **Example:**
    ```go
    import (
        "istio.io/istio/pkg/kube"
        "istio.io/istio/tests/integration/ambient/cni"
    )
    ```

### Export Style
- Use **named exports** for functions, types, and variables that need to be accessed outside their package.
  - **Example:**
    ```go
    // Exported function
    func RunIntegrationTests() { ... }

    // Exported type
    type TestHelper struct { ... }
    ```

- Unexported (private) identifiers start with a lowercase letter.
  - **Example:**
    ```go
    func runHelper() { ... }
    ```

### Commit Patterns
- Commit messages are freeform, sometimes prefixed with `tests`.
- Average commit message length: ~40 characters.

## Workflows

### Test Flake Fix and Skip Workflow
**Trigger:** When a test is flaky, fails intermittently, or needs to be skipped under certain environments (e.g., older Kubernetes versions, missing CRDs).  
**Command:** `/fix-flaky-test`

1. **Identify the Flaky Test**
   - Locate the test that fails intermittently or under specific CI conditions.
   - Example: A test fails only on Kubernetes <1.26 or when a CRD is missing.

2. **Diagnose the Root Cause**
   - Check for race conditions, insufficient timeouts, or environment mismatches.
   - Example: Test fails due to race in resource creation.

3. **Modify the Test Code**
   - Fix race conditions by using retry logic or helper functions.
   - Increase timeouts if the test is timing out.
   - **Example:**
     ```go
     // Use EventuallyEqual to wait for the condition
     util.EventuallyEqual(t, expected, actual, retryTimeout)
     ```

4. **Add Skip Logic**
   - Skip the test if prerequisites are not met (e.g., missing CRD, unsupported K8s version).
   - **Example:**
     ```go
     if !HasGatewayAPI() {
         t.Skip("Gateway API CRD not installed")
     }
     ```

5. **Refactor or Extract Helpers**
   - If skip logic is repeated, extract it into a helper function for reuse.
   - **Example:**
     ```go
     func SkipIfNoGatewayAPI(t *testing.T) {
         if !HasGatewayAPI() {
             t.Skip("Gateway API CRD not installed")
         }
     }
     ```

6. **Review and Adjust**
   - Optionally, revert or adjust previous fixes based on reviewer feedback.

7. **Commit the Changes**
   - Commit with a clear message describing the fix or skip logic.

**Files Commonly Involved:**
- `tests/integration/ambient/cni/main_test.go`
- `tests/integration/ambient/cnirepair/main_test.go`
- `tests/integration/ambient/cniupgrade/main_test.go`
- `tests/integration/ambient/pqc/main_test.go`
- `tests/integration/helm/install_test.go`
- `tests/integration/pilot/common/traffic.go`
- `tests/integration/pilot/gateway_httproute_grpcroute_test.go`
- `pkg/kube/multicluster/secretcontroller_test.go`

## Testing Patterns

- **Test File Naming:** All test files end with `_test.go`.
- **Test Framework:** No explicit framework detected, but Go's standard `testing` package is used.
- **Test Structure:** 
  - Functions start with `Test` and take `*testing.T` as a parameter.
  - Skipping tests is done via `t.Skip()`.
  - Retrying or waiting for conditions is handled with helper functions like `EventuallyEqual`.

**Example Test:**
```go
func TestFeatureX(t *testing.T) {
    if !HasPrerequisite() {
        t.Skip("Prerequisite not met")
    }
    // Test logic here
}
```

## Commands

| Command          | Purpose                                                |
|------------------|--------------------------------------------------------|
| /fix-flaky-test  | Initiate the workflow for fixing or skipping flaky tests |

```