/* Exercise: rot13Reader */

package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13.r.Read(b)
	
	for i := 0; i < n; i++ {
		if (b[i] <= 'M' && b[i] >= 'A') || (b[i] <= 'm' && b[i] >= 'a') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
