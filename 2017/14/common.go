package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

const (
	size  = 128
	input = "jzgqcdpd"
)

type Grid [size][size]bool

func NewGrid() *Grid {
	var g Grid
	for row := 0; row < size; row++ {
		s := input + "-" + strconv.Itoa(row)
		s = hash(s)
		buf, err := hex.DecodeString(s)
		if err != nil {
			log.Fatal(err)
		}
		bytemap := ""
		for i := 0; i < len(buf); i++ {
			bytemap += fmt.Sprintf("%8b", buf[i])
		}

		for i := 0; i < size; i++ {
			if bytemap[i] == '1' {
				g[row][i] = true
			}
		}
	}
	return &g
}
