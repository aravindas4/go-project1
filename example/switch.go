package example

import (
	"fmt"
	"reflect"
)

func switch_caller(value interface{}) interface{} {
	switch t := value.(type) {
	case bool:
		fmt.Print(reflect.TypeOf(true))
		return reflect.TypeOf(true)
	case int:
		return reflect.TypeOf(1)
	default:
		return reflect.TypeOf(t)
	}
}
