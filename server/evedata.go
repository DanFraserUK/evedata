package evedata

import (
	"database/sql"
	"evedata/config"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)

// appContext provides access to handles throughout the app.
type AppContext struct {
	Conf  *config.Config
	Db    *sqlx.DB
	Store *sessions.Store

	Bridge struct {
		HistoryUpdate *sql.Stmt
		OrderMark     *sql.Stmt
		OrderUpdate   *sql.Stmt
	}
	/*
		templates map[string]*template.Template
		decoder   *schema.Decoder*/
}

func GoServer() {

	var err error

	// Make a new app context.8
	ctx := &AppContext{}

	// Read configuation.
	ctx.Conf, err = config.ReadConfig()

	if err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	// Build Connection Pool
	ctx.Db, err = sqlx.Connect(ctx.Conf.Database.Driver, ctx.Conf.Database.Spec)
	if err != nil {
		log.Fatalf("Cannot build database pool: %v", err)
	}

	// Check we can connect
	err = ctx.Db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	// Allocate the routes
	rtr := NewRouter(ctx)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if ctx.Conf.EMDRCrestBridge.Enabled {
		log.Println("Starting EMDR <- Crest Bridge")
		go goEMDRCrestBridge(ctx)
	}

	log.Println("EveData Listening port 3000...")
	http.ListenAndServe(":3000", context.ClearHandler(rtr))
}
