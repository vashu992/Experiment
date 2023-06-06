package main

import "fmt"

func main() {


	// intger
	var digit[12] int
	digit[0] = 10
	digit[1] = 20
	digit[2] = 30
	digit[3] = 40
	digit[4] = 50
	digit[5] = 60
	digit[6] = 70
	digit[7] = 80
	digit[8] = 90
	digit[9] = 100
	digit[10] = 110
	digit[11] = 120
	

	fmt.Println(digit,digit)

	digit[5] = 1000
	fmt.Println(digit,digit)


	// string
	var doctor [20]string

	doctor[0] = "Dr.Rathi"
	doctor[1] = "Dr.Vivek Sharma"
	doctor[2] = "Dr.P S Dubey"
	doctor[3] = "Dr.Amol Gahwad"
	doctor[4] = "Dr.Vinod kumar"
	doctor[5] = "Dr.K K Gupta"
	doctor[6] = "Dr.Santosh "
	doctor[7] = "Dr.Pranav"
	doctor[8] = "Dr.Pankaj"
	doctor[9] = "Dr.Pawan"
	doctor[10] = "Dr.Ratnesh Pal"
	doctor[11] = "Dr.Ashis Yadav"
	doctor[12] = "Dr.P P Ojha "
	doctor[13] = "Dr.Raghav"
	doctor[14] = "Dr.Jhataka"
	doctor[15] = "Dr.Ashwin"
	doctor[16] = "Dr.Anand Madhav"
	doctor[17] = "Dr.Aryan Singh"
	doctor[18] = "Dr.B B Sinha"
	doctor[19] = "Dr. D K Tripadhi"

	doctor[1] = "Dr.Avadh Ojha"

	fmt.Println(" doctor = ",doctor)


	// interface
	var place [7] interface{}

	place [0] = "Varanasi"
	place [1] = true
	place [2] = 14
	place [3] = "UUDCPG"
	place [4] = "18"
	place [5] = 18
	place [6] = false

	fmt.Println("place = ",place)

	// bool

	var condition[5] bool

	condition [0] = true
	condition [1] = false
	condition [2] = true
	condition [3] = false
	condition [4] = true

	fmt.Println("condition =",condition)


	// SLICE 

	var student [] string

	student = append(student, "Ashutosh Pratap")
	student = append(student, "Rudra Bhai")
	student = append(student, "Vivek Sharma")
	
	
	fmt.Println("student = ",student[1])

	var number [] int

	number = append(number, 1000)
	number = append(number, 2000)
	number = append(number, 3000)
	number = append(number, 4000)

	fmt.Println(number[1])

	var Action [] bool

	Action = append(Action, false)
	Action = append(Action, true)

	fmt.Println(Action[0])

	var human [] interface {}

	human = append(human, "Ahutosh Pratap")
	human = append(human, 12)
	human = append(human, false)
	human = append(human, true)
	human = append(human, 12.12)
	human = append(human, "BABA VISHWANATH")
	
	fmt.Println(human[5])


}