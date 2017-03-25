package common

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// ParseLoansCSVFile - reads our loans input report and converts it into a multidimentional array
func ParseLoansCSVFile(filename string) [][]string {
	LoansInputReport, err := os.Open(filename)

	// release the file once done
	defer LoansInputReport.Close()

	// If something went wrong
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	reader := csv.NewReader(LoansInputReport)

	reader.Comma = ','
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return rows
}

// CheckIfStringInSplice - checks if a particular string exits in a slice of strings
func CheckIfStringInSplice(needle string, haystack []string) bool {

	/**

		GO does not have an implementation for..in like Python since its a low level language,
		I have therefore added my own implementation of this behaviour for :
		if string in otherstring:
		    true clause
	**/

	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

//StripQuotes - strips the single quotes from strings used for our CSV imports
func StripQuotes(s string) string {
	return strings.Replace(s, "'", "", -1)

}
