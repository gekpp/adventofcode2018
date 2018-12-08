package main

import (
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

type node struct {
	val      int
	children []*node
	meta     []int
}

func main() {
	stdin := os.Stdin
	var s scanner.Scanner
	s.Init(stdin)
	root := buildTree(&s)
	fmt.Println(root.val)
}

func buildTree(s *scanner.Scanner) *node {
	s.Scan()
	chCnt, _ := strconv.ParseInt(s.TokenText(), 10, 32)
	s.Scan()
	mtCnt, _ := strconv.ParseInt(s.TokenText(), 10, 32)

	children := make([]*node, chCnt, chCnt)
	for i := 0; i < int(chCnt); i++ {
		children[i] = buildTree(s)
	}

	meta := make([]int, mtCnt, mtCnt)
	for i := 0; i < int(mtCnt); i++ {
		s.Scan()
		metaV, _ := strconv.ParseInt(s.TokenText(), 10, 32)
		meta[i] = int(metaV)
	}

	val := calculateValue(children, meta)
	return &node{val, children, meta}
}

func calculateValue(children []*node, meta []int) (res int) {
	if len(children) == 0 {
		return sum(meta)
	}

	for _, im := range meta {
		if im == 0 || im > len(children) {
			continue
		}
		child := children[im-1]
		res += calculateValue(child.children, child.meta)
	}
	return res
}

func sum(list []int) (res int) {
	for _, v := range list {
		res += v
	}
	return res
}
