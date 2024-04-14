package main

import (
	"fmt"
	"math/big"
	"strings"
	"os"
)

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func findStrongNumber(name string) int {
	var asciiValues []string
	for _, letter := range name {
		asciiValues = append(asciiValues, fmt.Sprintf("%d", letter))  // konwertujemy imię na tablicę znaków ASCII
	}

	counter := 1   // przeszukujemy wszystkie możliwe silnie, poczynając od 1
	for {
		found := true // flaga
		factorialStr := factorial(counter).String() // konwertujemy silnię na string

		for _, asciiLetter := range asciiValues {
			index := strings.Index(factorialStr, asciiLetter)  // szukamy kodu ASCII naszej litery w silni
			if index == -1 { // jeżeli nie znajdziemy, wychodzimy z pętli
				found = false
				break
			}
			factorialStr = factorialStr[index+len(asciiLetter):] // jeżeli znajdziemy, to ucinamy początek naszej silni i sprawdzamy kolejną literę
		}

		if found { // jeżeli wszystkie litery zawierają się po kolei w naszej silni, kończymy działanie funkcji
			return counter
		}

		counter++
	}
}

func main() {
	name := "dawkal"
	strongNumber := findStrongNumber(name)
	fmt.Printf("Twoja Silna Liczba: %d\n", strongNumber)
	fac := factorial(strongNumber)
	content := fmt.Sprintf("%d\n%s", strongNumber, fac.String())
	file, err := os.Create("silnia.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Silnia Twojej Silnej Liczby została zapisana w pliku silnia.txt")

}
