package backend

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServePage(writer http.ResponseWriter, reqest *http.Request) {

	if reqest.Method != "GET" {
		log.Fatal(http.ErrNotSupported)
	}
	io.WriteString(writer, "Hello world!")

}

func ServeAnother(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Fatal(http.ErrNotSupported)
	}
	io.WriteString(w, "Post request had been made")
}

func Actual() {

	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./backend/html-docs"))
	mux.HandleFunc("/hello", ServePage)
	mux.HandleFunc("/content", ServeAnother)
	mux.Handle("/login", http.StripPrefix("/login", fileserver))

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal("Server has been corrupted bro")
	}

}

func GinServer() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
