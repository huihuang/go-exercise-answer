/* Exercise: Errors */

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

var delta float64 = 1e-12



func (e ErrNegativeSqrt) ToString() string {
	return fmt.Sprint("Convert %g to Stirng: ", e) + fmt.Sprint(float64(e))
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ") + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	
	if x >= 0 {
		var tmp float64 = x/2
	
		for math.Abs(x - tmp * tmp) >= delta {
		
			tmp = tmp - (tmp * tmp - x) / (2 * x)
			//fmt.Println(math.Abs(x*x - tmp * tmp))
		}
	
		return tmp, nil
	} else {
		return x, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
