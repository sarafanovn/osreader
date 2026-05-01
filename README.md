# osreader

A Go library for reading Linux system configuration files into typed structs.

**Platform:** Linux only (Windows support planned)

## Installation

```bash
go get github.com/sarafanovn/osreader
```

## Packages

### osrelease

Reads `/etc/os-release` (falls back to `/usr/lib/os-release`) and parses it into an `OSRelease` struct.

```go
import "github.com/sarafanovn/osreader/osrelease"

// Read from standard system path
info, err := osrelease.Read()
if err != nil {
    log.Fatal(err)
}
fmt.Println(info.Name, info.VersionID)

// Read from a custom path
info, err = osrelease.ReadFrom("/path/to/os-release")
```

**`OSRelease` fields:** `Name`, `PrettyName`, `Version`, `VersionID`, `ID`, `IDLike`, `HomeURL`, `SupportURL`, `BugReportURL`, `BuildID`, `Variant`, `VariantID`

### limits

Reads `/etc/security/limits.conf` and parses each rule into a `LimitRule` struct.

```go
import "github.com/sarafanovn/osreader/limits"

// Read from standard system path
rules, err := limits.Read()
if err != nil {
    log.Fatal(err)
}
for _, r := range rules {
    fmt.Printf("%s %s %s %s\n", r.User, r.Type, r.Parameter, r.Value)
}

// Read from a custom path
rules, err = limits.ReadFrom("/path/to/limits.conf")
```

**`LimitRule` fields:** `User`, `Type`, `Parameter`, `Value`

## Error Handling

Both packages wrap underlying OS errors, so you can use standard predicates:

```go
if errors.Is(err, os.ErrNotExist) {
    // file not found
}
if errors.Is(err, os.ErrPermission) {
    // no read permission
}
```

## License

[MIT](Licence.md)
