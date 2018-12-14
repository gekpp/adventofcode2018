package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	state, rules := readInput()
	state = ".." + state + "...."
	for i := 0; i < 20; i++ {

		fmt.Printf("%2d: %s\n", i, state)
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
	}

	res := 0
	for i, c := range state {
		if c == '#' {
			res += i - 20
		}
	}

	fmt.Println(res)
}

func readInput() (string, map[string]string) {
	stdin := strings.NewReader(`initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`)
	_ = stdin
	f, _ := os.Open("day12/input-ark.txt")
	r := bufio.NewReader(f)
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
