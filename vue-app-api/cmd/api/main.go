package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vue-app-api/internal/data"
	"vue-app-api/internal/driver"
)

// config is the type for all application configuration
type config struct {
	port int
}

// application is the type for all data we want to share with the
// various parts of our application. We will share this information in most
// cases by using this type as the receiver for functions
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	// db       *driver.DB		// will be from type data.Models
	models      data.Models
	environment string
}

func main() {

	var cfg config
	cfg.port = 8081
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// data source name
	// dsn := "host=localhost port=5432 user=postgres password=password dbname=vue_app_api sslmode=disable timezone=utc connect_timeout=5"

	// Change to env variable
	dsn := os.Getenv("DSN")
	environment := os.Getenv("ENV")

	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	// When application stops, before it stops, it will close the resource first
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		// db:       db,
		models:      data.New(db.SQL),
		environment: environment,
	}
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}
