package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	boxes := [256]Box{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, seq := range strings.Split(scanner.Text(), ",") {

		if strings.HasSuffix(seq, "-") {
			label := seq[:len(seq)-1]

			boxes[hash(label)].Remove(label)
		}

		if strings.Contains(seq, "=") {
			split := strings.Split(seq, "=")
			label := split[0]
			focalLen, _ := strconv.Atoi(split[1])

			boxes[hash(label)].Replace(Lens{label, focalLen})
		}
	}

	fmt.Println(focusingPower(boxes))
}

type Lens struct {
	Label    string
	FocalLen int
}

type Box []Lens

func (b *Box) Remove(label string) {
	for i := 0; i < len(*b); i++ {
		if (*b)[i].Label == label {
			*b = append((*b)[:i], (*b)[i+1:]...)
			return
		}
	}
}

func (b *Box) Replace(lens Lens) {
	for i := 0; i < len(*b); i++ {
		if (*b)[i].Label == lens.Label {
			(*b)[i].FocalLen = lens.FocalLen
			return
		}
	}
	*b = append(*b, lens)
}

func focusingPower(boxes [256]Box) int {
	sum := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < len(boxes[i]); j++ {
			sum += (i + 1) * (j + 1) * boxes[i][j].FocalLen
		}
	}
	return sum
}
