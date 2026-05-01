# Contributing

Contributions are welcome. By contributing you agree to license y
our work under the project's [MIT licence](Licence.md).

## Prerequisites

- Go 1.18 or later
- Linux (required to run `limits` and `osrelease` tests against real system files)

Cross-compilation from macOS or Windows works for building, but tests that call `Read()` require a Linux host.

## Running tests

```bash
# All tests (Linux only — both packages have //go:build linux)
go test ./...

# Single package
go test ./osrelease/
go test ./limits/

# With verbose output
go test ./... -v
```

## Submitting a pull request

1. Fork and clone the repository
2. Create a branch: `git checkout -b my-change`
3. Make your changes and add tests
4. Verify tests pass: `go test ./...`
5. Verify the linter is clean: `go vet ./...`
6. Open a pull request

## Guidelines

- Every new reader must expose `Read() ` and `ReadFrom(path string)` — consistent with existing packages
- All code must be covered by tests; use `testdata/` fixtures so tests run without a real system file
- New Linux-only packages must carry `//go:build linux` on every `.go` file including tests
- Follow standard Go style — `gofmt`, no unnecessary comments, no exported symbols without a purpose
- Keep changes focused; unrelated improvements belong in a separate PR

## Roadmap

The goal is a cross-platform library covering the most common system config files. Planned next steps:

- Windows support (registry-backed equivalents where applicable)
- Additional Linux readers: `/etc/hostname`, `/etc/fstab`, `/proc/meminfo`
