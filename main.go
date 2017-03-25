package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"mobile_loans/packages/common"
	"mobile_loans/packages/models"
	"sort"
	"strconv"
	"strings"
)

/**
   I have followed goes naming convention strickly throughout this application,
   see below a summary of most common rules:
	1. All module level variables must start with small letters hence why all these have
	   lowercase first letter names.
	2. Use camelCase rather then underscored_names
	3. Use var when outside a method or declaring maps
	4. Document all Public methods in following format // MethodName - description
	   I have further extended this and provided documentation for private variables and methods
	   as well.
**/

// monthHeadings - fixed list of months mapping to months in imported CSV
var monthHeadings = []string{
	`Jan`,
	`Feb`,
	`Mar`,
	`Apr`,
	`May`,
	`Jun`,
	`Jul`,
	`Aug`,
	`Sep`,
	`Oct`,
	`Nov`,
	`Dec`,
}

// allNetworkTotals - counter keeps track of total revenue across all networks
var allNetworkTotals = big.NewFloat(0.00)

//productHeadings  - stores product names to create the CSV file header and allow for sorting
var productHeadings []string

//networkHeadings  - stores network names so that they can be sorted by name
var networkHeadings []string

// networks - will store every single network imported from the CSV file
var networks = make(map[string]*models.Network)

//currencySymbol - currency symbol to use when displaying money values
var currencySymbol = "R"

// performCalculations - Takes in a CSV row in the form off a slice/array and
// stores the relevant figures in our network struct

func performCalculations(row []string, index int) {

	// Each row must contain 5 columns
	if len(row) == 5 {

		// capture each column from the row and clean off quotes
		colNetwork := common.StripQuotes(row[1])
		colMonth := row[2][4:7]
		colProduct := common.StripQuotes(row[3])
		colAmount := common.StripQuotes(row[4])

		// if the network doesn't exist create it and add to both the networks list
		// and network headings list.
		_, doesNetworkExist := networks[colNetwork]
		if doesNetworkExist == false {
			networks[colNetwork] = models.NewNetwork(colNetwork)
			networkHeadings = append(networkHeadings, colNetwork)
		}

		// currentNetwork - stores a pointer to the current network in our networks list
		currentNetwork := networks[colNetwork]

		// add amount to relevant month and perform relevant calculations
		networks[colNetwork].AddMonthPrice(colMonth, colAmount)

		/**
		    totalForCurrentProduct - contains the colAmount parsed as a bigfloat
		    the current product is also added to our network and relevant
		    calculations performed.
		**/
		totalForCurrentProduct := currentNetwork.AddProduct(colProduct, colAmount)

		// allNetworkTotals - increment total revenue for all networks by adding
		// this products total to the counter.
		allNetworkTotals = big.NewFloat(0.00).Add(allNetworkTotals, totalForCurrentProduct)

		// if product not in productHeadings add it
		if common.CheckIfStringInSplice(colProduct, productHeadings) == false {
			productHeadings = append(productHeadings, colProduct)
		}

	} else {
		fmt.Println(`Warning: row `, index+1, ` contains invalid data`)
	}
}

func generateReport() {

	// Sort headings alphabetically
	sort.StringsAreSorted(networkHeadings)
	sort.StringsAreSorted(productHeadings)

	// output - stores the entire outputted CSV's data before writing to file
	// build headings for the CSV
	output := `Network, ` + strings.Join(monthHeadings[:], " (Revenue) ,") + ` (Revenue) ,` + strings.Join(productHeadings[:], " (Revenue) ,")
	output += ` (Revenue) ,Total Loan Products,` + `Total Revenue From Network, ` + `Total Revenue From ALL Networks`
	output = `'` + strings.Replace(output, `,`, `','`, -1) + `'` + "\n"

	// Loop through all networks in networkHeadings - this should contain every single network imported
	for _, network := range networkHeadings {
		currentNetwork := networks[network]
		output = output + `'` + network + `'`

		// Now loop through months alphabetically - if a value is set add it to the report
		for _, m := range monthHeadings {
			output += "," + `'` + currencySymbol + currentNetwork.GetTotalProductPriceByMonth(m).Text('f', 2) + `'`
		}

		// Now loop through all products alphabetically  and add to the report
		// if this network does not have the current product simply add 0.00 to the report
		for _, p := range productHeadings {

			if currentNetwork.DoesProductExistOnNetwork(p) {

				currentProduct := currentNetwork.GetLoanProducts()[p]
				output += "," + `'` + currencySymbol + currentProduct.Text('f', 2) + `'`
			} else {
				output += "," + currencySymbol + `0.00`
			}

		}
		// totalLoanProducts - Calculate and add the total number of loan products to our report
		totalLoanProducts := strconv.Itoa(len(currentNetwork.GetLoanProducts()))

		// Add totals to file
		output += `,'` + totalLoanProducts + `'` + `,'`
		output += currencySymbol + currentNetwork.GetTotalProductPrice().Text('f', 2) + `'`
		output += `,'` + currencySymbol + allNetworkTotals.Text('f', 2) + `'`
		output += "\n"

	}

	// Finally write the output variables contents to our report file
	err := ioutil.WriteFile("reports/Output.csv", []byte(output), 0755)
	if err != nil {
		fmt.Println(`Error: could not create output file -> reports/Output.csv`)
	} else {
		fmt.Println(`Success: The results of your input have been written to reports/Output.csv`)
	}

}

func main() {

	// rows - read our loans file and parse the data into a [][]string splice
	rows := common.ParseLoansCSVFile("reports/Loans.csv")

	// totalRows - store total number of rows to process
	totalRows := len(rows) - 1

	if rows == nil {
		fmt.Println("Error:", `Cannot continue as errors have occured while trying to 
			import your CSV file. Please correct the errors and try again`)
		return
	}

	//Loop through each row and perform relevant calculations
	//rows[1:] because the first row are headers
	for index, row := range rows[1:] {
		fmt.Println(`Processing Row `, index+1, `of `, totalRows)
		performCalculations(row, index)
	}

	// Finally save generated results to reports/Output.csv
	fmt.Println(`Generating report file ....`)
	generateReport()

}
