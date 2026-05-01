//go:build linux

package limits

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestParse_valid(t *testing.T) {
	input := "# comment\n*\tsoft\tnofile\t1024\n*\thard\tnofile\t65536\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len: got %d, want 2", len(got))
	}
	if got[0].User != "*" || got[0].Type != "soft" || got[0].Parameter != "nofile" || got[0].Value != "1024" {
		t.Errorf("rule[0]: got %+v", got[0])
	}
	if got[1].User != "*" || got[1].Type != "hard" || got[1].Parameter != "nofile" || got[1].Value != "65536" {
		t.Errorf("rule[1]: got %+v", got[1])
	}
}

func TestParse_commentsAndBlanks(t *testing.T) {
	input := "# comment\n\n# another comment\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("len: got %d, want 0", len(got))
	}
}

func TestParse_malformed(t *testing.T) {
	input := "*\tsoft\tnofile\t1024\n*\tsoft\tnofile\n@users\tsoft\tcpu\t300\n"
	got, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len: got %d, want 2", len(got))
	}
}

func TestReadFrom_valid(t *testing.T) {
	got, err := ReadFrom("testdata/valid.conf")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 4 {
		t.Fatalf("len: got %d, want 4", len(got))
	}
	if got[0].User != "*" || got[0].Type != "soft" || got[0].Parameter != "nofile" || got[0].Value != "1024" {
		t.Errorf("rule[0]: got %+v", got[0])
	}
	if got[1].User != "*" || got[1].Type != "hard" || got[1].Parameter != "nofile" || got[1].Value != "65536" {
		t.Errorf("rule[1]: got %+v", got[1])
	}
	if got[2].User != "root" || got[2].Type != "hard" || got[2].Parameter != "nproc" || got[2].Value != "unlimited" {
		t.Errorf("rule[2]: got %+v", got[2])
	}
	if got[3].User != "@users" || got[3].Type != "soft" || got[3].Parameter != "cpu" || got[3].Value != "300" {
		t.Errorf("rule[3]: got %+v", got[3])
	}
}

func TestReadFrom_comments(t *testing.T) {
	got, err := ReadFrom("testdata/comments.conf")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("len: got %d, want 0", len(got))
	}
}

func TestReadFrom_malformed(t *testing.T) {
	got, err := ReadFrom("testdata/malformed.conf")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len: got %d, want 2", len(got))
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
