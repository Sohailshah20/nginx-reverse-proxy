package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "bad request", 404)
			return
		}
		w.Write([]byte("hello world!!"))
		return
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Printf("server started on port %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}
}
