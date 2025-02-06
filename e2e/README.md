# End-to-end tests

These tests run CLI commands like a user would,
built on top of the [`go-internal/testscript`](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript)
package.

To run these tests, use `go test ./... -tags=e2e`.
