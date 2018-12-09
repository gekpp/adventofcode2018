package main

import (
	"fmt"
	"os"
)

type elem struct {
	v    int
	next *elem
	prev *elem
}

func main() {
	n, worth := readInput()
	scores := getScores(n, worth)
	var (
		maxI, maxS int
	)
	for i, s := range scores {
		if s > maxS {
			maxS, maxI = s, i
		}
	}
	fmt.Printf("Winner is the Elf #%d with score %d\n", maxI+1, maxS)
}

func getScores(n int, max int) []int {
	scores := make([]int, n, n)
	cur := &elem{v: 0}
	cur.next = cur
	cur.prev = cur

	for i := 1; i <= max; i++ {
		if i%23 == 0 {
			plN := (i - 1) % n
			removed := remove7Back(cur)
			cur = removed.next
			scores[plN] += removed.v + i
		} else {
			cur = placeBetweenM1M2(cur, i)
		}
	}
	return scores
}

func remove7Back(cur *elem) *elem {
	for i := 0; i < 7; i++ {
		cur = cur.prev
	}
	cur.prev.next, cur.next.prev = cur.next, cur.prev
	return cur
}

func placeBetweenM1M2(cur *elem, val int) *elem {
	t := &elem{v: val}
	m1 := cur.next
	m2 := cur.next.next
	m1.next, t.prev = t, m1
	t.next, m2.prev = m2, t
	return t
}

func readInput() (int, int) {
	var (
		plN, maxP int
	)
	_, _ = fmt.Fscanf(os.Stdin, "%d", &plN)
	_, _ = fmt.Fscanf(os.Stdin, "%d", &maxP)
	return plN, maxP
}
