package main

import "fmt"

type User struct {
	Name string
	Class int
	Email string
	Stutus bool
}

type Student struct {
	Name string
	Class int
	Rollnumber int
	District string
	Fathername string
}

type Laptop struct {
	Name string
	Configuration string
	Modelnumber int
	Manufactured bool
	Weight float64
	Displayname string
	County string
	Ceo string
	Owner bool
	Customer string
	
}

func main() {
	fmt.Println("struct")
	ashutosh := User{"Ashutosh",5,"ashutoshpratap992@gmail.com",true}
	fmt.Println(ashutosh)
	ayan:= Student{"Ayan",12,03,"varanasi","elonmask"}
	fmt.Println(ayan)
	dell := Laptop{"Dell","good",5,true,1.23,"Retina","India","kk singh",false,"PP Saini"}
	fmt.Println(dell)
}