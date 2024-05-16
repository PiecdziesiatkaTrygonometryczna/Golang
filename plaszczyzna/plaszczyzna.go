package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	MAX_ILOSC_ITERACJI = 100
	ILOSC_PROB         = 10000
)

func pitagoras(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func avg(arr []float64) float64 {
	var sum float64
	for _, v := range arr {
		sum += v
	}
	return math.Round((sum / float64(len(arr))) * 100) / 100
}

func main() {
	var srednie []float64

	start := time.Now()

	for i := 1; i <= MAX_ILOSC_ITERACJI; i++ {
		var odleglosci []float64

		for licznik_prob := 0; licznik_prob < ILOSC_PROB; licznik_prob++ {
			var punkt [2]int
			for licznik_iteracji := 0; licznik_iteracji < i; licznik_iteracji++ {
				x_or_y := rand.Intn(2)
				plus_or_minus := rand.Intn(2)*2 - 1
				punkt[x_or_y] += plus_or_minus
			}
			odleglosc := pitagoras(float64(punkt[0]), float64(punkt[1]))
			odleglosci = append(odleglosci, odleglosc)
		}
		srednie = append(srednie, avg(odleglosci))
	}

	elapsed := time.Since(start)
	fmt.Printf("Czas wykonania: %s\n", elapsed)

	fmt.Println(srednie)
}
