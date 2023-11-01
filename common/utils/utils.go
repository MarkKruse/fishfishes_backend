/*
 *  utils.go
 *  Created on 22.02.2021
 *  Copyright (C) 2021 Volkswagen AG, All rights reserved.
 */

// Package utils provides a variety of helper functions, for example for the conversion of primitive datatypes to pointers and vice versa.
package utils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// ToString takes a string pointer and returns the string.
// A nil-pointer will result in an empty string.
func ToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

// ToStringPtrOrNil takes a string and returns the pointer to it.
// An empty string will result in a nil-pointer.
func ToStringPtrOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

// ToStringPtr takes a string and returns the pointer to it.
func ToStringPtr(value string) *string {
	return &value
}

// ToStringArray takes a pointer to a string array and returns the array.
// A nil-pointer will result in nil.
func ToStringArray(value *[]string) []string {
	if value == nil {
		return nil
	}
	return *value
}

// ToStringArrayPtr takes a string array and returns the pointer to it.
// Nil will result in a nil-pointer.
func ToStringArrayPtr(arr []string) *[]string {
	if arr == nil {
		return nil
	}
	return &arr
}

// ToBool takes a bool pointer and returns a bool.
// A nil-pointer will result in false.
func ToBool(value *bool) bool {
	if value == nil {
		return false
	}
	return *value
}

// ToBoolPtr takes a bool and returns the pointer to it.
func ToBoolPtr(value bool) *bool {
	return &value
}

// ToInt takes an int pointer and returns the int.
// A nil-pointer will result in 0.
func ToInt(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}

// ToIntPtr takes an int and returns the pointer to it.
func ToIntPtr(value int) *int {
	return &value
}

// ToInt32 takes an int32 pointer and returns the int32.
// A nil-pointer will result in 0.
func ToInt32(value *int32) int32 {
	if value == nil {
		return 0
	}
	return *value
}

// ToInt32Ptr takes an int32 and returns the pointer to it.
func ToInt32Ptr(value int32) *int32 {
	return &value
}

// ToInt64 takes an int64 pointer and returns the int64.
// A nil-pointer will result in 0.
func ToInt64(value *int64) int64 {
	if value == nil {
		return 0
	}
	return *value
}

// ToInt64Ptr takes an int64 and returns the pointer to it.
func ToInt64Ptr(value int64) *int64 {
	return &value
}

// ToFloat32Ptr takes an float32 and returns the pointer to it.
func ToFloat32Ptr(value float32) *float32 {
	return &value
}

// ToFloat32 takes an float32 pointer and returns the float32.
// A nil-pointer will result in 0.0.
func ToFloat32(value *float32) float32 {
	if value == nil {
		return 0.0
	}
	return *value
}

// ToFloat64 takes an float64 pointer and returns the float64.
// A nil-pointer will result in 0.0.
func ToFloat64(value *float64) float64 {
	if value == nil {
		return 0.0
	}
	return *value
}

// ToFloat64Ptr takes an float64 and returns the pointer to it.
func ToFloat64Ptr(value float64) *float64 {
	return &value
}

// DeepCopy copies the in value by marshaling/unmarshaling the source value
func DeepCopy(in, out interface{}) error {

	bytes, err := json.Marshal(in)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal the deep copy input object")
	}

	_ = json.Unmarshal(bytes, out)

	return nil
}

// ScaleHalfUp scales the given number up, rounds it up/down to the nearest int and scales it back.
func ScaleHalfUp(value float64, scale int) float64 {

	var round float64
	pow := math.Pow(10, float64(scale))
	digit := pow * value
	round = math.Round(digit)
	return round / pow
}

// ScaleUp scales the given number up, rounds it up to the nearest int and scales it back.
func ScaleUp(value float64, scale int) float64 {

	var round float64
	pow := math.Pow(10, float64(scale))
	digit := pow * value
	round = math.Ceil(digit)
	return round / pow
}

// ScaleDown scales the given number up, rounds it down to the nearest int and scales it back.
func ScaleDown(value float64, scale int) float64 {

	var round float64
	pow := math.Pow(10, float64(scale))
	digit := pow * value
	round = math.Floor(digit)
	return round / pow
}

// IsNumber returns true if the given string is a number. Leading and trailing whitespaces will be ignored.
func IsNumber(value string) bool {
	if _, err := strconv.Atoi(strings.TrimSpace(value)); err != nil {
		return false
	}
	return true
}

// IsNil returns true if the given parameter is nil or a pointer to a nil value.
func IsNil(value interface{}) bool {
	return value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}
