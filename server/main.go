package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("h")

	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./website/dist")))

	r.HandleFunc("/sables.exe", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./download/sables.exe")
	})

	http.ListenAndServe(":8080", r)
}
