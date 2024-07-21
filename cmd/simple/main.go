package main

import (
	"fmt"
	"strings"

	ms "github.com/xyproto/missingstdlib"
)

func main() {
	// Capitalize
	fmt.Println(ms.Capitalize("hello world")) // Output: Hello world

	// Reverse
	fmt.Println(ms.Reverse([]rune("Hello")))      // Output: [o l l e H]
	fmt.Println(ms.Reverse([]int{1, 2, 3, 4, 5})) // Output: [5 4 3 2 1]
	fmt.Println(ms.Reverse([]float64{1.1, 2.2, 3.3, 4.4, 5.5})) // Output: [5.5 4.4 3.3 2.2 1.1]
	fmt.Println(ms.Reverse([]string{"apple", "banana", "cherry"})) // Output: [cherry banana apple]

	// ReplaceFirst
	fmt.Println(ms.ReplaceFirst("hello world", "l", "x")) // Output: hexlo world

	// ReplaceLast
	fmt.Println(ms.ReplaceLast("hello world", "l", "x")) // Output: hello worxd

	// Partition
	arr := []int{1, 2, 3, 4, 5}
	even, odd := ms.Partition(arr, func(n int) bool { return n%2 == 0 })
	fmt.Println(even, odd) // Output: [2 4] [1 3 5]

	// Intersection
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	fmt.Println(ms.Intersection(a, b)) // Output: [3 4]

	// Difference
	fmt.Println(ms.Difference(a, b)) // Output: [1 2]

	// Flatten
	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(ms.Flatten(nested)) // Output: [1 2 3 4 5 6]
	
	nestedFloat := [][]float64{{1.1, 2.2}, {3.3, 4.4}, {5.5, 6.6}}
	fmt.Println(ms.Flatten(nestedFloat)) // Output: [1.1 2.2 3.3 4.4 5.5 6.6]

	// Unique
	arr = []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(ms.Unique(arr)) // Output: [1 2 3 4 5]
	
	strArr := []string{"apple", "banana", "apple", "cherry"}
	fmt.Println(ms.Unique(strArr)) // Output: [apple banana cherry]

	// Shuffle
	arr = []int{1, 2, 3, 4, 5}
	ms.Shuffle(arr)
	fmt.Println(arr) // Output: [3 5 4 1 2] (example, actual output may vary)

	// Zip int
	a = []int{1, 2, 3}
	b = []int{4, 5, 6}
	fmt.Println(ms.Zip(a, b)) // Output: [{1 4} {2 5} {3 6}]

	// Zip int and string
	a = []int{1, 2, 3}
	s := []string{"a", "b", "c"}
	fmt.Println(ms.Zip(a, s)) // Output: [{1 a} {2 b} {3 c}]
	
	// Contains
	fmt.Println(ms.Contains(arr, 3)) // Output: true

	// Filter
	arr = []int{1, 2, 3, 4, 5}
	even = ms.Filter(arr, func(n int) bool { return n%2 == 0 })
	fmt.Println(even) // Output: [2 4]

	// Map
	strArr = []string{"hello", "world"}
	upperStrArr := ms.Map(strArr, strings.ToUpper)
	fmt.Println(upperStrArr) // Output: [HELLO WORLD]

	// Map float64
	floatArr := []float64{1.1, 2.2, 3.3}
	intArr := ms.Map(floatArr, func(f float64) int { return int(f) })
	fmt.Println(intArr) // Output: [1 2 3]

	// Reduce
	sum := ms.Reduce(arr, 0, func(acc, n int) int { return acc + n })
	fmt.Println(sum) // Output: 15

	// Reduce float64
	floatSum := ms.Reduce(floatArr, 0.0, func(acc, n float64) float64 { return acc + n })
	fmt.Println(floatSum) // Output: 6.6
}
