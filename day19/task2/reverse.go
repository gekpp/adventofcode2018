package main

import "fmt"

func main1() {
	var (
		A, B, C, D, E, F int
		//0  1  2  3   4  5
	)
	A = 1
	{
		E = E + 2
		E = E * E
		//E = ip*4
		E = E * 19
		E = E * 11
		B = B + 6
		D = 22
		B = B * D
		B = B + 10
		E = E + B
		if A == 0 {
			fmt.Println("instruction 25")
		}

		B = 27
		B = B * 28
		B = 29 + B
		B = 30 * B
		B = B * 14
		B = B * 32
		E = E + B
		A = 0
	}

loop:
	for {
		F = 1
		for {
			C = 1
			for {
				B = F * C
				if B == E {
					A = A + F
					break
				}

				if B > E {
					break
				}

				C = C + 1
				if C > E {
					break
				}
			}

			F = F + 1
			if F > E {
				break loop
			}
		}
	}

	fmt.Println(A)
}
