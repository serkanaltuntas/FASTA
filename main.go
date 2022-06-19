package main

import (
	"bufio"
	"fmt"
	"os"
)

// create a type for fasta records	and define its fields
type fasta struct {
	header   string
	sequence string
}

// read fasta file and save its content as fasta records	in a slice
func readFasta(filename string) ([]fasta, error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create a slice to hold fasta records
	var records []fasta
	// display each record in the slice	as it is read
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '>' {
			// create a new fasta record
			record := fasta{header: line[1:]}
			// add it to the slice
			records = append(records, record)
		} else {
			// add the sequence to the last record in the slice
			records[len(records)-1].sequence += line
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil

}

// create main fucntion and call readFasta function
func main() {
	// get the filename from the command line	and store it in a variable if file is not evailable warn the user	and exit
	if len(os.Args) != 2 {
		println("Usage:", os.Args[0], "filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	records, err := readFasta(filename)
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		fmt.Println(record.header)
		fmt.Println(record.sequence)

		// if sequence include characters other than A,a,T,t,C,c,G,g, warn the user and pass to the next record
		if !(record.sequence[0] == 'A' || record.sequence[0] == 'a' || record.sequence[0] == 'T' || record.sequence[0] == 't' || record.sequence[0] == 'C' || record.sequence[0] == 'c' || record.sequence[0] == 'G' || record.sequence[0] == 'g') {
			fmt.Println("Warning: sequence contains characters other than A,a,T,t,C,c,G,g, skipping record for GC content calculation.")
			continue
		}

		// calculate GC content and print it
		gc, at, nt := record.GCcontent()
		fmt.Println("-----Summary-----")
		fmt.Println("All:", nt)
		fmt.Printf("GC content: %.2f\n", gc)
		fmt.Printf("AT content: %.2f\n", at)
	}
}

// create a function that calculates number of G and C contents separately in a fasta record
func (f fasta) GCcontent() (float64, float64, int) {
	var nt, gcCount, atCount int
	for _, nt := range f.sequence {
		switch nt {
		case 'G', 'g', 'C', 'c':
			gcCount++
		case 'A', 'a', 'T', 't':
			atCount++
		}
	}
	nt = gcCount + atCount
	return float64(gcCount) / float64(nt), float64(atCount) / float64(nt), nt
}
