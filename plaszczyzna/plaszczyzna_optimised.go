package main

import (
    "fmt"
    "math"
    "math/rand"
    "os"
    "time"
)

const (
    MAX_ILOSC_ITERACJI = 10000
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
    return sum / float64(len(arr))
}

func main() {
    var srednie []float64
    var punkty [][]int
    odleglosci := make([][]float64, ILOSC_PROB)

    start := time.Now()

    for i := 1; i <= MAX_ILOSC_ITERACJI; i++ {
        if i == 1 {
            punkty = make([][]int, ILOSC_PROB)
            for j := 0; j < ILOSC_PROB; j++ {
                punkty[j] = []int{0, 0}
                for k := 0; k < i; k++ {
                    if rand.Intn(2) == 0 {
                        punkty[j][0] += rand.Intn(2)*2 - 1
                    } else {
                        punkty[j][1] += rand.Intn(2)*2 - 1
                    }
                }
                odleglosci[j] = []float64{pitagoras(float64(punkty[j][0]), float64(punkty[j][1]))}
            }
        } else {
            for j := 0; j < ILOSC_PROB; j++ {
                if rand.Intn(2) == 0 {
                    punkty[j][0] += rand.Intn(2)*2 - 1
                } else {
                    punkty[j][1] += rand.Intn(2)*2 - 1
                }
                odleglosci[j] = append(odleglosci[j], pitagoras(float64(punkty[j][0]), float64(punkty[j][1])))
            }
        }

        var suma float64
        for j := 0; j < ILOSC_PROB; j++ {
            suma += odleglosci[j][i-1]
        }
        srednie = append(srednie, suma/float64(ILOSC_PROB))
    }

    fmt.Printf("Czas wykonania: %s\n", time.Since(start))

    file, _ := os.Create("srednie.csv")
    for i, srednia := range srednie {
        file.WriteString(fmt.Sprintf("%d,%f\n", i+1, srednia))
    }
	file.Close()

    fmt.Println("Średnie zostały zapisane do pliku srednie.csv.")
}
