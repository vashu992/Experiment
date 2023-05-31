package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int
	a = 10

	var b int16

	b = int16 (a)

	fmt.Println("b =",b)
	var c string
	c = fmt.Sprintf("%d",b)
	fmt.Println("c = ",c)

	num := "221301"
	numint ,err := strconv.Atoi(num)
	fmt.Println("numint = ",numint,"err = ",err)

	num = "Vishal"
	numint ,err = strconv.Atoi(num)
	fmt.Println("numint = ",numint,"err = ",err)

	var x float64
	x = 10.5

	var y float64

	y = (x)

	fmt.Println("y = ",y)

	var t bool
	t = false

	var s bool
	s = (s)

	fmt.Println("t = ",t)


    
}