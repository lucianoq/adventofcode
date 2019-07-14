package part1

import (
	"log"
	"strconv"
)

func prereq(ss ...string) bool {
	for _, s := range ss {
		if !available(s) && !isNum(s) {
			return false
		}
	}
	return true
}

func assignment(in, out string) bool {
	if !prereq(in) {
		return false
	}

	if isNum(in) {
		mem[out] = parseNum(in)
	} else {
		mem[out] = mem[in]
	}

	return true
}

func not(in, out string) bool {
	if !prereq(in) {
		return false
	}

	if isNum(in) {
		mem[out] = ^parseNum(in)
	} else {
		mem[out] = ^mem[in]
	}

	return true
}

func and(in1, in2, out string) bool {
	if !prereq(in1, in2) {
		return false
	}

	in1Num, in2Num := isNum(in1), isNum(in2)

	switch {
	case in1Num && in2Num:
		mem[out] = parseNum(in1) & parseNum(in2)
	case in1Num:
		mem[out] = parseNum(in1) & mem[in2]
	case in2Num:
		mem[out] = mem[in1] & parseNum(in2)
	default:
		mem[out] = mem[in1] & mem[in2]
	}

	return true
}

func or(in1, in2, out string) bool {
	if !prereq(in1, in2) {
		return false
	}

	in1Num, in2Num := isNum(in1), isNum(in2)

	switch {
	case in1Num && in2Num:
		mem[out] = parseNum(in1) | parseNum(in2)
	case in1Num:
		mem[out] = parseNum(in1) | mem[in2]
	case in2Num:
		mem[out] = mem[in1] | parseNum(in2)
	default:
		mem[out] = mem[in1] | mem[in2]
	}

	return true
}

func lshift(in, pos, out string) bool {
	if !prereq(in) {
		return false
	}

	if isNum(in) {
		mem[out] = parseNum(in) << parseNum(pos)
	} else {
		mem[out] = mem[in] << parseNum(pos)
	}
	return true
}

func rshift(in, pos, out string) bool {
	if !prereq(in) {
		return false
	}

	if isNum(in) {
		mem[out] = parseNum(in) >> parseNum(pos)
	} else {
		mem[out] = mem[in] >> parseNum(pos)
	}

	return true
}

func parseNum(s string) uint16 {
	num, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		log.Fatal(err)
	}
	return uint16(num)
}
