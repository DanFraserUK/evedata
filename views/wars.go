package views

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/antihax/evedata/models"
	"github.com/antihax/evedata/server"
	"github.com/antihax/evedata/templates"
)

func init() {
	evedata.AddRoute("wars", "GET", "/activeWars", activeWarsPage)
	evedata.AddRoute("wars", "GET", "/J/activeWars", activeWars)
}

func activeWarsPage(w http.ResponseWriter, r *http.Request) {
	p := newPage(r, "Active Wars")

	templates.Templates = template.Must(template.ParseFiles("templates/wars.html", templates.LayoutPath))
	err := templates.Templates.ExecuteTemplate(w, "base", p)

	if err != nil {
		httpErr(w, err)
		return
	}
}

func activeWars(w http.ResponseWriter, r *http.Request) {
	setCache(w, 60*60)
	v, err := models.GetActiveWarList()
	if err != nil {
		httpErr(w, err)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(v)
}
