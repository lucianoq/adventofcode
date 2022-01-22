package main

import "fmt"

func main() {
	rules := parse()

	cuboids := []Cuboid{}

	for _, rule := range rules {
		toAdd := []Cuboid{}

		for _, a := range cuboids {
			if intersection, intersect := rule.intersect(a); intersect {

				switch {

				// intersection is counted double as positive, subtract
				case rule.state && a.state:
					intersection.state = false

				// intersection is counted double as negative, add
				case !rule.state && !a.state:
					intersection.state = true

				default:
					// last applied rule wins,
					// so A∩B count as last state
					intersection.state = rule.state
				}

				toAdd = append(toAdd, intersection)
			}
		}

		if rule.state {
			toAdd = append(toAdd, rule)
		}

		cuboids = append(cuboids, toAdd...)
	}

	volume := 0
	for _, c := range cuboids {
		if c.state {
			volume += c.volume()
		} else {
			volume -= c.volume()
		}
	}

	fmt.Println(volume)
}

func (c Cuboid) intersect(c2 Cuboid) (Cuboid, bool) {
	xMin := max(c.minX, c2.minX)
	xMax := min(c.maxX, c2.maxX)
	if xMin > xMax {
		return Cuboid{}, false
	}

	yMin := max(c.minY, c2.minY)
	yMax := min(c.maxY, c2.maxY)
	if yMin > yMax {
		return Cuboid{}, false
	}

	zMin := max(c.minZ, c2.minZ)
	zMax := min(c.maxZ, c2.maxZ)
	if zMin > zMax {
		return Cuboid{}, false
	}

	return Cuboid{
		minX: xMin,
		maxX: xMax,
		minY: yMin,
		maxY: yMax,
		minZ: zMin,
		maxZ: zMax,
	}, true
}

func (c Cuboid) volume() int {
	return (c.maxX + 1 - c.minX) * (c.maxY + 1 - c.minY) * (c.maxZ + 1 - c.minZ)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
