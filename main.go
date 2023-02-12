package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const webPort = ":4000"

func home(w http.ResponseWriter, r *http.Request) {

	// non-existent routes are caught by the wildcard "/" path
	if r.URL.Path != "/" {
		http.Error(w, "The page you're looking for doesn't exist", http.StatusNotFound)
		return
	}

	_, _ = w.Write([]byte("You have hit the home page!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	_, _ = fmt.Fprintf(w, "This is the page displaying snippet with id=%d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Only POST method is allowed for snippet creation."))
	}

	_, _ = w.Write([]byte("You hit the snippet create route."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Server started on port %s", webPort)

	err := http.ListenAndServe(webPort, mux)
	log.Panic(err)
}
