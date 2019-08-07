package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	parseAndRun()
	fmt.Println(output[0] * output[1] * output[2])
}

func (b *bot) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		b.Store[0] = <-b.C
		b.Store[1] = <-b.C

		sort.Ints(b.Store[:])

		switch b.LowType {
		case "bot":
			bots[b.Low].C <- b.Store[0]
		case "output":
			output[b.Low] = b.Store[0]
		}

		switch b.HighType {
		case "bot":
			bots[b.High].C <- b.Store[1]
		case "output":
			output[b.High] = b.Store[1]
		}

		wg.Done()
	}()
}
