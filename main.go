package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

// TR1	CHR1	3	8M7D6M2I2M11D7M
type Transcript struct {
	Name  string
	Chrom string
	Pos   int
	Cigar string
}

// TR1	4
type Query struct {
	Name string
	Pos  int
}

func loadTranscripts(filename string) (map[string]Transcript, error) {
	transcripts := map[string]Transcript{}

	fin, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	reader := csv.NewReader(fin)
	reader.Comma = '\t'
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		pos, err := strconv.Atoi(rec[2])
		if err != nil {
			log.Fatal(err)
		}

		transcripts[rec[0]] = Transcript{rec[0], rec[1], pos, rec[3]}
	}
	return transcripts, nil
}

func loadQueries(filename string) ([]Query, error) {
	queries := []Query{}

	fin, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	reader := csv.NewReader(fin)
	reader.Comma = '\t'
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		pos, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal(err)
		}

		queries = append(queries, Query{rec[0], pos})
	}

	return queries, nil
}

func ConvertCoordinate(query Query, transcripts map[string]Transcript) (string, int, error) {
	fmt.Printf("query: %v, transcripts: %v\n", query, transcripts)
	// check if the Query transcript name is in the transcript map
	_, ok := transcripts[query.Name]
	// If the key does not exist, return error
	if !ok {
		return "NA", 0, errors.New("Missing transcript")
	}

	// get the transcript map entry
	transcript, _ := transcripts[query.Name]

	// cigar_string := transcript.Cigar // "8M7D6M2I2M11D7M"
	re := regexp.MustCompile(`(\d+)([MDISHX])`)
	matches := re.FindAllStringSubmatch(transcript.Cigar, -1) // [ [8M 8 M] , [7D 7 D], [6M 6 M], ... ]
	for _, match := range matches {                           // [8M 8 M]
		// convert the position match value to int
		cigarPos, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		cigarOperation := match[2]
		fmt.Printf("%v, %v\n", cigarPos, cigarOperation)

		// check if cigarOperation is M (Match) or Mismatch(X)
		matchMismatchOps := map[string]bool{"M": true, "X": true}
		_, okMatchMismatch := matchMismatchOps[cigarOperation]

		// check if cigarOperation is Insertion (I) or Softclip (S)
		insertionSoftclipOps := map[string]bool{"I": true, "S": true}
		_, okInsertionSoftclipOps := insertionSoftclipOps[cigarOperation]

		// check if cigarOperation is Deletion (D) or Hardclip (H)
		deletionHardClipOps := map[string]bool{"D": true, "H": true}
		_, okDeletionHardClipOps := deletionHardClipOps[cigarOperation]

		if okMatchMismatch {
			fmt.Printf("its M or X; %v\n", cigarOperation)
		} else if okInsertionSoftclipOps {
			fmt.Printf("its I or S; %v\n", cigarOperation)
		} else if okDeletionHardClipOps {
			fmt.Printf("its D or H; %v\n", cigarOperation)
		}

	}

	return "NA", 0, nil
}

func main() {
	fmt.Println("hello")
	transcripts, err := loadTranscripts("tests/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", transcripts)

	queries, err := loadQueries("tests/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, query := range queries {
		trName, trPos, err := ConvertCoordinate(query, transcripts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v %v\n", trName, trPos)
	}
}
