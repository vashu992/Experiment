package main

import (
	"fmt"
)

func main() {
	count := 0
	number := 2

	

	for count < 25 {
		isPrime := true

		for i := 2; i < number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Printf("%d ", number)
			count++
		}

		number++
	}

	fmt.Println("\nTotal count:", count)
}
