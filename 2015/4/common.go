package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func generator(ch chan<- string) {
	for i := 0; ; i++ {
		ch <- input + strconv.Itoa(i)
	}
}

func worker(ch <-chan string, out chan int, prefix string) {
	for str := range ch {
		buf := md5.Sum([]byte(str))
		res := hex.EncodeToString(buf[:])
		if strings.HasPrefix(res, prefix) {
			i, _ := strconv.Atoi(strings.TrimPrefix(str, input))
			out <- i
			return
		}
	}
}
