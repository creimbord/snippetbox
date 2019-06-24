package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Значение флага будет храниться в переменной addr в runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Парсит значение флага из командной строки
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
