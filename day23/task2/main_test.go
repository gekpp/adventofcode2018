package main

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	bots := readInput("../input-test-2.txt")
	bc := 0

	//c := buildBrick(point{20, 10, 10}, 1, 49, 49)
	//c := buildCube(point{10, 10, 10}, 3)
	//c := buildCube(point{10, 10, 10}, 6)
	c := buildBrick(point{12, 12, 12}, 1, 1, 1)
	//c := buildBrick(point{12, 10, 10}, 1, 39, 39)
	fmt.Printf("%+v\n", c)

	bots = bots[0:]
	for _, b := range bots {
		inter := b.interfere(c)
		fmt.Printf("%+v => %v\n", b, inter)

		if inter {
			bc++
		}
	}

	if bc != 5 {
		t.Fatalf("Expected 5 got %d", bc)
	}
}

func TestDouble(t *testing.T) {
	bots := readInput("../input-ark.txt")
	//p := point{18306794, 40419702, 46145344}
	//p := point{18306794, 40419702, 46145344}
	p := point{12296533,31093692,42829595}
	var res int
	for _, b := range bots {
		if dist(b.point, p) <= b.r {
			res++
		}
	}

	fmt.Println(res)
}
