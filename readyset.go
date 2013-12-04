// Package readyset provides a simple set implementation.
package readyset

import "fmt"

type Set map[interface{}]struct{}

// NewSet creates a new Set containing the given elements.
func NewSet(in ...interface{}) Set {
	set := Set{}
	set.Add(in...)
	return set
}

// Contains returns whether the given element is in the set.
func (set Set) Contains(i interface{}) bool {
	_, in := set[i]
	return in
}

// Add adds the given elements to the set.
func (set Set) Add(is ...interface{}) {
	for _, i := range is {
		set[i] = struct{}{}
	}
}

// Remove removes the given elements from the set.
func (set Set) Remove(is ...interface{}) {
	for _, i := range is {
		delete(set, i)
	}
}

// Slice returns the elements of the set as a slice.
func (set Set) Slice() []interface{} {
	s := []interface{}{}
	for k, _ := range set {
		s = append(s, k)
	}
	return s
}

// String renders a string representation of the set.
func (set Set) String() string {
	if len(set) == 0 {
		return "[]"
	}
	s := "["
	for k, _ := range set {
		s += fmt.Sprintf("%v ", k)
	}
	return s[:len(s)-1] + "]"
}

// Len returns the size of the set.
func (set Set) Len() int {
	return len(set)
}

// Intersection returns a set containing all elements contained in both given sets.
func Intersection(a, b Set) Set {
	i := NewSet()
	for k, _ := range a {
		if b.Contains(k) {
			i.Add(k)
		}
	}
	for k, _ := range b {
		if a.Contains(k) {
			i.Add(k)
		}
	}
	return i
}

// Intersection returns a set containing all elements contained in either of the given sets.
func Union(a, b Set) Set {
	i := a
	i.Add(b.Slice()...)
	return i
}