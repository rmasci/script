package script_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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
func TestFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		inDelim  string
		outDelim string
		fields   []int
		input    string
		want     string
	}{
		{
			name:     "Test 1",
			inDelim:  " ",
			outDelim: ",",
			fields:   []int{1, 3},
			input:    "",
			want:     "",
		}, {
			name:     "Test 2",
			inDelim:  " ",
			outDelim: ",",
			fields:   []int{1, 3},
			input:    " ",
			want:     " ",
		}, {
			name:     "Test 3",
			inDelim:  " ",
			outDelim: ",",
			fields:   []int{1, 3},
			input:    "field1 field2 field3 field4",
			want:     "field1,field3",
		},
		{
			name:     "Test 4",
			inDelim:  "|",
			outDelim: ",",
			fields:   []int{1, 4},
			input:    "field1|field2|field3|field4",
			want:     "field1,field4",
		},
		{
			name:     "Test 5",
			inDelim:  ";",
			outDelim: ",",
			fields:   []int{1, 4, 3, 2},
			input:    "field1;     field2;field3; field4",
			want:     "field1,field4,field3,field2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := script.NewPipe().Echo(tt.input)
			got, _ := p.Fields(tt.inDelim, tt.outDelim, tt.fields...).String()

			if diff := cmp.Diff(strings.TrimSpace(tt.want), strings.TrimSpace(got)); diff != "" {
				t.Errorf("Fields() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
