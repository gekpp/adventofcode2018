package main

import (
	"fmt"
)

func conversion(startA int) {
	var (
		_, A, B, C, D, E, F int
		// 0  1  2  3  4  5
	)
	ds := make(map[int]struct{})
	var (
		prev  int
		first = true
	)

	A = startA
	_ = E

	D = 123
	// _L0
	for {
		D = D & 456
		if D == 72 {
			break
		}
	}
	D = 0

	//mainLoop:
	for {
		C = D | 65536
		D = 10736359
		for {
			B = C & 255
			D = D + B
			D = D & 16777215
			D = D * 65899
			D = D & 16777215
			if 256 > C {
				if _, ok := ds[D]; ok {
					fmt.Println("Maximum steps when A=", prev)
					return
				}

				prev = D
				ds[D] = struct{}{}
				if first {
					fmt.Println("First break if A=", D)
					first = false
				}
				if D == A {
					fmt.Println("D=", D)
					return
				} else {
					break
				}
			} else {
				B = 0

				for {
					F = B + 1 // ip=18
					F = F * 256
					if F > C {
						C = B
						break
					} else {
						B = B + 1
						continue
					}
				}
			}
		}
	}
}
