package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "CCCC Hello Worlgo run")
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}
