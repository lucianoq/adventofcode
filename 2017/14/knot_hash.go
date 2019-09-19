package main

import "fmt"

const (
	inputSize = 256
	blockSize = 16
)

func hash(inputStr string) string {
	list := createList(inputSize)

	lengths := []uint8(inputStr)
	lengths = append(lengths, []uint8{17, 31, 73, 47, 23}...)

	var pos, skipSize int
	for i := 0; i < 64; i++ {
		for _, l := range lengths {
			list = knot(list, pos, int(l))

			pos += int(l) + skipSize
			skipSize++
		}
	}

	return toString(dense(list))
}

func createList(length int) []uint8 {
	list := make([]uint8, length)
	for i := 0; i < length; i++ {
		list[i] = uint8(i)
	}
	return list
}

func knot(list []uint8, pos int, l int) []uint8 {
	ls := make([]uint8, len(list))

	j := pos
	for i := l + pos - 1; j < l+pos; i-- {
		ls[j%len(list)] = list[i%len(list)]
		j++
	}

	for i := l + pos; j < pos+len(list); i++ {
		ls[j%len(list)] = list[i%len(list)]
		j++
	}
	return ls
}

func dense(list []uint8) [blockSize]uint8 {
	var res [blockSize]uint8

	for i := 0; i < blockSize; i++ {
		var acc uint8
		for j := 0; j < blockSize; j++ {
			acc ^= list[i*blockSize+j]
		}
		res[i] = acc
	}

	return res
}

func toString(list [blockSize]uint8) string {
	s := ""
	for _, i := range list {
		s += fmt.Sprintf("%02x", i)
	}
	return s
}
