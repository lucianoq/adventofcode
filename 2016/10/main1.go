package main

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

func main() {
	parseAndRun()
}

func (b *bot) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		b.Store[0] = <-b.C
		b.Store[1] = <-b.C

		sort.Ints(b.Store[:])

		if b.Store[0] == 17 && b.Store[1] == 61 {
			fmt.Println(b.ID)
			os.Exit(0)
		}

		if b.LowType == "bot" {
			bots[b.Low].C <- b.Store[0]
		}

		if b.HighType == "bot" {
			bots[b.High].C <- b.Store[1]
		}

		wg.Done()
	}()
}
