package internal

import "reflect"

func NewOf[Type any]() (*Type, []any) {
	var result Type
	theValue := reflect.ValueOf(&result).Elem()
	theType := reflect.TypeOf(result)

	if theType.Kind() == reflect.Ptr {
		result = reflect.New(theType.Elem()).Interface().(Type)
		theValue = reflect.ValueOf(result).Elem()
		theType = theType.Elem()
	}

	var fields []any
	fields = make([]any, theType.NumField())

	for i := 0; i < theType.NumField(); i++ {
		fields[i] = theValue.Field(i).Addr().Interface()
	}

	return &result, fields
}
