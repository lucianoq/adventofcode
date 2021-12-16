package main

import "fmt"

func (p *Packet) SumVersion() int {
	sum := p.Version
	for _, subP := range p.SubPackets {
		sum += subP.SumVersion()
	}
	return sum
}

func main() {
	s := parse()
	packet, _ := readPacket(s)
	fmt.Println(packet.SumVersion())
}
