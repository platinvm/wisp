package wisp_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/platinvm/wisp"
)

func TestParseAndVisit_Examples(t *testing.T) {
	root := "../../test/resources/examples"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".wisp") {
			return nil
		}

		content, readErr := os.ReadFile(path)
		if readErr != nil {
			t.Errorf("Failed to read %s: %v", path, readErr)
			return nil
		}

		visited, parseErr := wisp.ParseAndVisit(string(content))
		if parseErr != nil {
			t.Errorf("ParseAndVisit failed for %s: %v", path, parseErr)
			return nil
		}

		if len(visited) == 0 {
			t.Errorf("No rules visited for %s", path)
		}

		t.Logf("successfully parsed %s\n", path)

		return nil
	})

	if err != nil {
		t.Fatalf("Error walking example files: %v", err)
	}
}
