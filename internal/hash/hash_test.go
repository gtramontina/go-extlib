package hash_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/internal/assert"
	"github.com/gtramontina/go-extlib/internal/hash"
)

func TestHash(t *testing.T) {
	t.Run("unknown type", func(t *testing.T) {
		assert.Panic(t, func() { hash.Calc(uintptr(0)) }, `can't calculate hash for "uintptr": 0`)
	})

	t.Run("nil", func(t *testing.T) {
		assert.Eq(t, hash.Calc(nil), hash.Calc(nil))
	})

	t.Run("bool", func(t *testing.T) {
		assert.Eq(t, hash.Calc(true), hash.Calc(true))
		assert.Eq(t, hash.Calc(false), hash.Calc(false))
		assert.NotEq(t, hash.Calc(true), hash.Calc(false))
	})

	t.Run("int", func(t *testing.T) {
		assert.Eq(t, hash.Calc(-1), hash.Calc(-1))
		assert.Eq(t, hash.Calc(0), hash.Calc(0))
		assert.Eq(t, hash.Calc(1), hash.Calc(1))
		assert.NotEq(t, hash.Calc(-1), hash.Calc(0))
		assert.NotEq(t, hash.Calc(0), hash.Calc(1))
		assert.NotEq(t, hash.Calc(-1), hash.Calc(1))
	})

	t.Run("int8", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int8(-1)), hash.Calc(int8(-1)))
		assert.Eq(t, hash.Calc(int8(0)), hash.Calc(int8(0)))
		assert.Eq(t, hash.Calc(int8(1)), hash.Calc(int8(1)))
		assert.NotEq(t, hash.Calc(int8(-1)), hash.Calc(int8(0)))
		assert.NotEq(t, hash.Calc(int8(0)), hash.Calc(int8(1)))
		assert.NotEq(t, hash.Calc(int8(-1)), hash.Calc(int8(1)))
	})

	t.Run("int16", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int16(-1)), hash.Calc(int16(-1)))
		assert.Eq(t, hash.Calc(int16(0)), hash.Calc(int16(0)))
		assert.Eq(t, hash.Calc(int16(1)), hash.Calc(int16(1)))
		assert.NotEq(t, hash.Calc(int16(-1)), hash.Calc(int16(0)))
		assert.NotEq(t, hash.Calc(int16(0)), hash.Calc(int16(1)))
		assert.NotEq(t, hash.Calc(int16(-1)), hash.Calc(int16(1)))
	})

	t.Run("int32", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int32(-1)), hash.Calc(int32(-1)))
		assert.Eq(t, hash.Calc(int32(0)), hash.Calc(int32(0)))
		assert.Eq(t, hash.Calc(int32(1)), hash.Calc(int32(1)))
		assert.NotEq(t, hash.Calc(int32(-1)), hash.Calc(int32(0)))
		assert.NotEq(t, hash.Calc(int32(0)), hash.Calc(int32(1)))
		assert.NotEq(t, hash.Calc(int32(-1)), hash.Calc(int32(1)))
	})

	t.Run("rune", func(t *testing.T) {
		assert.Eq(t, hash.Calc('a'), hash.Calc('a'))
		assert.Eq(t, hash.Calc('b'), hash.Calc('b'))
		assert.Eq(t, hash.Calc('c'), hash.Calc('c'))
		assert.NotEq(t, hash.Calc('a'), hash.Calc('b'))
		assert.NotEq(t, hash.Calc('b'), hash.Calc('c'))
		assert.NotEq(t, hash.Calc('a'), hash.Calc('c'))
	})

	t.Run("int64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int64(-1)), hash.Calc(int64(-1)))
		assert.Eq(t, hash.Calc(int64(0)), hash.Calc(int64(0)))
		assert.Eq(t, hash.Calc(int64(1)), hash.Calc(int64(1)))
		assert.NotEq(t, hash.Calc(int64(-1)), hash.Calc(int64(0)))
		assert.NotEq(t, hash.Calc(int64(0)), hash.Calc(int64(1)))
		assert.NotEq(t, hash.Calc(int64(-1)), hash.Calc(int64(1)))
	})

	t.Run("uint", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint(0)), hash.Calc(uint(0)))
		assert.Eq(t, hash.Calc(uint(1)), hash.Calc(uint(1)))
		assert.Eq(t, hash.Calc(uint(2)), hash.Calc(uint(2)))
		assert.NotEq(t, hash.Calc(uint(0)), hash.Calc(uint(1)))
		assert.NotEq(t, hash.Calc(uint(1)), hash.Calc(uint(2)))
		assert.NotEq(t, hash.Calc(uint(0)), hash.Calc(uint(1)))
	})

	t.Run("uint8", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint8(0)), hash.Calc(uint8(0)))
		assert.Eq(t, hash.Calc(uint8(1)), hash.Calc(uint8(1)))
		assert.Eq(t, hash.Calc(uint8(2)), hash.Calc(uint8(2)))
		assert.NotEq(t, hash.Calc(uint8(0)), hash.Calc(uint8(1)))
		assert.NotEq(t, hash.Calc(uint8(1)), hash.Calc(uint8(2)))
		assert.NotEq(t, hash.Calc(uint8(0)), hash.Calc(uint8(1)))
	})

	t.Run("uint16", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint16(0)), hash.Calc(uint16(0)))
		assert.Eq(t, hash.Calc(uint16(1)), hash.Calc(uint16(1)))
		assert.Eq(t, hash.Calc(uint16(2)), hash.Calc(uint16(2)))
		assert.NotEq(t, hash.Calc(uint16(0)), hash.Calc(uint16(1)))
		assert.NotEq(t, hash.Calc(uint16(1)), hash.Calc(uint16(2)))
		assert.NotEq(t, hash.Calc(uint16(0)), hash.Calc(uint16(1)))
	})

	t.Run("uint32", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint32(0)), hash.Calc(uint32(0)))
		assert.Eq(t, hash.Calc(uint32(1)), hash.Calc(uint32(1)))
		assert.Eq(t, hash.Calc(uint32(2)), hash.Calc(uint32(2)))
		assert.NotEq(t, hash.Calc(uint32(0)), hash.Calc(uint32(1)))
		assert.NotEq(t, hash.Calc(uint32(1)), hash.Calc(uint32(2)))
		assert.NotEq(t, hash.Calc(uint32(0)), hash.Calc(uint32(1)))
	})

	t.Run("uint64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint64(0)), hash.Calc(uint64(0)))
		assert.Eq(t, hash.Calc(uint64(1)), hash.Calc(uint64(1)))
		assert.Eq(t, hash.Calc(uint64(2)), hash.Calc(uint64(2)))
		assert.NotEq(t, hash.Calc(uint64(0)), hash.Calc(uint64(1)))
		assert.NotEq(t, hash.Calc(uint64(1)), hash.Calc(uint64(2)))
		assert.NotEq(t, hash.Calc(uint64(0)), hash.Calc(uint64(1)))
	})

	t.Run("float32", func(t *testing.T) {
		assert.Eq(t, hash.Calc(float32(-1)), hash.Calc(float32(-1)))
		assert.Eq(t, hash.Calc(float32(0)), hash.Calc(float32(0)))
		assert.Eq(t, hash.Calc(float32(1)), hash.Calc(float32(1)))
		assert.NotEq(t, hash.Calc(float32(-1)), hash.Calc(float32(0)))
		assert.NotEq(t, hash.Calc(float32(0)), hash.Calc(float32(1)))
		assert.NotEq(t, hash.Calc(float32(-1)), hash.Calc(float32(1)))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(float64(-1)), hash.Calc(float64(-1)))
		assert.Eq(t, hash.Calc(float64(0)), hash.Calc(float64(0)))
		assert.Eq(t, hash.Calc(float64(1)), hash.Calc(float64(1)))
		assert.NotEq(t, hash.Calc(float64(-1)), hash.Calc(float64(0)))
		assert.NotEq(t, hash.Calc(float64(0)), hash.Calc(float64(1)))
		assert.NotEq(t, hash.Calc(float64(-1)), hash.Calc(float64(1)))
	})

	t.Run("complex64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(complex64(-1)), hash.Calc(complex64(-1)))
		assert.Eq(t, hash.Calc(complex64(0)), hash.Calc(complex64(0)))
		assert.Eq(t, hash.Calc(complex64(1)), hash.Calc(complex64(1)))
		assert.NotEq(t, hash.Calc(complex64(-1)), hash.Calc(complex64(0)))
		assert.NotEq(t, hash.Calc(complex64(0)), hash.Calc(complex64(1)))
		assert.NotEq(t, hash.Calc(complex64(-1)), hash.Calc(complex64(1)))
	})

	t.Run("complex128", func(t *testing.T) {
		assert.Eq(t, hash.Calc(complex128(-1)), hash.Calc(complex128(-1)))
		assert.Eq(t, hash.Calc(complex128(0)), hash.Calc(complex128(0)))
		assert.Eq(t, hash.Calc(complex128(1)), hash.Calc(complex128(1)))
		assert.NotEq(t, hash.Calc(complex128(-1)), hash.Calc(complex128(0)))
		assert.NotEq(t, hash.Calc(complex128(0)), hash.Calc(complex128(1)))
		assert.NotEq(t, hash.Calc(complex128(-1)), hash.Calc(complex128(1)))
	})

	t.Run("string", func(t *testing.T) {
		assert.Eq(t, hash.Calc("A"), hash.Calc("A"))
		assert.Eq(t, hash.Calc("B"), hash.Calc("B"))
		assert.Eq(t, hash.Calc("C"), hash.Calc("C"))
		assert.NotEq(t, hash.Calc("A"), hash.Calc("B"))
		assert.NotEq(t, hash.Calc("B"), hash.Calc("C"))
		assert.NotEq(t, hash.Calc("A"), hash.Calc("C"))
	})

	t.Run("array", func(t *testing.T) {
		assert.Eq(t, hash.Calc([9]int{}), hash.Calc([9]int{}))
		assert.Eq(t, hash.Calc([9]int{0}), hash.Calc([9]int{0}))
		assert.Eq(t, hash.Calc([9]int{1}), hash.Calc([9]int{1}))
		assert.Eq(t, hash.Calc([9]int{0, 1}), hash.Calc([9]int{0, 1}))
		assert.NotEq(t, hash.Calc([9]int{0}), hash.Calc([9]int{1}))
		assert.NotEq(t, hash.Calc([9]int{1, 0}), hash.Calc([9]int{0, 1}))
		assert.NotEq(t, hash.Calc([9]int{}), hash.Calc([9]string{}))
		assert.NotEq(t, hash.Calc([9]int{}), hash.Calc([8]int{}))
	})

	t.Run("func", func(t *testing.T) {
		arity0coarity0A := func() {}
		arity0coarity0B := func() {}
		assert.Eq(t, hash.Calc(arity0coarity0A), hash.Calc(arity0coarity0A))
		assert.Eq(t, hash.Calc(arity0coarity0B), hash.Calc(arity0coarity0B))
		assert.NotEq(t, hash.Calc(arity0coarity0A), hash.Calc(arity0coarity0B))

		arity1coarity0A := func(int) {}
		arity1coarity0B := func(int) {}
		assert.Eq(t, hash.Calc(arity1coarity0A), hash.Calc(arity1coarity0A))
		assert.Eq(t, hash.Calc(arity1coarity0B), hash.Calc(arity1coarity0B))
		assert.NotEq(t, hash.Calc(arity1coarity0A), hash.Calc(arity1coarity0B))

		arity0coarity1A := func() int { return 0 }
		arity0coarity1B := func() int { return 0 }
		assert.Eq(t, hash.Calc(arity0coarity1A), hash.Calc(arity0coarity1A))
		assert.Eq(t, hash.Calc(arity0coarity1B), hash.Calc(arity0coarity1B))
		assert.NotEq(t, hash.Calc(arity0coarity1A), hash.Calc(arity0coarity1B))

		arity1coarity1A := func(int) string { return "!" }
		arity1coarity1B := func(int) string { return "!" }
		assert.Eq(t, hash.Calc(arity1coarity1A), hash.Calc(arity1coarity1A))
		assert.Eq(t, hash.Calc(arity1coarity1B), hash.Calc(arity1coarity1B))
		assert.NotEq(t, hash.Calc(arity1coarity1A), hash.Calc(arity1coarity1B))

		arityNcoarityMA := func(int, string, bool) (uint64, error) { return 1, nil }
		arityNcoarityMB := func(int, string, bool) (uint64, error) { return 1, nil }
		assert.Eq(t, hash.Calc(arityNcoarityMA), hash.Calc(arityNcoarityMA))
		assert.Eq(t, hash.Calc(arityNcoarityMB), hash.Calc(arityNcoarityMB))
		assert.NotEq(t, hash.Calc(arityNcoarityMA), hash.Calc(arityNcoarityMB))
	})

	t.Run("slice", func(t *testing.T) {
		assert.Eq(t, hash.Calc([]int{}), hash.Calc([]int{}))
		assert.Eq(t, hash.Calc([]int{0}), hash.Calc([]int{0}))
		assert.Eq(t, hash.Calc([]int{1}), hash.Calc([]int{1}))
		assert.Eq(t, hash.Calc([]int{0, 1}), hash.Calc([]int{0, 1}))
		assert.NotEq(t, hash.Calc([]int{0}), hash.Calc([]int{1}))
		assert.NotEq(t, hash.Calc([]int{1, 0}), hash.Calc([]int{0, 1}))
		assert.NotEq(t, hash.Calc([]int{}), hash.Calc([]string{}))
	})

	t.Run("struct", func(t *testing.T) {
		assert.Eq(t, hash.Calc(struct{}{}), hash.Calc(struct{}{}))
		assert.Eq(t, hash.Calc(struct{ int }{}), hash.Calc(struct{ int }{}))
		assert.Eq(t, hash.Calc(struct{ int }{0}), hash.Calc(struct{ int }{}))
		assert.NotEq(t, hash.Calc(struct{ int }{0}), hash.Calc(struct{ int }{1}))
		assert.NotEq(t, hash.Calc(struct{ field int }{}), hash.Calc(struct{ int }{}))

		type emptyA struct{}
		type emptyB struct{}
		assert.Eq(t, hash.Calc(emptyA{}), hash.Calc(emptyA{}))
		assert.Eq(t, hash.Calc(emptyB{}), hash.Calc(emptyB{}))
		assert.NotEq(t, hash.Calc(emptyA{}), hash.Calc(emptyB{}))

		type unexported1FieldA struct{ field1 int }
		type unexported1FieldB struct{ field1 int }
		assert.Eq(t, hash.Calc(unexported1FieldA{}), hash.Calc(unexported1FieldA{}))
		assert.Eq(t, hash.Calc(unexported1FieldA{0}), hash.Calc(unexported1FieldA{}))
		assert.NotEq(t, hash.Calc(unexported1FieldA{0}), hash.Calc(unexported1FieldA{1}))
		assert.Eq(t, hash.Calc(unexported1FieldA{1}), hash.Calc(unexported1FieldA{1}))
		assert.NotEq(t, hash.Calc(unexported1FieldA{}), hash.Calc(unexported1FieldB{}))
		assert.NotEq(t, hash.Calc(unexported1FieldA{1}), hash.Calc(unexported1FieldB{1}))

		type unexported2FieldsA struct {
			field1 int
			field2 string
		}
		type unexported2FieldsB struct {
			field1 int
			field2 string
		}
		assert.Eq(t, hash.Calc(unexported2FieldsA{}), hash.Calc(unexported2FieldsA{}))
		assert.Eq(t, hash.Calc(unexported2FieldsA{0, ""}), hash.Calc(unexported2FieldsA{}))
		assert.NotEq(t, hash.Calc(unexported2FieldsA{0, ""}), hash.Calc(unexported2FieldsA{1, ""}))
		assert.NotEq(t, hash.Calc(unexported2FieldsA{1, ""}), hash.Calc(unexported2FieldsA{1, "!"}))
		assert.Eq(t, hash.Calc(unexported2FieldsA{1, "!"}), hash.Calc(unexported2FieldsA{1, "!"}))
		assert.NotEq(t, hash.Calc(unexported2FieldsA{}), hash.Calc(unexported2FieldsB{}))
		assert.NotEq(t, hash.Calc(unexported2FieldsA{1, "!"}), hash.Calc(unexported2FieldsB{1, "!"}))

		type anonymousUnexported1FieldA struct{ int }
		type anonymousUnexported1FieldB struct{ int }
		assert.Eq(t, hash.Calc(anonymousUnexported1FieldA{}), hash.Calc(anonymousUnexported1FieldA{}))
		assert.Eq(t, hash.Calc(anonymousUnexported1FieldA{0}), hash.Calc(anonymousUnexported1FieldA{}))
		assert.NotEq(t, hash.Calc(anonymousUnexported1FieldA{0}), hash.Calc(anonymousUnexported1FieldA{1}))
		assert.Eq(t, hash.Calc(anonymousUnexported1FieldA{1}), hash.Calc(anonymousUnexported1FieldA{1}))
		assert.NotEq(t, hash.Calc(anonymousUnexported1FieldA{}), hash.Calc(anonymousUnexported1FieldB{}))
		assert.NotEq(t, hash.Calc(anonymousUnexported1FieldA{1}), hash.Calc(anonymousUnexported1FieldB{1}))

		type anonymousUnexported2FieldsA struct {
			int
			string
		}
		type anonymousUnexported2FieldsB struct {
			int
			string
		}
		assert.Eq(t, hash.Calc(anonymousUnexported2FieldsA{}), hash.Calc(anonymousUnexported2FieldsA{}))
		assert.Eq(t, hash.Calc(anonymousUnexported2FieldsA{0, ""}), hash.Calc(anonymousUnexported2FieldsA{}))
		assert.NotEq(t, hash.Calc(anonymousUnexported2FieldsA{0, ""}), hash.Calc(anonymousUnexported2FieldsA{1, ""}))
		assert.NotEq(t, hash.Calc(anonymousUnexported2FieldsA{1, ""}), hash.Calc(anonymousUnexported2FieldsA{1, "!"}))
		assert.Eq(t, hash.Calc(anonymousUnexported2FieldsA{1, "!"}), hash.Calc(anonymousUnexported2FieldsA{1, "!"}))
		assert.NotEq(t, hash.Calc(anonymousUnexported2FieldsA{}), hash.Calc(anonymousUnexported2FieldsB{}))
		assert.NotEq(t, hash.Calc(anonymousUnexported2FieldsA{1, "!"}), hash.Calc(anonymousUnexported2FieldsB{1, "!"}))
	})

	t.Run("channel", func(t *testing.T) {
		chanA := make(chan int)
		chanB := make(chan int)
		chanC := make(chan string)
		assert.Eq(t, hash.Calc(chanA), hash.Calc(chanA))
		assert.Eq(t, hash.Calc(chanB), hash.Calc(chanB))
		assert.NotEq(t, hash.Calc(chanA), hash.Calc(chanB))
		assert.NotEq(t, hash.Calc(chanA), hash.Calc(chanC))
		assert.NotEq(t, hash.Calc(chanB), hash.Calc(chanC))
	})

	t.Run("map", func(t *testing.T) {
		assert.Eq(t, hash.Calc(map[int]int{}), hash.Calc(map[int]int{}))
		assert.NotEq(t, hash.Calc(map[string]int{}), hash.Calc(map[int]int{}))
		assert.NotEq(t, hash.Calc(map[int]string{}), hash.Calc(map[int]int{}))
		assert.NotEq(t, hash.Calc(map[int]int{0: 0}), hash.Calc(map[int]int{1: 0}))
		assert.NotEq(t, hash.Calc(map[int]int{0: 0}), hash.Calc(map[int]int{0: 1}))
		assert.NotEq(t, hash.Calc(map[int]int{0: 1, 1: 0}), hash.Calc(map[int]int{0: 0, 1: 1}))
		assert.Eq(t, hash.Calc(map[int]int{0: 0, 1: 1}), hash.Calc(map[int]int{1: 1, 0: 0}))
		assert.NotEq(t, hash.Calc(map[int]int{0: 0, 1: 1, 2: 2}), hash.Calc(map[int]int{1: 1, 0: 0}))
	})

	t.Run("pointer", func(t *testing.T) {
		blankStructA := struct{}{}
		blankStructB := struct{}{}
		assert.Eq(t, hash.Calc(&blankStructA), hash.Calc(&blankStructA))
		assert.Eq(t, hash.Calc(&blankStructB), hash.Calc(&blankStructB))
		assert.Eq(t, hash.Calc(&blankStructA), hash.Calc(&blankStructB))

		structA := struct{ int }{}
		structB := struct{ int }{}
		assert.Eq(t, hash.Calc(&structA), hash.Calc(&structA))
		assert.Eq(t, hash.Calc(&structB), hash.Calc(&structB))
		assert.NotEq(t, hash.Calc(&structA), hash.Calc(&structB))

		sliceA := []int{}
		sliceB := []int{}
		assert.Eq(t, hash.Calc(&sliceA), hash.Calc(&sliceA))
		assert.Eq(t, hash.Calc(&sliceB), hash.Calc(&sliceB))
		assert.NotEq(t, hash.Calc(&sliceA), hash.Calc(&sliceB))

		arrayA := [1]int{}
		arrayB := [1]int{}
		assert.Eq(t, hash.Calc(&arrayA), hash.Calc(&arrayA))
		assert.Eq(t, hash.Calc(&arrayB), hash.Calc(&arrayB))
		assert.NotEq(t, hash.Calc(&arrayA), hash.Calc(&arrayB))

		mapA := map[int]string{}
		mapB := map[int]string{}
		assert.Eq(t, hash.Calc(&mapA), hash.Calc(&mapA))
		assert.Eq(t, hash.Calc(&mapB), hash.Calc(&mapB))
		assert.NotEq(t, hash.Calc(&mapA), hash.Calc(&mapB))
	})
}
