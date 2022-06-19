package main

import (
	"bufio"
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
		println(record.header)
		println(record.sequence)
	}
}
