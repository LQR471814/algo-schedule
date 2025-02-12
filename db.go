package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

func OpenAndMigrateDB(schema, path string) (*sql.DB, error) {
	// to ensure that the db actually exists
	db, err := openDB(path)
	if err != nil {
		return nil, err
	}
	err = db.Close()
	if err != nil {
		return nil, err
	}

	_, err = exec.LookPath("atlas")
	if os.IsNotExist(err) {
		return db, fmt.Errorf(
			"could not find 'atlas' executable on path, is it installed? skipping migrations...",
		)
	}

	err = os.WriteFile("temp_migration_schema.sql", []byte(schema), 0666)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = os.Remove("temp_migration_schema.sql")
		if err != nil {
			slog.Warn("could not delete temp_migration_schema.sql", "err", err)
		}
	}()

	dbUrl := url.URL{
		Scheme: "sqlite",
		Path:   path,
	}
	cmd := exec.Command(
		"atlas", "schema", "apply",
		"--auto-approve",
		"--url", dbUrl.String(),
		"--to", "file://temp_migration_schema.sql",
		"--dev-url", "sqlite://file?mode=memory",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(string(out))
		return nil, err
	}

	return openDB(path)
}

func openDB(path string) (*sql.DB, error) {
	if path != ":memory:" {
		os.MkdirAll(filepath.Dir(path), 0777)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// see this stackoverflow post for information on why the following
	// lines exist: https://stackoverflow.com/questions/35804884/sqlite-concurrent-writing-performance
	db.SetMaxOpenConns(1)
	_, err = db.Exec("PRAGMA journal_mode=WAL")
	if err != nil {
		return nil, err
	}

	return db, nil
}
