package json_test

import (
	"testing"

	"github.com/gtramontina/go-extlib/json"
	"github.com/gtramontina/go-extlib/testing/assert"
)

func ptr[T any](it T) *T { return &it }

func TestJSONUnmarshal(t *testing.T) {
	t.Run("primitives", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[bool]("false"), false)
		assert.DeepEqual(t, json.Unmarshal[bool]("true"), true)
		assert.DeepEqual(t, json.Unmarshal[int]("0"), 0)
		assert.DeepEqual(t, json.Unmarshal[int]("100"), 100)
		assert.DeepEqual(t, json.Unmarshal[int]("-100"), -100)
		assert.DeepEqual(t, json.Unmarshal[float64]("0.0"), 0.0)
		assert.DeepEqual(t, json.Unmarshal[float64]("100.0"), 100.0)
		assert.DeepEqual(t, json.Unmarshal[float64]("-100.0"), -100.0)
		assert.DeepEqual(t, json.Unmarshal[string](`""`), "")
		assert.DeepEqual(t, json.Unmarshal[string](`"hello"`), "hello")
	})

	t.Run("pointers to primitives", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[*bool]("false"), ptr(false))
		assert.DeepEqual(t, json.Unmarshal[*bool]("true"), ptr(true))
		assert.DeepEqual(t, json.Unmarshal[*int]("0"), ptr(0))
		assert.DeepEqual(t, json.Unmarshal[*int]("100"), ptr(100))
		assert.DeepEqual(t, json.Unmarshal[*int]("-100"), ptr(-100))
		assert.DeepEqual(t, json.Unmarshal[*float64]("0.0"), ptr(0.0))
		assert.DeepEqual(t, json.Unmarshal[*float64]("100.0"), ptr(100.0))
		assert.DeepEqual(t, json.Unmarshal[*float64]("-100.0"), ptr(-100.0))
		assert.DeepEqual(t, json.Unmarshal[*string](`""`), ptr(""))
		assert.DeepEqual(t, json.Unmarshal[*string](`"hello"`), ptr("hello"))
	})

	t.Run("structs", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[Person](`{"Name":"John", "Age":30}`), Person{Name: "John", Age: 30})
		assert.DeepEqual(t, json.Unmarshal[Person](`{"Name":"John", "Age":30, "Extra": "field"}`), Person{Name: "John", Age: 30})
	})

	t.Run("pointers to structs", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[*Person](`{"Name":"John", "Age":30}`), &Person{Name: "John", Age: 30})
		assert.DeepEqual(t, json.Unmarshal[*Person](`{"Name":"John", "Age":30, "Extra": "field"}`), &Person{Name: "John", Age: 30})
	})

	t.Run("slices", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[[]bool]("[true, false]"), []bool{true, false})
		assert.DeepEqual(t, json.Unmarshal[[]int]("[0, 1, 2]"), []int{0, 1, 2})
		assert.DeepEqual(t, json.Unmarshal[[]float64]("[0.0, 1.0, 2.0]"), []float64{0.0, 1.0, 2.0})
		assert.DeepEqual(t, json.Unmarshal[[]string](`["hello", "world"]`), []string{"hello", "world"})
		assert.DeepEqual(t, json.Unmarshal[[]Person](`[{"Name":"John", "Age":30}, {"Name":"Jane", "Age":25}]`), []Person{{Name: "John", Age: 30}, {Name: "Jane", Age: 25}})
	})

	t.Run("slices of pointers", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[[]*bool]("[true, false]"), []*bool{ptr(true), ptr(false)})
		assert.DeepEqual(t, json.Unmarshal[[]*int]("[0, 1, 2]"), []*int{ptr(0), ptr(1), ptr(2)})
		assert.DeepEqual(t, json.Unmarshal[[]*float64]("[0.0, 1.0, 2.0]"), []*float64{ptr(0.0), ptr(1.0), ptr(2.0)})
		assert.DeepEqual(t, json.Unmarshal[[]*string](`["hello", "world"]`), []*string{ptr("hello"), ptr("world")})
		assert.DeepEqual(t, json.Unmarshal[[]*Person](`[{"Name":"John", "Age":30}, {"Name":"Jane", "Age":25}]`), []*Person{{Name: "John", Age: 30}, {Name: "Jane", Age: 25}})
	})

	t.Run("maps", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[map[string]bool](`{"a":true, "b":false}`), map[string]bool{"a": true, "b": false})
		assert.DeepEqual(t, json.Unmarshal[map[string]int](`{"a":0, "b":1, "c":2}`), map[string]int{"a": 0, "b": 1, "c": 2})
		assert.DeepEqual(t, json.Unmarshal[map[string]float64](`{"a":0.0, "b":1.0, "c":2.0}`), map[string]float64{"a": 0.0, "b": 1.0, "c": 2.0})
		assert.DeepEqual(t, json.Unmarshal[map[string]string](`{"a":"hello", "b":"world"}`), map[string]string{"a": "hello", "b": "world"})
		assert.DeepEqual(t, json.Unmarshal[map[string]Person](`{"a":{"Name":"John", "Age":30}, "b":{"Name":"Jane", "Age":25}}`), map[string]Person{"a": {Name: "John", Age: 30}, "b": {Name: "Jane", Age: 25}})
		assert.DeepEqual(t, json.Unmarshal[map[int]bool](`{"0":true, "1":false}`), map[int]bool{0: true, 1: false})
		assert.DeepEqual(t, json.Unmarshal[map[int]int](`{"0":0, "1":1, "2":2}`), map[int]int{0: 0, 1: 1, 2: 2})
		assert.DeepEqual(t, json.Unmarshal[map[int]float64](`{"0":0.0, "1":1.0, "2":2.0}`), map[int]float64{0: 0.0, 1: 1.0, 2: 2.0})
		assert.DeepEqual(t, json.Unmarshal[map[int]string](`{"0":"hello", "1":"world"}`), map[int]string{0: "hello", 1: "world"})
		assert.DeepEqual(t, json.Unmarshal[map[int]Person](`{"0":{"Name":"John", "Age":30}, "1":{"Name":"Jane", "Age":25}}`), map[int]Person{0: {Name: "John", Age: 30}, 1: {Name: "Jane", Age: 25}})
	})

	t.Run("maps of pointers", func(t *testing.T) {
		assert.DeepEqual(t, json.Unmarshal[map[string]*bool](`{"a":true, "b":false}`), map[string]*bool{"a": ptr(true), "b": ptr(false)})
		assert.DeepEqual(t, json.Unmarshal[map[string]*int](`{"a":0, "b":1, "c":2}`), map[string]*int{"a": ptr(0), "b": ptr(1), "c": ptr(2)})
		assert.DeepEqual(t, json.Unmarshal[map[string]*float64](`{"a":0.0, "b":1.0, "c":2.0}`), map[string]*float64{"a": ptr(0.0), "b": ptr(1.0), "c": ptr(2.0)})
		assert.DeepEqual(t, json.Unmarshal[map[string]*string](`{"a":"hello", "b":"world"}`), map[string]*string{"a": ptr("hello"), "b": ptr("world")})
		assert.DeepEqual(t, json.Unmarshal[map[string]*Person](`{"a":{"Name":"John", "Age":30}, "b":{"Name":"Jane", "Age":25}}`), map[string]*Person{"a": {Name: "John", Age: 30}, "b": {Name: "Jane", Age: 25}})
		assert.DeepEqual(t, json.Unmarshal[map[int]*bool](`{"0":true, "1":false}`), map[int]*bool{0: ptr(true), 1: ptr(false)})
		assert.DeepEqual(t, json.Unmarshal[map[int]*int](`{"0":0, "1":1, "2":2}`), map[int]*int{0: ptr(0), 1: ptr(1), 2: ptr(2)})
		assert.DeepEqual(t, json.Unmarshal[map[int]*float64](`{"0":0.0, "1":1.0, "2":2.0}`), map[int]*float64{0: ptr(0.0), 1: ptr(1.0), 2: ptr(2.0)})
		assert.DeepEqual(t, json.Unmarshal[map[int]*string](`{"0":"hello", "1":"world"}`), map[int]*string{0: ptr("hello"), 1: ptr("world")})
	})

	t.Run("complex", func(t *testing.T) {
		type People []Person
		assert.DeepEqual(t, json.Unmarshal[People](`[{"Name":"John", "Age":30}, {"Name":"Jane", "Age":25}]`), People{Person{Name: "John", Age: 30}, Person{Name: "Jane", Age: 25}})
		assert.DeepEqual(t, json.Unmarshal[map[string]People](`{"a":[{"Name":"John", "Age":30}, {"Name":"Jane", "Age":25}]}`), map[string]People{"a": {Person{Name: "John", Age: 30}, Person{Name: "Jane", Age: 25}}})
		assert.DeepEqual(t, json.Unmarshal[map[string]*People](`{"a":[{"Name":"John", "Age":30}, {"Name":"Jane", "Age":25}]}`), map[string]*People{"a": {Person{Name: "John", Age: 30}, Person{Name: "Jane", Age: 25}}})
		assert.DeepEqual(t, json.Unmarshal[map[string]*map[string]*People](`{"a":{"b":[{"Name":"John", "Age":30}, {"Name":"Jane", "Age":25}]}}`), map[string]*map[string]*People{"a": {"b": &People{Person{Name: "John", Age: 30}, Person{Name: "Jane", Age: 25}}}})
		assert.DeepEqual(t, json.Unmarshal[map[string]*map[string]*map[string]*People](`{"a":{"b":{"c":[{"Name":"John", "Age":30}, {"Name":"Jane", "Age":25}]}}}`), map[string]*map[string]*map[string]*People{"a": {"b": {"c": &People{Person{Name: "John", Age: 30}, Person{Name: "Jane", Age: 25}}}}})
	})

	t.Run("panics on malformed json", func(t *testing.T) {
		assert.Panics(t, func() { json.Unmarshal[bool](``) })
		assert.Panics(t, func() { json.Unmarshal[int](``) })
		assert.Panics(t, func() { json.Unmarshal[string](`"`) })
	})
}

func TestJSONTryUnmarshal(t *testing.T) {
	{
		_, err := json.TryUnmarshal[bool](``)
		assert.Error(t, err)
	}
	{
		_, err := json.TryUnmarshal[int](``)
		assert.Error(t, err)
	}
	{
		_, err := json.TryUnmarshal[string](``)
		assert.Error(t, err)
	}

	{
		out, err := json.TryUnmarshal[bool]("true")
		assert.NoError(t, err)
		assert.DeepEqual(t, out, true)
	}
	{
		out, err := json.TryUnmarshal[[]string](`["hello", "world"]`)
		assert.NoError(t, err)
		assert.DeepEqual(t, out, []string{"hello", "world"})
	}
}

type Person struct {
	Name string
	Age  int
}
