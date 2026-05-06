# onyx-mob-sync-field

`onyx-mob-sync-field` explores mobile workflows in Go. The repository keeps the core rule set compact, then surrounds it with examples that show how the decisions move.

## Onyx Mob Sync Field Notes

The quickest review path is the verifier first, then the fixtures, then the operations note. That order makes it easy to see whether the code, data, and explanation still agree.

## Why This Exists

The repository exists to keep a technical idea small enough to reason about. The implementation avoids external dependencies where possible, then uses fixtures to make changes easy to review.

## Feature Notes

- Models local state with deterministic scoring and explicit review decisions.
- Uses fixture data to keep sync pressure changes visible in code review.
- Includes extended examples for form constraints, including `surge` and `degraded`.
- Documents offline paths tradeoffs in `docs/operations.md`.
- Runs locally with a single verification command and no external credentials.

## Implementation Notes

The design is intentionally direct: parse or construct a signal, score it, classify it, and verify the expected branch. This makes the repository useful for studying mobile workflows behavior without needing a service or database unless the language project itself is SQL. The Go layout uses small packages and table-oriented tests so the behavior stays easy to follow.

## Code Tour

- `policy`: Go package with the core model
- `cmd`: small command entry point
- `fixtures`: compact golden scenarios
- `examples`: expanded scenario set
- `metadata`: project constants and verification metadata
- `docs`: operations and extension notes
- `scripts`: local verification and audit commands
- `go.mod`: Go module metadata

## Local Setup

Clone the repository, enter the directory, and run the verifier. No database server, cloud account, or token is required.

## Try It

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

This runs the language-level build or test path against the compact fixture set.

## Tests

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/audit.ps1
```

The audit command checks repository structure and README constraints before it delegates to the verifier.

## Example Scenarios

The extended cases are not random smoke tests. `degraded` keeps pressure on the review path, while `surge` shows the model when capacity and weight are strong enough to clear the threshold.

## Boundaries

The examples cover useful edges, not every edge. A larger version would add malformed-input tests, richer reports, and deeper domain parsers.

## Roadmap

- Add malformed input fixtures so the failure path is as visible as the happy path.
- Split the scoring constants into a typed configuration object and validate it before use.
- Add a comparison mode that shows how decisions change when one signal is adjusted.
- Add one more mobile workflows fixture that focuses on a malformed or borderline input.
