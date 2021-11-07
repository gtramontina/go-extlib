package set

import (
	"fmt"
	"github.com/gtramontina/go-extlib/internal/hash"
	"reflect"
	"sort"
	"strings"
)

// Set is a finite collection that contains no duplicate members. As implied by
// its name, this type aims to model the mathematical concept of sets.
type Set[Type any] struct {
	members map[uint64]Type
}

// New creates a Set containing the given members.
func New[Type any](members ...Type) Set[Type] {
	newMembers := make(map[uint64]Type, len(members))
	for _, member := range members {
		newMembers[hash.Calc(member)] = member
	}

	return Set[Type]{newMembers}
}

// Add creates a Set containing all members of this Set plus the given new
// member.
func (s Set[Type]) Add(newMember Type) Set[Type] {
	newMembers := make(map[uint64]Type, len(s.members)+1)
	newMembers[hash.Calc(newMember)] = newMember
	for h, member := range s.members {
		newMembers[h] = member
	}

	return Set[Type]{newMembers}
}

// Remove creates a Set containing all members of this Set minus the given
// member.
func (s Set[Type]) Remove(existingMember Type) Set[Type] {
	newMembers := make(map[uint64]Type, len(s.members)-1)
	existingMemberHash := hash.Calc(existingMember)
	for h, member := range s.members {
		if h != existingMemberHash {
			newMembers[h] = member
		}
	}

	return Set[Type]{newMembers}
}

// Cardinality returns the number of members of this finite Set.
//  |A| or #A
func (s Set[Type]) Cardinality() int {
	return len(s.members)
}

// Equals asserts whether this Set contains the exact same members as the other
// Set.
func (s Set[Type]) Equals(other Set[Type]) bool {
	for h := range other.members {
		if _, contains := s.members[h]; !contains {
			return false
		}
	}

	return s.Cardinality() == other.Cardinality()
}

// Contains checks whether the given element is a member os this Set.
//	A┌─────────────┐
//	 │             │
//	 │         a   │  a ∈ A
//	 │             │
//	 └─────────────┘
func (s Set[Type]) Contains(member Type) bool {
	_, contains := s.members[hash.Calc(member)]

	return contains
}

// SuperSetOf checks whether this Set is a super set of the given Set.
// A ⊇ B
func (s Set[Type]) SuperSetOf(other Set[Type]) bool {
	for h := range other.members {
		if _, contains := s.members[h]; !contains {
			return false
		}
	}

	return true
}

// Union creates a Set of all values that are a member of A, or B, or both.
//	A┌─────────────┐
//	 │#############│
//	 │####┌────────┼────┐
//	 │####│########│####│  A ∪ B
//	 └────┼────────┘####│
//	      │#############│
//	      └─────────────┘B
func (s Set[Type]) Union(other Set[Type]) Set[Type] {
	newMembers := make(map[uint64]Type, len(s.members)+len(other.members))
	for h, member := range other.members {
		newMembers[h] = member
	}
	for h, member := range s.members {
		newMembers[h] = member
	}

	return Set[Type]{newMembers}
}

// Intersection creates a Set of all values that are members of both A and B.
//	A┌─────────────┐
//	 │             │
//	 │    ┌────────┼────┐
//	 │    │########│    │  A ∩ B
//	 └────┼────────┘    │
//	      │             │
//	      └─────────────┘B
func (s Set[Type]) Intersection(other Set[Type]) Set[Type] {
	newMembers := map[uint64]Type{}
	for h, otherMember := range other.members {
		if _, contains := s.members[h]; contains {
			newMembers[h] = otherMember
		}
	}

	return Set[Type]{newMembers}
}

// Difference creates a Set of all values of A that are not members of B.
//	A┌─────────────┐
//	 │#############│
//	 │####┌────────┼────┐
//	 │####│        │    │  A \ B or
//	 └────┼────────┘    │  A - B
//	      │             │
//	      └─────────────┘B
func (s Set[Type]) Difference(other Set[Type]) Set[Type] {
	newMembers := map[uint64]Type{}
	for h, member := range s.members {
		if _, contains := other.members[h]; !contains {
			newMembers[h] = member
		}
	}

	return Set[Type]{newMembers}
}

// SymmetricDifference creates a Set of all values which are of one of the sets,
// but not both.
//	A┌─────────────┐
//	 │#############│
//	 │####┌────────┼────┐
//	 │####│        │####│  A ∆ B
//	 └────┼────────┘####│
//	      │#############│
//	      └─────────────┘B
func (s Set[Type]) SymmetricDifference(other Set[Type]) Set[Type] {
	return s.Union(other).Difference(s.Intersection(other))
}

// Filter calls the provided predicate function once for each member of the Set,
// and constructs a new Set of all the members for which the predicate returns
// true.
func (s Set[Type]) Filter(predicate func(Type) bool) Set[Type] {
	newMembers := make(map[uint64]Type, len(s.members))
	for h, member := range s.members {
		if predicate(member) {
			newMembers[h] = member
		}
	}

	return Set[Type]{newMembers}
}

// String renders itself as a string containing all members.
func (s Set[Type]) String() string {
	members := make([]string, 0, len(s.members))
	for _, member := range s.members {
		members = append(members, fmt.Sprintf("%+v", member))
	}

	sort.Slice(members, func(a, z int) bool {
		return members[a] < members[z]
	})

	kind := reflect.TypeOf(s.members).Elem().String()
	return "Set(" + kind + "){" + strings.Join(members, ", ") + "}"
}
