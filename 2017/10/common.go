package main

const (
	size  = 256
	block = 16
)

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
