package main

import (
	
	"fmt"
)

func main() {

	Buyer := "Apache"

	if Buyer == "Glemour"{
		fmt.Println("This Bike manufactured by Hero Motercorp.")
	}else if Buyer == "Apache" {
		fmt.Println("This Bike manufactured by TVS Company ")
	}else{
		fmt.Println("Please bro purchase the car Beacause car is very luxrious and comfortable for travelling",Buyer)
	}

	num := 10

	if( num == 0 || num > 0)&&(num % 2 ==0){
		fmt.Println(" Value is positive++++++++")

    }else if( num == 0 || num > 0)&&(num % 2 !=0){
			fmt.Println(" Value is positive odd")
	}else if num < 0 {
		fmt.Println("Value is negative ")
	}else {
		fmt.Println("Positive value hai --------",num)
	}
}