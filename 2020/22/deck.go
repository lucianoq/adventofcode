package main

// Deck is an enhanced Queue
type Deck []int

func (d *Deck) Len() int {
	return len(*d)
}

// Dequeue
func (d *Deck) Draw() int {
	var head int
	head, *d = (*d)[0], (*d)[1:]
	return head
}

// Enqueue
func (d *Deck) Append(x int) {
	*d = append(*d, x)
}

func (d *Deck) Score() int {
	s := 0
	ln := len(*d)
	for i := 0; i < ln; i++ {
		s += (*d)[i] * (ln - i)
	}
	return s
}

//
// only for part2
//

func (d *Deck) ToList() []int {
	return *d
}

func (d *Deck) CopyN(n int) *Deck {
	cp := make(Deck, n)
	copy(cp, *d)
	return &cp
}
