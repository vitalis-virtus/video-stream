package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hlsDir := "./videos"

	_, err := os.Stat(hlsDir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("HLS Directory doesn't exist")
		}
	}

	fs := http.FileServer(http.Dir(hlsDir))
	http.Handle("/hls/", http.StripPrefix("/hls/", fs))

	port := 8080

	fmt.Printf("Starting HLS server on http://localhost:%d/hls/\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
