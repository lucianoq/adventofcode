package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	cardPublicKey, doorPublicKey := parse()
	doorLoopSize := bruteForce(7, doorPublicKey)
	encryptionKey := transform(cardPublicKey, doorLoopSize)
	fmt.Println(encryptionKey)
}

func parse() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	card, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	door, _ := strconv.Atoi(scanner.Text())
	return card, door
}

func bruteForce(subject, goal int) int {
	val := 1
	for i := 1; ; i++ {
		val *= subject
		val %= 20201227

		if val == goal {
			return i
		}
	}
}

func transform(subject, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= subject
		val %= 20201227
	}
	return val
}
