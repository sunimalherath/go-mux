package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize method takes the data required to connect to the database
func (a *App) Initialize(user, password, dbname string) {

}

// Run method starts the application
func (a *App) Run(addr string) {

}
