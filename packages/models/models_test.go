package models

import (
	"fmt"
	"math/big"
	"strings"

	"testing"
)

//TestAddMonthPrice - Tests that the AddMonthPrice method of our Network model works correctly
func TestAddMonthPrice(t *testing.T) {
	network := NewNetwork(`Test Network`)
	sample := []string{
		`Mar:2000`,
		`Mar:3000`,
		`Jun:1450`,
		`Jun:8000`,
	}

	for _, s := range sample {
		data := strings.Split(s, `:`)
		network.AddMonthPrice(data[0], data[1])
	}

	march := network.GetTotalProductPriceByMonth(`Mar`)

	if big.NewFloat(5000.00).Cmp(march) != 0 {
		t.Error(
			`Network AddMonthPrice not working correctly for March, Got : `,
			march, ` Expected: 5000`,
		)

	}

	june := network.GetTotalProductPriceByMonth(`Jun`)
	if big.NewFloat(9450.00).Cmp(june) != 0 {
		t.Error(
			`Network AddMonthPrice not working correctly for June, Got : `,
			june, ` Expected: 9450`,
		)

	}

	totalMonthPrice := big.NewFloat(0.00)
	for _, v := range network.totalMonthPrice {
		totalMonthPrice = big.NewFloat(0.00).Add(totalMonthPrice, v)
	}
	if big.NewFloat(14450.00).Cmp(totalMonthPrice) != 0 {
		t.Error(
			`Total month calculation not working correctly, Got: `,
			network.totalMonthPrice, `Expected: 14450`)
	}

}

//TestAddProduct - Tests that the AddProduct method of our Network model works correctly
func TestAddProduct(t *testing.T) {
	network := NewNetwork(`TestNetwork`)
	sampleProducts := make(map[string]string)
	sampleProducts[`Product 1`] = `500.50`
	sampleProducts[`Product 2`] = `800.89`
	sampleProducts[`Product 3`] = `699.99`
	sampleProducts[`Product 4`] = `400.52`

	totalProductPrice := big.NewFloat(2401.8999999999996)

	for productName, productPrice := range sampleProducts {
		fmt.Println(productPrice)
		network.AddProduct(productName, productPrice)
	}

	if network.totalProductPrice.Cmp(totalProductPrice) != 0 {
		t.Error(
			`Network AddProduct not working correctly Got : `,
			network.totalProductPrice, ` Expected: `,
			totalProductPrice,
		)

	}

	network.AddProduct(`Product 1`, `500.50`)
	network.AddProduct(`Product 2`, `800.89`)

	loanProducts := network.GetLoanProducts()
	if loanProducts[`Product 1`].Cmp(big.NewFloat(1001.0)) != 0 {
		t.Error(
			`Network AddProduct product level calculations not working : `,
			loanProducts[`Product 1`], ` Expected: 1001.0`,
		)
	}

	if loanProducts[`Product 2`].Cmp(big.NewFloat(1601.78)) != 0 {
		t.Error(
			`Network AddProduct product level calculations not working : `,
			loanProducts[`Product 2`], ` Expected: 1601.78`,
		)
	}
}
