package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Sample struct {
	before []int
	after  []int
	opcode int
	A      int
	B      int
	C      int
}

type Instruction struct {
	opcode int
	A      int
	B      int
	C      int
}

func addr(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] + before[B]
	return after
}

func addi(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] + B
	return after
}

func mulr(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] * before[B]
	return after
}

func muli(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] * B
	return after
}

func banr(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] & before[B]
	return after
}

func bani(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] & B
	return after
}

func borr(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] | before[B]
	return after
}

func bori(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] | B
	return after
}

func setr(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A]
	return after
}

func seti(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = A
	return after
}

func gtir(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if A > before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func gtri(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] > B {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func gtrr(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] > before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func eqir(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if A == before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func eqri(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] == B {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func eqrr(A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] == before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

type fs func(A, B, C int, before []int) []int

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

func parse(filename string) ([]Sample, []Instruction) {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	data := strings.Trim(string(raw_data), "\n")
	parts := strings.Split(data, "\n\n\n\n")
	raw_samples := parts[0]
	samples_str_list := strings.Split(raw_samples, "\n\n")

	re_before := regexp.MustCompile(`Before: \[(\d+), (\d+), (\d+), (\d+)\]`)
	re_after := regexp.MustCompile(`After:  \[(\d+), (\d+), (\d+), (\d+)\]`)
	re_instructions := regexp.MustCompile(`(\d+) (\d+) (\d+) (\d+)`)

	samples := []Sample{}

	for _, sample := range samples_str_list {
		lines := strings.Split(sample, "\n")

		// registers before
		before_str := lines[0]
		before_matches := re_before.FindStringSubmatch(before_str)
		register1, _ := strconv.Atoi(before_matches[1])
		register2, _ := strconv.Atoi(before_matches[2])
		register3, _ := strconv.Atoi(before_matches[3])
		register4, _ := strconv.Atoi(before_matches[4])
		before_registers := []int{register1, register2, register3, register4}

		// instructions
		instruction_str := lines[1]
		instr_matches := re_instructions.FindStringSubmatch(instruction_str)
		opcode, _ := strconv.Atoi(instr_matches[1])
		A, _ := strconv.Atoi(instr_matches[2])
		B, _ := strconv.Atoi(instr_matches[3])
		C, _ := strconv.Atoi(instr_matches[4])

		// registers before
		after_str := lines[2]
		after_matches := re_after.FindStringSubmatch(after_str)
		register1, _ = strconv.Atoi(after_matches[1])
		register2, _ = strconv.Atoi(after_matches[2])
		register3, _ = strconv.Atoi(after_matches[3])
		register4, _ = strconv.Atoi(after_matches[4])
		after_registers := []int{register1, register2, register3, register4}

		samples = append(samples, Sample{before_registers, after_registers, opcode, A, B, C})
	}

	// parse program
	raw_program := parts[1]
	instructions_str_list := strings.Split(raw_program, "\n")
	instructions := []Instruction{}
	for _, instruction := range instructions_str_list {
		instr_matches := re_instructions.FindStringSubmatch(instruction)
		opcode, _ := strconv.Atoi(instr_matches[1])
		A, _ := strconv.Atoi(instr_matches[2])
		B, _ := strconv.Atoi(instr_matches[3])
		C, _ := strconv.Atoi(instr_matches[4])

		instructions = append(instructions, Instruction{opcode, A, B, C})
	}

	return samples, instructions
}

func regs_are_equal(regs1, regs2 []int) bool {
	for i := 0; i < 4; i++ {
		if regs1[i] != regs2[i] {
			return false
		}
	}
	return true
}

func map_instructions(samples []Sample) map[int]string {
	matches := make(map[int]map[string]bool)
	for _, sample := range samples {
		for name, function := range FUNCTIONS {
			after := function(sample.A, sample.B, sample.C, sample.before)
			if regs_are_equal(after, sample.after) {
				_, is_in_matches := matches[sample.opcode]
				if is_in_matches {
					matches[sample.opcode][name] = true
				} else {
					new_map := map[string]bool{name: true}
					matches[sample.opcode] = new_map
				}
			}
		}
	}
	correct_matches := make(map[int]string)
	for len(correct_matches) < 16 {

		// find a solution
		var index int = -1
		for opcode := 0; opcode < 16; opcode++ {
			if len(matches[opcode]) == 1 {
				index = opcode
				break
			}
		}
		if index != -1 {
			for k, _ := range matches[index] {
				correct_matches[index] = k
			}
		}
		// remove found solution from others
		for opcode := 0; opcode < 16; opcode++ {
			_, is_in_map := matches[opcode][correct_matches[index]]
			if is_in_map {
				delete(matches[opcode], correct_matches[index])
			}
		}

	}
	return correct_matches
}

func solve(instructions []Instruction, oti map[int]string) int {
	before := []int{0, 0, 0, 0}
	after := []int{0, 0, 0, 0}
	for _, instruction := range instructions {
		function := FUNCTIONS[oti[instruction.opcode]]
		after = function(instruction.A, instruction.B, instruction.C, before)
		before = after
	}
	return after[0]
}

func solution(filename string) int {
	samples, instructions := parse(filename)
	opcode_to_instructions := map_instructions(samples)
	return solve(instructions, opcode_to_instructions)
}

func main() {
	fmt.Println(solution("input.txt")) // 496
}
