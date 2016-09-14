package main

import "fmt"

var LIMIT int64 = 2000002 - 2

type prime struct {
	n        int64
	is_prime int64
}

func eratosthenes(lst []prime) {
	var nb int64
	count := 0
	sum := 0
	for i := 0; i < len(lst); i++ {
		for i < len(lst) && lst[i].is_prime != 0 {
			i++
		}
		if i >= len(lst) {
			return
		}
		nb = lst[i].n
		count++
		if nb >= LIMIT {
			fmt.Println("The sum of all primes below", LIMIT, "is", sum)
			return
		}
		sum += nb
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
