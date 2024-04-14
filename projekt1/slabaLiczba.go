package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "math"
)

var maxFibonacci = 30 // wartość, dla której liczymy ciąg fibonacciego
var counters = make([]int, maxFibonacci+1) // tablica przechowująca ilość wykonań funkcji fubonacci od danej liczby

func fibonacci(n int) int {
    counters[n]++ // jeżeli funkcja się wykonuje, dodajemy jedno wykonanie dla "n" do tablicy counters
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fibonacci_result := fibonacci(maxFibonacci)
    fmt.Printf("Fibonacci od %d: %d\n", maxFibonacci, fibonacci_result)
    for i := 0; i <= maxFibonacci; i++ {
        fmt.Printf("%d %d\n", i, counters[i]) // wyświetlamy liczbę wykonań funkcji dla każdej liczby
    }


    file, err := os.Open("silnia.txt") 
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    silna_liczba, err := strconv.Atoi(scanner.Text()) // wczytujemy silną liczbę z pliku silnia.txt
    if err != nil {
        fmt.Println(err)
        return
    }

    closest := math.Abs(float64(silna_liczba - counters[0])) // sprawdzamy która wartość funkcji fibonacciego jest najbliższa naszej Silnej Liczbie
    closest_index := 0
    for i := 0; i < len(counters); i++ { 
        if diff := math.Abs(float64(silna_liczba - counters[i])); diff < closest {
            closest = diff
            closest_index = i
        }
    }

    fmt.Printf("Najbliższa liczba Twojej Silnej Liczbie to %d, czyli Twoja Słaba Liczba to %d.\n", counters[closest_index], closest_index)
}