package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func httpHandle(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("HelloWorld")) }).Methods("GET")
}

func main() {
	router := mux.NewRouter()
	httpHandle(router)
	fmt.Println("Hello World")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"Accept", "multipart/form-data", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
