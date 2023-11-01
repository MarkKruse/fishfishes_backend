package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToEmptyStringArray(t *testing.T) {

	t.Parallel()

	type test struct {
		given    []string
		expected []string
	}

	cases := map[string]test{
		"non empty list": {
			given:    []string{"foo", "bar"},
			expected: []string{"foo", "bar"},
		},
		"empty list": {
			given:    []string{},
			expected: []string{},
		},
		"nil list": {
			given:    nil,
			expected: []string{},
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := ToEmptyStringArray(tc.given)
			assert.EqualValues(t, tc.expected, result)
		})
	}
}

func TestDeleteEmpty(t *testing.T) {

	t.Parallel()

	type test struct {
		input    []string
		expected []string
	}

	cases := map[string]test{
		"nil array": {
			input:    nil,
			expected: nil,
		},
		"empty array": {
			input:    []string{},
			expected: nil,
		},
		"array with one empty element": {
			input:    []string{""},
			expected: nil,
		},
		"array with elements and one empty element": {
			input:    []string{"a", "", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := DeleteEmpty(tc.input)
			assert.EqualValues(t, tc.expected, result)
		})
	}
}

func TestTrimStringArray(t *testing.T) {

	t.Parallel()

	type test struct {
		input    []string
		expected []string
	}

	cases := map[string]test{
		"nil array": {
			input:    nil,
			expected: nil,
		},
		"empty array": {
			input:    []string{},
			expected: nil,
		},
		"array with one white space element": {
			input:    []string{" "},
			expected: []string{""},
		},
		"array with elements": {
			input:    []string{" a ", "\t\r\n", " b\t", "\tc\r\n"},
			expected: []string{"a", "", "b", "c"},
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := TrimStringArray(tc.input)
			assert.EqualValues(t, tc.expected, result)
		})
	}
}

func TestJoinWithSlash(t *testing.T) {

	t.Parallel()

	type test struct {
		given    []string
		expected string
	}

	cases := map[string]test{
		"happy path": {
			given: []string{
				"foo",
				"bar",
			},
			expected: "foo/bar",
		},
		"with slashes": {
			given: []string{
				"/foo/",
				"/bar/",
			},
			expected: "foo/bar",
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := JoinWithSlash(tc.given...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestAtLeastOneStringNotEmpty(t *testing.T) {
	t.Parallel()

	type test struct {
		input    []string
		expected bool
	}

	cases := map[string]test{
		"no strings to check": {
			input:    []string{},
			expected: false,
		},
		"all strings empty": {
			input:    []string{"", "  ", " "},
			expected: false,
		},
		"one string not empty": {
			input:    []string{"", "  ", " a"},
			expected: true,
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := AtLeastOneStringNotEmpty(tc.input...)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestAtLeastOneStringPtrNotEmpty(t *testing.T) {
	t.Parallel()

	type test struct {
		input    []*string
		expected bool
	}

	cases := map[string]test{
		"no strings to check": {
			input:    []*string{},
			expected: false,
		},
		"all strings empty": {
			input:    []*string{ToStringPtr(""), ToStringPtr("  "), ToStringPtr(" ")},
			expected: false,
		},
		"one string not empty": {
			input:    []*string{ToStringPtr(""), ToStringPtr("  "), ToStringPtr(" a")},
			expected: true,
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := AtLeastOneStringPtrNotEmpty(tc.input...)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestAllStringPtrGiven(t *testing.T) {
	t.Parallel()

	type test struct {
		input    []*string
		expected bool
	}

	cases := map[string]test{
		"no string ptrs to check": {
			input:    nil,
			expected: true, // Edge case - if no string was passed, function is used in the wrong way, so it does not matter what is returned.
		},
		"one string ptr is nil": {
			input:    []*string{ToStringPtr("Value"), nil},
			expected: false,
		},
		"one string ptr contains only whitespaces": {
			input:    []*string{ToStringPtr("Value"), ToStringPtr("   ")},
			expected: false,
		},
		"all string ptrs have values": {
			input:    []*string{ToStringPtr("Value"), ToStringPtr("another Value")},
			expected: true,
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := AllStringPtrGiven(tc.input...)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestGetFirstNotEmpty(t *testing.T) {
	t.Parallel()

	type test struct {
		input    []*string
		expected *string
	}

	cases := map[string]test{
		"nothing to map": {
			input:    nil,
			expected: nil,
		},
		"all strings empty": {
			input:    []*string{nil, ToStringPtr("   "), ToStringPtr("")},
			expected: nil,
		},
		"first string given": {
			input:    []*string{ToStringPtr("given"), ToStringPtr("   "), nil},
			expected: ToStringPtr("given"),
		},
		"middle string given": {
			input:    []*string{ToStringPtr("   "), ToStringPtr("given"), nil},
			expected: ToStringPtr("given"),
		},
	}

	for name, tc := range cases {

		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := GetFirstNotEmptyStringPtr(tc.input...)
			assert.EqualValues(t, tc.expected, got)
		})
	}
}

func ExampleJoinWithSlash_withoutSlashesInInput() {
	fmt.Println(JoinWithSlash("a", "b", "c"))
	// Output: a/b/c
}

func ExampleJoinWithSlash_withSlashesInInput() {
	fmt.Println(JoinWithSlash("a/", "/b", "/c/"))
	// Output: a/b/c
}
