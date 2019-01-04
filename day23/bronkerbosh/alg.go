package main

type set map[Vertex]interface{}

func copySet(s set) set {
	res := make(set)
	for k := range s {
		res[k] = struct {
		}{}
	}
	return res
}
func (s set) add(b Vertex) {
	s[b] = struct{}{}
}

func (s set) contains(b Vertex) bool {
	_, ok := s[b]
	return ok
}

func (s set) intersect(other []Vertex) set {
	result := make(set)
	for _, v := range other {
		if s.contains(v) {
			result.add(v)
		}
	}
	return result
}

func buildSet(verts []Vertex) set {
	res := make(set)
	for _, v := range verts {
		res.add(v)
	}
	return res
}

func (s set) sub(neighbours []Vertex) set {
	res := make(set)
	ns := buildSet(neighbours)
	for v := range s {
		if !ns.contains(v) {
			res.add(v)
		}
	}
	return res
}

type AdjacentMatrix map[Vertex][]Vertex
type Vertex int

type Graph struct {
	adjMatrix map[Vertex][]Vertex
	verts     []Vertex
}

func NewGraph(adjM AdjacentMatrix) *Graph {
	vertices := make([]Vertex, 0, len(adjM))
	for k := range adjM {
		vertices = append(vertices, k)
	}
	return &Graph{adjM, vertices}
}

func (g *Graph) Neighbours(v Vertex) []Vertex {
	return g.adjMatrix[v]
}

func BronKerbosh(g Graph, r, p, x set) []set {
	if len(p) == 0 && len(x) == 0 {
		return []set{r}
	}

	u, _ := getPivot(g.adjMatrix, p)
	//fmt.Printf("chosen pivot element %v with degeneracy level %v\n", u, dl)
	res := make([]set, 0)
	pivN := g.Neighbours(u)
	pMinusNeighb := p.sub(pivN)
	for v := range pMinusNeighb {
		cpR := copySet(r)
		cpR.add(v)
		subres := BronKerbosh(g, cpR, p.intersect(g.Neighbours(v)), x.intersect(g.Neighbours(v)))
		res = append(res, subres...)
		delete(p, v)
		x.add(v)
	}
	return res
}

func getPivot(am map[Vertex][]Vertex, p set) (Vertex, int) {
	max := -1
	var mxv Vertex
	for k := range p {
		if d := len(am[k]); d > max {
			max = d
			mxv = k
		}
	}
	return mxv, max
}

func buildAdjacentMatrix(bots []bot) (res AdjacentMatrix) {
	res = make(AdjacentMatrix)
	for i, b1 := range bots {
		b1 := b1
		b1r := make([]Vertex, 0)
		for j, b2 := range bots {
			if i == j {
				continue
			}
			b2 := b2
			if b1.r+b2.r >= dist(b1.point, b2.point) {
				b1r = append(b1r, Vertex(j))
			}
		}
		res[Vertex(i)] = b1r
	}
	return res
}
