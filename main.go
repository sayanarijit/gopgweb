package main

/*
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopgweb#gopgweb
*/

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sayanarijit/gopassgen"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/static/", static)
	http.HandleFunc("/api/generate/", generateApi)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func generateApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := gopassgen.NewPolicy()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password := gopassgen.Generate(p)

	resp := map[string]string{"result": password}

	jsonResp, _ := json.Marshal(resp)

	w.Write(jsonResp)
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "."+r.URL.Path)
}
