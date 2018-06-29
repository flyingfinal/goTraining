package main

import (
	"fmt"
	"time"
)


func ping(c chan string) {
	for i := 0;;i++{
		c <- "ping..."
	}
}

func pong(c chan string) {
	for i := 0;;i++{
		c <- "pong..."
	}
}


func receiver(c chan string){
	for {
		fmt.Println(<- c)
		time.Sleep(time.Second)
	}

}


func main() {
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go receiver(c)

	var input string
	fmt.Scanln(&input)


}