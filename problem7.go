package main

import "fmt"

var LIMIT int = 2000000 - 2

type prime struct {
	n        int
	is_prime int
}

func eratosthenes(lst []prime) {
	var nb int
	count := 0
	for i := 0; i < len(lst); i++ {
		for i < len(lst) && lst[i].is_prime != 0 {
			i++
		}
		if i >= len(lst) {
			return
		}
		nb = lst[i].n
		count++
		if count == 10001 {
			fmt.Println("The 10001 nth Prime number is :", lst[i].n)
			return
		}
		lst[i].is_prime = 1
		for j := i; j < len(lst); j++ {
			if lst[j].is_prime != 1 && (lst[j].n%nb) == 0 {
				lst[j].is_prime = -1
			}
		}
	}
}

func main() {
	var lst []prime
	lst = make([]prime, LIMIT)
	for i := 0; i < LIMIT; i++ {
		lst[i] = prime{n: i + 2, is_prime: 0}
	}
	eratosthenes(lst)
}
