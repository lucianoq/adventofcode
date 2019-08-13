package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func hashSingle(salt string, i int) string {
	hash := md5.Sum([]byte(salt + strconv.Itoa(i)))
	return hex.EncodeToString(hash[:])
}

func hashRecycleBuffer(salt string, x int) string {
	buf := make([]byte, 32)

	first := md5.Sum([]byte(salt + strconv.Itoa(x)))
	hex.Encode(buf, first[:])

	for i := 0; i < 2016; i++ {
		tmp := md5.Sum(buf)
		hex.Encode(buf, tmp[:])
	}
	return string(buf)
}

func hashByteSlice(salt string, x int) string {
	sum := []byte(salt + strconv.Itoa(x))
	for i := 0; i < 2017; i++ {
		tmp := md5.Sum(sum)
		tmp2 := make([]byte, 32)
		hex.Encode(tmp2, tmp[:])
		sum = tmp2
	}
	return string(sum)
}

func hashString(salt string, x int) string {
	sum := salt + strconv.Itoa(x)
	for i := 0; i < 2017; i++ {
		hash := md5.Sum([]byte(sum))
		sum = hex.EncodeToString(hash[:])
	}
	return sum
}
