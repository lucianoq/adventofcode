package main

import "fmt"

func main() {
	ingredients, foodToIngredients, foodToAllergens := parse()
	_, ingredientToAllergen := findMatches(ingredients, foodToIngredients, foodToAllergens)

	count := map[string]int{}
	for ing := range ingredients {
		if _, ok := ingredientToAllergen[ing]; !ok {

			// ing is not an allergenic ingredient
			// count it
			for _, ings := range foodToIngredients {
				if _, ok := ings[ing]; ok {
					count[ing]++
				}
			}
		}
	}

	sum := 0
	for _, v := range count {
		sum += v
	}
	fmt.Println(sum)
}
