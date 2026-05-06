# Review Journal

The review surface for `onyx-mob-sync-field` is deliberately narrow: one fixture, one scoring rule, and one local check.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 194, lane `ship`
- `stress`: `sync drift`, score 209, lane `ship`
- `edge`: `local state`, score 209, lane `ship`
- `recovery`: `conflict cost`, score 233, lane `ship`
- `stale`: `form pressure`, score 209, lane `ship`

## Note

A future change should add new cases before it changes the scoring rule.
