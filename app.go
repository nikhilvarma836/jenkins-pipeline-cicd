// app.go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("2006-01-02 15:04:05")
    fmt.Fprintf(w, "Current Date and Time: %s", currentTime)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at :8080")
    http.ListenAndServe(":8080", nil)
}

