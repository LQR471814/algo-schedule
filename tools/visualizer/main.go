package main

import (
	"algo-schedule/internal/templates"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/{path...}", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Root(UI()).Render(r.Context(), w)
	})

	http.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		tbjson := r.Form.Get("timeblocks")
		tbjson = strings.Trim(tbjson, " \t\n")

		var timeblocks []templates.TimeBlock
		err = json.Unmarshal([]byte(tbjson), &timeblocks)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		templates.DayList(time.Local, timeblocks).Render(r.Context(), w)
	})

	log.Println("listening on port 3001...")

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
