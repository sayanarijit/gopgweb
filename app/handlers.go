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
		case "MinCapsAlpha":
			p.MinCapsAlpha, _ = strconv.Atoi(v[0])
		case "MinSmallAlpha":
			p.MinSmallAlpha, _ = strconv.Atoi(v[0])
		case "MinDigits":
			p.MinDigits, _ = strconv.Atoi(v[0])
		case "MinSpclChars":
			p.MinSpclChars, _ = strconv.Atoi(v[0])
		case "CapsAlphaPool":
			p.CapsAlphaPool = v[0]
		case "SmallAlphaPool":
			p.SmallAlphaPool = v[0]
		case "DigitPool":
			p.DigitPool = v[0]
		case "SpclCharPool":
			p.SpclCharPool = v[0]
		}
	}

	password := gopassgen.Generate(p)

	resp := map[string]string{"result": password}
	jsonResp, _ := json.Marshal(resp)

	w.Write(jsonResp)
}

// Static files
func Static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "."+r.URL.Path)
}
