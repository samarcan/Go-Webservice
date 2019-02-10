package main

import (
	"fmt"
    "net/http"
    "log"
	"github.com/gorilla/mux"
	"os"
)

func yourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {
    router := mux.NewRouter()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	
	fmt.Println("Serving in 127.0.0.1:" + port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatalln(err)
	}
}