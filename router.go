package main

import (
	"algo-schedule/db"
	"algo-schedule/templates"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"embed"
)

//go:embed static
var static embed.FS

type Router struct {
	qry *db.Queries
}

func (router Router) Attach(mux *http.ServeMux) {
	mux.HandleFunc("/", router.Root)
	mux.HandleFunc("/create_task", router.CreateTask)
	mux.HandleFunc("/start_edit_task/{id}", router.StartEditTask)
	mux.HandleFunc("/end_edit_task/{id}", router.EndEditTask)
	mux.HandleFunc("/delete_task/{id}", router.DeleteTask)
}

func (router Router) Root(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := router.qry.ListTasks(ctx)
	if err != nil {
		writeError(500, err, w)
		return
	}
	templates.Root(templates.Dashboard(tasks)).Render(ctx, w)
}

func parseSize(text string) (db.Size, error) {
	switch text {
	case "small":
		return db.SIZE_SMALL, nil
	case "medium":
		return db.SIZE_MEDIUM, nil
	}
	return -1, fmt.Errorf("unknown size '%s'", text)
}

func parseChallenge(text string) (db.Challenge, error) {
	switch text {
	case "easy":
		return db.CHALLENGE_EASY, nil
	case "medium":
		return db.CHALLENGE_MEDIUM, nil
	case "hard":
		return db.CHALLENGE_HARD, nil
	}
	return -1, fmt.Errorf("unknown challenge '%s'", text)
}

func parseTask(form url.Values) (name, description string, size db.Size, challenge db.Challenge, deadline time.Time, err error) {
	name = form.Get("name")
	if name == "" {
		err = fmt.Errorf("required field: name")
		return
	}
	description = form.Get("description")
	size, err = parseSize(form.Get("size"))
	if err != nil {
		err = fmt.Errorf("invalid size: %w", err)
		return
	}
	challenge, err = parseChallenge(form.Get("challenge"))
	if err != nil {
		err = fmt.Errorf("invalid challenge: %w", err)
		return
	}
	deadline, err = time.Parse(time.DateOnly, form.Get("deadline"))
	if err != nil {
		err = fmt.Errorf("invalid deadline: %w", err)
		return
	}
	deadline = time.Date(
		deadline.Year(),
		deadline.Month(),
		deadline.Day(),
		23,
		59,
		59,
		999,
		time.Local,
	)
	return
}

func (router Router) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(405, fmt.Errorf("unsupported method: %s", r.Method), w)
		return
	}

	ctx := r.Context()

	err := r.ParseForm()
	if err != nil {
		writeError(400, fmt.Errorf("invalid form data: %w", err), w)
		return
	}
	name, description, size, challenge, deadline, err := parseTask(r.Form)
	if err != nil {
		writeError(400, err, w)
		return
	}

	id, err := router.qry.CreateTask(r.Context(), db.CreateTaskParams{
		Name:        name,
		Size:        size,
		Challenge:   challenge,
		Deadline:    deadline,
		Description: description,
	})
	if err != nil {
		writeError(500, err, w)
		return
	}

	templates.AfterCreateTask(db.Task{
		ID:          id,
		Name:        name,
		Size:        size,
		Challenge:   challenge,
		Deadline:    deadline,
		Description: description,
	}).Render(ctx, w)
}

func (router Router) StartEditTask(w http.ResponseWriter, r *http.Request) {
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
	task, err := router.qry.ReadTask(ctx, id)
	if err != nil {
		writeError(500, err, w)
		return
	}
	templates.EditTask(task).Render(ctx, w)
}

func (router Router) EndEditTask(w http.ResponseWriter, r *http.Request) {
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
	name, description, size, challenge, deadline, err := parseTask(r.Form)
	if err != nil {
		writeError(400, err, w)
		return
	}

	err = router.qry.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          id,
		Name:        name,
		Size:        size,
		Challenge:   challenge,
		Deadline:    deadline,
		Description: description,
	})
	if err != nil {
		writeError(500, err, w)
		return
	}
	task, err := router.qry.ReadTask(ctx, id)
	if err != nil {
		writeError(500, err, w)
		return
	}
	templates.Task(task).Render(ctx, w)
}

func (router Router) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeError(405, fmt.Errorf("unsupported method: %s", r.Method), w)
		return
	}

	ctx := r.Context()
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(400, fmt.Errorf("invalid id: %w", err), w)
		return
	}

	err = router.qry.DeleteTask(ctx, id)
	if err != nil {
		writeError(500, err, w)
		return
	}
	w.Write([]byte(""))
}

func writeError(status int, err error, w http.ResponseWriter) {
	w.WriteHeader(status)
	err = fmt.Errorf("status %d: %w", status, err)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		panic(err)
	}
}
