package missingstdlib

import (
	"strings"
	"testing"
)

func TestCapitalize(t *testing.T) {
	result := Capitalize("hello world")
	expected := "Hello world"
	if result != expected {
		t.Errorf("Capitalize() = %v; want %v", result, expected)
	}
}

func TestReverse(t *testing.T) {
	result := Reverse([]rune("Hello"))
	expected := []rune("olleH")
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Reverse() = %v; want %v", string(result), string(expected))
			break
		}
	}

	intResult := Reverse([]int{1, 2, 3, 4, 5})
	intExpected := []int{5, 4, 3, 2, 1}
	for i, v := range intResult {
		if v != intExpected[i] {
			t.Errorf("Reverse() = %v; want %v", intResult, intExpected)
			break
		}
	}
}

func TestReplaceFirst(t *testing.T) {
	// Test with strings
	result := ReplaceFirst("hello world", "l", "x")
	expected := "hexlo world"
	if result != expected {
		t.Errorf("ReplaceFirst() = %v; want %v", result, expected)
	}

	// Test with slices of strings
	resultSlice := ReplaceFirstInSlice([]string{"hello", "world", "hello"}, "hello", "hi")
	expectedSlice := []string{"hi", "world", "hello"}
	if !slicesEqual(resultSlice, expectedSlice) {
		t.Errorf("ReplaceFirstInSlice() = %v; want %v", resultSlice, expectedSlice)
	}

	// Test with slices of ints
	resultInt := ReplaceFirstInSlice([]int{1, 2, 3, 1, 4}, 1, 9)
	expectedInt := []int{9, 2, 3, 1, 4}
	if !slicesEqual(resultInt, expectedInt) {
		t.Errorf("ReplaceFirstInSlice() = %v; want %v", resultInt, expectedInt)
	}
}

func TestReplaceLast(t *testing.T) {
	// Test with strings
	result := ReplaceLast("hello world", "l", "x")
	expected := "hello worxd"
	if result != expected {
		t.Errorf("ReplaceLast() = %v; want %v", result, expected)
	}

	result = ReplaceLast("hello world world", "world", "planet")
	expected = "hello world planet"
	if result != expected {
		t.Errorf("ReplaceLast() = %v; want %v", result, expected)
	}

	result = ReplaceLast("hello", "z", "x")
	expected = "hello"
	if result != expected {
		t.Errorf("ReplaceLast() = %v; want %v", result, expected)
	}

	// Test with slices of strings
	resultSlice := ReplaceLastInSlice([]string{"hello", "world", "hello"}, "hello", "hi")
	expectedSlice := []string{"hello", "world", "hi"}
	if !slicesEqual(resultSlice, expectedSlice) {
		t.Errorf("ReplaceLastInSlice() = %v; want %v", resultSlice, expectedSlice)
	}

	// Test with slices of ints
	resultInt := ReplaceLastInSlice([]int{1, 2, 3, 1, 4}, 1, 9)
	expectedInt := []int{1, 2, 3, 9, 4}
	if !slicesEqual(resultInt, expectedInt) {
		t.Errorf("ReplaceLastInSlice() = %v; want %v", resultInt, expectedInt)
	}
}

func TestPartition(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	even, odd := Partition(arr, func(n int) bool { return n%2 == 0 })
	expectedEven := []int{2, 4}
	expectedOdd := []int{1, 3, 5}
	if !slicesEqual(even, expectedEven) || !slicesEqual(odd, expectedOdd) {
		t.Errorf("Partition() = %v, %v; want %v, %v", even, odd, expectedEven, expectedOdd)
	}
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	result := Intersection(a, b)
	expected := []int{3, 4}
	if !slicesEqual(result, expected) {
		t.Errorf("Intersection() = %v; want %v", result, expected)
	}
}

func TestDifference(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	result := Difference(a, b)
	expected := []int{1, 2}
	if !slicesEqual(result, expected) {
		t.Errorf("Difference() = %v; want %v", result, expected)
	}
}

func TestFlatten(t *testing.T) {
	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
	result := Flatten(nested)
	expected := []int{1, 2, 3, 4, 5, 6}
	if !slicesEqual(result, expected) {
		t.Errorf("Flatten() = %v; want %v", result, expected)
	}
}

func TestUnique(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	result := Unique(arr)
	expected := []int{1, 2, 3, 4, 5}
	if !slicesEqual(result, expected) {
		t.Errorf("Unique() = %v; want %v", result, expected)
	}
}

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Shuffle(arr)
	// Check if the result is a permutation of the input
	expected := []int{1, 2, 3, 4, 5}
	if !isPermutation(arr, expected) {
		t.Errorf("Shuffle() = %v; not a permutation of %v", arr, expected)
	}
}

func TestZip(t *testing.T) {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}
	result := Zip(a, b)
	expected := []struct{ First int; Second string }{{1, "a"}, {2, "b"}, {3, "c"}}
	if !tuplesEqual(result, expected) {
		t.Errorf("Zip() = %v; want %v", result, expected)
	}
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Contains(arr, 3) {
		t.Errorf("Contains() = false; want true")
	}
	if Contains(arr, 6) {
		t.Errorf("Contains() = true; want false")
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := Filter(arr, func(n int) bool { return n%2 == 0 })
	expected := []int{2, 4}
	if !slicesEqual(result, expected) {
		t.Errorf("Filter() = %v; want %v", result, expected)
	}
}

func TestMap(t *testing.T) {
	strArr := []string{"hello", "world"}
	result := Map(strArr, strings.ToUpper)
	expected := []string{"HELLO", "WORLD"}
	if !slicesEqual(result, expected) {
		t.Errorf("Map() = %v; want %v", result, expected)
	}
}

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	sum := Reduce(arr, 0, func(acc, n int) int { return acc + n })
	expected := 15
	if sum != expected {
		t.Errorf("Reduce() = %v; want %v", sum, expected)
	}
}

// Helper functions for tests
func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isPermutation[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[T]int)
	for _, v := range a {
		counts[v]++
	}
	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}

func tuplesEqual(a, b []struct{ First int; Second string }) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].First != b[i].First || a[i].Second != b[i].Second {
			return false
		}
	}
	return true
}
