package main

import "fmt"

func main() {
	for i := 2; i <= 100; i++ {
		isPrime := true // Assume i is prime initially

		// Check if i is divisible by any number from 2 to i-1
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false // i is divisible, so it is not prime
				break
			}
		}

		if isPrime {
			fmt.Println(i)
		}
	}
}
