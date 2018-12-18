package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type regs [4]int
type fn func(a, b, c int, reg *regs)

type sample struct {
	before regs
	op     regs
	after  regs
}

var ops = []fn{
	addr,
	addi,
	mulr,
	muli,
	banr,
	bani,
	borr,
	bori,
	setr,
	seti,
	gtir,
	gtri,
	gtrr,
	eqir,
	eqri,
	eqrr,
}

func main() {
	samples := readSamples()
	var manyOptionSamples int
	// opcode -> []fnNum
	opCodes := make(map[int]map[int]struct{})
	for i := 0; i < len(samples); i++ {
		s := samples[i%len(samples)]
		var passCnt int

		var nums []int

		if len(opCodes[s.op[0]]) == 1 {
			continue
		}
		for i, op := range ops {
			reg := s.before
			op(s.op[1], s.op[2], s.op[3], &reg)
			if reg == s.after {
				passCnt ++
				nums = append(nums, i)
			}
		}

		prevNums, ok := opCodes[s.op[0]]
		if !ok {
			opCodes[s.op[0]] = make(map[int]struct{})
			for _, n := range nums {
				opCodes[s.op[0]][n] = struct{}{}
			}
			prevNums = opCodes[s.op[0]]
		} else {
			newNums := make(map[int]struct{})
			for _, n := range nums {
				if _, ok := prevNums[n]; ok {
					newNums[n] = struct{}{}
				}
			}
			opCodes[s.op[0]] = newNums
		}

		if passCnt >= 3 {
			manyOptionSamples ++
		}
	}

	oneQueue := make(map[int]struct{})
	for {
		startLen := len(oneQueue)
		for _, nums := range opCodes {
			if len(nums) == 1 {
				for k := range nums {
					oneQueue[k] = struct{}{}
				}
			}
		}
		if len(oneQueue) == startLen {
			break
		}

		for _, nums := range opCodes {
			if len(nums) == 1 {
				continue
			}

			for num := range nums {
				if _, ok := oneQueue[num]; ok {
					delete(nums, num)
				}
			}
		}
	}

	fmt.Println(manyOptionSamples)
	for i := 0; i < 16; i++ {
		fmt.Printf("op #%d: %v\n", i, opCodes[i])
	}

	finalOps := make(map[int]fn)
	for code, nums := range opCodes {
		for n := range nums {
			finalOps[code] = ops[n]
		}
	}

	instructions := readInstructions()

	var r regs
	for _, command := range instructions {
		finalOps[command[0]](command[1], command[2], command[3], &r)
	}
	fmt.Println(r)
}

func readSamples() (samples []sample) {
	f, err := os.Open("day16/input-ark-1.txt")
	if err != nil {
		panic(err)
	}
	_ = f

	s := bufio.NewScanner(f)
	for {
		s, err := readSample(s)
		if err != nil {
			break
		}
		samples = append(samples, s)
	}
	_ = f.Close()
	return samples
}

func readSample(s *bufio.Scanner) (sample, error) {
	var before, op, after []string
	if !s.Scan() {
		return sample{}, io.EOF
	}

	line := s.Text()
	before = strings.Split(strings.TrimRight(strings.TrimLeft(line, "Before: ["), "]"), ", ")

	s.Scan()
	line = s.Text()
	op = strings.Split(line, " ")

	s.Scan()
	line = s.Text()
	after = strings.Split(strings.TrimRight(strings.TrimLeft(line, "After: ["), "]"), ", ")
	s.Scan()
	line = s.Text()

	var res sample
	for i := 0; i < 4; i++ {
		_, _ = fmt.Sscanf(before[i], "%v", &res.before[i])
		_, _ = fmt.Sscanf(op[i], "%v", &res.op[i])
		_, _ = fmt.Sscanf(after[i], "%v", &res.after[i])
	}
	return res, nil
}

func readInstructions() (res []regs) {
	f, _ := os.Open("day16/input-ark-2.txt")
	s := bufio.NewScanner(f)
	for s.Scan() {
		var r regs
		fmt.Sscanf(s.Text(), "%v %v %v %v", &r[0], &r[1], &r[2], &r[3])
		res = append(res, r)
	}
	return res
}
