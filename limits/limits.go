// Package limits reads and parses /etc/security/limits.conf,
// which defines resource limits for users and groups on Linux systems.
package limits

import (
	"fmt"
	"os"
)

const defaultPath = "/etc/security/limits.conf"

// LimitRule represents a single rule from limits.conf.
// Fields correspond to the four columns: domain, type, item, and value.
// Lines with fewer than four fields are silently skipped.
type LimitRule struct {
	User      string
	Type      string
	Parameter string
	Value     string
}

// Read parses /etc/security/limits.conf and returns all valid rules.
// Comment lines (starting with #) and blank lines are ignored.
func Read() ([]LimitRule, error) {
	return ReadFrom(defaultPath)
}

// ReadFrom parses a limits.conf file at the given path.
func ReadFrom(path string) ([]LimitRule, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("limits: read %s: %w", path, err)
	}
	defer f.Close()
	return parse(f)
}
