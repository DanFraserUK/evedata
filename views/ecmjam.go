package views

import (
	"evedata/appContext"
	"evedata/server"
	"evedata/templates"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

func init() {
	evedata.AddRoute("ecmjam", "GET", "/ecmjam", ecmjamPage)
}

func ecmjamPage(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {
	setCache(w, 60*60)
	p := newPage(s, r, "EVE ECM Jam")
	templates.Templates = template.Must(template.ParseFiles("templates/ecmjam.html", templates.LayoutPath))
	err := templates.Templates.ExecuteTemplate(w, "base", p)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
