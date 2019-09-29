package main

type Rule struct {
	Left, Right Image
}

func (r Rule) Match(im Image) bool {
	if r.Left.Equal(im) {
		return true
	}

	for i := 0; i < 3; i++ {
		im = im.Rotate()
		if r.Left.Equal(im) {
			return true
		}
	}

	im = im.Flip()
	if r.Left.Equal(im) {
		return true
	}

	for i := 0; i < 3; i++ {
		im = im.Rotate()
		if r.Left.Equal(im) {
			return true
		}
	}

	return false
}

func (r Rule) Apply(im Image) Image {
	if r.Match(im) {
		return r.Right
	}
	return im
}
