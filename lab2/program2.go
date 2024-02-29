
package main
import ("fmt")
// import ("math")

func main() {

var a int = 11000
var b int = 12000

wyniki := []int{}
var max int = 0
	for i := a; i <= b; i++ {

		var n int = i
		var counter int = 0
		for n >= 4 {
			if n % 2 != 0 {
				n = n*3 + 1
			} else {
				n = n/2
			}
			counter = counter + 1
		}
		fmt.Println("Liczba wykonań pętli dla liczby " + fmt.Sprint(i) + ": " + fmt.Sprint(counter))
		if counter > max {
			max = counter
		}
		wyniki = append(wyniki, counter)

	}
	
	var sum int
	len := len(wyniki)
	
	for _, liczba := range wyniki {
	  sum += liczba
	}
	
	avg := float64(sum) / float64(len)
	

	fmt.Println("Zakes od " + fmt.Sprint(a) + " do " + fmt.Sprint(b) + ":")
	fmt.Println("Średnia: " + fmt.Sprint(avg))
	fmt.Println("Max: " + fmt.Sprint(max))



}