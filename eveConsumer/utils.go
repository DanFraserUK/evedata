package eveConsumer

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/antihax/evedata/models"
	"github.com/antihax/goesi"
	"golang.org/x/oauth2"
)

// Obtain an authenticated client from a stored access/refresh token.
func (c *EVEConsumer) getToken(characterID int64, tokenCharacterID int64) (oauth2.TokenSource, error) {
	tok, err := models.GetCRESTToken(characterID, tokenCharacterID)
	if err != nil {
		return nil, err
	}

	token := &goesi.CRESTToken{Expiry: tok.Expiry, AccessToken: tok.AccessToken, RefreshToken: tok.RefreshToken, TokenType: tok.TokenType}
	n, err := c.ctx.TokenAuthenticator.TokenSource(token)

	return n, err
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func tokenError(cid int64, tcid int64, r *http.Response, err error) {
	if r != nil {
		req, _ := httputil.DumpRequest(r.Request, true)
		res, _ := httputil.DumpResponse(r, true)
		e := models.SetTokenError(cid, tcid, r.StatusCode, r.Status, req, res)
		if e != nil {
			log.Println(e)
		}
	} else {
		e := models.SetTokenError(cid, tcid, 999, err.Error(), []byte{}, []byte{})
		if e != nil {
			log.Println(e)
		}
	}
}

func tokenSuccess(cid int64, tcid int64, code int, status string) {
	models.SetTokenError(cid, tcid, code, status, []byte{}, []byte{})
}
