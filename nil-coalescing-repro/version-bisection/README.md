# Version bisection (functional proof)

Each file is a self-contained test that exercises the *actual* `util/template` code of a
specific Argo release, confirming where `item.optionalKey ?? 'fallback'` (missing key, strict
path) starts failing.

| File | Checkout | Result |
|------|----------|--------|
| `v36_expression_test.go`   | any `v3.6.x` tag      | all guarded forms resolve to `fallback` (PASS) |
| `v3711_expression_test.go` | `v3.7.11` tag         | present resolves; missing-key forms rejected with `item.optionalKey is missing` |

How to run (example for v3.6):

```
git worktree add --detach /tmp/wt-v3.6 v3.6.19
cp version-bisection/v36_expression_test.go /tmp/wt-v3.6/util/template/
cd /tmp/wt-v3.6 && GOFLAGS=-mod=mod go test ./util/template/ -run Test_v36_NilCoalescing_StrictPath -v
```

The regression was introduced by #15442 (`d84169f`), first shipped in **v3.7.11**; absent from all
`v3.6.x` releases.
