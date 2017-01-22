package evedata

import (
	"crypto/rand"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/antihax/evedata/appContext"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type appFunc func(*appContext.AppContext, http.ResponseWriter, *http.Request, *sessions.Session) (int, error)
type appHandler struct {
	*appContext.AppContext
	h appFunc
}

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc appFunc
}

var availableRoutes []route
var notFoundHandler *route

func init() {
	// report correct
	mime.AddExtensionType(".svg", "image/svg+xml")
}

func AddRoute(name string, method string, pattern string, handlerFunc appFunc) {
	availableRoutes = append(availableRoutes, route{name, method, pattern, handlerFunc})
}

func AddNotFoundHandler(handlerFunc appFunc) {
	notFoundHandler = &route{"404", "GET", "", handlerFunc}
}

func (a appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	redisCon := ctx.Cache.Get()
	defer redisCon.Close()

	// Make a random hash to store the time to redis
	b := make([]byte, 32)
	rand.Read(b)

	redisCon.Do("ZADD", "EVEDATA_HTTPRequest", time.Now().UTC().Unix(), b)
	s, _ := a.AppContext.Store.Get(r, "session")

	status, err := a.h(a.AppContext, w, r, s)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		http.Error(w, err.Error(), status)
		return
	}
}

func NewRouter(ctx *appContext.AppContext) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	for _, route := range availableRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(appHandler{ctx, route.HandlerFunc})
	}

	// prometheus handler
	router.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("static/css"))))

	router.PathPrefix("/i/").Handler(http.StripPrefix("/i/",
		http.FileServer(http.Dir("static/i"))))

	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/",
		http.FileServer(http.Dir("static/images"))))

	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
		http.FileServer(http.Dir("static/js"))))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/",
		http.FileServer(http.Dir("static/fonts"))))

	if notFoundHandler != nil {
		router.NotFoundHandler = appHandler{ctx, notFoundHandler.HandlerFunc}
	}

	return router
}

const ContextKey int = 0
