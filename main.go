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
	fmt.Println(" 0 1 2 3 4 5 6")
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print("|", g.tableau[i][j])
		}
		fmt.Println("|")
	}
	fmt.Println("---------------")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (g *Game) AddJeton(colonne int) bool {
	if colonne < 0 || colonne >= 7 {
		return false
	}

	for i := 5; i >= 0; i-- {
		if g.tableau[i][colonne] == " " {
			g.tableau[i][colonne] = g.currentPlayer
			g.turnCount++
			return true
		}
	}
	return false
}

func verifierVictoire(g Game) bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if g.tableau[i][j] == g.tableau[i][j+1] && g.tableau[i][j] == g.tableau[i][j+2] && g.tableau[i][j] == g.tableau[i][j+3] {
				return true
			}
		}
	}
	for j := 0; j < 7; j++ {
		for i := 0; i < 3; i++ {
			if g.tableau[i][j] == g.tableau[i+1][j] && g.tableau[i][j] == g.tableau[i+2][j] && g.tableau[i][j] == g.tableau[i+3][j] {
				return true
			}
		}
	}
	return false
}

func verifierMatchNul(g Game) {
	if g.turnCount >= 42 {
		g.reset()
	}
}

func (g *Game) reset() {
	clearScreen()
	fmt.Println("========== REINITIALISATION DU JEU ===========")
	fmt.Println("Le jeu a été réinitialisé.")
	fmt.Println("Appuie sur Entrée pour continuer...")
	fmt.Scanln()
	*g = initGame()
}
func main() {
	game := initGame()
	game.afficherTableau()
}
