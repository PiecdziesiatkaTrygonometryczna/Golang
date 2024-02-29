
package main
import ("fmt")
// import ("time")

func main() {
	var n int
	fmt.Println("Witaj w programie rozwiązującym Problem Collatza.")
	fmt.Print("Podaj n: ")
	fmt.Scan(&n)
	var counter int = 0
	
for true {
	if n % 2 != 0 {
		n = n*3 + 1
	} else {
		n = n/2
	}

	fmt.Println(n)
	counter = counter + 1
	// time.Sleep(time.Second / 50)

	if n == 4 {
		fmt.Println("Liczba wykonań pętli: " + fmt.Sprint(counter))
		break
	}
}


}

// czy liczba jest parzysta
// jezeli nie to mnozysz przez 3 i dodajesz jeden
// jezeki tak to dzielisz przez 2

// ile wykonań pętli