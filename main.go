package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	c "gitlab.com/TheChertila/REST_API_TV_Market/controller"
)

func returnTVList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		c.GetTVsList("any", w)

	case "POST":
		c.AddTV(w, r)

	default:
		fmt.Fprintf(w, "Only GET and POST are supported")
	}

}

func returnTV(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/tvs/")
	switch r.Method {

	case "GET":
		c.GetTVsList(id, w)

	case "DELETE":
		c.RemoveTV(id, w)

	default:
		fmt.Fprint(w, "Only DET and DELETE are supported")
	}
}

func handleRequests() {
	http.HandleFunc("/api/tvs", returnTVList)
	http.HandleFunc("/api/tvs/", returnTV)
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}

func main() {
	handleRequests()
}
