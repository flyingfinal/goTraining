package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
	} ()

	panic("Panicc")

}

func fib(x int) (int) {
	if x == 0 {
		return 0;
	}
	if x == 1 {
		return 1;
	}
	return fib(x-1) + fib(x-2)
}