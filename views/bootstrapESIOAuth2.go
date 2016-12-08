package views

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"evedata/appContext"
	"evedata/server"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func init() {
	evedata.AddRoute("boostrap", "GET", "/boostrapEveAuth", boostrapEveSSO)
	evedata.AddRoute("boostrap", "GET", "/boostrapEveSSOAnswer", boostrapEveSSOAnswer)
}

func boostrapEveSSO(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	s.Values["BOOTSTRAPstate"] = state

	err := s.Save(r, w)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	url := c.ESIBootstrapAuthenticator.AuthorizeURL(state, true)
	http.Redirect(w, r, url, 302)
	return http.StatusMovedPermanently, nil
}

func boostrapEveSSOAnswer(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {

	code := r.FormValue("code")
	state := r.FormValue("state")

	if s.Values["BOOTSTRAPstate"] != state {

		return http.StatusInternalServerError, errors.New("Invalid State. It is possible that the session cookie is missing. Stop eating the cookies!")
	}

	tok, err := c.ESIBootstrapAuthenticator.TokenExchange(c.HTTPClient, code)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Failed Token Exchange")
	}

	cli := c.ESIBootstrapAuthenticator.GetClientFromToken(c.HTTPClient, tok)
	_, err = cli.Verify()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	s.Values["BOOTSTRAP"] = tok

	fmt.Fprintf(w, "%+v\n", tok)

	err = s.Save(r, w)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
