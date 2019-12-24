package main

import (
	"fmt"
	"time"
)

const Size = 50

func main() {
	vms := [Size]*VM{}
	natCh := make(chan int)

	// start all
	for i := 0; i < Size; i++ {
		vms[i] = NewVM("input", 10000)
		go vms[i].Run()
		vms[i].Input <- i  // address ID
		vms[i].Input <- -1 // noop necessary for routers that are creating packets
	}

	for i := 0; i < Size; i++ {
		go router(vms, i, natCh)
	}

	fmt.Println(nat(vms, natCh))
}

func router(vms [Size]*VM, id int, natCh chan<- int) {
	vm := vms[id]
	for {
		address := <-vm.Output
		x := <-vm.Output
		y := <-vm.Output
		if address == 255 {
			natCh <- x
			natCh <- y
		} else {
			vms[address].Input <- x
			vms[address].Input <- y
		}
	}
}

func nat(vms [Size]*VM, natCh <-chan int) int {
	x := -1
	y := -1

	// continuously overwrite package
	go func() {
		for {
			x = <-natCh
			y = <-natCh
		}
	}()

	// check idle
	lastY := -1
	for {
		time.Sleep(10 * time.Millisecond)
		if x != -1 && idle(vms) {
			vms[0].Input <- x
			vms[0].Input <- y
			if y == lastY {
				return y
			}
			lastY = y
		}
	}
}

func idle(vms [Size]*VM) bool {
	for i := 0; i < Size; i++ {
		if len(vms[i].Input) != 0 {
			return false
		}
	}
	return true
}
