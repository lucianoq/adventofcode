package main

import (
	"bufio"
	"os"
	"strings"
)

// Pulse

type Pulse bool

func (p Pulse) String() string {
	if p {
		return "high"
	}
	return "low"
}

const (
	HighPulse Pulse = true
	LowPulse  Pulse = false
)

// Message

type Message struct {
	Pulse    Pulse
	From, To string
}

var Button = Message{LowPulse, "button", "broadcaster"}

type Network map[string]Module

// Modules

type NetworkModule struct {
	Name         string
	Destinations []string
}

type Module interface {
	Process(Message) []Message
}

type FlipFlop struct {
	NetworkModule
	Status bool
}

func (f *FlipFlop) Process(msg Message) []Message {
	if msg.Pulse == LowPulse {
		f.Status = !f.Status

		var q []Message
		for _, d := range f.Destinations {
			q = append(q, Message{Pulse(f.Status), f.Name, d})
		}
		return q
	}
	return nil
}

type Conjunction struct {
	NetworkModule
	Memory map[string]Pulse
}

func (c *Conjunction) Process(msg Message) []Message {
	c.Memory[msg.From] = msg.Pulse

	allHigh := true
	for _, v := range c.Memory {
		if !v {
			allHigh = false
		}
	}

	var q []Message
	for _, d := range c.Destinations {
		if allHigh {
			q = append(q, Message{LowPulse, c.Name, d})
		} else {
			q = append(q, Message{HighPulse, c.Name, d})
		}
	}
	return q
}

type Broadcast struct {
	NetworkModule
}

func (b *Broadcast) Process(msg Message) []Message {
	var q []Message
	for _, d := range b.Destinations {
		q = append(q, Message{msg.Pulse, b.Name, d})
	}
	return q
}

// parse
func parse() Network {
	scanner := bufio.NewScanner(os.Stdin)

	nw := Network{}

	inputs := map[string][]string{}
	conjunctions := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		lr := strings.Split(line, " -> ")

		name := ""
		dest := strings.Split(lr[1], ", ")

		var m Module
		switch lr[0][0] {
		case '%':
			name = lr[0][1:]
			m = &FlipFlop{
				NetworkModule: NetworkModule{
					Name:         name,
					Destinations: dest,
				},
				Status: false,
			}
		case '&':
			name = lr[0][1:]
			m = &Conjunction{
				NetworkModule: NetworkModule{
					Name:         name,
					Destinations: dest,
				},
				Memory: make(map[string]Pulse),
			}
			conjunctions = append(conjunctions, name)
		case 'b':
			name = "broadcaster"
			m = &Broadcast{
				NetworkModule: NetworkModule{
					Name:         name,
					Destinations: dest,
				},
			}
		}

		for _, d := range dest {
			inputs[d] = append(inputs[d], name)
		}

		nw[name] = m
	}

	for _, c := range conjunctions {
		for _, in := range inputs[c] {
			con := nw[c].(*Conjunction)
			con.Memory[in] = false
		}
	}

	return nw
}
