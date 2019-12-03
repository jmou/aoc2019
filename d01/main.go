package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum1, sum2 := 0, 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		mass = mass/3 - 2
		sum1 += mass
		for ; mass > 0; mass = mass/3 - 2 {
			sum2 += mass
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
