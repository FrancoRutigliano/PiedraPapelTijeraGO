package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	// constante que busca carpeta template
	templateDir = "template/"
	// const que busca base.html
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "game.hml", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "game.html", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html", nil)
}

// Funcion Generica que recibe el response, la plantilla base y la plantilla que va a renderizar
func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	// renderizamos en el navegador
	if err := tpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
