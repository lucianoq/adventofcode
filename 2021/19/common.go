package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point [3]int32

func alignAll(scanners map[uint8]Scanner) map[uint8]Scanner {

	numScanners := len(scanners)

	overlapped := map[uint8]Scanner{0: scanners[0]}
	delete(scanners, 0)

Loop:
	for len(overlapped) < numScanners {
		for _, s1 := range overlapped {
			for _, s2 := range scanners {
				for _, tr := range Transformations {
					transformed := s2.Transform(tr)
					overlap, shift := s1.Overlap(transformed)
					if overlap {
						shifted := transformed.Shift(shift)
						delete(scanners, s2.id)
						overlapped[s2.id] = shifted
						continue Loop
					}
				}
			}
		}
	}
	return overlapped
}

func parse() map[uint8]Scanner {
	scanners := map[uint8]Scanner{}
	var beacons []Point
	var scannerID uint8

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			scanners[scannerID] = Scanner{
				id:      scannerID,
				beacons: beacons,
			}
			continue
		}

		if strings.HasPrefix(line, "---") {
			line = strings.Trim(line, "--- scanner")
			id, _ := strconv.ParseUint(line, 10, 8)
			scannerID = uint8(id)
			beacons = make([]Point, 0, 25)
			continue
		}

		ff := strings.Split(line, ",")
		c1, _ := strconv.ParseInt(ff[0], 10, 32)
		c2, _ := strconv.ParseInt(ff[1], 10, 32)
		c3, _ := strconv.ParseInt(ff[2], 10, 32)
		beacons = append(beacons, Point{int32(c1), int32(c2), int32(c3)})
	}
	scanners[scannerID] = Scanner{
		id:      scannerID,
		beacons: beacons,
	}
	return scanners
}
