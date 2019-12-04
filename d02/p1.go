package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	pieces := strings.Split(scanner.Text(), ",")
	program := make([]int, len(pieces))
	for i, piece := range pieces {
		var err error
		program[i], err = strconv.Atoi(piece)
		if err != nil {
			panic(err)
		}
	}
	if scanner.Scan() {
		panic("expected one line")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	program[1] = 12
	program[2] = 2

	position := 0
loop:
	for {
		var operation func(a, b int) int
		switch opcode := program[position]; opcode { // opcode
		case 1:
			operation = func(a, b int) int { return a + b }
		case 2:
			operation = func(a, b int) int { return a * b }
		case 99:
			break loop
		default:
			panic(fmt.Sprintf("unknown opcode %d at %d", opcode, position))
		}

		in1 := program[position+1]
		in2 := program[position+2]
		out := program[position+3]
		program[out] = operation(program[in1], program[in2])
		position += 4
	}

	fmt.Println(program[0])
}
