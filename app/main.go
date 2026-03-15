package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func stackHandler(w http.ResponseWriter, r *http.Request) {
	response := `{
		"proxy": "nginx",
		"api": "go",
		"cache": "redis",
		"orchestrated_by": "docker-compose"
	}`
	fmt.Fprintf(w, response)
}

func main() {

    rand.Seed(time.Now().UnixNano())

    err := loadIncidents()
    if err != nil {
        panic(err)
    }

    initRedis()   
	
    http.HandleFunc("/incident", incidentHandler)
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/stack", stackHandler)

    fmt.Println("Server running on port 5000")

    log.Fatal(http.ListenAndServe(":5000", nil))
}