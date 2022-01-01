package hash

import (
	"fmt"
	"hash/fnv"
	"reflect"
)

func Calc(subject any) uint64 {
	if subject == nil {
		return 0
	}

	value := reflect.ValueOf(subject)
	switch value.Kind() {
	case reflect.Bool:
		return hash(fmt.Sprintf("bool:%v", value.Bool()))
	case reflect.Int:
		return hash(fmt.Sprintf("int:%d", value.Int()))
	case reflect.Int8:
		return hash(fmt.Sprintf("int8:%d", value.Int()))
	case reflect.Int16:
		return hash(fmt.Sprintf("int16:%d", value.Int()))
	case reflect.Int32:
		return hash(fmt.Sprintf("int32:%d", value.Int()))
	case reflect.Int64:
		return hash(fmt.Sprintf("int64:%d", value.Int()))
	case reflect.Uint:
		return hash(fmt.Sprintf("uint:%d", value.Uint()))
	case reflect.Uint8:
		return hash(fmt.Sprintf("uint8:%d", value.Uint()))
	case reflect.Uint16:
		return hash(fmt.Sprintf("uint16:%d", value.Uint()))
	case reflect.Uint32:
		return hash(fmt.Sprintf("uint32:%d", value.Uint()))
	case reflect.Uint64:
		return hash(fmt.Sprintf("uint64:%d", value.Uint()))
	case reflect.Float32:
		return hash(fmt.Sprintf("float32:%f", value.Float()))
	case reflect.Float64:
		return hash(fmt.Sprintf("float64:%f", value.Float()))
	case reflect.Complex64:
		return hash(fmt.Sprintf("complex64:%f", value.Complex()))
	case reflect.Complex128:
		return hash(fmt.Sprintf("complex128:%f", value.Complex()))
	case reflect.Array:
		return hashArray(value)
	case reflect.Chan:
		return hash(fmt.Sprintf("chan:%d", value.UnsafePointer()))
	case reflect.Func:
		return hash(fmt.Sprintf("func:%d", value.UnsafePointer()))
	case reflect.Map:
		return hashMap(value)
	case reflect.Pointer:
		return hash(fmt.Sprintf("pointer:%d", value.UnsafePointer()))
	case reflect.Slice:
		return hashSlice(value)
	case reflect.String:
		return hash(fmt.Sprintf("string:%s", value.String()))
	case reflect.Struct:
		return hashStruct(value)
	default:
		panic(fmt.Sprintf(`can't calculate hash for "%s": %+v`, value.Type(), subject))
	}
}

func hashArray(value reflect.Value) uint64 {
	h := hash(fmt.Sprintf("array:%s", value.Type().Elem()))
	for i := 0; i < value.Len(); i++ {
		h = 31*h + Calc(value.Index(i).Interface())
	}
	return h
}

func hashMap(value reflect.Value) uint64 {
	h := hash(fmt.Sprintf("map:%s:%s", value.Type().Key(), value.Type().Elem()))
	iter := value.MapRange()
	for iter.Next() {
		pairH := Calc(iter.Key().Interface())
		pairH = 31*pairH + Calc(iter.Value().Interface())
		h ^= pairH
	}
	return h
}

func hashSlice(value reflect.Value) uint64 {
	h := hash(fmt.Sprintf("slice:%s", value.Type().Elem()))
	for i := 0; i < value.Len(); i++ {
		h = 31*h + Calc(value.Index(i).Interface())
	}
	return h
}

func hashStruct(value reflect.Value) uint64 {
	h := hash(fmt.Sprintf("struct:%s", value.Type()))

	addressableCopy := reflect.New(value.Type()).Elem()
	addressableCopy.Set(value)
	for i := 0; i < value.NumField(); i++ {
		field := addressableCopy.Field(i)
		field = reflect.NewAt(field.Type(), field.Addr().UnsafePointer()).Elem()
		h = 31*h + Calc(field.Interface())
	}

	return h
}

func hash(s string) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(s)) // fnv.sum64 never errors
	return h.Sum64()
}
