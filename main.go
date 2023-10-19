package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("hola mundo go")
	// routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	// start the server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	for i := 1000; i < 1100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
		w.Write([]byte("contact page"))
	}

}

// go run main.go
// go build main.go
// ./main
