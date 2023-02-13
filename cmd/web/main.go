package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog, infoLog *log.Logger
}

func main() {
	// parse address from command line
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// custom logger for leveled logging
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	// server static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// configure handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// server configuration
	// ErrorLog is added so that the server writes its errors to stderr
	srv := &http.Server{
		Handler:  mux,
		Addr:     *addr,
		ErrorLog: app.errorLog,
	}

	// start the server
	app.infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
