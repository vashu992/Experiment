package main

import "fmt"

func main() {
	fmt.Println("Welcome to my this code")
 // var
	var name string 
	fmt.Println(name)
	fmt.Println("Ashutosh ")

    pawan := "Pawan"
	fmt.Println(pawan)

	var number int
	fmt.Println(number,2323322323)

	var objective bool
	fmt.Println(objective,true)
// ifelse
	age := "128"

	if age == "20" {
         fmt.Println("You are Adult ")
	}else if age == "18" {
	      fmt.Println("You are below to Adult")
	}else {
		fmt.Println("You Are Already Man")
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

	// Loop

	for i:=1; i<=10; i++ {
		fmt.Println("hello ashu",i)
	}

	// defer
	defer fmt.Println("This will be printed last")


	//Array
	var class [5]int
	class [0] = 10
	class [1] = 20
	class [2] = 30
	class [3] = 40
	class [4] = 50

	fmt.Println(class ,class)

	var student [5]string
	student [0] = "Ayan suryavamshi"
	student [1] = "Ayan suryavamshi"
	student [2] = "Ayan suryavamshi"
	student [3] = "Ayan suryavamshi"
	student [4] = "Ayan suryavamshi"

	fmt.Print("student = ",student)

	// Slice

	var students[]string

	students = append(students, "ashutosh")

	fmt.Println("students = ",students)

	

	
}