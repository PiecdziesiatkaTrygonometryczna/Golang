package main

import (
    "fmt"
    "math"
    "math/rand"
    "os"
    "time"
)

const (
    MAX_ILOSC_KROKOW = 1_000
    ILOSC_SYMULACJI = 1_000
)

func pitagoras(a, b int) float64 {
    return math.Sqrt(float64(a*a + b*b))
}

func avg(arr []float64) float64 {
    var sum float64
    for _, v := range arr {
        sum += v
    }
    return sum / float64(len(arr))
}

func main() {
    var srednie []float64
    punkty := make([][2]int, ILOSC_SYMULACJI)

    start := time.Now()

    for i := 1; i <= MAX_ILOSC_KROKOW; i++ {
            var suma float64
            for j := 0; j < ILOSC_SYMULACJI; j++ {
                if rand.Intn(2) == 0 {
                    punkty[j][0] += rand.Intn(2)*2 - 1
                } else {
                    punkty[j][1] += rand.Intn(2)*2 - 1
                }
                suma += pitagoras(punkty[j][0], punkty[j][1])
            }
        srednie = append(srednie, suma/float64(ILOSC_SYMULACJI))
    }   

    fmt.Printf("Czas wykonania: %s\n", time.Since(start))

    file, _ := os.Create("srednie.csv")
    for i, srednia := range srednie {
        fmt.Fprintf(file, "%d,%f\n", i+1, srednia)
    }
	file.Close()

    fmt.Println("Średnie zostały zapisane do pliku srednie.csv.")
}
