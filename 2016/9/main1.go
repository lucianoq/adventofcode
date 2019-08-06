package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	str, _ := ioutil.ReadAll(os.Stdin)
	res := decompress(str)
	fmt.Println(len(res))
}

