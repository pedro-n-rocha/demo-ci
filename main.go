package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Print("request")
	fmt.Fprintf(w, "Ola, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
