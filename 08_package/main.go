package main

import "fmt"
import "github.com/flyingfinal/goTraining/08_package/math"
func main() {
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
}

