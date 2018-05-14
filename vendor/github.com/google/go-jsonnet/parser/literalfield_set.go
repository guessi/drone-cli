// Generated by: main
// TypeWriter: set
// Directive: +gen on LiteralField

package parser

// Set is a modification of https://github.com/deckarep/golang-set
// The MIT License (MIT)
// Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

// LiteralFieldSet is the primary type that represents a set
type LiteralFieldSet map[LiteralField]struct{}

// NewLiteralFieldSet creates and returns a reference to an empty set.
func NewLiteralFieldSet(a ...LiteralField) LiteralFieldSet {
	s := make(LiteralFieldSet)
	for _, i := range a {
		s.Add(i)
	}
	return s
}

// ToSlice returns the elements of the current set as a slice
func (set LiteralFieldSet) ToSlice() []LiteralField {
	var s []LiteralField
	for v := range set {
		s = append(s, v)
	}
	return s
}

// Add adds an item to the current set if it doesn't already exist in the set.
func (set LiteralFieldSet) Add(i LiteralField) bool {
	_, found := set[i]
	set[i] = struct{}{}
	return !found //False if it existed already
}

// Contains determines if a given item is already in the set.
func (set LiteralFieldSet) Contains(i LiteralField) bool {
	_, found := set[i]
	return found
}

// ContainsAll determines if the given items are all in the set
func (set LiteralFieldSet) ContainsAll(i ...LiteralField) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

// IsSubset determines if every item in the other set is in this set.
func (set LiteralFieldSet) IsSubset(other LiteralFieldSet) bool {
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSuperset determines if every item of this set is in the other set.
func (set LiteralFieldSet) IsSuperset(other LiteralFieldSet) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set LiteralFieldSet) Union(other LiteralFieldSet) LiteralFieldSet {
	unionedSet := NewLiteralFieldSet()

	for elem := range set {
		unionedSet.Add(elem)
	}
	for elem := range other {
		unionedSet.Add(elem)
	}
	return unionedSet
}

// Intersect returns a new set with items that exist only in both sets.
func (set LiteralFieldSet) Intersect(other LiteralFieldSet) LiteralFieldSet {
	intersection := NewLiteralFieldSet()
	// loop over smaller set
	if set.Cardinality() < other.Cardinality() {
		for elem := range set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set LiteralFieldSet) Difference(other LiteralFieldSet) LiteralFieldSet {
	differencedSet := NewLiteralFieldSet()
	for elem := range set {
		if !other.Contains(elem) {
			differencedSet.Add(elem)
		}
	}
	return differencedSet
}

// SymmetricDifference returns a new set with items in the current set or the other set but not in both.
func (set LiteralFieldSet) SymmetricDifference(other LiteralFieldSet) LiteralFieldSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clear clears the entire set to be the empty set.
func (set *LiteralFieldSet) Clear() {
	*set = make(LiteralFieldSet)
}

// Remove allows the removal of a single item in the set.
func (set LiteralFieldSet) Remove(i LiteralField) {
	delete(set, i)
}

// Cardinality returns how many items are currently in the set.
func (set LiteralFieldSet) Cardinality() int {
	return len(set)
}

// Iter returns a channel of type LiteralField that you can range over.
func (set LiteralFieldSet) Iter() <-chan LiteralField {
	ch := make(chan LiteralField)
	go func() {
		for elem := range set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

// Equal determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set LiteralFieldSet) Equal(other LiteralFieldSet) bool {
	if set.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Clone returns a clone of the set.
// Does NOT clone the underlying elements.
func (set LiteralFieldSet) Clone() LiteralFieldSet {
	clonedSet := NewLiteralFieldSet()
	for elem := range set {
		clonedSet.Add(elem)
	}
	return clonedSet
}
