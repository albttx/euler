package main

import (
	"log"
	"os"
	"strconv"
)

var LIMIT int = 2000002 - 2

type prime struct {
	n        int
	is_prime int
}

func eratosthenes(lst []prime) {
	var nb int
	count := 0

	f, err := os.Create("./primes.txt")
	if err != nil {
		log.Fatal("Error open primes.txt")
	}
	defer f.Close()

	for i := 0; i < len(lst); i++ {
		for i < len(lst) && lst[i].is_prime != 0 {
			i++
		}
		if i >= len(lst) {
			return
		}
		nb = lst[i].n
		count++
		lst[i].is_prime = 1
		f.WriteString(strconv.Itoa(nb) + "\n")
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
