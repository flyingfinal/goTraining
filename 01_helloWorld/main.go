package main

import "fmt"

func main() {

	x := make([]int, 5, 10)
	x[2] = 0

	defer f1(x...)

	Oddfunc := getOdd()
	i := 0
	for i < 10 {
		fmt.Println(Oddfunc())
		i++
	}
}

// This is comment

func f1(arr ...int) (int, int){
	x := 0
	y := 0
	for i,v := range arr {
		y += v
		fmt.Println(v, " ", i)
		x += i
	}


	return x,  y
}


func getOdd() func() (int){
	i := 1
	return func() (int){
		ret := i
		i += 2
		return ret
	}

}