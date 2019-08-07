package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type bot struct {
	ID       int
	C        chan int
	Store    [2]int
	Low      int
	LowType  string
	High     int
	HighType string
}

var (
	bots   = make(map[int]*bot)
	output = make(map[int]int)
)

func CreateBot(wg *sync.WaitGroup, id int, low int, lowType string, high int, highType string) {
	if _, ok := bots[id]; ok {
		return
	}

	bots[id] = &bot{
		ID:       id,
		C:        make(chan int),
		Store:    [2]int{},
		Low:      low,
		LowType:  lowType,
		High:     high,
		HighType: highType,
	}
	bots[id].Start(wg)
}

func parseAndRun() {
	type input struct{ chip, bot int }

	wg := &sync.WaitGroup{}
	inputs := make([]input, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "value") {
			var chip, bot int
			_, _ = fmt.Sscanf(line, "value %d goes to bot %d", &chip, &bot)
			inputs = append(inputs, input{chip, bot})
		} else {
			var bot, low, high int
			var lowType, highType string

			_, _ = fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &bot, &lowType, &low, &highType, &high)

			CreateBot(wg, bot, low, lowType, high, highType)
		}
	}

	for _, i := range inputs {
		bots[i.bot].C <- i.chip
	}

	wg.Wait()
}
