package main

import (
	"fmt"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	//Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments. The response writer is used to fill in the HTTP response. Here our simple response is just “hello\n”.
	
		query := req.URL.Query()
		filters, present := query["filters"]
		if !present || len(filters) == 0 {
			fmt.Println("filters not present")
		}
		w.WriteHeader(200)
      w.Write([]byte(strings.Join(filters, ",")))
}

func defHandler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Default")
}

func main(){

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", defHandler)
	http.ListenAndServe(":8080", nil)
}