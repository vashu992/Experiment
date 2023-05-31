package main

import "fmt"

// non primitive data type

type student struct{
	 Name string
	 Class int
	 Rollnumber int
	 studentAddress Address
} 

type Address struct{

	Lane1 string
	Lane2 string
	Post string
	District string
	Village string


}

type phone struct{
    BasicInfo
	Imail int
	Dispay string
	config configuration
}

type Laptop struct{
	BasicInfo
	Serialnumber int
	configuration configuration 
}

type BasicInfo struct{
	Model string
	Brand string
}

type configuration struct{
	RAM string
	ROM string
	Processor string
	Os string
}

func main() {

	// primitive aproach 
	// var name_vishal string
	// var class_vishal int
	// var rollnumber int

	// var name _shivam string
	// var class_shivam int
	// var rollnumber_ int

	// non- primitive
	


	var vishal student 
	vishal.Class = 123
	vishal.Rollnumber =65
	vishal.Name = "ashutosh pratap"
	vishal.studentAddress.District= "varanasi"
	vishal.studentAddress.Village = " ugapur"
	var  phone phone
	phone.Brand = "realme"
	phone.BasicInfo.Brand = "nice "
	phone.Dispay = "Retina"
	phone.Imail = 345634
	phone.Model = "narzo"
	var laptop Laptop
	laptop.configuration.RAM = "two"
	laptop.configuration.ROM = "four"
	laptop.Brand = "HP"
	laptop.Model = "inspiron"  
	laptop.configuration.Processor = "Octa core"
	laptop.BasicInfo.Model = "Latitude"

	shivam := student{
		Name: "Ashu singh",
		Class: 22,
		Rollnumber: 3323,
		studentAddress: Address{
			District:"Lucknow" ,
			Village: "ismailganj",
			Lane1: "Pushpa hospital Sector 8",
			Lane2: "Gomati Nagar",
		},
	}
	fmt.Println("Running go code...this is vishal struct",vishal)
	fmt.Println("Running go code...this is ashu singh struct",shivam)
	fmt.Println( "this code disigned by ashu",phone)
	fmt.Println("this code made is laptop",laptop)


	    
}