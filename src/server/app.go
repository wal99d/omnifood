package main

import (
	"log"
	"net/http"
)

func must(e error) {
	log.Fatal(e)
}
func main() {
	log.Printf("Starting App Server @ %s\n", "22331")
	must(http.ListenAndServe(":22331", http.FileServer(http.Dir("."))))
}
