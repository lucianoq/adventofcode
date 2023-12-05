package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Map []Transform

type Transform struct {
	Destination, Source, Length int
}

func parse() ([]int, [7]Map) {
	scanner := bufio.NewScanner(os.Stdin)

	seeds := []int{}
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i := 1; i < len(fields); i++ {
		num, _ := strconv.Atoi(fields[i])
		seeds = append(seeds, num)
	}

	return seeds, parseMaps(scanner)
}

func parseMaps(scanner *bufio.Scanner) [7]Map {
	maps := [7]Map{}
	t := -1
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			t++
			continue
		}

		fields := strings.Fields(line)

		dest, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		length, _ := strconv.Atoi(fields[2])

		maps[t] = append(maps[t], Transform{
			Destination: dest,
			Source:      source,
			Length:      length,
		})
	}
	return maps
}
