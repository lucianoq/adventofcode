package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Comp struct{ L, R int }

func parse() map[Comp]struct{} {
	comps := make(map[Comp]struct{}, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Split(line, "/")
		l, _ := strconv.Atoi(ff[0])
		r, _ := strconv.Atoi(ff[1])
		comps[Comp{l, r}] = struct{}{}
	}
	return comps
}

func strength(vals []int) int {
	s := 0
	for _, b := range vals {
		s += b
	}
	return s
}

func compatible(c Comp, soFar []int) (bool, []int) {
	last := soFar[len(soFar)-1]
	if c.L == last {
		return true, []int{c.L, c.R}
	}
	if c.R == last {
		return true, []int{c.R, c.L}
	}
	return false, nil
}

func removeComp(comps map[Comp]struct{}, c Comp) map[Comp]struct{} {
	res := make(map[Comp]struct{})
	for k := range comps {
		if k != c {
			res[k] = struct{}{}
		}
	}
	return res
}

func addComp(list []int, toAdd []int) []int {
	res := make([]int, len(list)+2)
	for i := 0; i < len(list); i++ {
		res[i] = list[i]
	}
	res[len(res)-2] = toAdd[0]
	res[len(res)-1] = toAdd[1]
	return res
}
