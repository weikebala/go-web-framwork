package cmp

import (
	"errors"
)

func In_array(item interface{}, arr interface{})(bool,error){
	//value := reflect.ValueOf(arr)
	//k := value.Kind()
	//if (k != reflect.Array || k != reflect.Map ||  k != reflect.Slice) {
	//	return false, errors.New("非法类型")
	//}

	return false, errors.New("非法类型")
}
