package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"algo-schedule/internal/db"

	"github.com/lmittmann/tint"
	_ "modernc.org/sqlite"
)

func registerRoutes(qry *db.Queries, dev bool) {
	router := Router{qry: qry}
	router.Attach(http.DefaultServeMux)

	if dev {
		slog.Debug("-dev is enabled! serving assets from local folder")
		fileserver := http.FileServer(http.FS(os.DirFS("./static")))
		handler := http.StripPrefix("/static/", fileserver)
		http.Handle("/static/{path...}", handler)
		return
	}
	http.Handle("/static/{path...}", http.FileServer(http.FS(static)))
}

func setupSlog(verbose bool) {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level: level,
		}),
	))
}

func main() {
	dbpath := flag.String("db", "database.db", "path to the database to use")
	dev := flag.Bool("dev", false, "enable developer mode. (note: this will automatically enable verbose logging)")
	verbose := flag.Bool("v", false, "enable verbose logging.")
	flag.Parse()

	setupSlog(*dev || *verbose)

	database, err := OpenAndMigrateDB(db.Schema, *dbpath)
	if err != nil {
		slog.Error("open db", "err", err)
		os.Exit(1)
	}
	defer database.Close()
	qry := db.New(database)

	registerRoutes(qry, *dev)

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
