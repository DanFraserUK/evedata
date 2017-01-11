package views

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/antihax/evedata/appContext"
	"github.com/antihax/evedata/esi"
	"github.com/antihax/evedata/server"
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

	tok, err := c.ESIBootstrapAuthenticator.TokenExchange(code)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Failed Token Exchange")
	}

	tokSrc, err := c.ESIBootstrapAuthenticator.TokenSource(tok)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	auth := context.WithValue(context.TODO(), esi.ContextOAuth2, tokSrc.Token)
	_, err = c.EVE.Verify(auth)
	if err != nil {
		return http.StatusInternalServerError, err
	}

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
