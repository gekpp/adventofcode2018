package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	kind_immune = iota
	kind_infect
)

type group struct {
	kind        int
	units       int
	hp          int
	weaknesses  []string
	immunities  []string
	damageType  string
	damagePower int
	initiative  int
}

func (g *group) effectivePower() int {
	return g.units * g.damagePower
}

func (g *group) isWeakTo(damageType string) bool {
	for _, w := range g.weaknesses {
		if w == damageType {
			return true
		}
	}
	return false
}

func (g *group) isImmuneTo(damageType string) bool {
	for _, w := range g.immunities {
		if w == damageType {
			return true
		}
	}
	return false
}

func (g *group) String() string {

}

func main() {
	groups := readInput("day24/input-test.txt")

	for !hasFinished(groups) {
		targets := targetPhase(groups)
		groups = attackPhase(groups, targets)
	}
}

func attackPhase(groups []*group, targets map[*group][]*group) []*group {

}

func targetPhase(groups []*group) map[*group][]*group {
	sort.Slice(groups, func(i, j int) bool {
		gi := groups[i]
		gj := groups[j]
		if gi.effectivePower() > gj.effectivePower() {
			return true
		}
		if gi.effectivePower() == gj.effectivePower() {
			return gi.initiative > gj.initiative
		}
		return false
	})

	imGrps, infGrps := splitGroups(groups)
	targets := make(map[*group][]*group)
	for _, g := range groups {
		if g.kind == kind_immune {
			if target, ok := findTargets(g, infGrps); ok {
				targets[g] = target
			}
		} else {
			if target, ok := findTargets(g, imGrps); ok {
				targets[g] = target
			}
		}
	}
	return targets
}

func findTargets(g *group, enemies map[*group]struct{}) ([]*group, bool) {
	var (
		maxG []*group
		maxD int
	)

	for eg := range enemies {
		k := 1
		if eg.isImmuneTo(g.damageType) {
			k = 0.5
		}
		if eg.isWeakTo(g.damageType) {
			k = 2
		}
		damage := g.effectivePower() * k
		if damage > maxD {
			maxD = damage
			maxG = []*group{eg}
		} else if damage == maxD {
			maxG = append(maxG, eg)
		}
	}

	return maxG, false
}

func splitGroups(groups []*group) (map[*group]struct{}, map[*group]struct{}) {
	imGrps := make(map[*group]struct{})
	infGrps := make(map[*group]struct{})
	for _, g := range groups {
		if g.kind == kind_immune {
			imGrps[g] = struct{}{}
		} else {
			infGrps[g] = struct{}{}
		}
	}
	return imGrps, infGrps
}

func hasFinished(groups []*group) bool {
	var immune, infect bool
	for _, g := range groups {
		if g.kind == kind_immune {
			immune = true
		}
		if g.kind == kind_infect {
			infect = true
		}

		if immune && infect {
			return true
		}
	}
	return false
}

func readInput(filename string) (groups []*group) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	s := bufio.NewScanner(f)
	s.Scan()
	s.Text()

	for s.Scan() {
		g, ok := readGroup(s)
		if ! ok {
			break
		}
		g.kind = kind_immune
		groups = append(groups, g)
	}

	s.Scan()
	_ = s.Text()
	for s.Scan() {
		g, ok := readGroup(s)
		if ! ok {
			break
		}
		g.kind = kind_infect
		groups = append(groups, g)
	}
	return groups
}

func readGroup(s *bufio.Scanner) (*group, bool) {
	line := s.Text()
	if line == "" {
		return nil, false
	}

	var g group

	split := strings.Split(line, " (")
	if _, err := fmt.Sscanf(split[0], "%v units each with %v hit points", &g.units, &g.hp); err != nil {
		log.Fatalf("%v", err)
	}

	split = strings.Split(split[1], ") with an attack that does ")
	props := strings.Split(split[0], "; ")
	for _, prop := range props {
		if strings.HasPrefix(prop, "immune to ") {
			g.immunities = strings.Split(strings.TrimLeft(prop, "immune to "), ", ")
			continue
		}

		if strings.HasPrefix(prop, "weak to ") {
			g.weaknesses = strings.Split(strings.TrimLeft(prop, "weak to "), ", ")
			continue
		}
	}

	if _, err := fmt.Sscanf(split[1], "%v %s damage at initiative %v", &g.damagePower, &g.damageType, &g.initiative); err != nil {
		log.Fatalf("%v", err)
	}
	return &g, true
}
