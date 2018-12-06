package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var ids []string

	stdin := os.Stdin
	r := bufio.NewReader(stdin)
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		ids = append(ids, string(bytes))
	}

	fmt.Println("Read ")

	for i := 0; i < len(ids)-1; i++ {

		for j := i + 1; j < len(ids); j++ {
			s1, s2 := ids[i], ids[j]
			if len(s1) != len(s2) {
				continue
			}

			var delta int
			for ri, c1 := range s1 {
				if c1 != []rune(s2)[ri] {
					delta ++
					if delta > 1 {
						break
					}
				}
			}

			if delta == 1 {
				fmt.Println(commonLetters(s1, s2))
			}
		}
	}
}

func commonLetters(s1, s2 string) string {
	var res strings.Builder
	for ri, c1 := range s1 {
		if c1 != []rune(s2)[ri] {
			continue
		}

		res.WriteRune(c1)
	}
	return res.String()
}
