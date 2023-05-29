package main

import (
	"net/http"
)

func main() {
	//Create the default mux
	mux := http.NewServeMux()

	//Handling the greatlife. the handler is a function
	mux.HandleFunc("/greatlife", lana)

	//Handling the poorlife.The handler is a type
	plife := poorLife{}
	mux.Handle("/poorLife", plife)

	//Create a server

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
func lana(res http.ResponseWriter, req *http.Request) {
	reply := []byte("This is Adelana")
	res.WriteHeader(200)
	res.Write(reply)
}

type poorLife struct{}

func (Y poorLife) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	reply := []byte("This guy is poor")
	res.WriteHeader(200)
	res.Write(reply)
}
