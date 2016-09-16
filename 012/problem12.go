package main

import (
	"fmt"
	"math"
)

func find_number_factor(n int) int {
	nb_factor := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			nb_factor++
			if i != (n / i) {
				nb_factor++
			}
		}
	}
	return nb_factor
}

func main() {
	triangle_nb := 0
	for i := 1; ; i++ {
		triangle_nb += i
		if find_number_factor(triangle_nb) >= 500 {
			fmt.Println("Solution is ", triangle_nb)
			return
		}
	}
}
