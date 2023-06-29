package calculator

import "fmt"

type User struct {
	Name string
	contact int
}

type student struct {
	Name string
	Rollno int	
}
func add(a, b int) int {
	fmt.Println("add call hua")
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) int {
	return a / b
}