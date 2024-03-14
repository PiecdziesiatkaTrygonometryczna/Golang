package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/skratchdot/open-golang/open"
	"github.com/wcharczuk/go-chart"
)

func collatz(n int) []int {
	result := []int{n}
	for true {
		if n == 1 {
			break
		}
		if n%2 != 0 {
			n = n*3 + 1
		} else {
			n = n / 2
		}
		result = append(result, n)
		
	}
	return result
}

func countOccurrences(allDigits []int) [10]int {
	var counts [10]int
	for _, digit := range allDigits {
		counts[digit]++
	}
	return counts
}

func main() {
	var allNumbers []int
	var allDigits []int

	for i := 1; i <= 100000; i++ {
		seq := collatz(i)
		allNumbers = append(allNumbers, seq...)
	}

	for _, num := range allNumbers {
		numStr := strconv.Itoa(num)
		for _, digit := range numStr {
			digitInt, _ := strconv.Atoi(string(digit))
			allDigits = append(allDigits, digitInt)
		}
	}

	digitCounts := countOccurrences(allDigits)


	var labels []string
	var values []float64
	for digit, count := range digitCounts {
		labels = append(labels, fmt.Sprintf("%d (%d)", digit, count))
		values = append(values, float64(count))
	}

	graph := chart.BarChart{
		Title:      "Occurrences of Digits",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		XAxis: chart.Style{
			Show: true,
		},
		Bars: []chart.Value{},
	}
	for i, label := range labels {
		graph.Bars = append(graph.Bars, chart.Value{Value: values[i], Label: label})
	}

	f, _ := os.Create("bar_chart.png")
	defer f.Close()
	graph.Render(chart.PNG, f)

	open.Run("bar_chart.png")
}
