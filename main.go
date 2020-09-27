package main

import (
	"fmt"
)

func main() {
	limit := 0
	var result float64 = 1 / 1
	fmt.Println("Ingrese un limite: ")
	fmt.Scan(&limit)
	for i := 1; i <= limit; i++ {
		fact := factorial(float64(i))
		result = result + float64(1)/fact
	}
	fmt.Println("result", result)
}

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * factorial(float64(n-1))
}
