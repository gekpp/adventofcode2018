package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func countLetters(id string) (twice []rune, three []rune) {
	cnt := make(map[rune]int)

	for _, r := range id {
		cnt[r] = cnt[r] + 1
	}

	for r, times := range cnt {
		if times == 2 {
			twice = append(twice, r)
		}
		if times == 3 {
			three = append(three, r)
		}
	}
	return twice, three
}

func main() {
	stdin := os.Stdin
	r := bufio.NewReader(stdin)
	var n2, n3 int
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF || len(bytes) == 0 {
			break
		}
		str := string(bytes)
		tw, tr := countLetters(str)
		if len(tw) > 0 {
			n2 ++
		}
		if len(tr) > 0 {
			n3 ++
		}
		fmt.Printf("%s: 2-%+v, 3-%+v\n", str, tw, tr)
	}
	fmt.Println(n2 * n3)
}
