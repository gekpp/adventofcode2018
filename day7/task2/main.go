package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

type graph map[rune][]rune

type worker struct {
	freeAtSec int32
	job       *rune
}

func main() {
	g := readInput()
	_, sec := assemble(g, 5)
	fmt.Println(": ", sec)
}

func assemble(g graph, wn int) ([]rune, int32) {
	workers := make([]*worker, wn, wn)
	for i := range workers {
		workers[i] = &worker{}
	}

	var (
		res []rune
		sec int32
	)

	doneSteps := make(map[rune]struct{})
	scheduledSteps := make(map[rune]struct{})

	for len(doneSteps) != len(g) {
		available := getAvailableSorted(g, doneSteps, scheduledSteps)
		if len(available) == 0 {
			sec = setWorkersFree(workers, sec, doneSteps, scheduledSteps)
			continue
		}

		w := getFreeWorker(sec, workers)
		if w == nil {
			sec = setWorkersFree(workers, sec, doneSteps, scheduledSteps)
			continue
		}

		job := available[0]
		scheduledSteps[job] = struct{}{}
		w.job = &job
		w.freeAtSec = sec + 60 + job - 64
	}
	return res, sec
}

func setWorkersFree(workers []*worker, sec int32, doneSteps map[rune]struct{}, scheduledSteps map[rune]struct{}) int32 {
	w := getFirstFreedWorker(workers)
	sec = w.freeAtSec
	for i, w := range workers {
		if workers[i].freeAtSec == sec {
			doneSteps[*w.job] = struct{}{}
			delete(scheduledSteps, *w.job)
			w.job = nil
		}
	}
	return sec
}

func getFreeWorker(sec int32, workers []*worker) *worker {
	for _, w := range workers {
		if w.freeAtSec <= sec {
			return w
		}
	}
	return nil
}

func getFirstFreedWorker(workers []*worker) (res *worker) {
	var min int32 = math.MaxInt32
	for _, w := range workers {
		if w.freeAtSec < min && w.job != nil {
			min = w.freeAtSec
			res = w
		}
	}
	return res
}

func getAvailableSorted(g graph, doneSteps, scheduled map[rune]struct{}) []rune {
	var res []rune
loop:
	for c, ps := range g {
		if _, ok := doneSteps[c]; ok {
			continue
		}
		if _, ok := scheduled[c]; ok {
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
	var stdin io.Reader = os.Stdin
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
