package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

var start = Image{false, true, false, false, false, true, true, true, true}

func parse() (rules []Rule) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, " => ")

		leftS := strings.ReplaceAll(ff[0], "/", "")
		left := make(Image, len(leftS))
		for i := range leftS {
			switch leftS[i] {
			case '.':
				left[i] = false
			case '#':
				left[i] = true
			default:
				log.Fatal("error")
			}
		}

		rightS := strings.ReplaceAll(ff[1], "/", "")
		right := make(Image, len(rightS))
		for i := range rightS {
			switch rightS[i] {
			case '.':
				right[i] = false
			case '#':
				right[i] = true
			default:
				log.Fatal("error")
			}
		}

		rules = append(rules, Rule{left, right})
	}
	return
}

func sqrt(i int) int {
	return int(math.Floor(math.Sqrt(float64(i))))
}

func PixelsOnAfter(times int) int {
	rules := parse()
	img := start
	for i := 0; i < times; i++ {
		img = iteration(rules, img)
	}
	return img.CountOn()
}

func iteration(rules []Rule, im Image) Image {
	var squares []Image

	switch {
	case im.Div2():
		squares = im.Split(2)
	case im.Div3():
		squares = im.Split(3)
	default:
		log.Fatal("not div by 2 or 3")
	}

	for i := range squares {
		squares[i] = ApplyOneRule(rules, squares[i])
	}

	return Merge(squares)
}

func ApplyOneRule(rules []Rule, im Image) Image {
	for _, r := range rules {
		if r.Match(im) {
			return r.Apply(im)
		}
	}
	log.Fatal("doesn't match with anything")
	return Image{}
}

func Merge(squares []Image) Image {
	bSize := sqrt(len(squares))
	sSize := sqrt(len(squares[0]))

	merged := make(Image, sSize*sSize*bSize*bSize)

	for I := 0; I < bSize; I++ {
		for J := 0; J < bSize; J++ {
			for i := 0; i < sSize; i++ {
				for j := 0; j < sSize; j++ {
					index := I*sSize*sSize*bSize + i*sSize*bSize + J*sSize + j
					merged[index] = squares[I*bSize+J][i*sSize+j]
				}
			}
		}
	}

	return merged
}
