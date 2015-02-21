/* Exercise: Loops and Functions */

package main

import (
	"fmt"
	"math"
)

var delta float64 = 1e-12

func Sqrt(x float64) float64 {
	
	var tmp float64 = x/2
	
	for math.Abs(x - tmp * tmp) >= delta {
		
		tmp = tmp - (tmp * tmp - x) / (2 * x)
		//fmt.Println(math.Abs(x*x - tmp * tmp))
	}
	
	return tmp
}

func main() {
	var i float64 = 2
	fmt.Println(Sqrt(i))
	fmt.Println(math.Sqrt(i))
}
