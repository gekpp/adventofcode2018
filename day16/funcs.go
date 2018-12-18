package main

func addr(a, b, c int, reg *regs) {
	reg[c] = reg[a] + reg[b]
}

func addi(a, b, c int, reg *regs) {
	reg[c] = reg[a] + b
}

func mulr(a, b, c int, reg *regs) {
	reg[c] = reg[a] * reg[b]
}

func muli(a, b, c int, reg *regs) {
	reg[c] = reg[a] * b
}

func banr(a, b, c int, reg *regs) {
	reg[c] = reg[a] & reg[b]
}

func bani(a, b, c int, reg *regs) {
	reg[c] = reg[a] & b
}

func borr(a, b, c int, reg *regs) {
	reg[c] = reg[a] | reg[b]
}

func bori(a, b, c int, reg *regs) {
	reg[c] = reg[a] | b
}

func setr(a, b, c int, reg *regs) {
	reg[c] = reg[a]
}

func seti(a, b, c int, reg *regs) {
	reg[c] = a
}

func gtir(a, b, c int, reg *regs) {
	if a > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func gtri(a, b, c int, reg *regs) {
	if reg[a] > b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func gtrr(a, b, c int, reg *regs) {
	if reg[a] > reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqir(a, b, c int, reg *regs) {
	if a == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqri(a, b, c int, reg *regs) {
	if reg[a] == b {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}

func eqrr(a, b, c int, reg *regs) {
	if reg[a] == reg[b] {
		reg[c] = 1
	} else {
		reg[c] = 0
	}
}
