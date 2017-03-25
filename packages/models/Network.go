package models

import (
	"fmt"
	"math/big"
)

/**
     Go does not have a data type for Decimal exactly as per other
	 programming languages, instead it provies a float64 type which
	 should be sufficient for most calcuations.

	 float64 should be fine for storing money type values however
	 as an added precaution - I have settled for using big.Float which allows
	 for far more precision and safer storage of Decimal types in GO.

**/

// Network - Holds a network with its associated products and calcuations
type Network struct {

	/*
		All variables start with lowercase letters therefore
		they can only be accessed within this module.This is done so to
		ensure proper encapsulation of data and allow interaction with
		this data only via accessors and mutators.
	*/

	name         string
	loanProducts map[string]*big.Float

	totalMonthPrice   map[string]*big.Float
	totalProductPrice *big.Float
}

// NewNetwork - Constructor method - sets up and returns a new Network struct
func NewNetwork(name string) *Network {
	return &Network{
		name:         name,
		loanProducts: make(map[string]*big.Float),

		totalMonthPrice:   make(map[string]*big.Float),
		totalProductPrice: big.NewFloat(0.00),
	}
}

// GetTotalProductPrice - returns total price for all products on current network
func (network *Network) GetTotalProductPrice() *big.Float {
	return network.totalProductPrice
}

// GetLoanProducts - returns loan products
func (network *Network) GetLoanProducts() map[string]*big.Float {
	return network.loanProducts
}

//DoesProductExistOnNetwork - checks if network has loan product and returns true or false
func (network *Network) DoesProductExistOnNetwork(name string) bool {
	_, exists := network.loanProducts[name]
	return exists

}

// AddMonthPrice - Increments total amount spent on product by month
func (network *Network) AddMonthPrice(month string, amount string) {
	bigFloatAmount, _, _ := big.NewFloat(0.00).Parse(amount, 10)

	_, exits := network.totalMonthPrice[month]

	if exits == false {
		network.totalMonthPrice[month] = big.NewFloat(0.00)
	}

	network.totalMonthPrice[month] = big.NewFloat(0.00).Add(
		network.totalMonthPrice[month], bigFloatAmount)

}

// GetTotalProductPriceByMonth - returns total product price by month
func (network *Network) GetTotalProductPriceByMonth(month string) *big.Float {

	_, exists := network.totalMonthPrice[month]
	if exists == true {

		return network.totalMonthPrice[month]
	}
	return big.NewFloat(0.00)
}

//AddProduct - Adds product to current network
func (network *Network) AddProduct(productName string, amount string) *big.Float {
	_, doesProductExist := network.loanProducts[productName]
	f, _, err := big.NewFloat(0.00).Parse(amount, 10)

	// only add the product to our products list once
	// if product exists simply continue and increment the existing value by amount
	if doesProductExist == false {
		network.loanProducts[productName] = big.NewFloat(0.00)
	}

	if err == nil {
		// bigFloat numbers don't support the + operator so we therefore need to use
		// it's add method to perform the addition
		network.totalProductPrice = big.NewFloat(0.00).Add(f, network.totalProductPrice)
		network.loanProducts[productName] = big.NewFloat(0.00).Add(f, network.loanProducts[productName])

	} else {
		fmt.Println(`Warning: `, ` Error reading amount as a money value:`, amount)
	}

	return f
}
