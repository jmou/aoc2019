package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntcodeComputer struct {
	memory []int
	ip     int
}

func (computer *IntcodeComputer) tick() bool {
	m := computer.memory
	var operation func(i []int) int
	switch opcode := m[computer.ip]; opcode { // opcode
	case 1:
		operation = func(i []int) int { m[i[2]] = m[i[0]] + m[i[1]]; return 3 }
	case 2:
		operation = func(i []int) int { m[i[2]] = m[i[0]] * m[i[1]]; return 3 }
	case 99:
		return false
	default:
		panic(fmt.Sprintf("unknown opcode %d at %d", opcode, computer.ip))
	}

	num_parameters := operation(m[computer.ip+1:])
	computer.ip += 1 + num_parameters
	return true
}

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

	computer := IntcodeComputer{memory: make([]int, len(program))}
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(computer.memory, program)
			computer.memory[1] = noun
			computer.memory[2] = verb
			computer.ip = 0
			for computer.tick() {
			}
			if computer.memory[0] == 19690720 {
				fmt.Println(100*noun + verb)
				break
			}
		}
	}
}
