package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	empty = "\U000026AB"  // ⚫
	tree  = "\U0001F333"  // 🌳
	fire  = "\U0001F525"  // 🔥
	ash   = "\U0001F32B "  // 🌫️
	flash  = "\U000026A1" // ⚡
)

func displayArray(arr [][]string) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%s ", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Przyjmuje współrzędne punktu, ustawia ogień dla każdego sąsiada, który jest drzewem
func spreadFire(arr [][]string, i, j int) {
	n := len(arr)

	if i > 0 && arr[i-1][j] == tree {
		arr[i-1][j] = fire
	}
	if i < n-1 && arr[i+1][j] == tree {
		arr[i+1][j] = fire
	}
	if j > 0 && arr[i][j-1] == tree {
		arr[i][j-1] = fire
	}
	if j < n-1 && arr[i][j+1] == tree {
		arr[i][j+1] = fire
	}
	if i > 0 && j > 0 && arr[i-1][j-1] == tree {
		arr[i-1][j-1] = fire
	}
	if i > 0 && j < n-1 && arr[i-1][j+1] == tree {
		arr[i-1][j+1] = fire
	}
	if i < n-1 && j > 0 && arr[i+1][j-1] == tree {
		arr[i+1][j-1] = fire
	}
	if i < n-1 && j < n-1 && arr[i+1][j+1] == tree {
		arr[i+1][j+1] = fire
	}
}


// iteruje przez każdy element, jeżeli element jest ogniem, to wywołuje dla niego funkję spreadFire
// oraz ustawia ten element na popiół
func simulateFire(arr [][]string) {
	n := len(arr)
	for {
		fireSpread := false // zmienna optymalizacyjna
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i][j] == fire {
					spreadFire(arr, i, j)
					arr[i][j] = ash
					fireSpread = true // jeżeli funkcja spreadfire się wywoła, to kontynuujemy rozprzestrzenianie
				}
			}
		}
		if !fireSpread {// jeżeli funkcja spreadfire się nie wywoła w ogóle, to znaczy, że pożar się zakończył
			break  			// więc przerywamy pętlę.
		}
		displayArray(arr)
		time.Sleep(300 * time.Millisecond)
	}
}



func calculateBurnedPercentage(arr [][]string) float64 {
	totalTrees := 0
	burnedTrees := 0

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == tree || arr[i][j] == fire || arr[i][j] == ash {
				totalTrees++
			}
			if arr[i][j] == ash {
				burnedTrees++
			}
		}
	}

	if totalTrees == 0 {
		return 0
	}

	return (float64(burnedTrees) / float64(totalTrees)) * 100
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var n = 10
	var numTrees = 50
	arr := make([][]string, n)
	for i := range arr {
		arr[i] = make([]string, n)
	}

	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = empty
		}
	}

	count := 0
	for count < numTrees {
		i := rand.Intn(n)
		j := rand.Intn(n)
		if arr[i][j] == empty {
			arr[i][j] = tree
			count++
		}
	}
	displayArray(arr)

	i := rand.Intn(n)	// Miejsce uderzenia pioruna
	j := rand.Intn(n)

	if arr[i][j] == tree {
		arr[i][j] = fire
		displayArray(arr)
		simulateFire(arr)
	} else {
		arr[i][j] = flash
	}

	displayArray(arr)

	percentage := calculateBurnedPercentage(arr)
	fmt.Printf("Procent spalonego lasu: %.2f%%\n", percentage)
}
