package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/news/{title}/topic/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		topic := vars["topic"]
		fmt.Fprintf(w, "You've requested the news: %s and the topic %s\n", title, topic)
	})

	http.ListenAndServe(":80", r)
}
