package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

type Pair struct {
	Ingredient Ingredient
	Spoons     int
}

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

func parse() []Ingredient {
	ingredients := make([]Ingredient, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			var i Ingredient
			n, _ := fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &i.Name, &i.Capacity, &i.Durability, &i.Flavor, &i.Texture, &i.Calories)
			if n != 6 {
				log.Fatal("parse failed")
			}
			i.Name = strings.TrimRight(i.Name, ":")
			ingredients = append(ingredients, i)
		}
	}
	return ingredients
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

// Returns a read-only channel of all combinations of 100 spoons.
// It will be closed after it ends.
func Generate(ingredients []Ingredient) <-chan []Pair {
	ch := make(chan []Pair)
	go func() {
		generator(ingredients, 100, nil, ch)
		close(ch)
	}()
	return ch
}

// Recursive function to find all cases
func generator(ingredients []Ingredient, spoonLeft int, accumulator []Pair, ch chan<- []Pair) {
	if len(ingredients) == 1 {
		ch <- append(accumulator, Pair{ingredients[0], spoonLeft})
		return
	}

	for a := 0; a <= spoonLeft; a++ {
		generator(ingredients[1:], spoonLeft-a, append(accumulator, Pair{ingredients[0], a}), ch)
	}
}
