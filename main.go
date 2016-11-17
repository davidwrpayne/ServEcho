package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling Request\n")
	fmt.Println("URL:" + r.URL.String())
	s, _ := httputil.DumpRequest(r,true)
	fmt.Printf("body = %s",string(s))

	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Printf("Starting ServEcho\n")
	fmt.Printf("This server just 200's for all requests and logs what happened")

	port := ":13330"
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)

}
