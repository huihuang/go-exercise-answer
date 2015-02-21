/* Exercise: Maps */

package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	
	var wordmap map[string]int
	wordmap = make(map[string]int)
	
	sl := strings.Fields(s)
	
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
		wordmap[sl[i]] += 1
	}
	
	return wordmap
}

func main() {
	wc.Test(WordCount)
}
