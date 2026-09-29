---
name: test-flake-fix-and-skip-workflow
description: Workflow command scaffold for test-flake-fix-and-skip-workflow in istio.
allowed_tools: ["Bash", "Read", "Write", "Grep", "Glob"]
---

# /test-flake-fix-and-skip-workflow

Use this workflow when working on **test-flake-fix-and-skip-workflow** in `istio`.

## Goal

Addresses flaky or failing tests by fixing race conditions, increasing timeouts, skipping tests under certain conditions, and refactoring test logic. This often involves modifying integration and unit test files to improve reliability, especially in CI environments or under specific Kubernetes versions.

## Common Files

- `tests/integration/ambient/cni/main_test.go`
- `tests/integration/ambient/cnirepair/main_test.go`
- `tests/integration/ambient/cniupgrade/main_test.go`
- `tests/integration/ambient/pqc/main_test.go`
- `tests/integration/helm/install_test.go`
- `tests/integration/pilot/common/traffic.go`

## Suggested Sequence

1. Understand the current state and failure mode before editing.
2. Make the smallest coherent change that satisfies the workflow goal.
3. Run the most relevant verification for touched files.
4. Summarize what changed and what still needs review.

## Typical Commit Signals

- Identify the flaky or failing test and its root cause (race, timeout, environment mismatch).
- Modify the test code to fix the race (e.g., use EventuallyEqual, add helper functions, increase retry timeouts).
- Add logic to skip the test if prerequisites are not met (e.g., check for Gateway API CRD support, Kubernetes version).
- Refactor or extract helpers for repeated skip logic.
- Optionally, revert or adjust previous fixes based on reviewer feedback.

## Notes

- Treat this as a scaffold, not a hard-coded script.
- Update the command if the workflow evolves materially.