package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
		// displayArray(arr)
		// time.Sleep(300 * time.Millisecond)
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

func runExperiment(n, trials int) map[int]float64 {
	results := make(map[int]float64)

	for percent := 1; percent <= 99; percent += 1 {
		totalBurnedPercentage := 0.0
		for trial := 0; trial < trials; trial++ {
			arr := generateForest(n, percent)
			i, j := generateLightningStrike(n)
			if arr[i][j] == tree {
				arr[i][j] = fire
				simulateFire(arr)
			}
			totalBurnedPercentage += calculateBurnedPercentage(arr)
		}
		averageBurnedPercentage := totalBurnedPercentage / float64(trials)
		results[percent] = averageBurnedPercentage
	}
	

	return results
}

func generateForest(n, percent int) [][]string {
	arr := make([][]string, n)
	for i := range arr {
		arr[i] = make([]string, n)
	}

	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = empty
		}
	}

	totalTrees := (n * n * percent) / 100
	count := 0
	for count < totalTrees {
		i := rand.Intn(n)
		j := rand.Intn(n)
		if arr[i][j] == empty {
			arr[i][j] = tree
			count++
		}
	}

	return arr
}


func generateLightningStrike(n int) (int, int) {
	i := rand.Intn(n)
	j := rand.Intn(n)
	return i, j
}

func main() {
	n := 10
	trials := 100000

	results := runExperiment(n, trials)

	file, _ := os.Create("results.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for percent, avgBurnedPercentage := range results {
		row := []string{strconv.Itoa(percent), fmt.Sprintf("%.2f", avgBurnedPercentage)}
		writer.Write(row)
	}

	fmt.Println("Wyniki zostały zapisane do pliku results.csv")
}