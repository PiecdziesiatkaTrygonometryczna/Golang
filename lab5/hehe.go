package main

import (
    // "bufio"
    "fmt"
    // "os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// os.Stdout.Write([]byte(text))


	var i int
	fmt.Print("Podaj liczbę: ")
    _, err := fmt.Scanf("%d", &i)
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)}
}	