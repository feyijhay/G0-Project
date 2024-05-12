package main

import (
	"fmt"
	"math"
	"time"
)

var (
	customerItems []string
	numberOfItems []int
	amountOfGoods []float64
	items         string
	date          = time.Now()
)

func main() {

	cashierInput()

	for index := 0; index < math.MaxInt64; index++ {
		fmt.Println("Customer's name:")
		var customerName string
		fmt.Scan(&customerName)
		fmt.Scanln(&customerName)

		fmt.Println("Please enter your items:")
		fmt.Scan(&items)
		_ = append(customerItems, items)

		fmt.Println("Please enter items quantity:")
		var itemsNumber int
		fmt.Scan(&itemsNumber)
		_ = append(numberOfItems, itemsNumber)

		fmt.Println("Enter amount:")
		var amount float64
		fmt.Scan(&amount)
		_ = append(amountOfGoods, amount)

		fmt.Println("Would you like to add more items?:")
		var moreItems string
		fmt.Scan(&moreItems)
		if moreItems == "no" || moreItems == "NO" || moreItems == "No" {
			break
		}
	}

}

func cashierInput() {
	fmt.Println("Cashier Name:")
	var cashierName string
	fmt.Scan(&cashierName)
	fmt.Scanln(&cashierName)

	fmt.Println("How much discount will the customer get:")
	var customerDiscount float64
	fmt.Scan(&customerDiscount)

}

func receipt() {
	fmt.Println("JHAY'S STORE")
	fmt.Println("Main Branch")
	fmt.Println("LOCATION: 312, Herbert Macaulay Way, Sabo Yaba, Lagos")
	fmt.Println("TEL: 09078480034")
	date.Format("12-03-1997")
	fmt.Println("Cashier:", cashierName)
}

var cashierName string
fmt.Print("Enter cashier's name: ")
fmt.Scanln(&cashierName)
fmt.Println("Cashier:", cashierName)