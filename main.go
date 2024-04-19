package main

import (
	"log"
	"net/http"
)

func main() {
	server := newPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":3000", server))
}
