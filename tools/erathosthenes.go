package primes

func eratosthenes(limit int) []int {
	tab := make([]bool, limit+1)
	ret := make([]int, 0)

	for i := 2; i <= limit; i++ {
		if tab[i] == false {
			ret = append(ret, i)
			for j := i; j <= limit; j += i {
				tab[j] = true
			}
		}
	}
	return ret
}
