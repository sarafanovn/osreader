package osrelease

import (
	"bufio"
	"io"
	"strings"
)

func parse(r io.Reader) (*OSRelease, error) {
	result := &OSRelease{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		if len(value) >= 2 {
			if (value[0] == '"' && value[len(value)-1] == '"') ||
				(value[0] == '\'' && value[len(value)-1] == '\'') {
				value = value[1 : len(value)-1]
			}
		}
		switch key {
		case "NAME":
			result.Name = value
		case "PRETTY_NAME":
			result.PrettyName = value
		case "VERSION":
			result.Version = value
		case "VERSION_ID":
			result.VersionID = value
		case "ID":
			result.ID = value
		case "ID_LIKE":
			result.IDLike = value
		case "HOME_URL":
			result.HomeURL = value
		case "SUPPORT_URL":
			result.SupportURL = value
		case "BUG_REPORT_URL":
			result.BugReportURL = value
		case "BUILD_ID":
			result.BuildID = value
		case "VARIANT":
			result.Variant = value
		case "VARIANT_ID":
			result.VariantID = value
		}
	}
	return result, scanner.Err()
}
