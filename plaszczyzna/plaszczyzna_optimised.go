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
	return math.Hypot(a, b) // Użyj math.Hypot do obliczenia długości wektora
}

func main() {
	var srednie []float64

	start := time.Now()

	for i := 1; i <= MAX_ILOSC_ITERACJI; i++ {
		var suma float64 // Przechowuje sumę odległości zamiast tablicy
		for licznik_prob := 0; licznik_prob < ILOSC_PROB; licznik_prob++ {
			var odleglosc float64
			for licznik_iteracji := 0; licznik_iteracji < i; licznik_iteracji++ {
				// Generowanie punktu w dwóch wymiarach
				x := float64(rand.Int31n(int32(i*2))) - float64(i)
				y := float64(rand.Int31n(int32(i*2))) - float64(i)
				odleglosc += pitagoras(x, y)
			}
			suma += odleglosc / float64(i) // Dodaj odległość do sumy i oblicz średnią w locie
		}
		srednie = append(srednie, suma/float64(ILOSC_PROB)) // Oblicz średnią na podstawie sumy i liczby prób
	}

	elapsed := time.Since(start)
	fmt.Printf("Czas wykonania: %s\n", elapsed)

	fmt.Println(srednie)
}
