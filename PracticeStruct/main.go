package main

import "fmt"

type Bike struct {
	Model int
	Name string
	manufactured int
	Design string
	Features string
}

type Car struct {
	Name string
	Airbags int
	Seat int 
	Lenght int
	Height int
	Look string
}

type BasicInfo struct{
	Engine string
	Speed int
	Break int
	Control string
}

type configuration struct{
	Power int
	Oil string
	Mileage int
	Sound string
}

func main() {
	

	var hero Bike
	hero.Name ="Super Splendor"
	hero.manufactured = 2023
	hero.Design = "Very good looking "
	hero.Features = "X Tech Technology"
	hero.Model = 5/12/2023
	
	var MarutiSuzuki Car
	MarutiSuzuki.Name = "Grand Vitara"
	MarutiSuzuki.Seat = 7
	MarutiSuzuki.Airbags = 6
	MarutiSuzuki.Height = 4
	MarutiSuzuki.Lenght = 3
	MarutiSuzuki.Look = "Sport wala look hai Bhai"



	BasicInfo := BasicInfo{
		Engine: "Very power full Engine",
		Speed: 120,
		Control: "ABS Technology",
		Break: 4,
	}
	configuration := configuration{
		Power: 180,
		Oil: "Petrol",
		Mileage: 456,
		Sound: "This sound Quality is boom boom",
	}
	fmt.Println("heri Bike is very good bike",hero)
	fmt.Println("This car is Looking is too good",MarutiSuzuki)
	fmt.Println("This is Configuration ",configuration)
	fmt.Println("This is all knoledge about basic info",BasicInfo)
}