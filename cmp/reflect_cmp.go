package cmp

import "reflect"

func Reflect_cmp(a interface{}, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
