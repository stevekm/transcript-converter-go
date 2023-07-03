package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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

func ConvertCoordinate(query Query, transcripts map[string]Transcript) (string, int) {
	return "NA", 0
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
		trName, trPos := ConvertCoordinate(query, transcripts)
		fmt.Printf("%+v %v\n", trName, trPos)
	}
}
