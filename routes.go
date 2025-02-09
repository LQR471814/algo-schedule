package main

import (
	"algo-schedule/components"
	"algo-schedule/db"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"embed"
)

//go:embed static
var static embed.FS

func writeError(status int, err error, w http.ResponseWriter) {
	w.WriteHeader(status)
	err = fmt.Errorf("status %d: %w", status, err)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		panic(err)
	}
}

func registerRoutes(qry *db.Queries) {
	http.Handle("/static", http.FileServer(http.FS(static)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tasks, err := qry.ListTasks(ctx)
		if err != nil {
			writeError(500, err, w)
			return
		}
		components.Root(components.Dashboard(tasks)).Render(ctx, w)
	})

	http.HandleFunc("/edit_task/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeError(405, fmt.Errorf("unsupported method: %s", r.Method), w)
			return
		}
		ctx := r.Context()
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			writeError(400, fmt.Errorf("invalid id: %w", err), w)
			return
		}
		task, err := qry.ReadTask(ctx, id)
		if err != nil {
			writeError(500, err, w)
			return
		}
		components.EditTask(task).Render(ctx, w)
	})

	http.HandleFunc("/apply_task_edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeError(405, fmt.Errorf("unsupported method: %s", r.Method), w)
			return
		}
		ctx := r.Context()
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			writeError(400, fmt.Errorf("invalid id: %w", err), w)
			return
		}
		err = r.ParseForm()
		if err != nil {
			writeError(400, fmt.Errorf("invalid form data: %w", err), w)
			return
		}

		var size int64
		switch r.PostForm.Get("size") {
		case "small":
			size = 0
		case "medium":
			size = 1
		}

		var challenge int64
		switch r.PostForm.Get("challenge") {
		case "easy":
			challenge = 0
		case "medium":
			challenge = 1
		case "hard":
			challenge = 2
		}

		deadline, err := time.Parse(time.DateOnly, r.PostForm.Get("deadline"))
		if err != nil {
			writeError(400, fmt.Errorf("invalid deadline: %w", err), w)
			return
		}

		err = qry.UpdateTask(ctx, db.UpdateTaskParams{
			ID:        id,
			Name:      r.PostForm.Get("name"),
			Size:      size,
			Challenge: challenge,
			Deadline:  deadline,
		})
		if err != nil {
			writeError(500, err, w)
			return
		}
		task, err := qry.ReadTask(ctx, id)
		if err != nil {
			writeError(500, err, w)
			return
		}
		components.Task(task).Render(ctx, w)
	})
}
