//go:build linux

package osrelease

import (
	"errors"
	"fmt"
	"os"
)

const etcPath = "/etc/os-release"
const usrPath = "/usr/lib/os-release"

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

func Read() (*OSRelease, error) {
	result, err := ReadFrom(etcPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return ReadFrom(usrPath)
	}
	return result, err
}

func ReadFrom(path string) (*OSRelease, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("osrelease: read %s: %w", path, err)
	}
	defer f.Close()
	return parse(f)
}
