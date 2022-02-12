package example

import (
	"fmt"
	"reflect"
)

func arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func CreateSlice(dataType reflect.Type, size int) []interface{} {
	return reflect.MakeSlice(reflect.SliceOf(dataType), size, size)
}

func AppendToSlice(slice []interface{}, item interface{}) []interface{} {
	slice = append(slice, item)
	return slice
}

func GetTypeOfSlice(slice []interface{}) reflect.Type {
	return reflect.TypeOf(slice[0])
}

func CopySlice(slice []interface{}) []interface{} {
	dataType := GetTypeOfSlice(slice)
	duplicate := reflect.MakeSlice(reflect.SliceOf(dataType), len(slice), len(slice)).Interface()
	copy(duplicate, slice)
	return duplicate
}
