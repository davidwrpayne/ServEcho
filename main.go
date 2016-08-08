package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling Request\n")
	fmt.Println("URL:" + r.URL.String())
	body, _ := ioutil.ReadAll(r.Body)
	for k, v := range r.Header {
		fmt.Print("\theaderKey = ")
		fmt.Print("\t" + k)
		fmt.Print("\theaderValue = ")
		fmt.Printf("\t%s\n",v)
	}
	fmt.Print("body = ")
	fmt.Println("\t",string(body))
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Printf("Starting ServEcho\n")
	fmt.Printf("This server just 200's for all requests and logs what happened")

	port := ":13330"
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)

}
