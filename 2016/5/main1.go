package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var input = "ugkcyxxp"

func main() {
	inputBuf := []byte(input)
	password := ""
	var buf = make([]byte, 0, len(input)+9)
	for i := 0; ; i++ {
		buf = strconv.AppendInt(inputBuf, int64(i), 10)
		md5Sum := md5.Sum(buf)

		// n hex string chars means n/2 bytes
		if md5Sum[0] == 0 && md5Sum[1] == 0 && md5Sum[2]>>4 == 0 {
			hexString := hex.EncodeToString(md5Sum[:])
			password += string(hexString[5])
			if len(password) == 8 {
				break
			}
		}
	}
	fmt.Println(password)
}
