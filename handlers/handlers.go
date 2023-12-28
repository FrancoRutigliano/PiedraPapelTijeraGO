package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/FrancoRutigliano/ppt"
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
	// Obtener el valor de "c" de los parámetros de la URL y convertirlo a un entero
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))

	// Llamar a la función PlayRound con la opción del jugador
	result := ppt.PlayRound(playerChoice)

	// Convertir el resultado a formato JSON con formato legible
	out, err := json.MarshalIndent(result, "", "   ")
	// manejo de error si sucede algo con la serializacion del json
	if err != nil {
		log.Println(err)
		return
	}

	// Establecer el encabezado de la respuesta HTTP para indicar que el contenido es de tipo JSON
	w.Header().Set("Content-Type", "application-json")

	// Escribir la respuesta en formato JSON en el cuerpo de la respuesta HTTP
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValues()
	RenderTemplate(w, "about.html", nil)
}

// Funcion Generica que recibe el response, la plantilla base y la plantilla que va a renderizar
func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	// renderizamos en el navegador
	if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// Funcion para reiniciar valores

func restartValues() {
	player.Name = ""
}
