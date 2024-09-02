package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	// Greet(os.Stdout, "Nihal")
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(MyGreeterHandler)))
}
