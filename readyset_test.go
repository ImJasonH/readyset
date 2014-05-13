package readyset

import "testing"

func TestContains(t *testing.T) {
	cs := []struct {
		s Set
		e []interface{}
	}{{
		NewSet(1, 2, 3),
		[]interface{}{1, 2, 3},
	}, {
		NewSet(true, "foo", struct{}{}),
		[]interface{}{true, "foo", struct{}{}},
	}}
	for _, c := range cs {
		for _, x := range c.e {
			if !c.s.Contains(x) {
				t.Errorf("set %v should contain %v", c.s, x)
			}
			if !sliceContains(c.s.Slice(), x) {
				t.Errorf("slice %v should contain %v", c.s.Slice(), x)
			}
		}
	}
}

func sliceContains(s []interface{}, e interface{}) bool {
	for _, se := range s {
		if se == e {
			return true
		}
	}
	return false
}

func TestIntersection(t *testing.T) {
	cs := []struct {
		a, b Set
		e    []interface{}
	}{{
		NewSet(1, 2, 3),
		NewSet(3, 4, 5),
		[]interface{}{3},
	}, {
		NewSet(true, false, "foo"),
		NewSet(false, "foo", "bar"),
		[]interface{}{false, "foo"},
	}}
	for _, c := range cs {
		i := Intersection(c.a, c.b)
		for _, e := range c.e {
			if !i.Contains(e) {
				t.Errorf("expected %v in %v", e, i)
			}
		}
	}
}

func TestUnion(t *testing.T) {
	cs := []struct {
		a, b Set
		e    []interface{}
	}{{
		NewSet(1, 2),
		NewSet(2, 3),
		[]interface{}{1, 2, 3},
	}, {
		NewSet(false, true, "foo"),
		NewSet(true, "boo", struct{}{}),
		[]interface{}{false, true, "foo", "boo", struct{}{}},
	}}
	for _, c := range cs {
		u := Union(c.a, c.b)
		for _, e := range c.e {
			if !u.Contains(e) {
				t.Errorf("expected %v in %v", e, u)
			}
		}
		if u.Len() != len(c.e) {
			t.Errorf("unexpected len, got %d, want %d", u.Len(), len(c.e))
		}
	}
}

func TestAddRemove(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("expected 1 to be removed from %v", s)
	}
	if s.Len() != 2 {
		t.Errorf("expected %v.Len() == 2", s)
	}

	// Add new element
	s.Add(4)
	if !s.Contains(4) {
		t.Errorf("expected 4 to be added to %v", s)
	}
	if s.Len() != 3 {
		t.Errorf("expected %v.Len() == 3", s)
	}

	// Readd existing element
	s.Add(2)
	if !s.Contains(2) {
		t.Errorf("expected 2 to be in %v", s)
	}
	if s.Len() != 3 {
		t.Errorf("expected %v.Len() == 3", s)
	}
}
