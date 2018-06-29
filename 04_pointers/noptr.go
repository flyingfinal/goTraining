package main

import "fmt"

func main() {
	var ptr [2]int

	assignValue(ptr)

	for _,value := range ptr {
		fmt.Println(value)
	}

}

func assignValue(x [2]int) {
	x[0] = 1
	x[1] = 3

	for _,value := range x {
		fmt.Println("Inside func: ", value)
	}
}