/* Exercise: Slices */

package main

import (
	"code.google.com/p/go-tour/pic"
	//"fmt"
	"math"
)

func foo(x, y int) int {
	
	return int(math.Pow(float64(x),float64(y^x)))//x^y//x*x*y + y*y*x + x*y + x + y + 1
}

func Pic(dx, dy int) [][]uint8 {
	slice0 := make([][]uint8, dy)
	
	for i := 0; i < dy; i++ {
		slice0[i] = make([]uint8, dx)
	}
	
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			slice0[i][j] = uint8(foo(i,j))
		}
	}
	
	return slice0
}

func main() {
	pic.Show(Pic)
}
