package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/app")))
	fmt.Println("Start listening port 9999")
	err := http.ListenAndServe(":9999", nil)
	if err == nil {
		log.Fatal("ListenAndServe", err)
		return
	}
}
