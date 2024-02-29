
package main
import ("fmt")
import ("math")

func main() {
  var a float64
  var b float64
  var c float64

  fmt.Println("Witaj w programie liczącym pierwiastki funkcji kwadratowej.")
  fmt.Print("Podaj a: ")
  fmt.Scan(&a)
  fmt.Print("Podaj b: ")
  fmt.Scan(&b)
  fmt.Print("Podaj c: ")
  fmt.Scan(&c)
  fmt.Println("Twoja funkcja: " + fmt.Sprint(a) + "x^2 + " + fmt.Sprint(b) + "x + " + fmt.Sprint(c))

  var delta float64 = b*b - 4*a*c

  fmt.Println("Delta wynosi " + fmt.Sprint(delta))
  
  var x1 float64 = (-b - math.Sqrt(delta))/2*a
  var x2 float64 = (-b + math.Sqrt(delta))/2*a

  fmt.Println("Rozwiązania: " + fmt.Sprint(x1) + ", " + fmt.Sprint(x2))
}