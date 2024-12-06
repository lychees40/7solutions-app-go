package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// X
func main() {
	http.HandleFunc("/", getHello)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	_, err := io.WriteString(w, "Hello "+name+"\n")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
