package ppt

import (
	"math/rand"
	"strconv"
)

// Lógica de juego

const (
	PIEDRA = 0
	PAPEL  = 1
	TIJERA = 2
)

// Estructura del resultado de cada ronda + dato de tipo json
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Bien hecho!",
	"¡Buen Trabajo!",
	"Increible ganaste",
}

var loseMessages = []string{
	"Que lastima!",
	"¡Intentalo de nuevo!",
	"Hoy no es tu día amigo",
}

var drawMessages = []string{
	"Empataste contra la mejor de las IAs",
	"Empate, intentalo de nuevo",
	"Esta vez, nadie gano",
}

// Variables para definir los puntajes
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	//Variables para almacenar lo que eligio la computadora
	var ComputerChoice, roundResult string
	var computerChoiceInt int

	// Vamos devolver un mensaje en base a lo que la computadora eligió
	switch computerValue {
	case PIEDRA:
		computerChoiceInt = PIEDRA
		ComputerChoice = "La computadora eligió piedra"

	case PAPEL:
		computerChoiceInt = PAPEL
		ComputerChoice = "La computadora eligió papel"

	case TIJERA:
		computerChoiceInt = TIJERA
		ComputerChoice = "La computadora eligió tijera"
	}

	// También para no estar generando siempre los mismo mensajes de resultados generamos un numero aleatorio
	messageRoundInt := rand.Intn(3)

	// declarar una variable que contenga el mensaje
	var messageValue string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		messageValue = drawMessages[messageRoundInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana!"

		messageValue = winMessages[messageRoundInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana!"

		messageValue = loseMessages[messageRoundInt]
	}

	//retornar resultado de tipo round
	return Round{
		Message:           messageValue,
		ComputerChoice:    ComputerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
