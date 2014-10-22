package stringset

import "strings"

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(elems []string) Set {
	n := Set(make(map[string]bool, len(elems)))
	for _, e := range elems {
		n.Add(e)
	}
	return n
}

func (s *Set) Add(a string)     { (*s)[a] = true }
func (s *Set) Delete(a string)  { delete(*s, a) }
func (s Set) IsEmpty() bool     { return len(s) == 0 }
func (s Set) Len() int          { return len(s) }
func (s Set) Has(a string) bool { return s[a] }

func (s Set) Slice() []string {
	elems := make([]string, 0, len(s))
	for k := range s {
		elems = append(elems, k)
	}
	return elems
}

func (s Set) String() string {
	elems := make([]string, 0, len(s))
	for k := range s {
		elems = append(elems, `"`+k+`"`)
	}
	return "{" + strings.Join(elems, ", ") + "}"
}

func Disjoint(s1, s2 Set) bool {
	for k1 := range s1 {
		if s2[k1] {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k1 := range s1 {
		if !s2[k1] {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	s3 := New()
	for k1 := range s1 {
		if s2[k1] {
			s3.Add(k1)
		}
	}
	return s3
}

func Union(s1, s2 Set) Set {
	s3 := make(Set, len(s1)+len(s2))
	for k := range s1 {
		s3.Add(k)
	}
	for k := range s2 {
		s3.Add(k)
	}
	return s3
}

func Difference(s1, s2 Set) Set {
	s3 := New()
	for k := range s1 {
		if !s2[k] {
			s3.Add(k)
		}
	}
	return s3
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2[k] {
			return false
		}
	}
	return true
}

func SymetricDifference(s1, s2 Set) Set {
	s3 := New()
	for k := range s1 {
		if !s2[k] {
			s3.Add(k)
		}
	}
	for k := range s2 {
		if !s1[k] {
			s3.Add(k)
		}
	}
	return s3
}
