package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type SetStr map[string]struct{}

func (s SetStr) Contains(x string) bool {
	_, ok := s[x]
	return ok
}

func (s SetStr) Add(x string) {
	s[x] = struct{}{}
}

func (s SetStr) Delete(x string) {
	delete(s, x)
}

func (s SetStr) Intersect(s2 SetStr) SetStr {
	result := SetStr{}
	for elem := range s {
		if s2.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Pop extract one element (random) from the set.
func (s SetStr) Pop() string {
	for k := range s {
		return k
	}
	return ""
}

type SetInt map[int]struct{}

func (s SetInt) Add(x int) {
	s[x] = struct{}{}
}

func parse() (SetStr, map[int]SetStr, map[int]SetStr) {
	allIngs := SetStr{}
	ingByFood := map[int]SetStr{}
	allgByFood := map[int]SetStr{}

	re := regexp.MustCompile(`^(.+) \(contains (.+)\)$`)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		sm := re.FindStringSubmatch(line)

		ingByFood[i] = SetStr{}
		allgByFood[i] = SetStr{}

		for _, ing := range strings.Fields(sm[1]) {
			ingByFood[i].Add(ing)
			allIngs.Add(ing)
		}

		for _, allg := range strings.Split(sm[2], ", ") {
			allgByFood[i].Add(allg)
		}
	}

	return allIngs, ingByFood, allgByFood
}

func findMatches(ingredients SetStr, foodToIngredients, foodToAllergens map[int]SetStr) (map[string]string, map[string]string) {
	// Allergen -> Set of food
	foodByAllergens := map[string]SetInt{}

	for foodID, allergens := range foodToAllergens {
		for a := range allergens {
			if _, ok := foodByAllergens[a]; !ok {
				foodByAllergens[a] = SetInt{}
			}
			foodByAllergens[a].Add(foodID)
		}
	}

	potentialIngredientsByAllergen := map[string]SetStr{}

	for a, foods := range foodByAllergens {
		remaining := ingredients
		for foodID := range foods {
			remaining = remaining.Intersect(foodToIngredients[foodID])
		}
		potentialIngredientsByAllergen[a] = remaining
	}

	// I'm sure allergen key is contained in ingredient value
	allergenToIngredient := map[string]string{}
	ingredientToAllergen := map[string]string{}

	for len(allergenToIngredient) != len(foodByAllergens) {
		for a, ingSet := range potentialIngredientsByAllergen {
			if len(ingSet) == 1 {

				// I'm sure about this one
				ingredient := ingSet.Pop()

				allergenToIngredient[a] = ingredient
				ingredientToAllergen[ingredient] = a

				// delete it from all the other sets
				for _, ingSet2 := range potentialIngredientsByAllergen {
					ingSet2.Delete(ingredient)
				}
			}
		}
	}

	return allergenToIngredient, ingredientToAllergen
}
