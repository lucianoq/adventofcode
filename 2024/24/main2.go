package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	facts, rules := parseInput()

	//GenerateDot(rules)
	//return

	// After generating the dot file, I could generate an image with
	//
	//   dot -Tsvg -O graph.dot
	//
	// and thanks to the colors of the arrows, it was easy to spot inconsistencies.
	// E.g.
	//   All zXX were reached by two red arrows (a XOR operation).
	//   `z13` was reached by two green, meaning it needed to be swapped with a node
	//      that is reached by two red arrows instead of green.
	//   `z19` and `z33` were incorrectly blue.
	//
	// The full list was:
	// - z13 green (supposed red)
	// - z19 blue (supposed red)
	// - z33 blue (supposed red)
	// - gws blue (supposed red)
	// - hgj red (supposed blue)
	// - npf red (supposed green)
	// - nnt red (supposed blue)
	// - cph red (supposed blue)

	list := []string{"z13", "z19", "z33", "npf", "hgj", "gws", "nnt", "cph"}
	sort.Strings(list)
	fmt.Println(strings.Join(list, ","))

	return

	// The assignment didn't ask to fix the gates (so the list above is enough
	// to generate the solution:
	//
	// Anyway it's pretty easy to spot the pairs:
	// - z13 <> npf  (the only green/red pair)
	// - gws <> nnt  (they are both in the z09 area: a different swap would change the graph massively)
	// - z19 <> cph  (they are both in the z19 area)
	// - z33 <> hgj  (they are both in the z33 area)

	swap := func(r1, r2 string) {
		rules[r1], rules[r2] = rules[r2], rules[r1]
	}
	swap("z13", "npf")
	swap("z33", "hgj")
	swap("gws", "nnt")
	swap("z19", "cph")

	x := GetNumber("x", facts, rules)
	y := GetNumber("y", facts, rules)
	z := GetNumber("z", facts, rules)
	fmt.Printf("  x: %d +\n  y: %d =\nx+y: %d\n  z: %d\n", x, y, x+y, z)
	fmt.Printf("Sum is working = %t\n", x+y == z)

	//   x: 22423510084409 +
	//   y: 19634559019535 =
	// x+y: 42058069103944
	//   z: 42058069103944
	// Sum is working = true
}

func GenerateDot(rules map[string]Rule) {
	f, err := os.Create("graph.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	color := map[string]string{
		"XOR": "red",
		"AND": "blue",
		"OR":  "green",
	}

	fmt.Fprint(f, "digraph {\n")
	for rName, r := range rules {
		fmt.Fprintf(f, "%s -> %s [color=\"%s\"];\n", r.Left, rName, color[r.Op])
		fmt.Fprintf(f, "%s -> %s [color=\"%s\"];\n", r.Right, rName, color[r.Op])
	}

	var xys, zs []string
	for i := 0; i <= 45; i++ {
		xys = append(xys, fmt.Sprintf("x%02d", i))
		xys = append(xys, fmt.Sprintf("y%02d", i))
		zs = append(zs, fmt.Sprintf("z%02d", i))
	}
	fmt.Fprintf(f, "{rank = min;\n %s ; \n};\n", strings.Join(xys, " -> "))
	fmt.Fprintf(f, "{rank = max;\n %s ; \n};\n", strings.Join(zs, " -> "))
	fmt.Fprint(f, "}\n")
}
