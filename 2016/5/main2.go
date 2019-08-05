package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
)

var input = "ugkcyxxp"

func main() {
	inputBuf := []byte(input)
	password := make(map[int]byte)

	var buf = make([]byte, 0, len(input)+9)
	for i := 0; ; i++ {
		buf = strconv.AppendInt(inputBuf, int64(i), 10)
		md5Sum := md5.Sum(buf)

		// n hex string chars means n/2 bytes
		// checking like this helps to lazy call EncodeToString
		if md5Sum[0]|md5Sum[1]|md5Sum[2]>>4 == 0 {
			hexString := hex.EncodeToString(md5Sum[:])

			position := int(hexString[5] - 48) // 48 ASCII code of '0'

			// ignore wrong position
			if position < 0 || position > 7 {
				continue
			}

			// ignore already found
			if _, ok := password[position]; ok {
				continue
			}

			password[position] = hexString[6]

			fmt.Fprint(os.Stderr, "\033[H\033[2J")
			printPassword(os.Stderr, password)

			if len(password) == 8 {
				break
			}
		}
	}

	printPassword(os.Stdout, password)
}

func printPassword(w io.Writer, password map[int]byte) {
	for i := 0; i < 8; i++ {
		if password[i] == 0 {
			_, _ = fmt.Fprint(w, "_")
		} else {
			_, _ = fmt.Fprint(w, string(password[i]))
		}
	}
	_, _ = fmt.Fprintln(w)
}
