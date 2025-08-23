package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "RunAt")
}

func main() {
	Greet(os.Stdout, "RunAt")
	http.ListenAndServe(":6969", http.HandlerFunc(MyGreetHandler))
}
