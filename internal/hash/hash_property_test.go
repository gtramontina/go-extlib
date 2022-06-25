package hash_test

import (
	"os"
	"reflect"
	"strconv"
	"testing"
	"testing/quick"

	"github.com/gtramontina/go-extlib/internal/hash"
	"github.com/gtramontina/go-extlib/testing/assert"
)

// maxCountScale reads the `MAX_COUNT_SCALE` environment variable and parses it
// as float64 if available. It defaults to 0.1 otherwise.
// FIXME: maxCountScale is a workaround to the `-quickchecks` flag.
func maxCountScale() float64 {
	maxCountScaleEnv := os.Getenv("MAX_COUNT_SCALE")
	if len(maxCountScaleEnv) > 0 {
		maxCountScale, err := strconv.ParseFloat(maxCountScaleEnv, 64)
		if err != nil {
			panic(err)
		}

		return maxCountScale
	}

	return 0.1
}

func TestHashProperties(t *testing.T) {
	config := &quick.Config{
		MaxCountScale: maxCountScale(),
	}

	sameHash := func(a, b any) bool {
		return hash.Calc(a) == hash.Calc(b)
	}

	differentHash := func(a, b any) bool {
		return !sameHash(a, b)
	}

	sameTypeHashCheck := func(a, b any) bool {
		return reflect.DeepEqual(a, b) || (sameHash(a, a) && sameHash(b, b) && differentHash(a, b))
	}

	check := func(fn any) func(*testing.T) {
		return func(t *testing.T) {
			t.Helper()
			assert.NoError(t, quick.Check(fn, config))
		}
	}

	t.Run("bool", check(func(a, b bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("int", check(func(a, b int) bool { return sameTypeHashCheck(a, b) }))
	t.Run("int8", check(func(a, b int8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("int16", check(func(a, b int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("int32", check(func(a, b int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("rune", check(func(a, b rune) bool { return sameTypeHashCheck(a, b) }))
	t.Run("int64", check(func(a, b int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("uint", check(func(a, b uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("uint8", check(func(a, b uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("uint16", check(func(a, b uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("uint32", check(func(a, b uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("uint64", check(func(a, b uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("float32", check(func(a, b float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("float64", check(func(a, b float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("complex64", check(func(a, b complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("complex128", check(func(a, b complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("string", check(func(a, b string) bool { return sameTypeHashCheck(a, b) }))

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

	t.Run("[n]int", check(func(a, b [25]int) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]int8", check(func(a, b [25]int8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]int16", check(func(a, b [25]int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]int32", check(func(a, b [25]int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]int64", check(func(a, b [25]int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]uint", check(func(a, b [25]uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]uint8", check(func(a, b [25]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]uint16", check(func(a, b [25]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]uint32", check(func(a, b [25]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]uint64", check(func(a, b [25]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]float32", check(func(a, b [25]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]float64", check(func(a, b [25]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]complex64", check(func(a, b [25]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]complex128", check(func(a, b [25]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]bool", check(func(a, b [25]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[n]string", check(func(a, b [25]string) bool { return sameTypeHashCheck(a, b) }))

	t.Run("[]int", check(func(a, b []int) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]int8", check(func(a, b []int8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]int16", check(func(a, b []int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]int32", check(func(a, b []int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]int64", check(func(a, b []int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]uint", check(func(a, b []uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]uint8", check(func(a, b []uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]uint16", check(func(a, b []uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]uint32", check(func(a, b []uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]uint64", check(func(a, b []uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]float32", check(func(a, b []float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]float64", check(func(a, b []float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]complex64", check(func(a, b []complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]complex128", check(func(a, b []complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]bool", check(func(a, b []bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("[]string", check(func(a, b []string) bool { return sameTypeHashCheck(a, b) }))

	t.Run("map[int]int", check(func(a, b map[int]int) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]int8", check(func(a, b map[int]int8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]int16", check(func(a, b map[int]int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]int32", check(func(a, b map[int]int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]int64", check(func(a, b map[int]int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]uint", check(func(a, b map[int]uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]uint8", check(func(a, b map[int]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]uint16", check(func(a, b map[int]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]uint32", check(func(a, b map[int]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]uint64", check(func(a, b map[int]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]float32", check(func(a, b map[int]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]float64", check(func(a, b map[int]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]complex64", check(func(a, b map[int]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]complex128", check(func(a, b map[int]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]bool", check(func(a, b map[int]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int]string", check(func(a, b map[int]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]int8", check(func(a, b map[int8]int8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]int16", check(func(a, b map[int8]int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]int32", check(func(a, b map[int8]int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]int64", check(func(a, b map[int8]int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]uint", check(func(a, b map[int8]uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]uint8", check(func(a, b map[int8]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]uint16", check(func(a, b map[int8]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]uint32", check(func(a, b map[int8]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]uint64", check(func(a, b map[int8]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]float32", check(func(a, b map[int8]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]float64", check(func(a, b map[int8]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]complex64", check(func(a, b map[int8]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]complex128", check(func(a, b map[int8]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]bool", check(func(a, b map[int8]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int8]string", check(func(a, b map[int8]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]int16", check(func(a, b map[int16]int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]int32", check(func(a, b map[int16]int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]int64", check(func(a, b map[int16]int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]uint", check(func(a, b map[int16]uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]uint8", check(func(a, b map[int16]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]uint16", check(func(a, b map[int16]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]uint32", check(func(a, b map[int16]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]uint64", check(func(a, b map[int16]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]float32", check(func(a, b map[int16]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]float64", check(func(a, b map[int16]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]complex64", check(func(a, b map[int16]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]complex128", check(func(a, b map[int16]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]bool", check(func(a, b map[int16]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int16]string", check(func(a, b map[int16]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]int32", check(func(a, b map[int32]int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]int64", check(func(a, b map[int32]int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]uint", check(func(a, b map[int32]uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]uint8", check(func(a, b map[int32]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]uint16", check(func(a, b map[int32]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]uint32", check(func(a, b map[int32]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]uint64", check(func(a, b map[int32]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]float32", check(func(a, b map[int32]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]float64", check(func(a, b map[int32]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]complex64", check(func(a, b map[int32]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]complex128", check(func(a, b map[int32]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]bool", check(func(a, b map[int32]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int32]string", check(func(a, b map[int32]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]int64", check(func(a, b map[int64]int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]uint", check(func(a, b map[int64]uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]uint8", check(func(a, b map[int64]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]uint16", check(func(a, b map[int64]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]uint32", check(func(a, b map[int64]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]uint64", check(func(a, b map[int64]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]float32", check(func(a, b map[int64]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]float64", check(func(a, b map[int64]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]complex64", check(func(a, b map[int64]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]complex128", check(func(a, b map[int64]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]bool", check(func(a, b map[int64]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[int64]string", check(func(a, b map[int64]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]uint", check(func(a, b map[uint]uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]uint8", check(func(a, b map[uint]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]uint16", check(func(a, b map[uint]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]uint32", check(func(a, b map[uint]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]uint64", check(func(a, b map[uint]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]float32", check(func(a, b map[uint]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]float64", check(func(a, b map[uint]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]complex64", check(func(a, b map[uint]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]complex128", check(func(a, b map[uint]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]bool", check(func(a, b map[uint]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint]string", check(func(a, b map[uint]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]uint8", check(func(a, b map[uint8]uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]uint16", check(func(a, b map[uint8]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]uint32", check(func(a, b map[uint8]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]uint64", check(func(a, b map[uint8]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]float32", check(func(a, b map[uint8]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]float64", check(func(a, b map[uint8]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]complex64", check(func(a, b map[uint8]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]complex128", check(func(a, b map[uint8]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]bool", check(func(a, b map[uint8]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint8]string", check(func(a, b map[uint8]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]uint16", check(func(a, b map[uint16]uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]uint32", check(func(a, b map[uint16]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]uint64", check(func(a, b map[uint16]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]float32", check(func(a, b map[uint16]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]float64", check(func(a, b map[uint16]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]complex64", check(func(a, b map[uint16]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]complex128", check(func(a, b map[uint16]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]bool", check(func(a, b map[uint16]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint16]string", check(func(a, b map[uint16]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]uint32", check(func(a, b map[uint32]uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]uint64", check(func(a, b map[uint32]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]float32", check(func(a, b map[uint32]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]float64", check(func(a, b map[uint32]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]complex64", check(func(a, b map[uint32]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]complex128", check(func(a, b map[uint32]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]bool", check(func(a, b map[uint32]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint32]string", check(func(a, b map[uint32]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint64]uint64", check(func(a, b map[uint64]uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint64]float32", check(func(a, b map[uint64]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint64]float64", check(func(a, b map[uint64]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint64]complex64", check(func(a, b map[uint64]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint64]complex128", check(func(a, b map[uint64]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint64]bool", check(func(a, b map[uint64]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[uint64]string", check(func(a, b map[uint64]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float32]float32", check(func(a, b map[float32]float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float32]float64", check(func(a, b map[float32]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float32]complex64", check(func(a, b map[float32]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float32]complex128", check(func(a, b map[float32]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float32]bool", check(func(a, b map[float32]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float32]string", check(func(a, b map[float32]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float64]float64", check(func(a, b map[float64]float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float64]complex64", check(func(a, b map[float64]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float64]complex128", check(func(a, b map[float64]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float64]bool", check(func(a, b map[float64]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[float64]string", check(func(a, b map[float64]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[complex64]complex64", check(func(a, b map[complex64]complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[complex64]complex128", check(func(a, b map[complex64]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[complex64]bool", check(func(a, b map[complex64]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[complex64]string", check(func(a, b map[complex64]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[complex128]complex128", check(func(a, b map[complex128]complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[complex128]bool", check(func(a, b map[complex128]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[complex128]string", check(func(a, b map[complex128]string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[bool]bool", check(func(a, b map[bool]bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[bool]string", check(func(a, b map[bool]string) bool { return sameTypeHashCheck(a, b) }))

	t.Run("map[*int]*int", check(func(a, b map[*int]*int) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*int8", check(func(a, b map[*int]*int8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*int16", check(func(a, b map[*int]*int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*int32", check(func(a, b map[*int]*int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*int64", check(func(a, b map[*int]*int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*uint", check(func(a, b map[*int]*uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*uint8", check(func(a, b map[*int]*uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*uint16", check(func(a, b map[*int]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*uint32", check(func(a, b map[*int]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*uint64", check(func(a, b map[*int]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*float32", check(func(a, b map[*int]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*float64", check(func(a, b map[*int]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*complex64", check(func(a, b map[*int]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*complex128", check(func(a, b map[*int]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*bool", check(func(a, b map[*int]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int]*string", check(func(a, b map[*int]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*int8", check(func(a, b map[*int8]*int8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*int16", check(func(a, b map[*int8]*int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*int32", check(func(a, b map[*int8]*int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*int64", check(func(a, b map[*int8]*int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*uint", check(func(a, b map[*int8]*uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*uint8", check(func(a, b map[*int8]*uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*uint16", check(func(a, b map[*int8]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*uint32", check(func(a, b map[*int8]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*uint64", check(func(a, b map[*int8]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*float32", check(func(a, b map[*int8]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*float64", check(func(a, b map[*int8]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*complex64", check(func(a, b map[*int8]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*complex128", check(func(a, b map[*int8]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*bool", check(func(a, b map[*int8]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int8]*string", check(func(a, b map[*int8]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*int16", check(func(a, b map[*int16]*int16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*int32", check(func(a, b map[*int16]*int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*int64", check(func(a, b map[*int16]*int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*uint", check(func(a, b map[*int16]*uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*uint8", check(func(a, b map[*int16]*uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*uint16", check(func(a, b map[*int16]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*uint32", check(func(a, b map[*int16]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*uint64", check(func(a, b map[*int16]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*float32", check(func(a, b map[*int16]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*float64", check(func(a, b map[*int16]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*complex64", check(func(a, b map[*int16]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*complex128", check(func(a, b map[*int16]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*bool", check(func(a, b map[*int16]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int16]*string", check(func(a, b map[*int16]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*int32", check(func(a, b map[*int32]*int32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*int64", check(func(a, b map[*int32]*int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*uint", check(func(a, b map[*int32]*uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*uint8", check(func(a, b map[*int32]*uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*uint16", check(func(a, b map[*int32]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*uint32", check(func(a, b map[*int32]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*uint64", check(func(a, b map[*int32]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*float32", check(func(a, b map[*int32]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*float64", check(func(a, b map[*int32]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*complex64", check(func(a, b map[*int32]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*complex128", check(func(a, b map[*int32]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*bool", check(func(a, b map[*int32]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int32]*string", check(func(a, b map[*int32]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*int64", check(func(a, b map[*int64]*int64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*uint", check(func(a, b map[*int64]*uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*uint8", check(func(a, b map[*int64]*uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*uint16", check(func(a, b map[*int64]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*uint32", check(func(a, b map[*int64]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*uint64", check(func(a, b map[*int64]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*float32", check(func(a, b map[*int64]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*float64", check(func(a, b map[*int64]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*complex64", check(func(a, b map[*int64]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*complex128", check(func(a, b map[*int64]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*bool", check(func(a, b map[*int64]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*int64]*string", check(func(a, b map[*int64]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*uint", check(func(a, b map[*uint]*uint) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*uint8", check(func(a, b map[*uint]*uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*uint16", check(func(a, b map[*uint]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*uint32", check(func(a, b map[*uint]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*uint64", check(func(a, b map[*uint]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*float32", check(func(a, b map[*uint]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*float64", check(func(a, b map[*uint]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*complex64", check(func(a, b map[*uint]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*complex128", check(func(a, b map[*uint]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*bool", check(func(a, b map[*uint]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint]*string", check(func(a, b map[*uint]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*uint8", check(func(a, b map[*uint8]*uint8) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*uint16", check(func(a, b map[*uint8]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*uint32", check(func(a, b map[*uint8]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*uint64", check(func(a, b map[*uint8]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*float32", check(func(a, b map[*uint8]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*float64", check(func(a, b map[*uint8]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*complex64", check(func(a, b map[*uint8]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*complex128", check(func(a, b map[*uint8]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*bool", check(func(a, b map[*uint8]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint8]*string", check(func(a, b map[*uint8]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*uint16", check(func(a, b map[*uint16]*uint16) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*uint32", check(func(a, b map[*uint16]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*uint64", check(func(a, b map[*uint16]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*float32", check(func(a, b map[*uint16]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*float64", check(func(a, b map[*uint16]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*complex64", check(func(a, b map[*uint16]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*complex128", check(func(a, b map[*uint16]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*bool", check(func(a, b map[*uint16]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint16]*string", check(func(a, b map[*uint16]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*uint32", check(func(a, b map[*uint32]*uint32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*uint64", check(func(a, b map[*uint32]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*float32", check(func(a, b map[*uint32]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*float64", check(func(a, b map[*uint32]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*complex64", check(func(a, b map[*uint32]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*complex128", check(func(a, b map[*uint32]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*bool", check(func(a, b map[*uint32]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint32]*string", check(func(a, b map[*uint32]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint64]*uint64", check(func(a, b map[*uint64]*uint64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint64]*float32", check(func(a, b map[*uint64]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint64]*float64", check(func(a, b map[*uint64]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint64]*complex64", check(func(a, b map[*uint64]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint64]*complex128", check(func(a, b map[*uint64]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint64]*bool", check(func(a, b map[*uint64]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*uint64]*string", check(func(a, b map[*uint64]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float32]*float32", check(func(a, b map[*float32]*float32) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float32]*float64", check(func(a, b map[*float32]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float32]*complex64", check(func(a, b map[*float32]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float32]*complex128", check(func(a, b map[*float32]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float32]*bool", check(func(a, b map[*float32]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float32]*string", check(func(a, b map[*float32]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float64]*float64", check(func(a, b map[*float64]*float64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float64]*complex64", check(func(a, b map[*float64]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float64]*complex128", check(func(a, b map[*float64]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float64]*bool", check(func(a, b map[*float64]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*float64]*string", check(func(a, b map[*float64]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*complex64]*complex64", check(func(a, b map[*complex64]*complex64) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*complex64]*complex128", check(func(a, b map[*complex64]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*complex64]*bool", check(func(a, b map[*complex64]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*complex64]*string", check(func(a, b map[*complex64]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*complex128]*complex128", check(func(a, b map[*complex128]*complex128) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*complex128]*bool", check(func(a, b map[*complex128]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*complex128]*string", check(func(a, b map[*complex128]*string) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*bool]*bool", check(func(a, b map[*bool]*bool) bool { return sameTypeHashCheck(a, b) }))
	t.Run("map[*bool]*string", check(func(a, b map[*bool]*string) bool { return sameTypeHashCheck(a, b) }))

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
	t.Run("with nested structs", check(func(a, b ExportedNested) bool { return sameTypeHashCheck(a, b) }))
}
