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

// Estructura de jugador
type Player struct {
	Name string
}

// creando jugador
var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restartValues()
	RenderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValues()
	RenderTemplate(w, "new_game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// Leer los datos del formulario
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parseando el form", http.StatusBadRequest)
			return
		}

		player.Name = r.Form.Get("name")
	}

	// Redirecíon a otra ruta
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	RenderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValues()
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

// Funcion para reiniciar valores

func restartValues() {
	player.Name = ""
}
