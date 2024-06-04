package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rockpaperscissors/rockpaperscissors"
	"strconv"
)

const (
	templateDir  = "templates/"
	baseTemplate = templateDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderTemplate(w, "new-game.html", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rockpaperscissors.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(out)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
		}
		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
		return
	}

	renderTemplate(w, "game.html", player)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValues()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, templateData any) {
	t := template.Must(template.ParseFiles(baseTemplate, templateDir+page))

	err := t.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func restartValues() {
	player.Name = ""
	rockpaperscissors.PlayerScore = 0
	rockpaperscissors.ComputerScore = 0
}
