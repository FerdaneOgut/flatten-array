package main

import (
	"fmt"
	"reflect"
)

func main() {
	mixed := []interface{}{"foo", []interface{}{2, 3, []interface{}{"a", "b", []float64{1.3, 5.6}}, ""}, "test", nil}
	res := flatten(mixed)
	fmt.Println(res)
}

//to flatten giving arbitrary array
//reflect library used to be able to determine the types
func flatten(arr []interface{}) []interface{} {
	//slice of result
	s := make([]interface{}, 0)
	//function variable for recursive calls
	var iterate func(val interface{})

	iterate = func(val interface{}) {
		typ := reflect.TypeOf(val).Kind()
		//to make sure we get the correct type that is array
		if typ == reflect.Slice || typ == reflect.Array {
			r := reflect.ValueOf(val)
			for i := 0; i < r.Len(); i++ {
				v := r.Index(i).Interface()
				//to not let app panic we check if there is nil value in array
				//if there is append to slice
				if v == nil {
					s = append(s, nil)
					continue
				}
				//check the type if array or slice make recursive call
				//if not add to result array
				switch reflect.TypeOf(v).Kind() {
				case reflect.Slice:
					iterate(v)
				case reflect.Array:
					iterate(v)
				default:
					s = append(s, v)
				}
			}
		}
	}
	iterate(arr)
	return s
}
