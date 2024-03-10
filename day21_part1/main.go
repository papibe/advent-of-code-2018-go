package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	opcode string
	A      int
	B      int
	C      int
}

func addr(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] + (*registers)[B]
}

func addi(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] + B
}

func mulr(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] * (*registers)[B]
}

func muli(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] * B
}

func banr(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] & (*registers)[B]
}

func bani(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] & B
}

func borr(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] | (*registers)[B]
}

func bori(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A] | B
}

func setr(A, B, C int, registers *[]int) {
	(*registers)[C] = (*registers)[A]
}

func seti(A, B, C int, registers *[]int) {
	(*registers)[C] = A
}

func gtir(A, B, C int, registers *[]int) {
	if A > (*registers)[B] {
		(*registers)[C] = 1
	} else {
		(*registers)[C] = 0
	}
}

func gtri(A, B, C int, registers *[]int) {
	if (*registers)[A] > B {
		(*registers)[C] = 1
	} else {
		(*registers)[C] = 0
	}
}

func gtrr(A, B, C int, registers *[]int) {
	if (*registers)[A] > (*registers)[B] {
		(*registers)[C] = 1
	} else {
		(*registers)[C] = 0
	}
}

func eqir(A, B, C int, registers *[]int) {
	if A == (*registers)[B] {
		(*registers)[C] = 1
	} else {
		(*registers)[C] = 0
	}
}

func eqri(A, B, C int, registers *[]int) {
	if (*registers)[A] == B {
		(*registers)[C] = 1
	} else {
		(*registers)[C] = 0
	}
}

func eqrr(A, B, C int, registers *[]int) {
	if (*registers)[A] == (*registers)[B] {
		(*registers)[C] = 1
	} else {
		(*registers)[C] = 0
	}
}

type fs func(A, B, C int, before *[]int)

// var FUNCTIONS = []fs{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

var FUNCTIONS = map[string]fs{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

func parse(filename string) ([]Instruction, int) {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	data := strings.Trim(string(raw_data), "\n")
	whole_program := strings.Split(data, "\n")

	// get register and program
	register_instruction := whole_program[0]
	register, _ := strconv.Atoi(strings.Split(register_instruction, " ")[1])
	program := whole_program[1:]

	instructions := []Instruction{}

	for _, instruction := range program {
		items := strings.Split(instruction, " ")

		// instructions
		opcode := items[0]
		A, _ := strconv.Atoi(items[1])
		B, _ := strconv.Atoi(items[2])
		C, _ := strconv.Atoi(items[3])

		instructions = append(instructions, Instruction{opcode, A, B, C})
	}

	return instructions, register
}

type State struct {
	reg0 int
	reg1 int
	reg2 int
	reg3 int
	reg4 int
	reg5 int
}

func solve(instructions []Instruction, intruction_pointer, register0, max_instructions int) bool {
	registers := []int{register0, 0, 0, 0, 0, 0}
	program_pointer := registers[intruction_pointer]
	state := make(map[State]bool)
	for ninstuction := 0; ninstuction < max_instructions; ninstuction++ {
		flag := false
		if program_pointer > 25 {
			fmt.Print("pointer", program_pointer, registers)
			flag = true
		}
		current_state := State{registers[0], registers[1], registers[2], registers[3], registers[4], registers[5]}
		_, is_already_in_state := state[current_state]
		if is_already_in_state {
			fmt.Println("loop detected at ", ninstuction)
		} else {
			state[current_state] = true
		}
		if program_pointer < 0 || program_pointer > len(instructions)-1 {
			return true
		}
		instruction := instructions[program_pointer]
		FUNCTIONS[instruction.opcode](instruction.A, instruction.B, instruction.C, &registers)
		registers[intruction_pointer] += 1
		program_pointer = registers[intruction_pointer]
		if flag {
			fmt.Println("\t", registers)
		}
	}
	return false
}

func solution(filename string) int {
	instructions, instruction_register := parse(filename)
	max_instructions := 10000
	for reg0 := 15823996; reg0 < 15823996+1; reg0++ {
		if solve(instructions, instruction_register, reg0, max_instructions) {
			fmt.Println(reg0, "finishes")
		}
	}
	return 0
}

func main() {
	fmt.Println(solution("input.txt")) //
}
