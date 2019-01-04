package main

import (
	"fmt"
	"log"
	"math"
	"testing"
)

func TestBronKerbosh(t *testing.T) {
	//bots := readInput("../input-test-2.txt")
	bots := readInput("../input-ark-2.txt")
	m := buildAdjacentMatrix(bots)
	g := NewGraph(m)
	var (
		r = make(set)
		p = make(set)
		x = make(set)
	)
	for _, v := range g.verts {
		p.add(v)
	}

	groups := BronKerbosh(*g, r, p, x)

	var (
		ml       int
		maxGroup set
	)

	for _, g := range groups {
		if l := len(g); l > ml {
			ml = l
			maxGroup = g
		}
	}

	for i := range bots {
		if _, ok := maxGroup[Vertex(i)]; !ok {
			fmt.Println(i)
		}
	}
	//Validate that clique
	for v := range maxGroup {
		diff := maxGroup.sub(m[v])

		if d := len(diff); d != 1 {
			log.Fatalf("Expected vertex %d to be adjacent with %d vertexes but it is actualy with %d", v, ml, d)
		}
	}

	var (
		mnr = math.MaxInt32
		mnb bot
		mnd = math.MaxInt32
	)

	for v := range maxGroup {
		if b := bots[int(v)]; b.r < mnr || (b.r == mnr && dist(b.point, point{}) < mnd) {
			mnr = b.r
			mnb = b
			mnd = dist(b.point, point{})
		}
	}

	wx := 2*mnb.r
	wy := 2*mnb.r
	wz := 2*mnb.r
	//wx,wy,wz=3116975, 3116975, 6233950
	searchArea := buildBrick(point{mnb.x - wx/2, mnb.y - wy/2, mnb.z - wz/2}, wx, wy, wz)
	queue := []brick{searchArea}

	{

		queue, max := searchDown(queue, bots)
		_ = queue
		_ = max
	}

	fmt.Printf("Expected to count %d bots\n", len(maxGroup))
	var max int
	for wx > 1 || wy > 1 || wz > 1 {

		longestAxis := maxAxis(wx, wy, wz)
		oldQ := queue
		queue = make([]brick, 0, len(oldQ)*2)
		for _, b := range oldQ {
			queue = append(queue, b.double(longestAxis)...)
		}

		b0 := queue[0]
		wx, wy, wz = b0.width(axX), b0.width(axY), b0.width(axZ)
		fmt.Printf("wx,wy,wz=%d, %d, %d. Len=%d.", wx, wy, wz, len(queue))
		queue, max = searchDown(queue, bots)
		fmt.Printf(" Max=%d\n", max)
	}

	fmt.Println(queue[0].points[0])
	fmt.Println(dist(point{},queue[0].points[0]))
}