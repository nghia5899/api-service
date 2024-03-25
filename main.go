package main

import (
	"apiService/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	router := router.CreateRouter()

	go func() {
		addr := fmt.Sprintf("0.0.0.0:%d", 8000)
		log.Printf("Starting HTTP server on %s\n", addr)
		srv := &http.Server{
			Handler:      router,
			Addr:         addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())
	}()

	for {
		time.Sleep(time.Second)
	}
}
