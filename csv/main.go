package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// filterFile copies from one CSV file with price data to another, skipping entries where the price
// is zero. It returns an error if an I/O error occurs (e.g.  file not found) or if the input
// doesn't have the right format.
//
// See README.md for the CSV file format.
func filterFile(inFilename, outFilename string) error {
	incsv, err := os.Open(inFilename)
	if err != nil {
		return fmt.Errorf("issue opening inFilename: %s", err)
	}
	defer incsv.Close()
	outcsv, err := os.Create(outFilename)
	if err != nil {
		return fmt.Errorf("issue creating outFilename: %s", err)
	}
	defer outcsv.Close()
	f := filter(incsv, outcsv)
	if f != nil {
		return f
	}
	return nil
}

// filter copies from one CSV file with price data to another, skipping entries where the price is
// zero. It returns an error if an I/O error occurs (e.g.  file not found) or if the input doesn't
// have the right format.
//
// See README.md for the CSV file format.
func filter(in io.Reader, out io.Writer) error {
	reader := csv.NewReader(in)
	writer := csv.NewWriter(out)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error with reader: %s", err)
		}
		v, err := skipRecord(record)
		if err != nil {
			return fmt.Errorf("error skipping record: %s", err)
		}
		if v == true {
			continue
		}
		err = writer.Write(record)
		if err != nil {
			return fmt.Errorf("error writing file: %s", err)
		}
		writer.Flush()
	}
	return nil
}

// skipRecord takes one line of price data from a CSV file and returns true if the price is zero. It
// returns an error if the line doesn't have the right format.
//
func skipRecord(record []string) (bool, error) {

	if (record[1] == "0") == true {
		return true, nil
	}

	_, err := strconv.Atoi(record[1])
	if err != nil {
		return false, err
	}

	// TODO implement this function; make sure to return an error if the price is not an integer
	return false, nil
}
func main() {
	// parse command-line arguments
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s input.csv output.csv\n", os.Args[0])
		os.Exit(2)
	}
	inFilename, outFilename := args[0], args[1]

	// filter file
	err := filterFile(inFilename, outFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
