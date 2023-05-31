package main

import "fmt"

func main() {

	for i := 0; i < 26; i++ {
		if i%2 == 0 {
			fmt.Printf("%c ", 'A'+i)
		} else {
			fmt.Printf("%c ", 'a'+i)
		}
	}

	// for i:=0 ;i <5 ; i++ {
	// 	if i%2==0 {
	// 		fmt.Printf("%c",'u' +i)
	// 	}
	// }
}
