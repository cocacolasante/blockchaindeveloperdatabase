package application

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	repository "github.com/cocacolasante/blockchaindeveloperdatabase/internal/database"
	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/web3"
)

type Application struct {
	Port     int
	Domain   string
	DSN      string
	DB       repository.DatabaseRepo
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Web3 web3.Web3Connect
}

func NewApplication(domain, dsn string, port int) *Application {
	return &Application{
		Port:   port,
		Domain: domain,
		DSN:    dsn,
	}
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)

	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func (app *Application) ConnectToDb() (*sql.DB, error) {
	conn, err := openDb(app.DSN)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Database")

	return conn, nil
}

func (app *Application) Start() {
	app.InfoLog.Printf("Connecting to domain %s at port %d\n", app.Domain, app.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())
	if err != nil {
		log.Fatal(err)
	}

}
