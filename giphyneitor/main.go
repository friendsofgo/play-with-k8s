package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/friendsofgo/play-with-k8s/giphyneitor/giphy"
	"github.com/gorilla/mux"
)

const (
	port = 8080
)

func main () {
	giphyAPIKey := os.Getenv("GIPHY_API_KEY")

	httpAddr := fmt.Sprintf(":%d", + port)
	r := mux.NewRouter()
	r.HandleFunc("/", giphyHandler(giphyAPIKey))
	http.Handle("/", r)

	log.Printf("giphyneitor server tap on http://localhost%s...\n", httpAddr)
	if err := http.ListenAndServe(httpAddr, nil); err != nil {
		panic(err)
	}
}

type gif struct {
	Name string
	URL string
}

func giphyHandler(giphyAPIKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g := giphy.NewClient(giphyAPIKey)
		gif, err := g.RandomGif()
		if err != nil {
			log.Fatal(err)
		}

		t, err := template.ParseFiles("templates/single.html")
		if err != nil {
			log.Fatalf("error parsing template files: %v", err)
		}

		err = t.Execute(w, gif)
		if err != nil {
			log.Fatalf("error trying to process the template: %v", err)
		}
		w.Header().Set("Content-Type", "text/html")
	}
}