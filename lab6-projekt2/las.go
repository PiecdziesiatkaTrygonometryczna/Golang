package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) 

	
	var n int = 10
	arr := make([][]uint8, n)
	for i := range arr {
		arr[i] = make([]uint8, n)
	}

	count := 0
	for count < n {
		i := rand.Intn(n)
		j := rand.Intn(n)
		if arr[i][j] == 0 {
			arr[i][j] = 1
			count++
		}
	}


	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println()

	var arrOfTrees [][]int
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == 1 {
				cordsArr := []int{i, j}				
				arrOfTrees = append(arrOfTrees, cordsArr)
			}


		}
	}
	fmt.Println(arrOfTrees)
}
