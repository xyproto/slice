package slice

import (
	"math/rand"
	"unicode"
)

// Capitalize returns a new slice of runes with the first letter capitalized.
func Capitalize(runes []rune) []rune {
	if len(runes) == 0 {
		return runes
	}
	newRunes := make([]rune, len(runes))
	copy(newRunes, runes)
	newRunes[0] = unicode.ToUpper(newRunes[0])
	return newRunes
}

// Reverse returns a new slice that is the reverse of the input slice.
func Reverse[T any](s []T) []T {
	length := len(s)
	reversed := make([]T, length)
	for i, v := range s {
		reversed[length-1-i] = v
	}
	return reversed
}

// ReverseInPlace reverses the input slice in place.
func ReverseInPlace[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ReplaceFirst replaces the first occurrence of needle with replacement in a slice.
func ReplaceFirst[T comparable](haystack []T, needle, replacement T) []T {
	for i, v := range haystack {
		if v == needle {
			haystack[i] = replacement
			break
		}
	}
	return haystack
}

// ReplaceLast replaces the last occurrence of needle with replacement in a slice.
func ReplaceLast[T comparable](haystack []T, needle, replacement T) []T {
	for i := len(haystack) - 1; i >= 0; i-- {
		if haystack[i] == needle {
			haystack[i] = replacement
			break
		}
	}
	return haystack
}

// Partition partitions the input slice into two slices based on the predicate function.
func Partition[T any](arr []T, pred func(T) bool) ([]T, []T) {
	truePart, falsePart := []T{}, []T{}
	for _, v := range arr {
		if pred(v) {
			truePart = append(truePart, v)
		} else {
			falsePart = append(falsePart, v)
		}
	}
	return truePart, falsePart
}

// Intersection returns the intersection of two slices.
func Intersection[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range a {
		m[item] = true
	}
	var result []T
	for _, item := range b {
		if m[item] {
			result = append(result, item)
		}
	}
	return result
}

// Difference returns the difference between two slices.
func Difference[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range b {
		m[item] = true
	}
	var result []T
	for _, item := range a {
		if !m[item] {
			result = append(result, item)
		}
	}
	return result
}

// Flatten flattens a nested slice into a single-level slice.
func Flatten[T any](nested [][]T) []T {
	var flat []T
	for _, inner := range nested {
		flat = append(flat, inner...)
	}
	return flat
}

// Unique removes duplicate elements from a slice.
func Unique[T comparable](arr []T) []T {
	m := make(map[T]bool)
	var result []T
	for _, item := range arr {
		if !m[item] {
			m[item] = true
			result = append(result, item)
		}
	}
	return result
}

// Shuffle randomly shuffles the elements of a slice in place.
func Shuffle[T any](arr []T) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

// Zip combines two slices into a slice of tuples.
func Zip[T any, U any](a []T, b []U) []struct {
	First  T
	Second U
} {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	zipped := make([]struct {
		First  T
		Second U
	}, length)
	for i := 0; i < length; i++ {
		zipped[i] = struct {
			First  T
			Second U
		}{a[i], b[i]}
	}
	return zipped
}

// Contains checks if a slice contains a specific element.
func Contains[T comparable](arr []T, elem T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

// Filter filters a slice based on a predicate function.
func Filter[T any](arr []T, pred func(T) bool) []T {
	var result []T
	for _, v := range arr {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map applies a function to each element in a slice and returns a new slice.
func Map[T any, U any](arr []T, fn func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// Reduce reduces a slice to a single value using a reduction function.
func Reduce[T any, U any](arr []T, init U, fn func(U, T) U) U {
	result := init
	for _, v := range arr {
		result = fn(result, v)
	}
	return result
}
