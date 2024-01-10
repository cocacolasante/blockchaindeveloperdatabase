package main

import (
	"flag"
	"log"
	"os"

	"github.com/cocacolasante/blockchaindeveloperdatabase/application"
	"github.com/cocacolasante/blockchaindeveloperdatabase/initialize"
	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/database/postgresrepo"
	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/web3"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	dsn := flag.String("dsn", "host=localhost port=5432 user=postgres password=postgres dbname=bcdevdatabase sslmode=disable timezone=utc connect_timeout=5", "db dsn")
	domain := flag.String("domain", "localhost", "domain url")
	port := flag.Int("port", 8080, "port number")

	flag.Parse()
	initialize.Init()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application.NewApplication(*domain, *dsn, *port)
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	conn, err := app.ConnectToDb()
	if err != nil {
		app.ErrorLog.Fatalf("cannot connect to db %s\n", err)
	}

	app.DB = &postgresrepo.PostgresDb{Db: conn}

	defer app.DB.Connection().Close()

	app.InfoLog.Printf("Connecting to blockchain")
	web3Conn, err := web3.GetClient()
	if err != nil {
		app.ErrorLog.Println("Unable to connect to blockchain")
		return
	}
	app.Web3.Client = web3Conn
	defer app.Web3.Client.Close()
	app.InfoLog.Println("connected to blockchain")

	m := app.CreateMail()
	app.Mailer = &m
	

	app.InfoLog.Printf("Application configured\n")
	app.InfoLog.Printf("starting application\n")
	app.Start()

}
