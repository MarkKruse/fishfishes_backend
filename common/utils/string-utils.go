package utils

import "strings"

// ToEmptyStringArray checks if a string array is nil. If so, it will return an empty array.
// Otherwise, the function returns the input array.
func ToEmptyStringArray(arr []string) []string {
	if arr == nil {
		return make([]string, 0)
	}
	return arr
}

// DeleteEmpty removes all empty strings within the passed string array and returns an array with the remaining strings.
func DeleteEmpty(arr []string) []string {
	var result []string
	for _, str := range arr {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// TrimStringArray removes leading and trailing whitespaces from each array element and returns the string array.
func TrimStringArray(arr []string) []string {
	var result []string
	for _, str := range arr {
		result = append(result, strings.TrimSpace(str))
	}
	return result
}

// JoinWithSlash returns a string consisting of all parameters joined with slashes.
// For parameters which already start or end with a slash, it will be removed first.
func JoinWithSlash(args ...string) string {

	for i := range args {
		args[i] = strings.TrimSuffix(strings.TrimPrefix(args[i], "/"), "/")
	}

	return strings.Join(args, "/")
}

// AtLeastOneStringNotEmpty returns true if at least one of the given strings is not empty after whitespaces where removed.
func AtLeastOneStringNotEmpty(elements ...string) bool {
	for _, elem := range elements {
		if len(strings.TrimSpace(elem)) != 0 {
			return true
		}
	}
	return false
}

// AtLeastOneStringPtrNotEmpty is equivalent to AtLeastOneStringNotEmpty, but expected string pointers.
// It returns true if at least one of the string pointers points to a non-empty value (whitespaces count as empty).
func AtLeastOneStringPtrNotEmpty(elements ...*string) bool {
	for _, elem := range elements {
		if len(strings.TrimSpace(ToString(elem))) != 0 {
			return true
		}
	}
	return false
}

// AllStringPtrGiven returns true if each of the given string pointers points to a non-empty value (whitespaces count as empty).
func AllStringPtrGiven(elements ...*string) bool {
	for _, elem := range elements {
		if len(strings.TrimSpace(ToString(elem))) == 0 {
			return false
		}
	}
	return true
}

// GetFirstNotEmptyStringPtr returns the first string pointer which points to a value that is not empty (whitespaces count as empty).
func GetFirstNotEmptyStringPtr(elements ...*string) *string {
	for _, elem := range elements {
		if len(strings.TrimSpace(ToString(elem))) != 0 {
			return elem
		}
	}
	return nil
}
