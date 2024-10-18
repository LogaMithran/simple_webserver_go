package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(response http.ResponseWriter, request *http.Request) {
	log.Printf("Form handler")

	if request.Method != "POST" {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if request.URL.Path != "/form" {
		http.Error(response, "404 not found", http.StatusNotFound)
	}
	name := request.PostFormValue("username")
	mobile := request.PostFormValue("mobile_number")
	log.Printf(name, mobile)

	fmt.Fprintf(response, "Username ::: %v and user's Mobile number ::: %v", name, mobile)
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	log.Printf("Hello handler")
	if request.Method != "GET" {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if request.URL.Path != "/hello" {
		http.Error(response, "404 Not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(response, "Hello from webserver")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	http.Handle("/pattern", http.NotFoundHandler())
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Printf("Oops some error in starting the server")
		log.Fatal(err)
	}
}
