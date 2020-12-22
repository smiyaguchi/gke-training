package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("failed: The number of arguments should be two.")
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	log.Println("start listen and serve")
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
