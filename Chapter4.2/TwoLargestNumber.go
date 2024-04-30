package main

import "fmt"

func main() {
	var (
		number1, number2, holder int
	)

	fmt.Println("Enter 10 numbers:")
	fmt.Scan(&number1)

	if number1 < number2 {
		holder = number1
		number1 = number2
		number2 = holder
	}

	for check := 2; check <= 10; check++ {
		fmt.Printf("Number %d: ", check)
		fmt.Scan(&holder)

		if holder > number1 {
			number2 = number1
			number1 = holder
		} else if holder > number2 && holder != number1 {

		}
	}

	fmt.Println("largest number is : ", number1)
	fmt.Println("second largest number is : ", number2)
}
