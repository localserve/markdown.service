package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/markdown", ProcessMarkdown)
	port := os.Getenv("PORT")
	if port == "" {
		port = "40001"
	}
	address := ":" + port
	http.ListenAndServe(address, nil)
}

func ProcessMarkdown(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		insecure := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
		securer := bluemonday.UGCPolicy().SanitizeBytes(insecure)
		w.Write(securer)
	default:
		fmt.Fprintf(w, "Only POST supported.\n")
	}
}
