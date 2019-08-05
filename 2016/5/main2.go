package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

var input = "ugkcyxxp"

func main() {
	inputBuf := []byte(input)
	password := make(map[uint8]byte)

	var buf = make([]byte, 0, len(input)+9)
	for i := 0; ; i++ {
		buf = strconv.AppendInt(inputBuf, int64(i), 10)
		md5Sum := md5.Sum(buf)

		// n hex string chars means n/2 bytes
		if md5Sum[0] == 0 && md5Sum[1] == 0 && md5Sum[2]>>4 == 0 {
			hexString := hex.EncodeToString(md5Sum[:])
			log.Println(hexString)

			position := hexString[5] - 48
			if position < 0 || position > 7 {
				log.Println("ignored because position = ", position)
				continue
			}

			if _, ok := password[position]; ok {
				log.Println("ignored because already found ", position)
				continue
			}

			password[position] = hexString[6]

			for i := uint8(0); i < 8; i++ {
				if password[i] == 0 {
					fmt.Print("_")
				} else {
					fmt.Print(string(password[i]))
				}
			}
			fmt.Println()

			if len(password) == 8 {
				break
			}
		}
	}

	for i := uint8(0); i < 8; i++ {
		fmt.Print(string(password[i]))
	}
	fmt.Println()
}
