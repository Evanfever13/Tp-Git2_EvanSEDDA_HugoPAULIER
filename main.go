package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	player1       string
	player2       string
	tableau       [6][7]string
	currentPlayer string
	turnCount     int
}

func initGame() Game {
	clearScreen()

	fmt.Println("========== BIENVENUE DANS LE PUISSANCE 4 ===========")
	fmt.Println("CE SOIR, C'EST JEU !")
	fmt.Println("LE BUT DU JEU EST D'ALIGNER 4 JETONS DE SA COULEUR")
	fmt.Println("LE JOUEUR 1 AURA LES X, LE JOUEUR 2 AURA LES O")
	fmt.Println("Appuie sur Entrée pour continuer...")
	fmt.Scanln()

	// La grille vide
	var tableau [6][7]string
	for i := range tableau {
		for j := range tableau[i] {
			tableau[i][j] = " "
		}
	}

	// Tirage aléatoire
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(2)

	var player1, player2, current string
	if n == 0 {
		player1, player2 = "X", "O"
		current = player1
		fmt.Println("LE JOUEUR 1 COMMENCE en PREMIER")
	} else {
		player1, player2 = "O", "X"
		current = player2
		fmt.Println("LE JOUEUR 2 COMMENCE en PREMIER")
	}

	return Game{
		player1:       player1,
		player2:       player2,
		tableau:       tableau,
		currentPlayer: current,
		turnCount:     0,
	}
}

func (g *Game) afficherTableau() {
	fmt.Println("Joueur 1 :", g.player1)
	fmt.Println("Joueur 2 :", g.player2)
	fmt.Println("Joueur courant :", g.currentPlayer)
	fmt.Println("====================================================")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	game := initGame()
	game.afficherTableau()
}
