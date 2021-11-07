package hash_test

import (
	"reflect"
	"testing"
	"testing/quick"

	"github.com/gtramontina/go-extlib/internal/assert"
	"github.com/gtramontina/go-extlib/internal/hash"
)

func TestHash(t *testing.T) {
	config := &quick.Config{
		MaxCountScale: maxCountScale(),
	}

	sameHash := func(a, b interface{}) bool {
		return hash.Calc(a) == hash.Calc(b)
	}

	differentHash := func(a, b interface{}) bool {
		return !sameHash(a, b)
	}

	sameTypeHashCheck := func(a, b interface{}) bool {
		return reflect.DeepEqual(a, b) || (sameHash(a, a) && sameHash(b, b) && differentHash(a, b))
	}

	check := func(fn interface{}) func(*testing.T) {
		return func(t *testing.T) {
			t.Helper()
			assert.NoError(t, quick.Check(fn, config))
		}
	}

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
		assert.NoError(t, quick.Check(func(a, b bool) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("int", func(t *testing.T) {
		assert.Eq(t, hash.Calc(-1), hash.Calc(-1))
		assert.Eq(t, hash.Calc(0), hash.Calc(0))
		assert.Eq(t, hash.Calc(1), hash.Calc(1))
		assert.NotEq(t, hash.Calc(-1), hash.Calc(0))
		assert.NotEq(t, hash.Calc(0), hash.Calc(1))
		assert.NotEq(t, hash.Calc(-1), hash.Calc(1))
		assert.NoError(t, quick.Check(func(a, b int) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("int8", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int8(-1)), hash.Calc(int8(-1)))
		assert.Eq(t, hash.Calc(int8(0)), hash.Calc(int8(0)))
		assert.Eq(t, hash.Calc(int8(1)), hash.Calc(int8(1)))
		assert.NotEq(t, hash.Calc(int8(-1)), hash.Calc(int8(0)))
		assert.NotEq(t, hash.Calc(int8(0)), hash.Calc(int8(1)))
		assert.NotEq(t, hash.Calc(int8(-1)), hash.Calc(int8(1)))
		assert.NoError(t, quick.Check(func(a, b int8) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("int16", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int16(-1)), hash.Calc(int16(-1)))
		assert.Eq(t, hash.Calc(int16(0)), hash.Calc(int16(0)))
		assert.Eq(t, hash.Calc(int16(1)), hash.Calc(int16(1)))
		assert.NotEq(t, hash.Calc(int16(-1)), hash.Calc(int16(0)))
		assert.NotEq(t, hash.Calc(int16(0)), hash.Calc(int16(1)))
		assert.NotEq(t, hash.Calc(int16(-1)), hash.Calc(int16(1)))
		assert.NoError(t, quick.Check(func(a, b int16) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("int32", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int32(-1)), hash.Calc(int32(-1)))
		assert.Eq(t, hash.Calc(int32(0)), hash.Calc(int32(0)))
		assert.Eq(t, hash.Calc(int32(1)), hash.Calc(int32(1)))
		assert.NotEq(t, hash.Calc(int32(-1)), hash.Calc(int32(0)))
		assert.NotEq(t, hash.Calc(int32(0)), hash.Calc(int32(1)))
		assert.NotEq(t, hash.Calc(int32(-1)), hash.Calc(int32(1)))
		assert.NoError(t, quick.Check(func(a, b int32) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("rune", func(t *testing.T) {
		assert.Eq(t, hash.Calc('a'), hash.Calc('a'))
		assert.Eq(t, hash.Calc('b'), hash.Calc('b'))
		assert.Eq(t, hash.Calc('c'), hash.Calc('c'))
		assert.NotEq(t, hash.Calc('a'), hash.Calc('b'))
		assert.NotEq(t, hash.Calc('b'), hash.Calc('c'))
		assert.NotEq(t, hash.Calc('a'), hash.Calc('c'))
		assert.NoError(t, quick.Check(func(a, b rune) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("int64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(int64(-1)), hash.Calc(int64(-1)))
		assert.Eq(t, hash.Calc(int64(0)), hash.Calc(int64(0)))
		assert.Eq(t, hash.Calc(int64(1)), hash.Calc(int64(1)))
		assert.NotEq(t, hash.Calc(int64(-1)), hash.Calc(int64(0)))
		assert.NotEq(t, hash.Calc(int64(0)), hash.Calc(int64(1)))
		assert.NotEq(t, hash.Calc(int64(-1)), hash.Calc(int64(1)))
		assert.NoError(t, quick.Check(func(a, b int64) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("uint", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint(0)), hash.Calc(uint(0)))
		assert.Eq(t, hash.Calc(uint(1)), hash.Calc(uint(1)))
		assert.Eq(t, hash.Calc(uint(2)), hash.Calc(uint(2)))
		assert.NotEq(t, hash.Calc(uint(0)), hash.Calc(uint(1)))
		assert.NotEq(t, hash.Calc(uint(1)), hash.Calc(uint(2)))
		assert.NotEq(t, hash.Calc(uint(0)), hash.Calc(uint(1)))
		assert.NoError(t, quick.Check(func(a, b uint) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("uint8", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint8(0)), hash.Calc(uint8(0)))
		assert.Eq(t, hash.Calc(uint8(1)), hash.Calc(uint8(1)))
		assert.Eq(t, hash.Calc(uint8(2)), hash.Calc(uint8(2)))
		assert.NotEq(t, hash.Calc(uint8(0)), hash.Calc(uint8(1)))
		assert.NotEq(t, hash.Calc(uint8(1)), hash.Calc(uint8(2)))
		assert.NotEq(t, hash.Calc(uint8(0)), hash.Calc(uint8(1)))
		assert.NoError(t, quick.Check(func(a, b uint8) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("uint16", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint16(0)), hash.Calc(uint16(0)))
		assert.Eq(t, hash.Calc(uint16(1)), hash.Calc(uint16(1)))
		assert.Eq(t, hash.Calc(uint16(2)), hash.Calc(uint16(2)))
		assert.NotEq(t, hash.Calc(uint16(0)), hash.Calc(uint16(1)))
		assert.NotEq(t, hash.Calc(uint16(1)), hash.Calc(uint16(2)))
		assert.NotEq(t, hash.Calc(uint16(0)), hash.Calc(uint16(1)))
		assert.NoError(t, quick.Check(func(a, b uint16) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("uint32", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint32(0)), hash.Calc(uint32(0)))
		assert.Eq(t, hash.Calc(uint32(1)), hash.Calc(uint32(1)))
		assert.Eq(t, hash.Calc(uint32(2)), hash.Calc(uint32(2)))
		assert.NotEq(t, hash.Calc(uint32(0)), hash.Calc(uint32(1)))
		assert.NotEq(t, hash.Calc(uint32(1)), hash.Calc(uint32(2)))
		assert.NotEq(t, hash.Calc(uint32(0)), hash.Calc(uint32(1)))
		assert.NoError(t, quick.Check(func(a, b uint32) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("uint64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(uint64(0)), hash.Calc(uint64(0)))
		assert.Eq(t, hash.Calc(uint64(1)), hash.Calc(uint64(1)))
		assert.Eq(t, hash.Calc(uint64(2)), hash.Calc(uint64(2)))
		assert.NotEq(t, hash.Calc(uint64(0)), hash.Calc(uint64(1)))
		assert.NotEq(t, hash.Calc(uint64(1)), hash.Calc(uint64(2)))
		assert.NotEq(t, hash.Calc(uint64(0)), hash.Calc(uint64(1)))
		assert.NoError(t, quick.Check(func(a, b uint64) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("float32", func(t *testing.T) {
		assert.Eq(t, hash.Calc(float32(-1)), hash.Calc(float32(-1)))
		assert.Eq(t, hash.Calc(float32(0)), hash.Calc(float32(0)))
		assert.Eq(t, hash.Calc(float32(1)), hash.Calc(float32(1)))
		assert.NotEq(t, hash.Calc(float32(-1)), hash.Calc(float32(0)))
		assert.NotEq(t, hash.Calc(float32(0)), hash.Calc(float32(1)))
		assert.NotEq(t, hash.Calc(float32(-1)), hash.Calc(float32(1)))
		assert.NoError(t, quick.Check(func(a, b float32) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(float64(-1)), hash.Calc(float64(-1)))
		assert.Eq(t, hash.Calc(float64(0)), hash.Calc(float64(0)))
		assert.Eq(t, hash.Calc(float64(1)), hash.Calc(float64(1)))
		assert.NotEq(t, hash.Calc(float64(-1)), hash.Calc(float64(0)))
		assert.NotEq(t, hash.Calc(float64(0)), hash.Calc(float64(1)))
		assert.NotEq(t, hash.Calc(float64(-1)), hash.Calc(float64(1)))
		assert.NoError(t, quick.Check(func(a, b float64) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("complex64", func(t *testing.T) {
		assert.Eq(t, hash.Calc(complex64(-1)), hash.Calc(complex64(-1)))
		assert.Eq(t, hash.Calc(complex64(0)), hash.Calc(complex64(0)))
		assert.Eq(t, hash.Calc(complex64(1)), hash.Calc(complex64(1)))
		assert.NotEq(t, hash.Calc(complex64(-1)), hash.Calc(complex64(0)))
		assert.NotEq(t, hash.Calc(complex64(0)), hash.Calc(complex64(1)))
		assert.NotEq(t, hash.Calc(complex64(-1)), hash.Calc(complex64(1)))
		assert.NoError(t, quick.Check(func(a, b complex64) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("complex128", func(t *testing.T) {
		assert.Eq(t, hash.Calc(complex128(-1)), hash.Calc(complex128(-1)))
		assert.Eq(t, hash.Calc(complex128(0)), hash.Calc(complex128(0)))
		assert.Eq(t, hash.Calc(complex128(1)), hash.Calc(complex128(1)))
		assert.NotEq(t, hash.Calc(complex128(-1)), hash.Calc(complex128(0)))
		assert.NotEq(t, hash.Calc(complex128(0)), hash.Calc(complex128(1)))
		assert.NotEq(t, hash.Calc(complex128(-1)), hash.Calc(complex128(1)))
		assert.NoError(t, quick.Check(func(a, b complex128) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("string", func(t *testing.T) {
		assert.Eq(t, hash.Calc("A"), hash.Calc("A"))
		assert.Eq(t, hash.Calc("B"), hash.Calc("B"))
		assert.Eq(t, hash.Calc("C"), hash.Calc("C"))
		assert.NotEq(t, hash.Calc("A"), hash.Calc("B"))
		assert.NotEq(t, hash.Calc("B"), hash.Calc("C"))
		assert.NotEq(t, hash.Calc("A"), hash.Calc("C"))
		assert.NoError(t, quick.Check(func(a, b string) bool { return sameTypeHashCheck(a, b) }, config))
	})

	t.Run("different hashes for different types", func(t *testing.T) {
		t.Run("int!=int8", check(func(a int, b int8) bool { return differentHash(a, b) }))
		t.Run("int!=int16", check(func(a int, b int16) bool { return differentHash(a, b) }))
		t.Run("int!=int32", check(func(a int, b int32) bool { return differentHash(a, b) }))
		t.Run("int!=int64", check(func(a int, b int64) bool { return differentHash(a, b) }))
		t.Run("int!=uint", check(func(a int, b uint) bool { return differentHash(a, b) }))
		t.Run("int!=uint8", check(func(a int, b uint8) bool { return differentHash(a, b) }))
		t.Run("int!=uint16", check(func(a int, b uint16) bool { return differentHash(a, b) }))
		t.Run("int!=uint32", check(func(a int, b uint32) bool { return differentHash(a, b) }))
		t.Run("int!=uint64", check(func(a int, b uint64) bool { return differentHash(a, b) }))
		t.Run("int!=float32", check(func(a int, b float32) bool { return differentHash(a, b) }))
		t.Run("int!=float64", check(func(a int, b float64) bool { return differentHash(a, b) }))
		t.Run("int!=complex64", check(func(a int, b complex64) bool { return differentHash(a, b) }))
		t.Run("int!=complex128", check(func(a int, b complex128) bool { return differentHash(a, b) }))
		t.Run("int!=bool", check(func(a int, b bool) bool { return differentHash(a, b) }))
		t.Run("int!=string", check(func(a int, b string) bool { return differentHash(a, b) }))
		t.Run("int8!=int16", check(func(a int8, b int16) bool { return differentHash(a, b) }))
		t.Run("int8!=int32", check(func(a int8, b int32) bool { return differentHash(a, b) }))
		t.Run("int8!=int64", check(func(a int8, b int64) bool { return differentHash(a, b) }))
		t.Run("int8!=uint", check(func(a int8, b uint) bool { return differentHash(a, b) }))
		t.Run("int8!=uint8", check(func(a int8, b uint8) bool { return differentHash(a, b) }))
		t.Run("int8!=uint16", check(func(a int8, b uint16) bool { return differentHash(a, b) }))
		t.Run("int8!=uint32", check(func(a int8, b uint32) bool { return differentHash(a, b) }))
		t.Run("int8!=uint64", check(func(a int8, b uint64) bool { return differentHash(a, b) }))
		t.Run("int8!=float32", check(func(a int8, b float32) bool { return differentHash(a, b) }))
		t.Run("int8!=float64", check(func(a int8, b float64) bool { return differentHash(a, b) }))
		t.Run("int8!=complex64", check(func(a int8, b complex64) bool { return differentHash(a, b) }))
		t.Run("int8!=complex128", check(func(a int8, b complex128) bool { return differentHash(a, b) }))
		t.Run("int8!=bool", check(func(a int8, b bool) bool { return differentHash(a, b) }))
		t.Run("int8!=string", check(func(a int8, b string) bool { return differentHash(a, b) }))
		t.Run("int16!=int32", check(func(a int16, b int32) bool { return differentHash(a, b) }))
		t.Run("int16!=int64", check(func(a int16, b int64) bool { return differentHash(a, b) }))
		t.Run("int16!=uint", check(func(a int16, b uint) bool { return differentHash(a, b) }))
		t.Run("int16!=uint8", check(func(a int16, b uint8) bool { return differentHash(a, b) }))
		t.Run("int16!=uint16", check(func(a int16, b uint16) bool { return differentHash(a, b) }))
		t.Run("int16!=uint32", check(func(a int16, b uint32) bool { return differentHash(a, b) }))
		t.Run("int16!=uint64", check(func(a int16, b uint64) bool { return differentHash(a, b) }))
		t.Run("int16!=float32", check(func(a int16, b float32) bool { return differentHash(a, b) }))
		t.Run("int16!=float64", check(func(a int16, b float64) bool { return differentHash(a, b) }))
		t.Run("int16!=complex64", check(func(a int16, b complex64) bool { return differentHash(a, b) }))
		t.Run("int16!=complex128", check(func(a int16, b complex128) bool { return differentHash(a, b) }))
		t.Run("int16!=bool", check(func(a int16, b bool) bool { return differentHash(a, b) }))
		t.Run("int16!=string", check(func(a int16, b string) bool { return differentHash(a, b) }))
		t.Run("int32!=int64", check(func(a int32, b int64) bool { return differentHash(a, b) }))
		t.Run("int32!=uint", check(func(a int32, b uint) bool { return differentHash(a, b) }))
		t.Run("int32!=uint8", check(func(a int32, b uint8) bool { return differentHash(a, b) }))
		t.Run("int32!=uint16", check(func(a int32, b uint16) bool { return differentHash(a, b) }))
		t.Run("int32!=uint32", check(func(a int32, b uint32) bool { return differentHash(a, b) }))
		t.Run("int32!=uint64", check(func(a int32, b uint64) bool { return differentHash(a, b) }))
		t.Run("int32!=float32", check(func(a int32, b float32) bool { return differentHash(a, b) }))
		t.Run("int32!=float64", check(func(a int32, b float64) bool { return differentHash(a, b) }))
		t.Run("int32!=complex64", check(func(a int32, b complex64) bool { return differentHash(a, b) }))
		t.Run("int32!=complex128", check(func(a int32, b complex128) bool { return differentHash(a, b) }))
		t.Run("int32!=bool", check(func(a int32, b bool) bool { return differentHash(a, b) }))
		t.Run("int32!=string", check(func(a int32, b string) bool { return differentHash(a, b) }))
		t.Run("int64!=uint", check(func(a int64, b uint) bool { return differentHash(a, b) }))
		t.Run("int64!=uint8", check(func(a int64, b uint8) bool { return differentHash(a, b) }))
		t.Run("int64!=uint16", check(func(a int64, b uint16) bool { return differentHash(a, b) }))
		t.Run("int64!=uint32", check(func(a int64, b uint32) bool { return differentHash(a, b) }))
		t.Run("int64!=uint64", check(func(a int64, b uint64) bool { return differentHash(a, b) }))
		t.Run("int64!=float32", check(func(a int64, b float32) bool { return differentHash(a, b) }))
		t.Run("int64!=float64", check(func(a int64, b float64) bool { return differentHash(a, b) }))
		t.Run("int64!=complex64", check(func(a int64, b complex64) bool { return differentHash(a, b) }))
		t.Run("int64!=complex128", check(func(a int64, b complex128) bool { return differentHash(a, b) }))
		t.Run("int64!=bool", check(func(a int64, b bool) bool { return differentHash(a, b) }))
		t.Run("int64!=string", check(func(a int64, b string) bool { return differentHash(a, b) }))
		t.Run("uint!=uint8", check(func(a uint, b uint8) bool { return differentHash(a, b) }))
		t.Run("uint!=uint16", check(func(a uint, b uint16) bool { return differentHash(a, b) }))
		t.Run("uint!=uint32", check(func(a uint, b uint32) bool { return differentHash(a, b) }))
		t.Run("uint!=uint64", check(func(a uint, b uint64) bool { return differentHash(a, b) }))
		t.Run("uint!=float32", check(func(a uint, b float32) bool { return differentHash(a, b) }))
		t.Run("uint!=float64", check(func(a uint, b float64) bool { return differentHash(a, b) }))
		t.Run("uint!=complex64", check(func(a uint, b complex64) bool { return differentHash(a, b) }))
		t.Run("uint!=complex128", check(func(a uint, b complex128) bool { return differentHash(a, b) }))
		t.Run("uint!=bool", check(func(a uint, b bool) bool { return differentHash(a, b) }))
		t.Run("uint!=string", check(func(a uint, b string) bool { return differentHash(a, b) }))
		t.Run("uint8!=uint16", check(func(a uint8, b uint16) bool { return differentHash(a, b) }))
		t.Run("uint8!=uint32", check(func(a uint8, b uint32) bool { return differentHash(a, b) }))
		t.Run("uint8!=uint64", check(func(a uint8, b uint64) bool { return differentHash(a, b) }))
		t.Run("uint8!=float32", check(func(a uint8, b float32) bool { return differentHash(a, b) }))
		t.Run("uint8!=float64", check(func(a uint8, b float64) bool { return differentHash(a, b) }))
		t.Run("uint8!=complex64", check(func(a uint8, b complex64) bool { return differentHash(a, b) }))
		t.Run("uint8!=complex128", check(func(a uint8, b complex128) bool { return differentHash(a, b) }))
		t.Run("uint8!=bool", check(func(a uint8, b bool) bool { return differentHash(a, b) }))
		t.Run("uint8!=string", check(func(a uint8, b string) bool { return differentHash(a, b) }))
		t.Run("uint16!=uint32", check(func(a uint16, b uint32) bool { return differentHash(a, b) }))
		t.Run("uint16!=uint64", check(func(a uint16, b uint64) bool { return differentHash(a, b) }))
		t.Run("uint16!=float32", check(func(a uint16, b float32) bool { return differentHash(a, b) }))
		t.Run("uint16!=float64", check(func(a uint16, b float64) bool { return differentHash(a, b) }))
		t.Run("uint16!=complex64", check(func(a uint16, b complex64) bool { return differentHash(a, b) }))
		t.Run("uint16!=complex128", check(func(a uint16, b complex128) bool { return differentHash(a, b) }))
		t.Run("uint16!=bool", check(func(a uint16, b bool) bool { return differentHash(a, b) }))
		t.Run("uint16!=string", check(func(a uint16, b string) bool { return differentHash(a, b) }))
		t.Run("uint32!=uint64", check(func(a uint32, b uint64) bool { return differentHash(a, b) }))
		t.Run("uint32!=float32", check(func(a uint32, b float32) bool { return differentHash(a, b) }))
		t.Run("uint32!=float64", check(func(a uint32, b float64) bool { return differentHash(a, b) }))
		t.Run("uint32!=complex64", check(func(a uint32, b complex64) bool { return differentHash(a, b) }))
		t.Run("uint32!=complex128", check(func(a uint32, b complex128) bool { return differentHash(a, b) }))
		t.Run("uint32!=bool", check(func(a uint32, b bool) bool { return differentHash(a, b) }))
		t.Run("uint32!=string", check(func(a uint32, b string) bool { return differentHash(a, b) }))
		t.Run("uint64!=float32", check(func(a uint64, b float32) bool { return differentHash(a, b) }))
		t.Run("uint64!=float64", check(func(a uint64, b float64) bool { return differentHash(a, b) }))
		t.Run("uint64!=complex64", check(func(a uint64, b complex64) bool { return differentHash(a, b) }))
		t.Run("uint64!=complex128", check(func(a uint64, b complex128) bool { return differentHash(a, b) }))
		t.Run("uint64!=bool", check(func(a uint64, b bool) bool { return differentHash(a, b) }))
		t.Run("uint64!=string", check(func(a uint64, b string) bool { return differentHash(a, b) }))
		t.Run("float32!=float64", check(func(a float32, b float64) bool { return differentHash(a, b) }))
		t.Run("float32!=complex64", check(func(a float32, b complex64) bool { return differentHash(a, b) }))
		t.Run("float32!=complex128", check(func(a float32, b complex128) bool { return differentHash(a, b) }))
		t.Run("float32!=bool", check(func(a float32, b bool) bool { return differentHash(a, b) }))
		t.Run("float32!=string", check(func(a float32, b string) bool { return differentHash(a, b) }))
		t.Run("float64!=complex64", check(func(a float64, b complex64) bool { return differentHash(a, b) }))
		t.Run("float64!=complex128", check(func(a float64, b complex128) bool { return differentHash(a, b) }))
		t.Run("float64!=bool", check(func(a float64, b bool) bool { return differentHash(a, b) }))
		t.Run("float64!=string", check(func(a float64, b string) bool { return differentHash(a, b) }))
		t.Run("complex64!=complex128", check(func(a complex64, b complex128) bool { return differentHash(a, b) }))
		t.Run("complex64!=bool", check(func(a complex64, b bool) bool { return differentHash(a, b) }))
		t.Run("complex64!=string", check(func(a complex64, b string) bool { return differentHash(a, b) }))
		t.Run("complex128!=bool", check(func(a complex128, b bool) bool { return differentHash(a, b) }))
		t.Run("complex128!=string", check(func(a complex128, b string) bool { return differentHash(a, b) }))
		t.Run("bool!=string", check(func(a bool, b string) bool { return differentHash(a, b) }))
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

		t.Run("int", check(func(a, b [25]int) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8", check(func(a, b [25]int8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16", check(func(a, b [25]int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32", check(func(a, b [25]int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64", check(func(a, b [25]int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint", check(func(a, b [25]uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8", check(func(a, b [25]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16", check(func(a, b [25]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32", check(func(a, b [25]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64", check(func(a, b [25]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32", check(func(a, b [25]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float64", check(func(a, b [25]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex64", check(func(a, b [25]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex128", check(func(a, b [25]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("bool", check(func(a, b [25]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("string", check(func(a, b [25]string) bool { return sameTypeHashCheck(a, b) }))
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

		t.Run("int", check(func(a, b []int) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8", check(func(a, b []int8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16", check(func(a, b []int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32", check(func(a, b []int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64", check(func(a, b []int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint", check(func(a, b []uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8", check(func(a, b []uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16", check(func(a, b []uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32", check(func(a, b []uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64", check(func(a, b []uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32", check(func(a, b []float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float64", check(func(a, b []float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex64", check(func(a, b []complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex128", check(func(a, b []complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("bool", check(func(a, b []bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("string", check(func(a, b []string) bool { return sameTypeHashCheck(a, b) }))
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

		type ExportedPrimitives struct {
			Int        int
			Int8       int8
			Int16      int16
			Int32      int32
			Int64      int64
			Uint       uint
			Uint8      uint8
			Uint16     uint16
			Uint32     uint32
			Uint64     uint64
			Float32    float32
			Float64    float64
			Complex64  complex64
			Complex128 complex128
			Bool       bool
			String     string
		}
		t.Run("with primitives", check(func(a, b ExportedPrimitives) bool { return sameTypeHashCheck(a, b) }))

		type ExportedPrimitiveSlices struct {
			SliceInt        []int
			SliceInt8       []int8
			SliceInt16      []int16
			SliceInt32      []int32
			SliceInt64      []int64
			SliceUint       []uint
			SliceUint8      []uint8
			SliceUint16     []uint16
			SliceUint32     []uint32
			SliceUint64     []uint64
			SliceFloat32    []float32
			SliceFloat64    []float64
			SliceComplex64  []complex64
			SliceComplex128 []complex128
			SliceBool       []bool
			SliceString     []string
		}
		t.Run("with primitive slices", check(func(a, b ExportedPrimitiveSlices) bool { return sameTypeHashCheck(a, b) }))

		type ExportedPrimitiveArrays struct {
			ArrayInt        [10]int
			ArrayInt8       [10]int8
			ArrayInt16      [10]int16
			ArrayInt32      [10]int32
			ArrayInt64      [10]int64
			ArrayUint       [10]uint
			ArrayUint8      [10]uint8
			ArrayUint16     [10]uint16
			ArrayUint32     [10]uint32
			ArrayUint64     [10]uint64
			ArrayFloat32    [10]float32
			ArrayFloat64    [10]float64
			ArrayComplex64  [10]complex64
			ArrayComplex128 [10]complex128
			ArrayBool       [10]bool
			ArrayString     [10]string
		}
		t.Run("with primitive arrays", check(func(a, b ExportedPrimitiveArrays) bool { return sameTypeHashCheck(a, b) }))

		type ExportedNested struct {
			Value           string
			AnotherExported *ExportedNested
			SliceExported   []*ExportedNested
			ArrayExported   [10]*ExportedNested
		}
		t.Run("with nested structs", check(func(a, b ExportedNested) bool {
			return sameTypeHashCheck(a, b)
		}))
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

		t.Run("int:int", check(func(a, b map[int]int) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:int8", check(func(a, b map[int]int8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:int16", check(func(a, b map[int]int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:int32", check(func(a, b map[int]int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:int64", check(func(a, b map[int]int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:uint", check(func(a, b map[int]uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:uint8", check(func(a, b map[int]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:uint16", check(func(a, b map[int]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:uint32", check(func(a, b map[int]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:uint64", check(func(a, b map[int]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:float32", check(func(a, b map[int]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:float64", check(func(a, b map[int]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:complex64", check(func(a, b map[int]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:complex128", check(func(a, b map[int]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:bool", check(func(a, b map[int]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int:string", check(func(a, b map[int]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:int8", check(func(a, b map[int8]int8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:int16", check(func(a, b map[int8]int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:int32", check(func(a, b map[int8]int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:int64", check(func(a, b map[int8]int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:uint", check(func(a, b map[int8]uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:uint8", check(func(a, b map[int8]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:uint16", check(func(a, b map[int8]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:uint32", check(func(a, b map[int8]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:uint64", check(func(a, b map[int8]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:float32", check(func(a, b map[int8]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:float64", check(func(a, b map[int8]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:complex64", check(func(a, b map[int8]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:complex128", check(func(a, b map[int8]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:bool", check(func(a, b map[int8]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int8:string", check(func(a, b map[int8]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:int16", check(func(a, b map[int16]int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:int32", check(func(a, b map[int16]int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:int64", check(func(a, b map[int16]int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:uint", check(func(a, b map[int16]uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:uint8", check(func(a, b map[int16]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:uint16", check(func(a, b map[int16]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:uint32", check(func(a, b map[int16]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:uint64", check(func(a, b map[int16]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:float32", check(func(a, b map[int16]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:float64", check(func(a, b map[int16]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:complex64", check(func(a, b map[int16]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:complex128", check(func(a, b map[int16]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:bool", check(func(a, b map[int16]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int16:string", check(func(a, b map[int16]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:int32", check(func(a, b map[int32]int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:int64", check(func(a, b map[int32]int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:uint", check(func(a, b map[int32]uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:uint8", check(func(a, b map[int32]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:uint16", check(func(a, b map[int32]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:uint32", check(func(a, b map[int32]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:uint64", check(func(a, b map[int32]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:float32", check(func(a, b map[int32]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:float64", check(func(a, b map[int32]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:complex64", check(func(a, b map[int32]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:complex128", check(func(a, b map[int32]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:bool", check(func(a, b map[int32]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int32:string", check(func(a, b map[int32]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:int64", check(func(a, b map[int64]int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:uint", check(func(a, b map[int64]uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:uint8", check(func(a, b map[int64]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:uint16", check(func(a, b map[int64]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:uint32", check(func(a, b map[int64]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:uint64", check(func(a, b map[int64]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:float32", check(func(a, b map[int64]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:float64", check(func(a, b map[int64]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:complex64", check(func(a, b map[int64]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:complex128", check(func(a, b map[int64]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:bool", check(func(a, b map[int64]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("int64:string", check(func(a, b map[int64]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:uint", check(func(a, b map[uint]uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:uint8", check(func(a, b map[uint]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:uint16", check(func(a, b map[uint]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:uint32", check(func(a, b map[uint]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:uint64", check(func(a, b map[uint]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:float32", check(func(a, b map[uint]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:float64", check(func(a, b map[uint]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:complex64", check(func(a, b map[uint]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:complex128", check(func(a, b map[uint]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:bool", check(func(a, b map[uint]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint:string", check(func(a, b map[uint]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:uint8", check(func(a, b map[uint8]uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:uint16", check(func(a, b map[uint8]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:uint32", check(func(a, b map[uint8]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:uint64", check(func(a, b map[uint8]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:float32", check(func(a, b map[uint8]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:float64", check(func(a, b map[uint8]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:complex64", check(func(a, b map[uint8]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:complex128", check(func(a, b map[uint8]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:bool", check(func(a, b map[uint8]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint8:string", check(func(a, b map[uint8]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:uint16", check(func(a, b map[uint16]uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:uint32", check(func(a, b map[uint16]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:uint64", check(func(a, b map[uint16]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:float32", check(func(a, b map[uint16]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:float64", check(func(a, b map[uint16]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:complex64", check(func(a, b map[uint16]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:complex128", check(func(a, b map[uint16]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:bool", check(func(a, b map[uint16]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint16:string", check(func(a, b map[uint16]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:uint32", check(func(a, b map[uint32]uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:uint64", check(func(a, b map[uint32]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:float32", check(func(a, b map[uint32]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:float64", check(func(a, b map[uint32]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:complex64", check(func(a, b map[uint32]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:complex128", check(func(a, b map[uint32]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:bool", check(func(a, b map[uint32]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint32:string", check(func(a, b map[uint32]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64:uint64", check(func(a, b map[uint64]uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64:float32", check(func(a, b map[uint64]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64:float64", check(func(a, b map[uint64]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64:complex64", check(func(a, b map[uint64]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64:complex128", check(func(a, b map[uint64]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64:bool", check(func(a, b map[uint64]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("uint64:string", check(func(a, b map[uint64]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32:float32", check(func(a, b map[float32]float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32:float64", check(func(a, b map[float32]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32:complex64", check(func(a, b map[float32]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32:complex128", check(func(a, b map[float32]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32:bool", check(func(a, b map[float32]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float32:string", check(func(a, b map[float32]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float64:float64", check(func(a, b map[float64]float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float64:complex64", check(func(a, b map[float64]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float64:complex128", check(func(a, b map[float64]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float64:bool", check(func(a, b map[float64]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("float64:string", check(func(a, b map[float64]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex64:complex64", check(func(a, b map[complex64]complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex64:complex128", check(func(a, b map[complex64]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex64:bool", check(func(a, b map[complex64]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex64:string", check(func(a, b map[complex64]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex128:complex128", check(func(a, b map[complex128]complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex128:bool", check(func(a, b map[complex128]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("complex128:string", check(func(a, b map[complex128]string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("bool:bool", check(func(a, b map[bool]bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("bool:string", check(func(a, b map[bool]string) bool { return sameTypeHashCheck(a, b) }))

		t.Run("*int:*int", check(func(a, b map[*int]*int) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*int8", check(func(a, b map[*int]*int8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*int16", check(func(a, b map[*int]*int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*int32", check(func(a, b map[*int]*int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*int64", check(func(a, b map[*int]*int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*uint", check(func(a, b map[*int]*uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*uint8", check(func(a, b map[*int]*uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*uint16", check(func(a, b map[*int]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*uint32", check(func(a, b map[*int]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*uint64", check(func(a, b map[*int]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*float32", check(func(a, b map[*int]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*float64", check(func(a, b map[*int]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*complex64", check(func(a, b map[*int]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*complex128", check(func(a, b map[*int]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*bool", check(func(a, b map[*int]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int:*string", check(func(a, b map[*int]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*int8", check(func(a, b map[*int8]*int8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*int16", check(func(a, b map[*int8]*int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*int32", check(func(a, b map[*int8]*int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*int64", check(func(a, b map[*int8]*int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*uint", check(func(a, b map[*int8]*uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*uint8", check(func(a, b map[*int8]*uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*uint16", check(func(a, b map[*int8]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*uint32", check(func(a, b map[*int8]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*uint64", check(func(a, b map[*int8]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*float32", check(func(a, b map[*int8]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*float64", check(func(a, b map[*int8]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*complex64", check(func(a, b map[*int8]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*complex128", check(func(a, b map[*int8]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*bool", check(func(a, b map[*int8]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8:*string", check(func(a, b map[*int8]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*int16", check(func(a, b map[*int16]*int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*int32", check(func(a, b map[*int16]*int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*int64", check(func(a, b map[*int16]*int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*uint", check(func(a, b map[*int16]*uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*uint8", check(func(a, b map[*int16]*uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*uint16", check(func(a, b map[*int16]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*uint32", check(func(a, b map[*int16]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*uint64", check(func(a, b map[*int16]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*float32", check(func(a, b map[*int16]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*float64", check(func(a, b map[*int16]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*complex64", check(func(a, b map[*int16]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*complex128", check(func(a, b map[*int16]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*bool", check(func(a, b map[*int16]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16:*string", check(func(a, b map[*int16]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*int32", check(func(a, b map[*int32]*int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*int64", check(func(a, b map[*int32]*int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*uint", check(func(a, b map[*int32]*uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*uint8", check(func(a, b map[*int32]*uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*uint16", check(func(a, b map[*int32]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*uint32", check(func(a, b map[*int32]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*uint64", check(func(a, b map[*int32]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*float32", check(func(a, b map[*int32]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*float64", check(func(a, b map[*int32]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*complex64", check(func(a, b map[*int32]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*complex128", check(func(a, b map[*int32]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*bool", check(func(a, b map[*int32]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32:*string", check(func(a, b map[*int32]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*int64", check(func(a, b map[*int64]*int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*uint", check(func(a, b map[*int64]*uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*uint8", check(func(a, b map[*int64]*uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*uint16", check(func(a, b map[*int64]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*uint32", check(func(a, b map[*int64]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*uint64", check(func(a, b map[*int64]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*float32", check(func(a, b map[*int64]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*float64", check(func(a, b map[*int64]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*complex64", check(func(a, b map[*int64]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*complex128", check(func(a, b map[*int64]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*bool", check(func(a, b map[*int64]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64:*string", check(func(a, b map[*int64]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*uint", check(func(a, b map[*uint]*uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*uint8", check(func(a, b map[*uint]*uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*uint16", check(func(a, b map[*uint]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*uint32", check(func(a, b map[*uint]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*uint64", check(func(a, b map[*uint]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*float32", check(func(a, b map[*uint]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*float64", check(func(a, b map[*uint]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*complex64", check(func(a, b map[*uint]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*complex128", check(func(a, b map[*uint]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*bool", check(func(a, b map[*uint]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint:*string", check(func(a, b map[*uint]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*uint8", check(func(a, b map[*uint8]*uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*uint16", check(func(a, b map[*uint8]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*uint32", check(func(a, b map[*uint8]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*uint64", check(func(a, b map[*uint8]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*float32", check(func(a, b map[*uint8]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*float64", check(func(a, b map[*uint8]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*complex64", check(func(a, b map[*uint8]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*complex128", check(func(a, b map[*uint8]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*bool", check(func(a, b map[*uint8]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8:*string", check(func(a, b map[*uint8]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*uint16", check(func(a, b map[*uint16]*uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*uint32", check(func(a, b map[*uint16]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*uint64", check(func(a, b map[*uint16]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*float32", check(func(a, b map[*uint16]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*float64", check(func(a, b map[*uint16]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*complex64", check(func(a, b map[*uint16]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*complex128", check(func(a, b map[*uint16]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*bool", check(func(a, b map[*uint16]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16:*string", check(func(a, b map[*uint16]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*uint32", check(func(a, b map[*uint32]*uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*uint64", check(func(a, b map[*uint32]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*float32", check(func(a, b map[*uint32]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*float64", check(func(a, b map[*uint32]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*complex64", check(func(a, b map[*uint32]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*complex128", check(func(a, b map[*uint32]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*bool", check(func(a, b map[*uint32]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32:*string", check(func(a, b map[*uint32]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64:*uint64", check(func(a, b map[*uint64]*uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64:*float32", check(func(a, b map[*uint64]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64:*float64", check(func(a, b map[*uint64]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64:*complex64", check(func(a, b map[*uint64]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64:*complex128", check(func(a, b map[*uint64]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64:*bool", check(func(a, b map[*uint64]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64:*string", check(func(a, b map[*uint64]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float32:*float32", check(func(a, b map[*float32]*float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float32:*float64", check(func(a, b map[*float32]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float32:*complex64", check(func(a, b map[*float32]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float32:*complex128", check(func(a, b map[*float32]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float32:*bool", check(func(a, b map[*float32]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float32:*string", check(func(a, b map[*float32]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float64:*float64", check(func(a, b map[*float64]*float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float64:*complex64", check(func(a, b map[*float64]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float64:*complex128", check(func(a, b map[*float64]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float64:*bool", check(func(a, b map[*float64]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float64:*string", check(func(a, b map[*float64]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex64:*complex64", check(func(a, b map[*complex64]*complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex64:*complex128", check(func(a, b map[*complex64]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex64:*bool", check(func(a, b map[*complex64]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex64:*string", check(func(a, b map[*complex64]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex128:*complex128", check(func(a, b map[*complex128]*complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex128:*bool", check(func(a, b map[*complex128]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex128:*string", check(func(a, b map[*complex128]*string) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*bool:*bool", check(func(a, b map[*bool]*bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*bool:*string", check(func(a, b map[*bool]*string) bool { return sameTypeHashCheck(a, b) }))
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

		t.Run("*int", check(func(a, b *int) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int8", check(func(a, b *int8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int16", check(func(a, b *int16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int32", check(func(a, b *int32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*int64", check(func(a, b *int64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint", check(func(a, b *uint) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint8", check(func(a, b *uint8) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint16", check(func(a, b *uint16) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint32", check(func(a, b *uint32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*uint64", check(func(a, b *uint64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float32", check(func(a, b *float32) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*float64", check(func(a, b *float64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex64", check(func(a, b *complex64) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*complex128", check(func(a, b *complex128) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*bool", check(func(a, b *bool) bool { return sameTypeHashCheck(a, b) }))
		t.Run("*string", check(func(a, b *string) bool { return sameTypeHashCheck(a, b) }))
	})
}
