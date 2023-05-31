package main

import "fmt"

func main() {
	 
	val := 1223
	val2 := "678908"


	var interfaceExample interface{}

	interfaceExample = val
	fmt.Println("Interface value",interfaceExample)

	interfaceExample = val2
	fmt.Println("Interface value",interfaceExample)

	interfaceExample = false
	fmt.Println("Interface value",interfaceExample)
}