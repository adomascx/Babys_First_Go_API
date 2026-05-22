package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is home :3")
	fmt.Println("turim login mamamia")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	err := http.ListenAndServe(":8000", mux)
	fmt.Println("Serving traffic @ localhost:8000")
	if err != nil {
		log.Fatal("nepasileido :(")
	}

}
