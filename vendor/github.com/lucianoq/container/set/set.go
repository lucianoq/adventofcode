package set

type Set[T comparable] map[T]struct{}

// Contains checks if the current set contains an element.
func (s Set[T]) Contains(elem T) bool {
	if _, ok := s[elem]; ok {
		return true
	}
	return false
}

// Add adds one or more elements to the set.
func (s Set[T]) Add(elements ...T) {
	for _, elem := range elements {
		s[elem] = struct{}{}
	}
}

// AddSets adds all the elements of one or more sets to the current set.
func (s Set[T]) AddSets(sets ...Set[T]) {
	for _, set := range sets {
		for elem := range set {
			s.Add(elem)
		}
	}
}

// Intersect returns a new set containing all the common elements between the
// current set and the other.
func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	intersection := New[T]()
	for elem := range s {
		if !s2.Contains(elem) {
			intersection.Add(elem)
		}
	}
	return intersection
}

// Sub returns a new Set with all the elements of the current set that are not
// contained in the other.
func (s Set[T]) Sub(s2 Set[T]) Set[T] {
	intersection := New[T]()
	for elem := range s {
		if !s2.Contains(elem) {
			intersection.Add(elem)
		}
	}
	return intersection
}

// Remove deletes an element from the current set.
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Len returns the number of elements in the current set.
func (s Set[T]) Len() int {
	return len(s)
}

// Pop returns a random element from the current set removing it.
func (s Set[T]) Pop() T {
	var x T
	for elem := range s {
		x = elem
		break
	}
	delete(s, x)
	return x
}

// Clear removes all the elements in the current set.
func (s Set[T]) Clear() {
	for elem := range s {
		delete(s, elem)
	}
}

// New creates a new Set.
func New[T comparable]() Set[T] {
	return Set[T]{}
}
