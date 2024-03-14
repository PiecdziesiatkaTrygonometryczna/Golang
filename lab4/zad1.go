package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateControllDigit(pesel string) int{
	s := strings.Split(pesel, "")
	var digits []int

	for _, digitStr := range s {
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			fmt.Println("Error during conversion")
			return -1
		}
		digits = append(digits, digit)
	}

	// fmt.Println(digits)
	var digitsMultipliedByWeights [10]int
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 4, 8:
			digitsMultipliedByWeights[i] = digits[i] * 1
		case 1, 5, 9:
			digitsMultipliedByWeights[i] = digits[i] * 3
		case 2, 6:
			digitsMultipliedByWeights[i] = digits[i] * 7
		case 3, 7:
			digitsMultipliedByWeights[i] = digits[i] * 9
		}

		if digitsMultipliedByWeights[i] >= 10 {
			digitsMultipliedByWeights[i] = digitsMultipliedByWeights[i]%10
		}
	}

	sum := 0
	for _, i := range digitsMultipliedByWeights {
		sum = sum + i
	}
	

	// fmt.Println(digitsMultipliedByWeights)
	// fmt.Println(sum)


	var control int
	if sum < 10 {
		control = 10 - sum
	} else {
		control = 10 - sum%10
	}
	// fmt.Println(control)
	return control

}

func main() {
	// var pesel string
	// fmt.Println("Podaj numer PESEL: ")
	// fmt.Scan(&pesel)

	fmt.Print(calculateControllDigit("02070803628"))
}
