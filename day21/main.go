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

	_ = regIp
	_ = program
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
