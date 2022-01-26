package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const Size = 119315717514047

// const Size = 10

type Deck struct {
	Step      int
	Direction int // +1 clockwise, -1 counterclockwise
	Top       int
}

func (d *Deck) DealIntoNewStack() {
	d.Top = (d.Top - d.Direction*d.Step + Size) % Size
	d.Direction *= -1
}

func (d *Deck) CutN(n int) {
	d.Top = (d.Top + (d.Direction * d.Step * n) + Size) % Size
}

func (d *Deck) DealWithIncrementN(n int) {
	inv := modinv(n, Size)
	d.Step *= inv
	d.Top *= inv
	log.Printf("deal with increment %d, modinv=%d, newStep=%d, newTop=%d", n, inv, d.Step, d.Top)
}

func (d *Deck) Pick(n int) int {
	current := d.Top
	for i := 0; i < n; i++ {
		current = ((current+d.Direction*d.Step)%Size + Size) % Size
	}
	return current
}

func NewDeck() *Deck {
	return &Deck{
		Step:      1,
		Direction: 1,
		Top:       0,
	}
}

func main() {
	size, iter := big.NewInt(Size), big.NewInt(101741582076661)
	offset, increment := big.NewInt(0), big.NewInt(1)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "deal into new stack" {
			increment.Mul(increment, big.NewInt(-1))
			offset.Add(offset, increment)
			continue
		}

		if strings.HasPrefix(line, "cut") {
			ff := strings.Fields(line)
			n, _ := strconv.Atoi(ff[1])

			offset.Add(offset, big.NewInt(0).Mul(big.NewInt(int64(n)), increment))
			continue
		}

		if strings.HasPrefix(line, "deal with increment") {

			ff := strings.Fields(line)
			n, _ := strconv.Atoi(ff[len(ff)-1])

			increment.Mul(increment, big.NewInt(0).Exp(big.NewInt(int64(n)), big.NewInt(0).Sub(size, big.NewInt(2)), size))
			continue
		}
	}

	finalIncr := big.NewInt(0).Exp(increment, iter, size)

	finalOffs := big.NewInt(0).Exp(increment, iter, size)
	finalOffs.Sub(big.NewInt(1), finalOffs)
	invmod := big.NewInt(0).Exp(big.NewInt(0).Sub(big.NewInt(1), increment), big.NewInt(0).Sub(size, big.NewInt(2)), size)
	finalOffs.Mul(finalOffs, invmod)
	finalOffs.Mul(finalOffs, offset)

	answer := big.NewInt(0).Mul(big.NewInt(2020), finalIncr)
	answer.Add(answer, finalOffs)
	answer.Mod(answer, size)

	fmt.Println(answer)
}

func egcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, y, x := egcd(b%a, a)
	return gcd, x - (b/a)*y, y
}

func modinv(a, m int) int {
	g, x, _ := egcd(a, m)
	if g != 1 {
		log.Fatal("modular inverse does not exist")
	}
	if x < 0 {
		x += m
	}
	return x % m
}
