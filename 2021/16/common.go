package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Packet struct {
	Version    int
	TypeID     int
	Number     int
	SubPackets []*Packet
}

func parse() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	buf := ""
	for i := 0; i < len(line)-1; i += 2 {
		b, err := strconv.ParseUint(line[i:i+2], 16, 8)
		if err != nil {
			log.Fatal(err)
		}
		buf += fmt.Sprintf("%08b", b)
	}
	return buf
}

func readPackets(s string) []*Packet {
	packets := []*Packet{}
	for len(s) >= 3+3+5 {
		p, rem := readPacket(s)
		packets = append(packets, p)
		s = rem
	}
	return packets
}

func readPacket(s string) (*Packet, string) {

	packet := &Packet{
		Version: mustParseBin(s[:3]),
		TypeID:  mustParseBin(s[3:6]),
	}
	s = s[6:]

	// literal packet
	if packet.TypeID == 4 {
		bits := ""
		for {
			lastGroup := s[0] == '0'
			bits += s[1:5]
			s = s[5:]

			if lastGroup {
				break
			}
		}
		packet.Number = mustParseBin(bits)
		return packet, s
	}

	// operator packet

	switch {

	// If the length type ID is 0, then the next 15 bits are a number
	// that represents the total length in bits of the sub-packets
	// contained by this packet.
	case s[0] == '0':
		length := mustParseBin(s[1 : 1+15])
		s = s[1+15:]

		subString := s[:length]
		packet.SubPackets = readPackets(subString)
		s = s[length:]

	// If the length type ID is 1, then the next 11 bits are a number
	// that represents the number of sub-packets immediately contained
	// by this packet.
	case s[0] == '1':
		numSubPackets := mustParseBin(s[1 : 1+11])
		s = s[1+11:]

		for i := 0; i < numSubPackets; i++ {
			p, rem := readPacket(s)
			packet.SubPackets = append(packet.SubPackets, p)
			s = rem
		}
	}
	return packet, s
}

func mustParseBin(s string) int {
	number, _ := strconv.ParseInt(s, 2, 64)
	return int(number)
}
