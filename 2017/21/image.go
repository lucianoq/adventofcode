package main

import (
	"fmt"
)

type Image []bool

func (im Image) Print() {
	size := im.Len()

	for i := 0; i < size; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if im[i*size+j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	for i := 0; i < size; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	fmt.Println()
}

func (im Image) Div2() bool {
	return len(im)&1 == 0
}

func (im Image) Div3() bool {
	return len(im)%3 == 0
}

func (im Image) Equal(oth Image) bool {
	if len(im) != len(oth) {
		return false
	}

	for i := 0; i < len(im); i++ {
		if im[i] != oth[i] {
			return false
		}
	}

	return true
}

func (im Image) Len() int {
	return sqrt(len(im))
}

func (im Image) Split(smallSize int) []Image {
	bigSize := im.Len() / smallSize

	squares := make([]Image, bigSize*bigSize)
	for I := 0; I < bigSize; I++ {
		for J := 0; J < bigSize; J++ {
			if squares[I*bigSize+J] == nil {
				squares[I*bigSize+J] = make(Image, smallSize*smallSize)
			}
			for i := 0; i < smallSize; i++ {
				for j := 0; j < smallSize; j++ {
					index := I*smallSize*smallSize*bigSize + i*smallSize*bigSize + J*smallSize + j
					squares[I*bigSize+J][i*smallSize+j] = im[index]
				}
			}
		}
	}
	return squares
}

func (im Image) Rotate() Image {
	size := im.Len()
	newIm := make(Image, len(im))
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newI, newJ := j, size-i-1
			newIm[newI*size+newJ] = im[i*size+j]
		}
	}
	return newIm
}

func (im Image) Flip() Image {
	size := im.Len()
	for i := 0; i < size/2; i++ {
		for j := 0; j < size; j++ {
			oppositeI := size - i - 1
			im[oppositeI*size+j], im[i*size+j] = im[i*size+j], im[oppositeI*size+j]
		}
	}
	return im
}

func (im Image) CountOn() int {
	size := im.Len()
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if im[i*size+j] {
				count++
			}
		}
	}
	return count
}
