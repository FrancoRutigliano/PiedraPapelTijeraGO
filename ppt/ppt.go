package ppt

// Lógica de juego

const (
	PIEDRA = 0
	PAPEL  = 1
	TIJERA = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}
