package main

import "fmt"

func main() {
	line := 5
	

	for i := 0; i < line; i++ { 
		for j := 0; j < line -i; j++ {
			if j == line-i-1 {
				fmt.Print("a")
			} else {
				fmt.Print("a ")
			}
		}
		fmt.Println()
	}
}
