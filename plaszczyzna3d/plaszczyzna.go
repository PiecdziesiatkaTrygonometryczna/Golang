package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	MAX_ILOSC_KROKOW = 1000
	ILOSC_SYMULACJI = 100
)

func pitagoras3stopnia(a, b, c float64) float64 {
	return math.Sqrt(a*a + b*b + c*c)
}

func avg(arr []float64) float64 {
	var sum float64
	for _, v := range arr {
		sum += v
	}
	return sum / float64(len(arr))
}

func main() {
	start := time.Now()

	var srednie []float64

	for i := 1; i <= MAX_ILOSC_KROKOW; i++ {
		var odleglosci []float64

		for j := 0; j < ILOSC_SYMULACJI; j++ {
			var punkt [3]int

			for k := 0; k < i; k++ {
				switch rand.Intn(3) {
				case 0:
					punkt[0] += rand.Intn(2)*2 - 1
				case 1:
					punkt[1] += rand.Intn(2)*2 - 1
				case 2:
					punkt[2] += rand.Intn(2)*2 - 1
				}
			}

			odleglosci = append(odleglosci, pitagoras3stopnia(float64(punkt[0]), float64(punkt[1]), float64(punkt[2])))
		}

		srednie = append(srednie, avg(odleglosci))
	}

	fmt.Printf("Czas wykonania: %s\n", time.Since(start))

	file, err := os.Create("srednie.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()


	for i, srednia := range srednie {
		file.WriteString(fmt.Sprintf("%d,%f\n", i+1, srednia))
	}

	fmt.Println("Średnie zostały zapisane do pliku srednie.csv.")
}
