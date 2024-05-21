package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	MAX_ILOSC_ITERACJI = 1000
	ILOSC_PROB         = 1000
)

func pitagoras(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func avg(arr []float64) float64 {
	var sum float64
	for _, v := range arr {
		sum += v
	}
	// return sum/float64(len(arr))
	return math.Round(sum/float64(len(arr))*100) / 100
}

func main() {
	var srednie []float64

	start := time.Now()

	for i := 1; i <= MAX_ILOSC_ITERACJI; i++ {
		// var i = 26384
		var odleglosci []float64
		for j := 0; j < ILOSC_PROB; j++ {
			var punkt [2]int
			for k := 0; k < i; k++ {
				if rand.Intn(2) == 0 {
					punkt[0] += rand.Intn(2)*2 - 1
				} else {
					punkt[1] += rand.Intn(2)*2 - 1
				}
			}
			odleglosci = append(odleglosci, pitagoras(float64(punkt[0]), float64(punkt[1])))
		}
		srednie = append(srednie, avg(odleglosci))
	}

	fmt.Printf("Czas wykonania: %s\n", time.Since(start))

	file, _ := os.Create("srednie.txt")
	for _, srednia := range srednie {
		file.WriteString(fmt.Sprintf("%.2f\n", srednia))
	}
	file.Close()

	fmt.Println("Średnie zostały zapisane do pliku srednie.txt.")
}
