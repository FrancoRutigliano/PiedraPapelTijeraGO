package main

import (
	"log"
	"net/http"

	"github.com/FrancoRutigliano/handlers"
)

// Router componente que se encarga de dirigir las peticiones http a la funciones handler correspondientes

func main() {
	// crea el enrutador
	router := http.NewServeMux()

	//configuramos la ruta
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)
	port := ":8080"

	log.Printf("Servidor escuchando en: http://localhost%s \n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
