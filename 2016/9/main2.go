package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	str, _ := ioutil.ReadAll(os.Stdin)

	i := 0

	for bytes.IndexByte(str, '(') != -1 {
		//for findP(str) != -1 {
		log.Printf("Starting iteration %d", i)
		log.Printf("str is now long %d", len(str))
		str = decompress(str)
		i++
	}

	fmt.Println(len(str))
	fmt.Println(string(str))
}
