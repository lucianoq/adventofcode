package main

type OpFunc func(a, b, c int)

type Instruction struct {
	OpCode int
	A      int
	B      int
	C      int
}

var (
	register [4]int
	solved   map[int]OpFunc

	opFuncs = [16]OpFunc{
		addi,
		addr,
		mulr,
		muli,
		banr,
		bani,
		borr,
		bori,
		setr,
		seti,
		gtir,
		gtri,
		gtrr,
		eqir,
		eqri,
		eqrr,
	}
	candidates [16]map[int]bool
)

func init() {
	for i := 0; i < 16; i++ {
		candidates[i] = make(map[int]bool)
	}
	solved = make(map[int]OpFunc)
}

func addi(a, b, c int) {
	register[c] = register[a] + b
}

func addr(a, b, c int) {
	register[c] = register[a] + register[b]
}

func mulr(a, b, c int) {
	register[c] = register[a] * register[b]
}

func muli(a, b, c int) {
	register[c] = register[a] * b
}

func banr(a, b, c int) {
	register[c] = register[a] & register[b]
}
func bani(a, b, c int) {
	register[c] = register[a] & b
}

func borr(a, b, c int) {
	register[c] = register[a] | register[b]
}

func bori(a, b, c int) {
	register[c] = register[a] | b
}

func setr(a, b, c int) {
	register[c] = register[a]
}
func seti(a, b, c int) {
	register[c] = a
}

func gtir(a, b, c int) {
	if a > register[b] {
		register[c] = 1
	} else {
		register[c] = 0
	}
}

func gtri(a, b, c int) {
	if register[a] > b {
		register[c] = 1
	} else {
		register[c] = 0
	}
}

func gtrr(a, b, c int) {
	if register[a] > register[b] {
		register[c] = 1
	} else {
		register[c] = 0
	}
}

func eqir(a, b, c int) {
	if a == register[b] {
		register[c] = 1
	} else {
		register[c] = 0
	}
}

func eqri(a, b, c int) {
	if register[a] == b {
		register[c] = 1
	} else {
		register[c] = 0
	}
}

func eqrr(a, b, c int) {
	if register[a] == register[b] {
		register[c] = 1
	} else {
		register[c] = 0
	}
}
