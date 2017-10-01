package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

    // adding debug header to test (strong/weak) ETags in combination with NGINX
    w.Header().Set("ETag", "HelloWorld")
    w.Header().Set("Custom-ETag", "HelloWorld")

    fmt.Fprintln(w, "<b>Look for the ETag & Custom-ETag response headers! :-)</b>")

}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}