package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Element struct {
	Name string
	Qty  int
}

type Reaction struct {
	Reagents []Element
	Product  Element
}

func costPerFuel(reactions map[string]Reaction, fuel int) int {
	needed := map[string]int{"FUEL": fuel}
	excess := map[string]int{}

	for len(needed) != 1 || needed["ORE"] == 0 {
		for k := range needed {
			if k != "ORE" {
				apply(reactions[k], needed, excess)
				break
			}
		}
	}
	return needed["ORE"]
}

func apply(r Reaction, needed, excess map[string]int) {
	obj := r.Product.Name

	if excess[obj] > 0 {
		if excess[obj] >= needed[obj] {
			excess[obj] -= needed[obj]
			if excess[obj] == 0 {
				delete(excess, obj)
			}
			delete(needed, obj)
			return
		}

		needed[obj] -= excess[obj]
		delete(excess, obj)
	}

	delete(excess, obj)

	if needed[obj] >= r.Product.Qty {
		howManyCompleteReactions := needed[obj] / r.Product.Qty

		needed[obj] -= howManyCompleteReactions * r.Product.Qty
		for _, i := range r.Reagents {
			needed[i.Name] += howManyCompleteReactions * i.Qty
		}
	}

	if needed[obj] > 0 {
		for _, i := range r.Reagents {
			needed[i.Name] += i.Qty
		}
		excess[obj] += r.Product.Qty - needed[obj]
	}
	delete(needed, obj)
}

func parse() map[string]Reaction {
	splitQtyName := func(s string) (int, string) {
		rs := strings.Split(s, " ")
		num, err := strconv.Atoi(rs[0])
		if err != nil || num == 0 {
			log.Fatal(err)
		}
		name := rs[1]
		return num, name
	}

	reactions := make(map[string]Reaction, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " => ")

		reagents := make([]Element, 0)
		reags := strings.Split(parts[0], ",")
		for _, reag := range reags {
			reagQty, reagName := splitQtyName(strings.TrimSpace(reag))
			reagents = append(reagents, Element{reagName, reagQty})
		}

		prodQty, prodName := splitQtyName(parts[1])
		reactions[prodName] = Reaction{
			Product:  Element{prodName, prodQty},
			Reagents: reagents,
		}
	}
	return reactions
}
