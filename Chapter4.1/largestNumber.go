package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var (
		counter = 0
		num     int
		largest int
		//smallest int
	)

	for counter <= 10 {
		fmt.Println("Enter an integer:")
		fmt.Scan(&num)

		if num > largest {
			largest = num
		}

		counter++
	}
	fmt.Println("largest number is:", largest)

}
