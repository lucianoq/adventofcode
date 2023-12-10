package main

import "fmt"

func main() {
	m, start := parse()

	mainPipe := map[P]struct{}{start: {}}
	s := State{start.E(), W}
	for s.Current != start {
		mainPipe[s.Current] = struct{}{}
		s = s.Next(m)
	}

	count := 0
	for r := 0; r < 140; r++ {

		inside := false
		temp := ""

		for c := 0; c < 140; c++ {

			p := P{r, c}

			if _, ok := mainPipe[p]; !ok {
				if inside {
					count++
				}
				continue
			}

			switch m[p] {
			case '|':
				inside = !inside
			case 'L', 'F', '7', 'J':
				temp += string(m[p])
				switch temp {
				case "LJ", "F7": // U-shaped, they cancel each other
					temp = ""
				case "L7", "FJ": // like a single vertical line, switch inside
					inside = !inside
					temp = ""
				}
			}
		}
	}
	fmt.Println(count)
}
