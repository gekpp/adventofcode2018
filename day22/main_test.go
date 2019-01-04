package main

import (
	"testing"
)

func TestOtherTool(t *testing.T) {
	tools := map[int]string{torch: "torch", climb: "climbing gear", neith: "neither"}
	tcs := []struct {
		typ   int
		one   int
		other int
	}{
		{0, 1, 2},
		{0, 2, 1},
		{1, 2, 4},
		{1, 4, 2},
		{2, 4, 1},
		{2, 1, 4},
	}

	for _, tc := range tcs {
		tc := tc
		actualOther := otherTool(tc.typ, tc.one)
		if actualOther != tc.other {
			t.Errorf("Wrong other tool for cost type=%v. Expected=%s, but found=%s", tc.typ, tools[tc.other], tools[actualOther])
		}
	}
}

func TestCan(t *testing.T) {
	tools := map[int]string{torch: "torch", climb: "climbing gear", neith: "neither"}
	tcs := []struct {
		typ int
		one int
		can bool
	}{
		{0, 1, true},
		{0, 2, true},
		{0, 4, false},
		{1, 1, false},
		{1, 2, true},
		{1, 4, true},
		{2, 1, true},
		{2, 2, false},
		{2, 4, true},
	}

	for _, tc := range tcs {
		tc := tc
		actualCan := can(tc.typ, tc.one)
		if actualCan != tc.can {
			t.Errorf("Wrong tool allowance for cost type=%v and tool%v. Expected=%v, but found=%v", tc.typ, tools[tc.one], tc.can, actualCan)
		}
	}
}
