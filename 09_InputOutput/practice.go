package main
import (
	"fmt"
	"io/ioutil"
)
func main() {
	bs, err := ioutil.ReadFile("G:\\gogo\\src\\github.com\\flyingfinal\\goTraining\\09_InputOutput\\test.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(bs)
	fmt.Println(str)
}