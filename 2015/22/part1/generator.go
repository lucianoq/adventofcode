package main

func Generate(n int) <-chan []Magic {
	ch := make(chan []Magic)
	go func() {
		generate(n, AllMagics, nil, ch)
		close(ch)
	}()
	return ch
}

func generate(n int, magics []Magic, accum []Magic, ch chan<- []Magic) {
	if n == 0 {
		ch <- accum
		return
	}

	for _, m := range magics {
		newAccum := make([]Magic, 0, len(accum)+1)
		newAccum = append(newAccum, accum...)
		newAccum = append(newAccum, m)
		generate(n-1, magics, newAccum, ch)
	}
}
