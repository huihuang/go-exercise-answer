/* Exercise: Fibonacci closure */

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	/*
	a := 0
	b := 1
	flag1 := true
	flag2 := true
	return func() int {
		
		if flag1 {
			flag1 = false
			return a
		}
		
		if flag2 {
			flag2 = false
			return b
		}
		
		temp := a + b
		a = b
		b = temp
		
		return a+b
	}
*/
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
