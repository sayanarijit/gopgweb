package app

/*
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopgweb#gopgweb
*/

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sayanarijit/gopassgen"
)

// Index page
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// GenerateAPI endpoint
func GenerateAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	p := gopassgen.NewPolicy()

	for k, v := range r.Form {
		switch k {
		case "MaxLength":
			p.MaxLength, _ = strconv.Atoi(v[0])
		case "MinLength":
			p.MinLength, _ = strconv.Atoi(v[0])
		case "MinUppers":
			p.MinUppers, _ = strconv.Atoi(v[0])
		case "MinLowers":
			p.MinLowers, _ = strconv.Atoi(v[0])
		case "MinDigits":
			p.MinDigits, _ = strconv.Atoi(v[0])
		case "MinSpclChars":
			p.MinSpclChars, _ = strconv.Atoi(v[0])
		case "UpperPool":
			p.UpperPool = v[0]
		case "LowerPool":
			p.LowerPool = v[0]
		case "DigitPool":
			p.DigitPool = v[0]
		case "SpclCharPool":
			p.SpclCharPool = v[0]
		}
	}

	resp := map[string]string{}
	password, err := gopassgen.Generate(p)
	if err != nil {
		resp["error"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp["result"] = password
	}
	jsonResp, _ := json.Marshal(resp)

	w.Write(jsonResp)
}

// Static files
func Static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "."+r.URL.Path)
}
