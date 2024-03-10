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

func solve(instructions []Instruction, intruction_pointer, register0 int) int {
	registers := []int{register0, 0, 0, 0, 0, 0}
	program_pointer := registers[intruction_pointer]
	for {
		// input _analisys.txt
		// instruction 28 uses register 0
		if program_pointer == 28 {
			return registers[4]
		}
		instruction := instructions[program_pointer]
		FUNCTIONS[instruction.opcode](instruction.A, instruction.B, instruction.C, &registers)
		registers[intruction_pointer] += 1
		program_pointer = registers[intruction_pointer]
	}
}

func solution(filename string) int {
	instructions, instruction_register := parse(filename)
	return solve(instructions, instruction_register, 0)
}

func main() {
	fmt.Println(solution("input.txt")) // 15823996
}
