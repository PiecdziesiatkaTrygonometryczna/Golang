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
// Silna liczba dla "dawkal" to 476
//[100, 97, 119, 107, 97, 108] - kody ASCII mojego nicku
// 476! = 35867867064398478289209288875534387019028416258811836097155671133214919145444260935805432100557203526635297136372252064698892609981421602639632038857887829475680340687505785927330492899589147639440355071247941210396418976790496116638139844315723889910751888763214906402206414338520669922794130323497648626571442758115898723512911338574508975674744299325488079349804387736838497922290989218810430929600444317990195672429289253146597780504680119710765307542463279546249869881122681659239945050603874329534230560986949687181508415056078842404709174342374251031568128240156579745710039907957665407214016795792446584465884461641692036758184309865915561118301965962774560047028427934154513836663823856366754722964717185684148702128819913812101327095717508673881082348629325961023547740711742985105209962397132625285663757717550333733499916097970929256584070049965155727074714934204283094829846721551122905670975430021850311566169589525408067135863332586651648000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000