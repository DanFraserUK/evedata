package views

import (
	"errors"
	"evedata/appContext"
	"evedata/models"
	"evedata/server"
	"evedata/strip"
	"evedata/templates"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func init() {
	evedata.AddRoute("entity", "GET", "/alliance", alliancePage)
	evedata.AddRoute("entity", "GET", "/corporation", corporationPage)
	evedata.AddRoute("entity", "GET", "/character", characterPage)
}

func alliancePage(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {
	setCache(w, 60*60)
	p := newPage(s, r, "Unknown Alliance")

	idStr := r.FormValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Invalid Alliance ID. Please provide an ?id=")
	}

	errc := make(chan error)

	// Get known wars. This takes the longest.
	go func() {
		ref, err := models.GetWarsForEntityByID(id)
		if err != nil {
			errc <- err
			return
		}
		p["Wars"] = ref
		errc <- err
	}()

	// Get known activity.
	go func() {
		ref, err := models.GetConstellationActivity(id, "alliance")
		if err != nil {
			errc <- err
			return
		}
		p["Activity"] = ref
		errc <- err
	}()

	// Get the alliance information
	go func() {
		ref, err := models.GetAlliance(id)
		if err != nil {
			errc <- err
			return
		}
		ref.Description = strip.StripTags(ref.Description)
		p["Alliance"] = ref
		p["Title"] = ref.AllianceName
		errc <- err
	}()

	// Get the alliance asset information
	go func() {
		ref, err := models.GetAllianceAssetsInSpace(id)
		if err != nil {
			errc <- err
			return
		}
		p["Assets"] = ref
		errc <- err
	}()

	// Get the alliance members
	go func() {
		ref, err := models.GetAllianceMembers(id)
		if err != nil {
			errc <- err
			return
		}
		p["AllianceMembers"] = ref
		errc <- err
	}()

	// Get known Ships.
	go func() {
		ref, err := models.GetKnownShipTypes(id, "alliance")
		if err != nil {
			errc <- err
			return
		}
		p["KnownShips"] = ref
		errc <- err
	}()

	// Get known Allies.
	go func() {
		ref, err := models.GetKnownAlliesByID(id)
		if err != nil {
			errc <- err
			return
		}
		p["KnownAllies"] = ref
		errc <- err
	}()

	// clear the error channel
	for i := 0; i < 7; i++ {
		err := <-errc
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	templates.Templates = template.Must(template.ParseFiles("templates/entities.html", templates.LayoutPath))
	err = templates.Templates.ExecuteTemplate(w, "base", p)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func corporationPage(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {
	setCache(w, 60*60)
	p := newPage(s, r, "Unknown Corporation")

	idStr := r.FormValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Invalid corporation ID. Please provide an ?id=")
	}

	errc := make(chan error)

	// Get known wars. This takes the longest.
	go func() {
		ref, err := models.GetWarsForEntityByID(id)
		if err != nil {
			errc <- err
			return
		}
		p["Wars"] = ref
		errc <- nil
	}()

	// Get known activity.
	go func() {
		ref, err := models.GetConstellationActivity(id, "corporation")
		if err != nil {
			errc <- err
			return
		}
		p["Activity"] = ref
		errc <- err
	}()

	// Get the corporation information
	go func() {
		ref, err := models.GetCorporation(id)
		if err != nil {
			errc <- err
			return
		}
		ref.Description = strip.StripTags(ref.Description)
		p["Corporation"] = ref
		p["Title"] = ref.CorporationName
		errc <- nil
	}()

	// Get the alliance asset information
	go func() {
		ref, err := models.GetCorporationAssetsInSpace(id)
		if err != nil {
			errc <- err
			return
		}
		p["Assets"] = ref
		errc <- err
	}()

	// Get known Ships.
	go func() {
		ref, err := models.GetKnownShipTypes(id, "corporation")
		if err != nil {
			errc <- err
			return
		}
		p["KnownShips"] = ref
		errc <- err
	}()

	// Get known Allies.
	go func() {
		ref, err := models.GetKnownAlliesByID(id)
		if err != nil {
			errc <- err
			return
		}
		p["KnownAllies"] = ref
		errc <- err
	}()

	// clear the error channel
	for i := 0; i < 6; i++ {
		err := <-errc
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	templates.Templates = template.Must(template.ParseFiles("templates/entities.html", templates.LayoutPath))
	err = templates.Templates.ExecuteTemplate(w, "base", p)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func characterPage(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {
	setCache(w, 60*60)
	p := newPage(s, r, "Unknown Character")

	idStr := r.FormValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Invalid character ID. Please provide an ?id=")
	}

	errc := make(chan error)

	// Get the character information
	go func() {
		ref, err := models.GetCharacter(id)

		p["Character"] = ref
		p["Title"] = ref.CharacterName
		errc <- err
	}()

	// Get known activity.
	go func() {
		ref, err := models.GetConstellationActivity(id, "character")
		if err != nil {
			errc <- err
			return
		}
		p["Activity"] = ref
		errc <- err
	}()

	// Get known Ships.
	go func() {
		ref, err := models.GetKnownShipTypes(id, "character")
		if err != nil {
			errc <- err
			return
		}
		p["KnownShips"] = ref
		errc <- err
	}()

	// clear the error channel
	for i := 0; i < 3; i++ {
		err := <-errc
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	templates.Templates = template.Must(template.ParseFiles("templates/entities.html", templates.LayoutPath))
	err = templates.Templates.ExecuteTemplate(w, "base", p)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
