# onyx-mob-sync-field

`onyx-mob-sync-field` is a compact Go repository for mobile workflows, centered on this goal: Create a Go reference implementation for sync workflows, centered on graph analysis, node-edge fixtures, and cycle and reachability reports.

## Project Rationale

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Onyx Mob Sync Field Review Notes

For a quick review, compare `conflict cost` with `form pressure` before reading the middle cases.

## Feature Set

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/onyx-mob-sync-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `conflict cost` and `form pressure`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Architecture

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `form pressure`, `sync drift`, `local state`, and `conflict cost`.

The Go addition stays small enough to inspect in one sitting.

## Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Test Command

That command is also the regression path. It verifies the domain cases and catches mismatches between the CSV, metadata, and code.

## Next Improvements

The repository is intentionally scoped to local checks. I would expand it by adding adversarial fixtures before adding features.
