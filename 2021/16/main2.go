package main

import "fmt"

func (p *Packet) Val() int {
	switch p.TypeID {

	case 0:
		sum := 0
		for _, sp := range p.SubPackets {
			sum += sp.Val()
		}
		return sum

	case 1:
		prod := 1
		for _, sp := range p.SubPackets {
			prod *= sp.Val()
		}
		return prod

	case 2:
		min := 1<<63 - 1
		for _, sp := range p.SubPackets {
			val := sp.Val()
			if val < min {
				min = val
			}
		}
		return min

	case 3:
		max := 0
		for _, sp := range p.SubPackets {
			val := sp.Val()
			if val > max {
				max = val
			}
		}
		return max

	case 4:
		return p.Number

	case 5:
		if p.SubPackets[0].Val() > p.SubPackets[1].Val() {
			return 1
		}
		return 0

	case 6:
		if p.SubPackets[0].Val() < p.SubPackets[1].Val() {
			return 1
		}
		return 0

	case 7:
		if p.SubPackets[0].Val() == p.SubPackets[1].Val() {
			return 1
		}
		return 0
	}

	return 0
}

func main() {
	s := parse()
	packet, _ := readPacket(s)
	fmt.Println(packet.Val())
}
