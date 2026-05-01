//go:build linux

package osrelease

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestParse_full(t *testing.T) {
	input := "NAME=\"Ubuntu\"\nPRETTY_NAME=\"Ubuntu 22.04 LTS\"\nVERSION_ID=\"22.04\"\nID=ubuntu\nID_LIKE=debian\nHOME_URL=\"https://www.ubuntu.com/\"\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "Ubuntu" {
		t.Errorf("Name: got %q, want %q", got.Name, "Ubuntu")
	}
	if got.PrettyName != "Ubuntu 22.04 LTS" {
		t.Errorf("PrettyName: got %q, want %q", got.PrettyName, "Ubuntu 22.04 LTS")
	}
	if got.VersionID != "22.04" {
		t.Errorf("VersionID: got %q, want %q", got.VersionID, "22.04")
	}
	if got.IDLike != "debian" {
		t.Errorf("IDLike: got %q, want %q", got.IDLike, "debian")
	}
}

func TestParse_commentsAndBlanks(t *testing.T) {
	input := "# comment\n\n# another comment\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "" || got.ID != "" {
		t.Errorf("expected empty result, got %+v", got)
	}
}

func TestParse_unquotedValues(t *testing.T) {
	input := "ID=arch\nVERSION_ID=rolling\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != "arch" {
		t.Errorf("ID: got %q, want %q", got.ID, "arch")
	}
	if got.VersionID != "rolling" {
		t.Errorf("VersionID: got %q, want %q", got.VersionID, "rolling")
	}
}

func TestParse_singleQuotedValues(t *testing.T) {
	input := "NAME='My OS'\nID=myos\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "My OS" {
		t.Errorf("Name: got %q, want %q", got.Name, "My OS")
	}
}

func TestParse_unknownKeysIgnored(t *testing.T) {
	input := "UNKNOWN_KEY=value\nID=myos\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != "myos" {
		t.Errorf("ID: got %q, want %q", got.ID, "myos")
	}
}

func TestReadFrom_ubuntu(t *testing.T) {
	got, err := ReadFrom("testdata/ubuntu.release")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "Ubuntu" {
		t.Errorf("Name: got %q, want %q", got.Name, "Ubuntu")
	}
	if got.ID != "ubuntu" {
		t.Errorf("ID: got %q, want %q", got.ID, "ubuntu")
	}
	if got.VersionID != "22.04" {
		t.Errorf("VersionID: got %q, want %q", got.VersionID, "22.04")
	}
}

func TestReadFrom_minimal(t *testing.T) {
	got, err := ReadFrom("testdata/minimal.release")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != "linux" {
		t.Errorf("ID: got %q, want %q", got.ID, "linux")
	}
	if got.Name != "" {
		t.Errorf("Name: got %q, want empty", got.Name)
	}
}

func TestReadFrom_quoted(t *testing.T) {
	got, err := ReadFrom("testdata/quoted.release")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "Arch Linux" {
		t.Errorf("Name: got %q, want %q", got.Name, "Arch Linux")
	}
	if got.VersionID != "rolling" {
		t.Errorf("VersionID: got %q, want %q", got.VersionID, "rolling")
	}
}

func TestReadFrom_notFound(t *testing.T) {
	_, err := ReadFrom("testdata/nonexistent")
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected ErrNotExist, got %v", err)
	}
}

func TestRead(t *testing.T) {
	_, err := Read()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("unexpected error: %v", err)
	}
}
