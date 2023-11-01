/*
 *  utils_test.go
 *  Created on 22.02.2021
 *  Copyright (C) 2021 Volkswagen AG, All rights reserved.
 */

package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {

	type test struct {
		value  *string
		result string
	}

	value := "MyValue"
	cases := map[string]test{
		"simple value": {
			value:  &value,
			result: value,
		},
		"nil value": {
			value:  nil,
			result: "",
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToString(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToStringPtrOrNil(t *testing.T) {

	type test struct {
		value  string
		result *string
	}

	value := "MyValue"
	cases := map[string]test{
		"simple value": {
			value:  value,
			result: &value,
		},
		"empty value": {
			value:  "",
			result: nil,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToStringPtrOrNil(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToStringPtr(t *testing.T) {

	type test struct {
		value  string
		result *string
	}

	value := "foo"
	empty := ""
	cases := map[string]test{
		"simple value": {
			value:  "foo",
			result: &value,
		},
		"empty value": {
			value:  "",
			result: &empty,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToStringPtr(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToStringArray(t *testing.T) {

	t.Parallel()

	type test struct {
		given    *[]string
		expected []string
	}

	cases := map[string]test{
		"regular array": {
			given: &[]string{
				"foo",
				"bar",
			},
			expected: []string{
				"foo",
				"bar",
			},
		},
		"nil": {
			given:    nil,
			expected: nil,
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ToStringArray(tc.given)
			assert.EqualValues(t, tc.expected, result)
		})
	}
}

func TestToStringArrayPtr(t *testing.T) {

	t.Parallel()

	type test struct {
		given    []string
		expected *[]string
	}

	cases := map[string]test{
		"regular array": {
			given: []string{
				"foo",
				"bar",
			},
			expected: &[]string{
				"foo",
				"bar",
			},
		},
		"nil": {
			given:    nil,
			expected: nil,
		},
		"empty": {
			given:    []string{},
			expected: &[]string{},
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ToStringArrayPtr(tc.given)
			assert.EqualValues(t, tc.expected, result)
		})
	}
}

func TestToBool(t *testing.T) {

	type test struct {
		value  *bool
		result bool
	}

	value := true
	cases := map[string]test{
		"true value": {
			value:  &value,
			result: value,
		},
		"nil value": {
			value:  nil,
			result: false,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToBool(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToBoolPtr(t *testing.T) {

	type test struct {
		value  bool
		result *bool
	}

	value := true
	cases := map[string]test{
		"simple value": {
			value:  true,
			result: &value,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToBoolPtr(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToInt(t *testing.T) {

	type test struct {
		value  *int
		result int
	}

	value := 123
	cases := map[string]test{
		"value": {
			value:  &value,
			result: value,
		},
		"nil value": {
			value:  nil,
			result: 0,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToInt(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToIntPtr(t *testing.T) {

	type test struct {
		value  int
		result *int
	}

	value := 123
	cases := map[string]test{
		"simple value": {
			value:  value,
			result: &value,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToIntPtr(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToInt32(t *testing.T) {

	type test struct {
		value  *int32
		result int32
	}

	value := int32(1234)
	cases := map[string]test{
		"simple value": {
			value:  &value,
			result: value,
		},
		"nil value": {
			value:  nil,
			result: 0,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToInt32(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToInt32Ptr(t *testing.T) {

	type test struct {
		value  int32
		result *int32
	}

	value := int32(123)
	cases := map[string]test{
		"simple value": {
			value:  value,
			result: &value,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToInt32Ptr(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToInt64(t *testing.T) {

	type test struct {
		value  *int64
		result int64
	}

	value := int64(1234)
	cases := map[string]test{
		"simple value": {
			value:  &value,
			result: value,
		},
		"nil value": {
			value:  nil,
			result: 0,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToInt64(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToInt64Ptr(t *testing.T) {

	type test struct {
		value  int64
		result *int64
	}

	value := int64(123)
	cases := map[string]test{
		"simple value": {
			value:  value,
			result: &value,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToInt64Ptr(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToFloat32(t *testing.T) {

	type test struct {
		value  *float32
		result float32
	}

	value := float32(123.45)
	cases := map[string]test{
		"simple value": {
			value:  &value,
			result: value,
		},
		"nil value": {
			value:  nil,
			result: 0.0,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToFloat32(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToFloat32Ptr(t *testing.T) {

	type test struct {
		value  float32
		result *float32
	}

	var value float32 = 123.45

	cases := map[string]test{
		"simple value": {
			value:  value,
			result: &value,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToFloat32Ptr(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToFloat64(t *testing.T) {

	type test struct {
		value  *float64
		result float64
	}

	value := float64(123.45)
	cases := map[string]test{
		"simple value": {
			value:  &value,
			result: value,
		},
		"nil value": {
			value:  nil,
			result: 0.0,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToFloat64(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestToFloat64Ptr(t *testing.T) {

	type test struct {
		value  float64
		result *float64
	}

	var value = 123.45

	cases := map[string]test{
		"simple value": {
			value:  value,
			result: &value,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ToFloat64Ptr(tc.value)
			assrt.Equal(tc.result, result, "Values do not match")
		})
	}
}

func TestDeepCopy(t *testing.T) {

	type obj struct {
		Value string
	}

	type test struct {
		in       interface{}
		out      interface{}
		expected interface{}
		hasError bool
	}

	cases := map[string]test{
		"obj copied": {
			in: &obj{
				Value: "foo",
			},
			out: &obj{},
			expected: &obj{
				Value: "foo",
			},
			hasError: false,
		},
		"in err": {
			in:       complex64(0),
			out:      &obj{},
			hasError: true,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			err := DeepCopy(tc.in, &tc.out)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err, "an error may not occur")
				assert.EqualValues(t, tc.expected, tc.out, "Values do not match")
			}
		})
	}
}

func TestScale(t *testing.T) {

	type test struct {
		value  float64
		scale  int
		result float64
	}

	cases := map[string]test{
		"scale zero, is zero": {
			value:  0.0,
			scale:  2,
			result: 0.0,
		},
		"scale up": {
			value:  1.555,
			scale:  2,
			result: 1.56,
		},
		"scale down": {
			value:  1.554,
			scale:  2,
			result: 1.55,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ScaleHalfUp(tc.value, tc.scale)

			assrt.InDelta(tc.result, result, 0.001, "scaled value does not match")
		})
	}
}

func TestScaleUp(t *testing.T) {

	type test struct {
		value  float64
		scale  int
		result float64
	}

	cases := map[string]test{
		"scale zero, is zero": {
			value:  0.0,
			scale:  2,
			result: 0.0,
		},
		"scale up with 5": {
			value:  1.555,
			scale:  2,
			result: 1.56,
		},
		"scale up with 6": {
			value:  1.556,
			scale:  2,
			result: 1.56,
		},
		"scale up with 4": {
			value:  1.554,
			scale:  2,
			result: 1.56,
		},
		"scale up with 0": {
			value:  1.550,
			scale:  2,
			result: 1.55,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ScaleUp(tc.value, tc.scale)

			assrt.InDelta(tc.result, result, 0.001, "scaled value does not match")
		})
	}
}

func TestScaleDown(t *testing.T) {

	type test struct {
		value  float64
		scale  int
		result float64
	}

	cases := map[string]test{
		"scale zero, is zero": {
			value:  0.0,
			scale:  2,
			result: 0.0,
		},
		"scale down with 5": {
			value:  1.555,
			scale:  2,
			result: 1.55,
		},
		"scale down with 6": {
			value:  1.556,
			scale:  2,
			result: 1.55,
		},
		"scale down with 4": {
			value:  1.554,
			scale:  2,
			result: 1.55,
		},
		"scale down with 0": {
			value:  1.550,
			scale:  2,
			result: 1.55,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {

			assrt := assert.New(t)

			result := ScaleDown(tc.value, tc.scale)

			assrt.InDelta(tc.result, result, 0.001, "scaled value does not match")
		})
	}
}

func TestIsNumber(t *testing.T) {

	t.Parallel()

	type test struct {
		value    string
		expected bool
	}

	cases := map[string]test{
		"valid number 1": {
			value:    "12345",
			expected: true,
		},
		"valid number 2": {
			value:    "  12345  ",
			expected: true,
		},
		"invalid number 1": {
			value:    "",
			expected: false,
		},
		"invalid number 2": {
			value:    "123f",
			expected: false,
		},
		"invalid number 3": {
			value:    "foo",
			expected: false,
		},
		"invalid number 4": {
			value:    "0x123",
			expected: false,
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result := IsNumber(tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestIsNil(t *testing.T) {

	t.Parallel()

	type test struct {
		value    interface{}
		expected bool
	}

	type foo struct {
		bar string
	}

	cases := map[string]test{
		"is nil": {
			value:    nil,
			expected: true,
		},
		"is not nil (ptr)": {
			value:    &foo{bar: "foobar"},
			expected: false,
		},
		"is not nil (reference)": {
			value:    foo{bar: "foobar"},
			expected: false,
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := IsNil(tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func ExampleScaleHalfUp_noChange() {
	fmt.Println(ScaleHalfUp(5.9, 1))
	fmt.Println(ScaleHalfUp(5.99, 2))
	// Output: 5.9
	// 5.99
}

func ExampleScaleHalfUp_roundUp() {
	fmt.Println(ScaleHalfUp(5.99, 1))
	fmt.Println(ScaleHalfUp(5.995, 2))
	// Output: 6
	// 6
}

func ExampleScaleHalfUp_roundDown() {
	fmt.Println(ScaleHalfUp(5.94, 1))
	fmt.Println(ScaleHalfUp(5.994, 2))
	// Output: 5.9
	// 5.99
}

func ExampleScaleUp_noChange() {
	fmt.Println(ScaleUp(5.9, 1))
	fmt.Println(ScaleUp(5.99, 2))
	// Output: 5.9
	// 5.99
}

func ExampleScaleUp_roundUp() {
	fmt.Println(ScaleUp(5.99, 1))
	fmt.Println(ScaleUp(5.993, 2))
	// Output: 6
	// 6
}

func ExampleScaleDown_noChange() {
	fmt.Println(ScaleDown(5.9, 1))
	fmt.Println(ScaleDown(5.99, 2))
	// Output: 5.9
	// 5.99
}

func ExampleScaleDown_roundDown() {
	fmt.Println(ScaleDown(5.99, 1))
	fmt.Println(ScaleDown(5.997, 2))
	// Output: 5.9
	// 5.99
}
