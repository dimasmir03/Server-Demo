package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++
	fmt.Fprintln(w, "Counter:", ct.counter)
}

func main() {
	th := &CounterHandler{counter: 0}

	http.Handle("/count", th)
	http.HandleFunc("/hello", Hello)

	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
