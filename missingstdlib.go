package missingstdlib

import (
	"math/rand"
	"strings"
)

// Capitalize capitalizes the first letter of a string.
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// Reverse reverses a slice.
func Reverse[T any](s []T) []T {
	length := len(s)
	reversed := make([]T, length)
	for i, v := range s {
		reversed[length-1-i] = v
	}
	return reversed
}

// ReplaceFirst replaces the first occurrence of old with new in a string.
func ReplaceFirst[T ~string](s, old, new T) T {
	return T(strings.Replace(string(s), string(old), string(new), 1))
}

// Partition partitions a slice into two slices based on a predicate function.
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

// Shuffle randomly shuffles the elements of a slice.
func Shuffle[T any](arr []T) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

// Zip combines multiple slices into a slice of tuples.
func Zip[T any, U any](a []T, b []U) [][2]interface{} {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	zipped := make([][2]interface{}, length)
	for i := 0; i < length; i++ {
		zipped[i] = [2]interface{}{a[i], b[i]}
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
