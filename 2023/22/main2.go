package main

import "fmt"

func main() {
	space := parse()
	_ = space.Fall()

	count := 0
	for i := 0; i < len(space.Bricks); i++ {
		s2 := space.DeepCopy()
		s2.Delete(space.Bricks[i])
		count += s2.Fall()
	}
	fmt.Println(count)
}

func (s *Space) Delete(b *Brick) {
	for x := b.Start.X; x <= b.End.X; x++ {
		for y := b.Start.Y; y <= b.End.Y; y++ {
			for z := b.Start.Z; z <= b.End.Z; z++ {
				delete(s.Map, P{x, y, z})
			}
		}
	}
	delete(s.Bricks, b.ID)
}

func (s *Space) DeepCopy() *Space {
	s2 := NewSpace()
	for k, v := range s.Bricks {
		s2.Bricks[k] = &Brick{
			ID:    v.ID,
			Start: v.Start,
			End:   v.End,
		}
	}
	for k, v := range s.Map {
		s2.Map[k] = v
	}
	return s2
}
