// Package osrelease reads and parses the os-release file, which describes
// the Linux distribution. It tries /etc/os-release first, falling back to
// /usr/lib/os-release as defined in the os-release(5) specification.
package osrelease

import (
	"errors"
	"fmt"
	"os"
)

const etcPath = "/etc/os-release"
const usrPath = "/usr/lib/os-release"

// OSRelease holds the fields defined in the os-release(5) specification.
// Unknown keys in the file are silently ignored.
// Values with single or double quotes have the surrounding quotes stripped.
type OSRelease struct {
	Name         string
	PrettyName   string
	Version      string
	VersionID    string
	ID           string
	IDLike       string
	HomeURL      string
	SupportURL   string
	BugReportURL string
	BuildID      string
	Variant      string
	VariantID    string
}

// Read parses the os-release file from its standard system location.
// It tries /etc/os-release first; if that file cannot be accessed, it falls back
// to /usr/lib/os-release. Returns os.ErrNotExist if neither file can be accessed.
func Read() (*OSRelease, error) {
	result, err := ReadFrom(etcPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return ReadFrom(usrPath)
	}
	return result, err
}

// ReadFrom parses an os-release file at the given path.
func ReadFrom(path string) (*OSRelease, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("osrelease: read %s: %w", path, err)
	}
	defer f.Close()
	return parse(f)
}
