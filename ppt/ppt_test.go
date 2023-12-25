package ppt

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		switch i {
		case 0:
			fmt.Println("El usuario eligió PIEDRA")
		case 1:
			fmt.Println("El usuario eligió PAPEL")
		case 2:
			fmt.Println("El usuario eligió Tijeras")
		}

		fmt.Printf("Mensaje -> %s \n", round.Message)
		fmt.Printf("Elección de la computadora -> %s \n", round.ComputerChoice)
		fmt.Printf("Resultado de la ronda -> %s \n", round.RoundResult)
		fmt.Printf("Elección de la computadora en int -> %d \n", round.ComputerChoiceInt)
		fmt.Printf("Puntaje de la computadora -> %s \n", round.ComputerScore)
		fmt.Printf("Puntaje del jugador -> %s \n", round.PlayerScore)

		fmt.Println("\n======================================================")
	}

}
