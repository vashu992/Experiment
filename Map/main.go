package main

import (
	"fmt"
	
)

func main() {

	// Map decleration

	m := make(map[string]int)

	// Save Data in map
	m["Ashutosh"] = 03
	m["Pratap"] = 05

	// Read the Data from map

	val, ok := m["Ashutosh"]
	fmt.Println("val = ",val ,"ok = ",ok)
	
	val , ok = m["abc"]
	if ok {
		fmt.Println("Value not found , val = ",val)
	}else{
		fmt.Println("value not found")
	}

	a := make(map[int]int)

	a[12] = 15
	a[22] = 28

	val, ok = a [12]
	fmt.Println(val ,  ok )

	val, ok = a [27]
	fmt.Println(val,ok)

	h := make(map[int]string)
	h[6] = "Bajaragbali"
	h[5] = "Siyaram"

	val, ok = h[9] 
	fmt.Println("val =", val, "ok =", ok)

	d := make(map[string] )

	d["jay Shree Ram"] = "Jai ho"
	d["Jay Shree RadheKrishna"] = "Vasudev"

	val , ok = d["jai Shree Ram"]
	fmt.Println(" val = ", val ,"ok = " , ok)

	z := make(map[interface{}]interface{})

	z[5] = "ashu"
	z[7] = 3

	val , ok := z[7]
	fmt.Println("val = ",val,"ok = ",ok)

	
}


	



