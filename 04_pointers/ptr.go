package main

import "fmt"

func main() {
	arr := make([]int, 5, 10)
	//ptr := &arr

	addValue(arr)
	for _,value := range arr {
		fmt.Println(value)
	}
}

func addValue(x []int) {
	(x)[0] = 1
	(x)[4] = 8
}

