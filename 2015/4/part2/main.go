package part2

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const input = "yzbqklnj"

func main() {
	inputs := make(chan string, 1000)
	output := make(chan int)

	go generator(inputs)
	for j := 0; j < 30; j++ {
		go worker(inputs, output)
	}

	i := <-output
	fmt.Println(i)
}

func generator(ch chan<- string) {
	for i := 0; ; i++ {
		ch <- input + strconv.Itoa(i)
	}
}

func worker(ch <-chan string, out chan int) {
	for str := range ch {
		buf := md5.Sum([]byte(str))
		res := hex.EncodeToString(buf[:])
		if strings.HasPrefix(res, "000000") {
			i, _ := strconv.Atoi(strings.TrimPrefix(str, input))
			out <- i
			return
		}
	}

}
