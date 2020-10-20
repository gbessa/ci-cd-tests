package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>CI/CD FullCycle course Modified 03!")
}

func main() {
	var port = ":8080"
	fmt.Printf("Listening at port %s...", port)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))

}
