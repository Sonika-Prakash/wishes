package base

import (
	"fmt"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

// IndexHandler is the handler to display the webpage
func (a *Application) IndexHandler(w http.ResponseWriter, r *http.Request) {
	name := getCapitalizedName(r.URL.Query().Get("n"))
	// TODO: check for empty names
	vars := make(jet.VarMap)
	vars.Set("name", name)
	vars.Set("nextLink", "")
	err := a.render(w, r, "index", vars)
	if err != nil {
		a.errLog.Println("error while rendering index page", err)
		a.serverErr(w, err)
		return
	}
}

// IndexPostHandler generates the link to share
func (a *Application) IndexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1024*2)

	err := r.ParseForm()
	if err != nil {
		a.errLog.Println(err)
		a.serverErr(w, err)
		return
	}

	f := New(r.PostForm)
	if !f.Valid() {
		vars := make(jet.VarMap)
		capName := getCapitalizedName(f.Get("prevName"))
		vars.Set("name", capName)
		vars.Set("nextLink", "")
		err := a.render(w, r, "index", vars)
		if err != nil {
			a.errLog.Println(err)
			a.serverErr(w, err)
			return
		}
	}

	name := getCapitalizedName(f.Get("name"))
	nextLink := a.generateURL(name)

	vars := make(jet.VarMap)
	vars.Set("nextLink", nextLink)
	if name == "" {
		fmt.Println("using prev name")
		prevName := getCapitalizedName(f.Get("prevName"))
		vars.Set("name", prevName)
	} else {
		fmt.Println("using new name")
		vars.Set("name", name)
	}
	err = a.render(w, r, "index", vars)
	if err != nil {
		a.errLog.Println(err)
		a.serverErr(w, err)
		return
	}
}
