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
		default:

		}
	}

}
