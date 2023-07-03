package main

import (
	// "fmt"
	"github.com/google/go-cmp/cmp"
	// "log"
	// "os"
	// "path/filepath"
	"testing"
)

func TestFindAllFiles(t *testing.T) {
	tests := map[string]struct {
		query       Query
		transcripts map[string]Transcript
		wantPos     int
		wantName    string
	}{
		"first": {
			query: Query{Name: "TR1", Pos: 4},
			transcripts: map[string]Transcript{
				"TR1": Transcript{Name: "TR1", Chrom: "CHR1", Pos: 3, Cigar: "8M7D6M2I2M11D7M"},
			},
			wantPos:  7,
			wantName: "TR1",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			trName, trPos := ConvertCoordinate(tc.query, tc.transcripts)
			// sizes, err := SubDirSizes(tempdir)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			got := trPos
			// fmt.Printf("sizes: %v\n", sizes)
			if diff := cmp.Diff(tc.wantPos, got); diff != "" {
				t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tc.wantName, trName); diff != "" {
				t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
