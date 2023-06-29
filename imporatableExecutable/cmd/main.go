package main

import (
	"fmt"

	"github.com/vashu992/Experiment/imporatableExecutable/calculator"
)

func main() {

	usr := calculator.User {}
	usr.Name = "adsgfdgf" 
	//fmt.Println(" sum = ",calculator.add(100,76))
	fmt.Println(" sub = ",calculator.Divide(100,77))
	fmt.Println(" sum = ",calculator.Multiply(100,76))
	fmt.Println(" sum = ",calculator.Sub(100,76))
	
}