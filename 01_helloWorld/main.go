package main

import "fmt"

func main() {
	fmt.Println(f1())
}

// This is comment

func f1() (int, float64){
	var x float64 = 1.6
	y := x
	return int(x),  y
}