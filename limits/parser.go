//go:build linux

package limits

import (
	"bufio"
	"io"
	"strings"
)

func parse(r io.Reader) ([]LimitRule, error) {
	var rules []LimitRule
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		rules = append(rules, LimitRule{
			User:      fields[0],
			Type:      fields[1],
			Parameter: fields[2],
			Value:     fields[3],
		})
	}
	return rules, scanner.Err()
}
