package backend

import (
	"io"
	"log"
	"net/http"
)

func ServePage(writer http.ResponseWriter, reqest *http.Request) {

	if reqest.Method != "GET" {
		log.Fatal(http.ErrNotSupported)
	}
	io.WriteString(writer, "Hello world!")

}

func ServeAnother(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Fatal(http.ErrNotSupported)
	}
	io.WriteString(w, "Post request had been made")
}

func Actual() {
	fileserver := http.FileServer(http.Dir("./myfiles"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", ServePage)
	http.HandleFunc("/content", ServeAnother)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("Server has been corrupted bro")
	}

}
