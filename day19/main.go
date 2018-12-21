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
	regIp, program := readInput("day19/input-ark.txt")

	var (
		reg = regs{1, 0, 0, 0, 0, 0}
		ip  = reg[regIp]
	)

	for ip < len(program) {
		reg[regIp] = ip

		command := program[ip]
		fn := funcs[command.op]
		args := command.args

		fmt.Printf("ip=%.2d %v %s %v => ", ip, reg, command.op, command.args)

		fn(args[0], args[1], args[2], &reg)
		ip = reg[regIp]
		ip++
		fmt.Printf("%v\n", reg)
	}

	fmt.Println(reg[0])
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
