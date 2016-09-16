package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	var sum, nb *big.Int

	sum = big.NewInt(0)
	nb = big.NewInt(0)
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		nb.SetString(txt, 10)
		sum = sum.Add(sum, nb)
	}
	fmt.Println(sum)
}
