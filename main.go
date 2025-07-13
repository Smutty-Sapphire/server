package main

import (
    "fmt"
    "net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! ")
    fmt.Println(r.Header)
}

func main() {

    http.HandleFunc("/", httpHandler)
    http.ListenAndServe(":8080", nil)

    fmt.Println("Hello")
}
