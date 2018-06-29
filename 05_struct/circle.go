package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
	x float64
	y float64
}
// (c *Circle) or (c Circle), there is difference
func (c *Circle) area() float64{
	c.r = 3
	return math.Pi * c.r * c.r
}

func main() {
	cir := Circle{2, 1, 2}

	fmt.Println("Area: ", cir.area())
	fmt.Println("Area: ", cir.r)

}
