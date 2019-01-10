package main

type OpFunc func(a, b, c int)

type Instruction struct {
	FName string
	F     OpFunc
	A     int
	B     int
	C     int
}

func (i *Instruction) Exec() {
	register[ipRegister] = ip
	i.F(i.A, i.B, i.C)
	ip = register[ipRegister]
	ip++
}

var (
	register   [6]int
	ipRegister int
	ip         int
	opFuncs    = map[string]OpFunc{
		"addi": addi,
		"addr": addr,
		"mulr": mulr,
		"muli": muli,
		"banr": banr,
		"bani": bani,
		"borr": borr,
		"bori": bori,
		"setr": setr,
		"seti": seti,
		"gtir": gtir,
		"gtri": gtri,
		"gtrr": gtrr,
		"eqir": eqir,
		"eqri": eqri,
		"eqrr": eqrr,
	}
)

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
