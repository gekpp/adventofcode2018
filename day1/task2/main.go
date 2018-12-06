package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	stdin := os.Stdin
	r := bufio.NewReader(stdin)
	input := make([]int64, 0)
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF || len(bytes) == 0 {
			break
		}
		str := string(bytes)

		t, _ := strconv.ParseInt(str, 10, 32)
		input = append(input, t)

	}

	fmt.Println(whatTwice(input))
}

func whatTwice(input []int64) int64 {
	resMap := make(map[int64]struct{})
	var (
		res int64
		i   int
	)
	for {
		t := input[i]
		res += t

		if _, ok := resMap[res]; ok {
			break
		}
		resMap[res] = struct{}{}
		i = (i + 1) % len(input)
	}
	return res
}
