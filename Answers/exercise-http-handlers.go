/* Exercise: HTTP Handlers */

package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (st *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, st.Greeting + st.Punct + st.Who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("<html><body><h1>I'm a frayed knot.</h1><body></html>"))
	http.Handle("/struct", &Struct{"Hello",":","Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
