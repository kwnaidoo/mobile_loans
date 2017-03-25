package common

import (
	"reflect"
	"testing"
)

// TestStripQuotes - will test the common.StripQuotes function to see if
// it accurately removes quotes from the sampel string.

func TestStripQuotes(t *testing.T) {
	testString := `'this is a test string'`
	result := StripQuotes(testString)
	if result != `this is a test string` {
		t.Error(`Failed to strip quotes from string: \n`, testString)
	}

}

// TestParseLoansCSVFile - tests to see if we can successfully open and parse a CSV file
func TestParseLoansCSVFile(t *testing.T) {
	result := ParseLoansCSVFile("fixtures/sample_loans.csv")

	if result == nil {
		t.Error(`Failed to load Loans.csv file`)
	}
	if reflect.TypeOf(result).Kind() != reflect.Slice {
		t.Error(`Result not a valid CSV file`)
	}
}
