package main

import "fmt"

func main() {
	list := parse()
	size := len(list)

	for i := 0; i < size-2; i++ {

		for j := i + 1; j < size-1; j++ {

			partial := list[i] + list[j]

			if partial > 2020 {
				continue
			}

			for k := j + 1; k < size; k++ {

				if partial+list[k] == 2020 {
					fmt.Println(list[i] * list[j] * list[k])
					return
				}
			}
		}
	}
}
