package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const input = "abcdefghijklmnop"

type Cmd struct {
	Op                 byte
	SpinSize           int
	Ex1, Ex2           int
	Partner1, Partner2 byte
}

func (c Cmd) Apply(s []byte) {
	switch c.Op {
	case 's':
		i := len(s) - c.SpinSize
		copy(s, append(s[i:], s[:i]...))
	case 'x':
		s[c.Ex1], s[c.Ex2] = s[c.Ex2], s[c.Ex1]
	case 'p':
		for i := 0; i < len(s); i++ {
			var found1, found2 bool
			switch s[i] {
			case c.Partner1:
				s[i] = c.Partner2
				found1 = true
			case c.Partner2:
				s[i] = c.Partner1
				found2 = true
			}
			if found1 && found2 {
				break
			}
		}
	}
}

func parse() []Cmd {
	buf, _ := ioutil.ReadAll(os.Stdin)
	buf = bytes.TrimSpace(buf)
	ff := strings.Split(string(buf), ",")

	cmds := make([]Cmd, 0)
	for _, f := range ff {
		cmd := Cmd{Op: f[0]}
		switch cmd.Op {
		case 's':
			cmd.SpinSize, _ = strconv.Atoi(f[1:])
		case 'x':
			args := strings.Split(f[1:], "/")
			cmd.Ex1, _ = strconv.Atoi(args[0])
			cmd.Ex2, _ = strconv.Atoi(args[1])
		case 'p':
			args := strings.Split(f[1:], "/")
			cmd.Partner1 = []byte(args[0])[0]
			cmd.Partner2 = []byte(args[1])[0]
		}
		cmds = append(cmds, cmd)
	}
	return cmds
}
