package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// parse address from command line
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// custom logger for leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERRO\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// server static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// configure handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// start the server
	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
