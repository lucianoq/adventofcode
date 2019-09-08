package main

import "fmt"

func main() {
	start := Building{}
	start.Floors[0].Add(
		PoloniumRTG,
		ThuliumRTG,
		ThuliumChip,
		PromethiumRTG,
		RutheniumRTG,
		RutheniumChip,
		CobaltRTG,
		CobaltChip,
		EleriumChip,
		EleriumRTG,
		DilithiumChip,
		DilithiumRTG,
	)
	start.Floors[1].Add(PoloniumChip, PromethiumChip)

	end := Building{}
	end.Floors[3] = 0x7f7f

	final := bfs(start, end)

	steps := countPath(final)

	fmt.Println(steps)
}
