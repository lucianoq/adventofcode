package main

type Ring struct {
	link map[int]int
	curr int
}

func NewRing() *Ring {
	return &Ring{
		link: map[int]int{},
	}
}

func (r *Ring) Insert(x int) {
	if len(r.link) == 0 {
		r.link[x] = x
		r.curr = x
		return
	}
	r.link[r.curr], r.link[x] = x, r.link[r.curr]
}

func (r *Ring) Next() {
	r.curr = r.link[r.curr]
}

func (r *Ring) Value() int {
	return r.curr
}

func (r *Ring) MoveTo(x int) {
	r.curr = x
}

func (r *Ring) PopNext() int {
	val := r.link[r.curr]
	r.link[r.curr] = r.link[val]
	delete(r.link, val)
	return val
}

func (r *Ring) PickUp3() [3]int {
	var res [3]int
	for i := 0; i < 3; i++ {
		res[i] = r.PopNext()
	}
	return res
}

func (r *Ring) Do(f func(x int)) {
	tmp := r.curr
	for i := 0; i < len(r.link); i++ {
		f(tmp)
		tmp = r.link[tmp]
	}
}
