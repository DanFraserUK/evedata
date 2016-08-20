package views

import (
	"encoding/json"
	"errors"
	"evedata/appContext"
	"evedata/eveapi"
	"evedata/models"
	"evedata/server"
	"evedata/templates"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func init() {
	evedata.AddRoute("account", "GET", "/account", accountPage)

	evedata.AddRoute("apiKeys", "GET", "/U/apiKeys", apiGetKeys)
	evedata.AddRoute("apiKeys", "DELETE", "/U/apiKeys", apiDeleteKey)
	evedata.AddRoute("apiKeys", "PUT", "/U/apiKeys", apiAddKey)

	evedata.AddRoute("crestTokens", "GET", "/U/crestTokens", apiGetCRESTTokens)
	evedata.AddRoute("crestTokens", "DELETE", "/U/crestTokens", apiDeleteCRESTToken)
}

func accountPage(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {

	p := newPage(s, r, "Account Information")
	templates.Templates = template.Must(template.ParseFiles("templates/account.html", templates.LayoutPath))

	if err := templates.Templates.ExecuteTemplate(w, "base", p); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func apiGetKeys(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {
	characterID := s.Values["characterID"].(int64)
	keys, err := models.GetAPIKeys(characterID)
	if err != nil {
		return http.StatusNotFound, err
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(keys)

	return 200, nil
}

func apiDeleteKey(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {

	keyID, err := strconv.Atoi(r.FormValue("keyID"))
	if err != nil {
		return http.StatusNotFound, errors.New("Invalid keyID")
	}
	characterID := s.Values["characterID"].(int64)
	if err := models.DeleteApiKey(characterID, keyID); err != nil {
		return http.StatusConflict, err
	}

	return 200, nil
}

func apiAddKey(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {

	type localApiKey struct {
		KeyID string
		VCode string
	}
	var key localApiKey

	if r.Body == nil {
		return http.StatusNotFound, errors.New("No Data Received")
	}
	err := json.NewDecoder(r.Body).Decode(&key)
	if err != nil {
		return http.StatusNotFound, err
	}
	keyID, err := strconv.Atoi(key.KeyID)
	if err != nil {
		return http.StatusNotFound, errors.New("Invalid keyID")
	}

	if eveapi.IsValidVCode(key.VCode) == false {
		return http.StatusConflict, errors.New("Invalid vCode")
	}

	if s.Values["characterID"] == nil {
		return http.StatusForbidden, nil
	}
	characterID := s.Values["characterID"].(int64)
	if err := models.AddApiKey(characterID, keyID, key.VCode); err != nil {
		return http.StatusConflict, err
	}

	return 200, nil
}

func apiGetCRESTTokens(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {
	if s.Values["characterID"] == nil {
		return http.StatusForbidden, nil
	}
	characterID := s.Values["characterID"].(int64)

	tokens, err := models.GetCRESTTokens(characterID)
	if err != nil {
		return http.StatusNotFound, err
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(tokens)

	return 200, nil
}

func apiDeleteCRESTToken(c *appContext.AppContext, w http.ResponseWriter, r *http.Request, s *sessions.Session) (int, error) {

	cid, err := strconv.Atoi(r.FormValue("tokenCharacterID"))
	if err != nil {
		return http.StatusNotFound, errors.New("Invalid tokenCharacterID")
	}
	characterID := s.Values["characterID"].(int64)
	if err := models.DeleteCRESTToken(characterID, cid); err != nil {
		return http.StatusConflict, err
	}

	return 200, nil
}
