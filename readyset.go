package readyset

import "fmt"

type Set map[interface{}]struct{}

func NewSet(in ...interface{}) Set {
	set := Set{}
	set.Add(in...)
	return set
}

func (set Set) Contains(i interface{}) bool {
	_, in := set[i]
	return in
}

func (set Set) Add(is ...interface{}) {
	for _, i := range is {
		set[i] = struct{}{}
	}
}

func (set Set) Remove(is ...interface{}) {
	for _, i := range is {
		delete(set, i)
	}
}

func (set Set) Slice() []interface{} {
	s := []interface{}{}
	for k, _ := range set {
		s = append(s, k)
	}
	return s
}

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

func (set Set) Len() int {
	return len(set)
}

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

func Union(a, b Set) Set {
	i := a
	i.Add(b.Slice()...)
	return i
}