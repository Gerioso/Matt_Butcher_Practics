package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	port, err := os.LookupEnv("PORT")
	if !err {
		fmt.Println("Not found")
	}
	fmt.Println(port)
	http.ListenAndServe("localhost:"+port, nil)
}
func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
