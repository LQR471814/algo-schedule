package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"algo-schedule/db"

	_ "modernc.org/sqlite"
)

func main() {
	dbpath := flag.String("db", "database.db", "path to the database to use")
	flag.Parse()

	database, err := OpenAndMigrateDB(db.Schema, *dbpath)
	if err != nil {
		slog.Error("open db", "err", err)
		os.Exit(1)
	}
	defer database.Close()
	qry := db.New(database)

	registerRoutes(qry)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		slog.Info("starting http server", "port", 3000)
		err = http.ListenAndServe(":3000", nil)
		if err != nil {
			slog.Error("start http server", "err", err)
			quit <- os.Interrupt
		}
	}()

	<-quit

	slog.Info("shutting down...")
}
