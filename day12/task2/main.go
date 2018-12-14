package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	left  = "left"
	right = "right"
	total = 50000000000
)

type step struct {
	n      int
	offset int
	dots   int
}

func main() {
	state, rules := readInput()
	lefts := make([]string, 0, len(rules))
	for left := range rules {
		lefts = append(lefts, left)
	}

	leftDots := maxDotsCount(left, lefts)
	rightDots := maxDotsCount(right, lefts)
	state = ".." + state + "...."
	offset := 2

	states := make(map[string]step)
	states[state] = step{0, offset, countDots(left, state)}

	var totalOffset int
	for i := 1; i <= total; i++ {

		if ld := countDots(left, state); ld < leftDots {
			offset += leftDots - ld
			state = strings.Repeat(".", leftDots-ld) + state
		}

		if rd := countDots(right, state); rd < rightDots {
			state = state + strings.Repeat(".", rightDots-rd)
		}

		newState := ""
	loop:
		for j := 0; j < len(state)-5+1; j++ {
			str := state[j : j+5]
			for left, right := range rules {
				if left == str {
					newState += right
					continue loop
				}
			}
			newState = newState + "."
		}
		state = ".." + newState + ".."

		trimmed := strings.Trim(state, ".")
		if prev, ok := states[trimmed]; ok {
			speed := countDots(left, state) - prev.dots
			totalOffset = offset - speed*(total-i)
			break
		}
		states[trimmed] = step{i, offset, countDots(left, state)}

	}

	res := 0
	for i, c := range state {
		if c == '#' {
			res += i - totalOffset
		}
	}

	fmt.Println(res)
}

func maxDotsCount(side string, strings []string) (max int) {
	for _, s := range strings {
		if dc := countDots(side, s); dc > max {
			max = dc
		}
	}
	return max
}

func countDots(side, payload string) (res int) {
	for res = 0; res < len(payload); res++ {
		var c uint8
		if side == "left" {
			c = payload[res]
		} else {
			c = payload[len(payload)-res-1]
		}

		if c == '#' {
			return res
		}
	}
	return res
}

func readInput() (string, map[string]string) {
	r := bufio.NewReader(os.Stdin)
	var (
		state string
		rules = make(map[string]string)
		t     string
	)

	bytes, _, _ := r.ReadLine()
	_, err := fmt.Sscanf(string(bytes), "%s%s%s", &t, &t, &state)
	if err != nil {
		panic(err)
	}

	_, _, _ = r.ReadLine()
	var ruleLeft, ruleRight string
	for {
		bytes, _, _ := r.ReadLine()
		line := string(bytes)

		_, err := fmt.Sscanf(line, "%5s%4s%1s", &ruleLeft, &t, &ruleRight)
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		rules[ruleLeft] = ruleRight
	}

	return state, rules
}
