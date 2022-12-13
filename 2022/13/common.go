package main

import "encoding/json"

func parse(line string) []any {
	ll := []any{}
	_ = json.Unmarshal([]byte(line), &ll)
	return ll
}

func compare(left any, right any) int {
	lInt, lList, lIsInt := typeOf(left)
	rInt, rList, rIsInt := typeOf(right)

	// both integer
	if lIsInt && rIsInt {
		return intCompare(lInt, rInt)
	}

	// both lists
	if !lIsInt && !rIsInt {
		lLen, rLen := len(lList), len(rList)
		for i := 0; i < min(lLen, rLen); i++ {
			// return if they're different
			if cmp := compare(lList[i], rList[i]); cmp != 0 {
				return cmp
			}
			// continue if they're equal
		}

		// when a list is a prefix of another
		// the shortest comes first
		return intCompare(lLen, rLen)
	}

	if lIsInt {
		return compare([]any{lInt}, right)
	}

	if rIsInt {
		return compare(left, []any{rInt})
	}

	panic("wrong types")
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func typeOf(x any) (int, []any, bool) {
	switch val := x.(type) {
	case int:
		return val, nil, true
	case float64:
		return int(val), nil, true
	case []any:
		return 0, val, false
	}
	panic("wrong type")
}

func intCompare(a, b int) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}
