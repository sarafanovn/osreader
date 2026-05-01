//go:build linux

package limits

import (
	"fmt"
	"os"
)

const defaultPath = "/etc/security/limits.conf"

type LimitRule struct {
	User      string
	Type      string
	Parameter string
	Value     string
}

func Read() ([]LimitRule, error) {
	return ReadFrom(defaultPath)
}

func ReadFrom(path string) ([]LimitRule, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("limits: read %s: %w", path, err)
	}
	defer f.Close()
	return parse(f)
}
