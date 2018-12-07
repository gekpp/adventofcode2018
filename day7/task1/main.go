package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

type graph map[rune][]rune

func main() {
	g := readInput()
	steps := assemble(g)
	for _, r := range steps {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}

func assemble(g graph) []rune {
	var res []rune
	doneSteps := make(map[rune]struct{})
	for len(doneSteps) != len(g) {
		available := getAvailableSorted(g, doneSteps)
		if len(available) == 0 {
			panic("no available steps for assembly")
		}

		doneSteps[available[0]] = struct{}{}
		res = append(res, available[0])
	}
	return res
}

func getAvailableSorted(g graph, doneSteps map[rune]struct{}) []rune {
	var res []rune
loop:
	for c, ps := range g {
		if _, ok := doneSteps[c]; ok {
			continue
		}

		for _, p := range ps {
			if _, ok := doneSteps[p]; !ok {
				continue loop
			}
		}
		res = append(res, c)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func readInput() map[rune][]rune {
	var stdin io.Reader = strings.NewReader(`
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`)

	stdin = os.Stdin
	r := bufio.NewReader(stdin)
	res := make(map[rune][]rune)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF && len(line) == 0 {
			break
		}

		if len(line) == 0 {
			continue
		}

		child, parent := parseLine(string(line))
		res[child] = append(res[child], parent)
		if _, ok := res[parent]; !ok {
			res[parent] = nil
		}
	}
	return res
}

func parseLine(line string) (child rune, parent rune) {
	split := strings.Split(line, " ")
	child, _ = utf8.DecodeRuneInString(split[7])
	parent, _ = utf8.DecodeRuneInString(split[1])
	return child, parent
}
