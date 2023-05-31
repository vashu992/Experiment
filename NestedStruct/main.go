package main

import "fmt"

type Personal struct{
	Name string
	Age int
	Salary float64
	Education string
	// Address
	address Address
}

type Address struct{
	City string
	State string
	Zipcode int
	Country string

}

type School struct{
	Name string
	Teacherno int
	Department string
	Class int
	Playground string
	Sport string

}

type SmartPhone struct{
	Name string
	Model string
	Manufactured int
	Display string
	Quality string
}

type Laptop struct{
	Name string
	ModelName string
	Battery int
	Size int
	Processor string

}

func main() {

	address := Address{City: "Bhadohi",State: "Utter Pradesh",Zipcode: 221301,Country: "India"}
	personal := Personal{Name: "Ashutosh Pratap",Age: 23,Salary: 21.54,Education: "PG Diploma in computer application",address: address}
	cps := School{Name: "ChandraPrabha Public School",Teacherno: 23,Department: "So Many Department",Class: 1-10,Playground: "Very high",Sport: "Cricket Hockey etc"}
	iphone :=SmartPhone{Name: "Iphone 14 pro MAX",Model:"14 Series",Manufactured: 12/12/2013,Display: "Retina",Quality: "This is very good" }
	Dell := Laptop{Name: "Dell Inspiron",ModelName: "YOGA",Battery: 5000,Size: 16,Processor: "Octacore"}

	fmt.Println(personal)
	fmt.Println(address)
	fmt.Println(cps)
	fmt.Println(iphone)
	fmt.Println(Dell)
}