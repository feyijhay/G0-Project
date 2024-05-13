package main

import (
	"fmt"
	"math"
	"time"
)

var (
	customerItems []string
	numberOfItems []float64
	amountOfGoods []float64
	items         string
	sub           float64
	subb          float64
	subTotal      []float64

	date = time.Now()
)

func main() {

	customerNameHere := ""

	for index := 0; index < math.MaxInt64; index++ {
		fmt.Println("Customer's name:")
		var customerName string
		fmt.Scan(&customerName)
		customerNameHere = customerName

		fmt.Println("Please enter your items:")
		fmt.Scan(&items)
		_ = append(customerItems, items)

		fmt.Println("Please enter items quantity:")
		var itemsNumber float64
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

	value := cashierInput()
	receipt(value, customerNameHere)
}

func cashierInput() string {
	fmt.Println("Cashier Name:")
	var cashierName string
	fmt.Scan(&cashierName)

	fmt.Println("How much discount will the customer get:")
	var customerDiscount float64
	fmt.Scan(&customerDiscount)

	return cashierName
}

func receipt(cashierName string, customerName string) {
	fmt.Println("JHAY'S STORE")
	fmt.Println("Main Branch")
	fmt.Println("LOCATION: 312, Herbert Macaulay Way, Sabo Yaba, Lagos")
	fmt.Println("TEL: 09078480034")
	date.Format("12-03-1997")
	fmt.Println("Cashier:", cashierName)
	fmt.Println("Customer Name:", customerName)

	fmt.Println("______________________________________________________________________________________________________________________________________________________")
	fmt.Println("ITEM         QTY         PRICE         TOTAL(NGN)")
	fmt.Println("______________________________________________________________________________________________________________________________________________________")
	for check := 0; check < len(customerItems); check++ {
		result := customerItems[check]
		answer := numberOfItems[check]
		total := amountOfGoods[check]
		sumOfQuantity := numberOfItems[check] * amountOfGoods[check]
		fmt.Printf("%s, %f, %f, %f", result, answer, total, sumOfQuantity)
		subTotal = append(subTotal, sumOfQuantity)
	}
	fmt.Println("_______________________________________________________________________________________________________________________________________________________")

	for check := 0; check < len(subTotal); check++ {
		var sub float64
		sub = sub + subTotal[check]
	}

	fmt.Printf("Sub Total: %f ", sub)

	discount := sub * 8 / 100
	fmt.Printf("Discount: %f", discount)

	VAT := sub * 7.5 / 100
	fmt.Printf("VAT @ 7.5: %f", VAT)

	fmt.Println("_________________________________________________________________________________________________________________________________________________________")
	billTotal := (sub + VAT) - discount
	fmt.Printf("Bill Total: %f", billTotal)
	fmt.Println("__________________________________________________________________________________________________________________________________________________________")

	fmt.Printf("KINDLY PAY TO THE CASHIER: %f", billTotal)

	fmt.Println("___________________________________________________________________________________________________________________________________________________________")

	fmt.Println("How much did the customer give you?:")
	var amountPaid float64
	fmt.Scan(&amountPaid)

	fmt.Println("   ")

	fmt.Println("JHAY'S STORE")
	fmt.Println("Main Branch")
	fmt.Println("LOCATION: 312, Herbert Macaulay Way, Sabo Yaba, Lagos")
	fmt.Println("TEL: 09078480034")
	date.Format("12-03-1997")
	fmt.Println("Cashier:", cashierName)
	fmt.Println("Customer Name:", customerName)
	fmt.Println("_____________________________________________________________________________________________________________________________________________________________")

	for check := 0; check < len(customerItems); check++ {
		result := customerItems[check]
		answer := numberOfItems[check]
		totall := amountOfGoods[check]
		sumQuantity := numberOfItems[check] * amountOfGoods[check]
		fmt.Printf("%s, %f, %f, %f", result, answer, totall, sumQuantity)
	}
	fmt.Println("_______________________________________________________________________________________________________________________________________________________________")
	for check := 0; check < len(subTotal); check++ {
		var subb float64
		subb = subb + subTotal[check]
	}

	fmt.Printf("Sub Total: %f", subb)

	discounted := subb * 8 / 100
	fmt.Printf("Discount: %f", discounted)

	VAAT := subb * 7.5 / 100
	fmt.Printf("VAT @ 7.5: %f", VAAT)
	fmt.Println("____________________________________________________________________________________________________________________________________________________________________")
	totalAmount := (subb + VAAT) - discounted
	fmt.Printf("Bill Total: %f", totalAmount)

	fmt.Printf("Amount paid: %f", amountPaid)

	balance := (amountPaid - totalAmount)
	fmt.Printf("Balance: %f", balance)
	fmt.Println("_______________________________________________________________________________________________________________________________________________________________________")

	fmt.Println("THANK YOU FOR YOUR PATRONAGE!!!")
	fmt.Println("________________________________________________________________________________________________________________________________________________________________________")
}
