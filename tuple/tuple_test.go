package tuple_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/testing/assert"
	"github.com/gtramontina/go-extlib/tuple"
)

func TestTuple(t *testing.T) {
	type sample struct{ x int }

	t.Run("of 1", func(t *testing.T) {
		monuple := tuple.Of1("a")
		assert.DeepEqual(t, monuple.Get1(), "a")

		val, ok := monuple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = monuple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), monuple.Get1())

		val, ok = monuple.GetN(2)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 2", func(t *testing.T) {
		couple := tuple.Of2("a", 2)
		assert.DeepEqual(t, couple.Get1(), "a")
		assert.DeepEqual(t, couple.Get2(), 2)

		val, ok := couple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = couple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), couple.Get1())

		val, ok = couple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), couple.Get2())

		val, ok = couple.GetN(3)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 3", func(t *testing.T) {
		triple := tuple.Of3("a", 2, sample{3})
		assert.DeepEqual(t, triple.Get1(), "a")
		assert.DeepEqual(t, triple.Get2(), 2)
		assert.DeepEqual(t, triple.Get3(), sample{3})

		val, ok := triple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = triple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), triple.Get1())

		val, ok = triple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), triple.Get2())

		val, ok = triple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), triple.Get3())

		val, ok = triple.GetN(4)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 4", func(t *testing.T) {
		quadruple := tuple.Of4("a", 2, sample{3}, 'd')
		assert.DeepEqual(t, quadruple.Get1(), "a")
		assert.DeepEqual(t, quadruple.Get2(), 2)
		assert.DeepEqual(t, quadruple.Get3(), sample{3})
		assert.DeepEqual(t, quadruple.Get4(), 'd')

		val, ok := quadruple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = quadruple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), quadruple.Get1())

		val, ok = quadruple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), quadruple.Get2())

		val, ok = quadruple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), quadruple.Get3())

		val, ok = quadruple.GetN(4)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(rune), quadruple.Get4())

		val, ok = quadruple.GetN(5)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 5", func(t *testing.T) {
		quintuple := tuple.Of5("a", 2, sample{3}, 'd', 5.5)
		assert.DeepEqual(t, quintuple.Get1(), "a")
		assert.DeepEqual(t, quintuple.Get2(), 2)
		assert.DeepEqual(t, quintuple.Get3(), sample{3})
		assert.DeepEqual(t, quintuple.Get4(), 'd')
		assert.DeepEqual(t, quintuple.Get5(), 5.5)

		val, ok := quintuple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = quintuple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), quintuple.Get1())

		val, ok = quintuple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), quintuple.Get2())

		val, ok = quintuple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), quintuple.Get3())

		val, ok = quintuple.GetN(4)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(rune), quintuple.Get4())

		val, ok = quintuple.GetN(5)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(float64), quintuple.Get5())

		val, ok = quintuple.GetN(6)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 6", func(t *testing.T) {
		sextuple := tuple.Of6("a", 2, sample{3}, 'd', 5.5, true)
		assert.DeepEqual(t, sextuple.Get1(), "a")
		assert.DeepEqual(t, sextuple.Get2(), 2)
		assert.DeepEqual(t, sextuple.Get3(), sample{3})
		assert.DeepEqual(t, sextuple.Get4(), 'd')
		assert.DeepEqual(t, sextuple.Get5(), 5.5)
		assert.DeepEqual(t, sextuple.Get6(), true)

		val, ok := sextuple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = sextuple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), sextuple.Get1())

		val, ok = sextuple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), sextuple.Get2())

		val, ok = sextuple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), sextuple.Get3())

		val, ok = sextuple.GetN(4)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(rune), sextuple.Get4())

		val, ok = sextuple.GetN(5)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(float64), sextuple.Get5())

		val, ok = sextuple.GetN(6)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(bool), sextuple.Get6())

		val, ok = sextuple.GetN(7)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 7", func(t *testing.T) {
		septuple := tuple.Of7("a", 2, sample{3}, 'd', 5.5, true, uint(7))
		assert.DeepEqual(t, septuple.Get1(), "a")
		assert.DeepEqual(t, septuple.Get2(), 2)
		assert.DeepEqual(t, septuple.Get3(), sample{3})
		assert.DeepEqual(t, septuple.Get4(), 'd')
		assert.DeepEqual(t, septuple.Get5(), 5.5)
		assert.DeepEqual(t, septuple.Get6(), true)
		assert.DeepEqual(t, septuple.Get7(), 7)

		val, ok := septuple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = septuple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), septuple.Get1())

		val, ok = septuple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), septuple.Get2())

		val, ok = septuple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), septuple.Get3())

		val, ok = septuple.GetN(4)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(rune), septuple.Get4())

		val, ok = septuple.GetN(5)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(float64), septuple.Get5())

		val, ok = septuple.GetN(6)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(bool), septuple.Get6())

		val, ok = septuple.GetN(7)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(uint), septuple.Get7())

		val, ok = septuple.GetN(8)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 8", func(t *testing.T) {
		octuple := tuple.Of8("a", 2, sample{3}, 'd', 5.5, true, uint(7), complex(8, 5))
		assert.DeepEqual(t, octuple.Get1(), "a")
		assert.DeepEqual(t, octuple.Get2(), 2)
		assert.DeepEqual(t, octuple.Get3(), sample{3})
		assert.DeepEqual(t, octuple.Get4(), 'd')
		assert.DeepEqual(t, octuple.Get5(), 5.5)
		assert.DeepEqual(t, octuple.Get6(), true)
		assert.DeepEqual(t, octuple.Get7(), 7)
		assert.DeepEqual(t, octuple.Get8(), complex(8, 5))

		val, ok := octuple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = octuple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), octuple.Get1())

		val, ok = octuple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), octuple.Get2())

		val, ok = octuple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), octuple.Get3())

		val, ok = octuple.GetN(4)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(rune), octuple.Get4())

		val, ok = octuple.GetN(5)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(float64), octuple.Get5())

		val, ok = octuple.GetN(6)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(bool), octuple.Get6())

		val, ok = octuple.GetN(7)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(uint), octuple.Get7())

		val, ok = octuple.GetN(8)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(complex128), octuple.Get8())

		val, ok = octuple.GetN(9)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 9", func(t *testing.T) {
		nonuple := tuple.Of9("a", 2, sample{3}, 'd', 5.5, true, uint(7), complex(8, 5), []int{9})
		assert.DeepEqual(t, nonuple.Get1(), "a")
		assert.DeepEqual(t, nonuple.Get2(), 2)
		assert.DeepEqual(t, nonuple.Get3(), sample{3})
		assert.DeepEqual(t, nonuple.Get4(), 'd')
		assert.DeepEqual(t, nonuple.Get5(), 5.5)
		assert.DeepEqual(t, nonuple.Get6(), true)
		assert.DeepEqual(t, nonuple.Get7(), 7)
		assert.DeepEqual(t, nonuple.Get8(), complex(8, 5))
		assert.DeepEqual(t, nonuple.Get9(), []int{9})

		val, ok := nonuple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = nonuple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), nonuple.Get1())

		val, ok = nonuple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), nonuple.Get2())

		val, ok = nonuple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), nonuple.Get3())

		val, ok = nonuple.GetN(4)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(rune), nonuple.Get4())

		val, ok = nonuple.GetN(5)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(float64), nonuple.Get5())

		val, ok = nonuple.GetN(6)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(bool), nonuple.Get6())

		val, ok = nonuple.GetN(7)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(uint), nonuple.Get7())

		val, ok = nonuple.GetN(8)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(complex128), nonuple.Get8())

		val, ok = nonuple.GetN(9)
		assert.True(t, ok)
		assert.DeepEqual(t, val.([]int), nonuple.Get9())

		val, ok = nonuple.GetN(10)
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("of 10", func(t *testing.T) {
		decuple := tuple.Of10("a", 2, sample{3}, 'd', 5.5, true, uint(7), complex(8, 5), []int{9}, map[int]string{10: "ten"})
		assert.DeepEqual(t, decuple.Get1(), "a")
		assert.DeepEqual(t, decuple.Get2(), 2)
		assert.DeepEqual(t, decuple.Get3(), sample{3})
		assert.DeepEqual(t, decuple.Get4(), 'd')
		assert.DeepEqual(t, decuple.Get5(), 5.5)
		assert.DeepEqual(t, decuple.Get6(), true)
		assert.DeepEqual(t, decuple.Get7(), 7)
		assert.DeepEqual(t, decuple.Get8(), complex(8, 5))
		assert.DeepEqual(t, decuple.Get9(), []int{9})
		assert.DeepEqual(t, decuple.Get10(), map[int]string{10: "ten"})

		val, ok := decuple.GetN(0)
		assert.False(t, ok)
		assert.Nil(t, val)

		val, ok = decuple.GetN(1)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(string), decuple.Get1())

		val, ok = decuple.GetN(2)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(int), decuple.Get2())

		val, ok = decuple.GetN(3)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(sample), decuple.Get3())

		val, ok = decuple.GetN(4)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(rune), decuple.Get4())

		val, ok = decuple.GetN(5)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(float64), decuple.Get5())

		val, ok = decuple.GetN(6)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(bool), decuple.Get6())

		val, ok = decuple.GetN(7)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(uint), decuple.Get7())

		val, ok = decuple.GetN(8)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(complex128), decuple.Get8())

		val, ok = decuple.GetN(9)
		assert.True(t, ok)
		assert.DeepEqual(t, val.([]int), decuple.Get9())

		val, ok = decuple.GetN(10)
		assert.True(t, ok)
		assert.DeepEqual(t, val.(map[int]string), decuple.Get10())

		val, ok = decuple.GetN(11)
		assert.False(t, ok)
		assert.Nil(t, val)
	})
}
