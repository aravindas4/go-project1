package tested

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Struct:
		for index := 0; index < val.NumField(); index++ {
			walkValue(val.Field(index))
		}
	case reflect.Slice, reflect.Array:
		for index := 0; index < val.Len(); index++ {
			walkValue(val.Index(index))
		}
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for value, ok := val.Recv(); ok; value, ok = val.Recv() {
			walkValue(value)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
