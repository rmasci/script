package script_test

import (
	"strings"
	"testing"

	"github.com/rmasci/script"
)

func TestZFile(t *testing.T) {
	lineCount, err := script.ZFile("testdata/releases.json.gz").CountLines()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if lineCount != 312 {
		t.Errorf("expected 312 lines, got %d", lineCount)
	}
}

func TestCat(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name     string
		filePath string
		want     string
	}{
		{
			name:     "Test with plain text file",
			filePath: "testdata/test.txt",
			want:     "Hello, world",
		},
		{
			name:     "Test with gzipped text file",
			filePath: "testdata/test.txt.gz",
			want:     "Hello, world",
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Use the Cat function to read a test file
			p := script.Cat(tc.filePath)

			// Get the contents of the file as a string
			got, err := p.String()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check if the output contains the expected line
			if !strings.Contains(got, tc.want) {
				t.Errorf("Cat() = %q, want it to contain %q", got, tc.want)
			}
		})
	}
}
