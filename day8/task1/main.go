package main

import (
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

func main() {
	stdin := os.Stdin
	var s scanner.Scanner
	s.Init(stdin)
	res := sumMeta(&s)
	fmt.Println(res)
}

func sumMeta(s *scanner.Scanner) (sum int) {
	s.Scan()
	txt := s.TokenText()
	chCnt, _ := strconv.ParseInt(txt, 10, 32)
	s.Scan()
	mtCnt, _ := strconv.ParseInt(s.TokenText(), 10, 32)

	for i := 0; i < int(chCnt); i++ {
		sum += sumMeta(s)
	}

	for i := 0; i < int(mtCnt); i++ {
		s.Scan()
		meta, _ := strconv.ParseInt(s.TokenText(), 10, 32)
		sum += int(meta)
	}
	return sum
}
