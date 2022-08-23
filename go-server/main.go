package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "Error 404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Unsupported Method", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request ){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	age := r.FormValue("age")
	fmt.Fprintf(w, "name: %v\n", name)
	fmt.Fprintf(w, "email: %v\n", email)
	fmt.Fprintf(w, "age: %v\n", age)
}


func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting the server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}