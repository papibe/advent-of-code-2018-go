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

func addr(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] + before[B]
	return after
}

func addi(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] + B
	return after
}

func mulr(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] * before[B]
	return after
}

func muli(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] * B
	return after
}

func banr(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] & before[B]
	return after
}

func bani(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] & B
	return after
}

func borr(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] | before[B]
	return after
}

func bori(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A] | B
	return after
}

func setr(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = before[A]
	return after
}

func seti(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	after[C] = A
	return after
}

func gtir(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if A > before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func gtri(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] > B {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func gtrr(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] > before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func eqir(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if A == before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func eqri(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] == B {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

func eqrr(opcode, A, B, C int, before []int) []int {
	after := make([]int, 4)
	copy(after, before)
	if before[A] == before[B] {
		after[C] = 1
	} else {
		after[C] = 0
	}
	return after
}

type fs func(opcode, A, B, C int, before []int) []int

var FUNCTIONS = []fs{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

func parse(filename string) []Sample {
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

	return samples
}

func regs_are_equal(regs1, regs2 []int) bool {
	for i := 0; i < 4; i++ {
		if regs1[i] != regs2[i] {
			return false
		}
	}
	return true
}

func solve(samples []Sample) int {
	three_or_more_opcodes := 0
	for _, sample := range samples {
		behaves_like_a_opcode := 0
		for _, function := range FUNCTIONS {
			after := function(sample.opcode, sample.A, sample.B, sample.C, sample.before)
			if regs_are_equal(after, sample.after) {
				behaves_like_a_opcode += 1
			}
		}
		if behaves_like_a_opcode >= 3 {
			three_or_more_opcodes += 1
		}
	}
	return three_or_more_opcodes
}

func solution(filename string) int {
	samples := parse(filename)
	return solve(samples)
}

func main() {
	fmt.Println(solution("input.txt")) // 509
}
