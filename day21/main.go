package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type instruction struct {
	op   string
	args [3]int
}

func main() {
	regIp, program := readInput("day21/input-ark.txt")
	conversion(0)

	var reg regs
	reg[0] = 1413889
	steps := asm(&reg, regIp, program)
	fmt.Println(steps)
}

func asm(reg *regs, regIp int, program []instruction) int {
	ip := reg[regIp]
	var instCnt int
	for ip < len(program) {
		reg[regIp] = ip

		command := program[ip]
		fn := funcs[command.op]
		args := command.args

		fn(args[0], args[1], args[2], reg)
		ip = reg[regIp]
		ip++
		instCnt++
		if ip > 30 {
			break
		}
	}
	return instCnt
}

func readInput(filename string) (ip int, res []instruction) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	if !s.Scan() {
		panic("empty input")
	}
	fmt.Sscanf(strings.Split(s.Text(), " ")[1], "%d", &ip)

	for s.Scan() {
		var instr instruction
		split := strings.Split(s.Text(), " ")

		instr.op = split[0]
		fmt.Sscanf(split[1], "%d", &instr.args[0])
		fmt.Sscanf(split[2], "%d", &instr.args[1])
		fmt.Sscanf(split[3], "%d", &instr.args[2])
		res = append(res, instr)
	}
	return ip, res
}
