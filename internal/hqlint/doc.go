// Package hqlint holds the structural lint for the soul-hq headquarters
// layout.
//
// It carries no runtime code: the checks live in hqlint_test.go and ride the
// standard quality gate via `go test ./...`, so `make test`, `make check`, and
// CI enforce the hq invariants automatically — the same way the five
// per-project hqs enforced theirs before they merged into this repository.
package hqlint
