package main


import "fmt"
import "net/http"



func handleFn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, GO lang ")
}

func main() {
	http.HandleFunc("/", handleFn)
	_ = http.ListenAndServe(":8888", nil)
}


