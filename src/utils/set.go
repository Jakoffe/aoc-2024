package utils

import "fmt"

type Set[E comparable] map[E]struct{}

func NewSet[E comparable](vals ...E) Set[E] {
	s := Set[E]{}
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[E]) Add(vals ...E) Set[E] {
	for _, v := range vals {
		s[v] = struct{}{}
	}

	return s
}

func (s Set[E]) Remove(vals ...E) Set[E] {
	for _, v := range vals {
		delete(s, v)
	}

	return s
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

func (s Set[E]) Members() []E {
	result := make([]E, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}

func (s Set[E]) Size() int {
	return len(s)
}

func (s Set[E]) String() string {
	return fmt.Sprintf("%v", s.Members())
}

func (s Set[E]) Print() {
	fmt.Printf("%v", s.Members())
}

func (s Set[E]) Union(s2 Set[E]) Set[E] {
	result := NewSet(s.Members()...)
	result.Add(s2.Members()...)
	return result
}

func (s Set[E]) Intersection(s2 Set[E]) Set[E] {
	result := NewSet[E]()
	for _, v := range s.Members() {
		if s2.Contains(v) {
			result.Add(v)
		}
	}
	return result
}
