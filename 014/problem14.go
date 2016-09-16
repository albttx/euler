package main

import "fmt"

func collatz_sequence(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return (3 * n) + 1
	}
}

func sequence_len(n int) int {
	len := 1
	for {
		n = collatz_sequence(n)
		len += 1
		if n == 1 {
			break
		}
	}
	return len
}

func main() {
	longest_size := 0
	longest_i := 0
	for i := 1; i < 1000000; i++ {
		ret := sequence_len(i)
		if ret > longest_size {
			longest_i = i
			longest_size = ret
		}
	}
	fmt.Println(longest_size, longest_i)
}
