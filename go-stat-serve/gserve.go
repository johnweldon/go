package main

import (
	"fmt"
	"github.com/johnweldon/go/go-stat"
	"log"
	"net/http"
	"os"
)

func main() {
	var defaultFn = func(w http.ResponseWriter, r *http.Request) {
		log.Print(r)
	}
	http.Handle("/projects/default/notes/", http.StripPrefix("/projects/default/notes/", go_stat.NoteHandler{}))
	http.Handle("/projects/default/events/", http.StripPrefix("/projects/default/events/", go_stat.EventHandler{}))
	http.Handle("/projects/default/", http.StripPrefix("/projects/default/", go_stat.ProjectHandler{}))
	http.HandleFunc("/", defaultFn)
	log.Fatal(http.ListenAndServe(":8818", nil))
	fmt.Fprintf(os.Stdout, "Hello\n")
}
