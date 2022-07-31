package main

import (
	"database/sql"
	"flag"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func openDB(dbt, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbt, dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	app := newDefaultApplication()

	pathToCfg := flag.String("cfg", "./conf.json", "Configuration file path.")
	flag.Parse()

	cfg := newConfig()
	cfg.loadConfigFromPath(*pathToCfg)

	app.infoLogger.Printf("Starting server on %s", cfg.getAddressAndPort())

	db, err := openDB(cfg.DBType, cfg.getDatabaseConnStr())
	if err != nil {
		app.errLogger.Println("Failed to establish connection with database.", err)
	}
	defer db.Close()

	app.setDatabase(db)

	hsrv := &http.Server{
		Addr:     cfg.getAddressAndPort(),
		Handler:  app.routes(cfg.StaticDir),
		ErrorLog: app.errLogger,
	}

	err = hsrv.ListenAndServe()
	app.errLogger.Fatal(err)
}
