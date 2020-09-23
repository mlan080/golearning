package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var in string = `laptop,1200
smartphone,600
free sample,0
monitor,200
`
var invalid string = `laptop,1200
smartphone,600
nokia phone,not available
monitor,200
`
var want string = `laptop,1200
smartphone,600
monitor,200
`

func TestFilterFile(t *testing.T) {
	// prepare temporary file with valid input
	inFile, err := ioutil.TempFile("", "input.csv")
	if err != nil {
		t.Fatalf("error creating tempfile: %v", err)
	}
	inFilename := inFile.Name()
	defer os.Remove(inFilename)
	if _, err := inFile.Write([]byte(in)); err != nil {
		t.Fatalf("error writing to tempfile: %v", err)
	}
	inFile.Close()

	// prepare temporary file with invalid input
	invalidFile, err := ioutil.TempFile("", "invalid.csv")
	if err != nil {
		t.Fatalf("error creating tempfile: %v", err)
	}
	invalidFilename := invalidFile.Name()
	defer os.Remove(invalidFilename)
	if _, err := invalidFile.Write([]byte(invalid)); err != nil {
		t.Fatalf("error writing to tempfile: %v", err)
	}
	inFile.Close()

	// prepare temporary file for output
	outFile, err := ioutil.TempFile("", "output.csv")
	if err != nil {
		t.Fatalf("error creating tempfile: %v", err)
	}
	outFilename := outFile.Name()
	defer os.Remove(outFilename)
	outFile.Close()

	// run filterFile with valid input
	err = filterFile(inFilename, outFilename)
	if err != nil {
		t.Fatalf("filterFile returned error: %v", err)
	}
	gotBytes, err := ioutil.ReadFile(outFilename)
	if err != nil {
		t.Fatalf("error reading tempfile: %v", err)
	}
	got := string(gotBytes)
	if got != want {
		t.Errorf("filter returned %#v, want %#v", got, want)
	}

	// run filterFile with invalid input
	err = filterFile(invalidFilename, outFilename)
	if err == nil {
		t.Error("filter didn't return error for invalid input")
	}
}

// invalidWriter is an implementation of io.Writer that always returns an error.
type invalidWriter struct{}

func (w *invalidWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("boom!")
}

func TestFilter(t *testing.T) {
	// call filter with valid parameters
	buf := new(bytes.Buffer)
	err := filter(strings.NewReader(in), buf)
	if err != nil {
		t.Fatalf("filter returned error: %v", err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("filter returned %#v, want %#v", got, want)
	}

	// call filter with invalid io.Writer
	err = filter(strings.NewReader(in), new(invalidWriter))
	if err == nil {
		t.Error("filter didn't return error for invalid writer")
	}
}

func TestSkipRecord(t *testing.T) {
	cases := []struct {
		record        []string
		want, wantErr bool
	}{
		{
			record: []string{"laptop", "1200"},
		},
		{
			record: []string{"free sample", "0"},
			want:   true,
		},
		{
			record:  []string{"nokia phone", "not available"},
			wantErr: true,
		},
	}
	for _, c := range cases {
		got, err := skipRecord(c.record)
		gotErr := err != nil
		if got != c.want {
			t.Errorf("skipRecord(%#v) returned %v, want %v", c.record, got, c.want)
		}
		if gotErr && !c.wantErr {
			t.Errorf("skipRecord(%#v) returned error", c.record)
		}
		if !gotErr && c.wantErr {
			t.Errorf("skipRecord(%#v) didn't return error", c.record)
		}
	}
}
