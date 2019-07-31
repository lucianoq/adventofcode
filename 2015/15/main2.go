package main

import "fmt"

func main() {
	ingredients := parse()
	ch := Generate(ingredients)

	maxScore := -1
	for x := range ch {
		sc := score(x)
		if sc > maxScore {
			maxScore = sc
		}
	}

	fmt.Println(maxScore)
}

func score(qty []Pair) int {
	var (
		capacity   int
		durability int
		flavor     int
		texture    int
		calories   int
	)

	for _, pair := range qty {
		capacity += pair.Ingredient.Capacity * pair.Spoons
		durability += pair.Ingredient.Durability * pair.Spoons
		flavor += pair.Ingredient.Flavor * pair.Spoons
		texture += pair.Ingredient.Texture * pair.Spoons
		calories += pair.Ingredient.Calories * pair.Spoons
	}

	if calories != 500 {
		return 0
	}

	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0
	}

	return capacity * durability * flavor * texture
}
