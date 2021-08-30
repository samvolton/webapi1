package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}

// https://github.com/cihanozhan/golang-webapi-samples/blob/master/00_webApi/main.go

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
	fmt.Println("Web Server")

}
