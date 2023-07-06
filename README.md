# Transcript Converter

Program to translate transcript coordinates to genomic coordinates.

Reads in a set of transcripts with CIGAR strings, and then converts a set of given coordinates supplied in the secondary query file.

```
$ go run main.go "tests/input1.txt" "tests/input2.txt
TR1	4	CHR1	7
TR2	0	CHR2	10
TR1	13	CHR1	23
TR2	10	CHR2	20
```

A Go implementation to a common Python bioinformatics practice problem.

# References

- https://github.com/varshini712/transcript_converter

- https://samtools.github.io/hts-specs/SAMv1.pdf

- https://www.drive5.com/usearch/manual/cigar.html

- https://github.com/monarin/bioinf
