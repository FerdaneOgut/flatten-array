package main

import (
	"reflect"
	"testing"
)

func Test_ShouldFlattenArray(t *testing.T) {
	testObject := struct {
		firstName string
		lastName  string
	}{"ferdane", "ogut"}

	arrMixed := []interface{}{"foo", []interface{}{2, 3, []interface{}{"a", "b", []float64{1.3, 5.6}}, ""}, "test", testObject}
	arrwithNil := []interface{}{"a", nil, []int{5, 6}}
	arrEmpty := []interface{}{}

	arrWithObject := []interface{}{testObject}

	testCases := []struct {
		expected []interface{}
		value    []interface{}
	}{
		{
			value:    arrMixed,
			expected: []interface{}{"foo", 2, 3, "a", "b", 1.3, 5.6, "", "test", testObject},
		},
		{
			value:    arrwithNil,
			expected: []interface{}{"a", nil, 5, 6},
		},
		{
			value:    arrEmpty,
			expected: []interface{}{},
		},
		{
			value:    arrWithObject,
			expected: []interface{}{testObject},
		},
	}
	for _, v := range testCases {
		res := flatten(v.value)
		expected := v.expected
		if len(expected) != len(expected) {
			t.Errorf("Couldn't flatten the array expected: %v, got :%v", expected, res)
		}
		isEqual := reflect.DeepEqual(res, expected)
		if !isEqual {
			t.Errorf("Couldn't flatten the array expected: %v, got :%v", expected, res)
		}
	}
}
