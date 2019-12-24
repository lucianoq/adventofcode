package main

import "fmt"

const Size = 50

func main() {
	done := make(chan struct{})

	vms := [Size]*VM{}

	// start all
	for i := 0; i < Size; i++ {
		vms[i] = NewVM("input", 10000)
		go vms[i].Run()
		vms[i].Input <- i  // address ID
		vms[i].Input <- -1 // noop necessary for routers that are creating packets
	}

	for i := 0; i < Size; i++ {
		go router(vms, i, done)
	}

	<-done
}

func router(vms [Size]*VM, id int, done chan<- struct{}) {
	vm := vms[id]
	for {
		address := <-vm.Output
		x := <-vm.Output
		y := <-vm.Output
		if address == 255 {
			fmt.Println(y)
			done <- struct{}{}
			return
		}

		vms[address].Input <- x
		vms[address].Input <- y
	}
}
