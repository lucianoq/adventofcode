package main

import "fmt"

var solution = []string{
	"west", "south", "take pointer",
	"north", "east", "east", "south", "south", "take space heater",
	"north", "north", "north", "take wreath",
	"north", "west", "take dehydrated water",
	"north", "east", "south",
}

func main() {
	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		NewVM("input", input, output).Run()
		close(output)
	}()

	// To play from stdin,
	// enable this goroutine and disable the next one

	// go func() {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	for {
	// 		text, err := reader.ReadString('\n')
	// 		if err != nil {
	// 			continue
	// 		}
	// 		if text != "" {
	// 			for _, c := range text {
	// 				input <- int(c)
	// 			}
	// 		}
	// 	}
	// }()
	go func() {
		for _, cmd := range solution {
			for _, c := range cmd {
				input <- int(c)
			}
			input <- '\n'
		}
	}()

	for x := range output {
		fmt.Print(string(x))
	}
}
