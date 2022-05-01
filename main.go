package main

import (
	"io"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	checkErr := func(err error) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	r, err := http.NewRequest(r.Method, "https://go.dev"+r.URL.RequestURI(), nil)
	checkErr(err)

	client := http.DefaultClient
	resp, err := client.Do(r)
	checkErr(err)

	if resp != nil {
		w.WriteHeader(resp.StatusCode)
		_, err := io.Copy(w, resp.Body)
		checkErr(err)
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
