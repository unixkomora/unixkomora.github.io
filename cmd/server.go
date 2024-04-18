package cmd

import (
	"html/template"
	"log"
	"net/http"
)

const (
	local = "http://localhost"
	port  = ":8000"

	Bold   = "\033[1m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func loadHtml(name string, html string, data interface{}) {
	log.Print(Yellow + "Loading: " + Reset + html + ".html...")
	http.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./web/" + html + ".html"))
		tmpl.Execute(w, data)
	})
}

func Server() {
	log.Print(Yellow+"Server started: "+Reset, local, port)

	text := map[string]interface{}{
		//Tags
		"Code":     "код",
		"Telegram": "телеграм",
		"SubMain":  "головна",
		"SubGoal":  "мета",
		// Text
		"SiteName": "Комора Лінуксоїда",
		"Slogan":   "ok",
	}

	loadHtml("", "index", text)

	log.Print(Yellow + "Loading: " + Reset + "/static/...")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	log.Print(Yellow + "Server listening on: " + Reset + local + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
