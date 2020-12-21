package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	allergenToIngredient, _ := findMatches(parse())

	ls := make([]string, 0, len(allergenToIngredient))
	for k := range allergenToIngredient {
		ls = append(ls, k)
	}
	sort.Strings(ls)

	ing := make([]string, 0, len(allergenToIngredient))
	for _, a := range ls {
		ing = append(ing, allergenToIngredient[a])
	}

	fmt.Println(strings.Join(ing, ","))
}
