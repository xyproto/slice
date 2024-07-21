# missingstdlib

`missingstdlib` is a Go package that provides a collection of utility functions that are often needed but missing from the Go standard library. This package includes functions for manipulating strings and slices, such as reversing, shuffling, finding unique elements, and more.

## Installation

To install the package, run:

```sh
go get github.com/xyproto/missingstdlib
```

## Usage

Here are some examples of how to use the functions provided by `missingstdlib`.

### Capitalize

Capitalizes the first letter of a string.

```go
import ms "github.com/xyproto/missingstdlib"

fmt.Println(ms.Capitalize("hello world")) // Output: Hello world
```

### Reverse

Creates and returns a new slice that is the reverse of the input slice.

```go
import ms "github.com/xyproto/missingstdlib"

fmt.Println(ms.Reverse([]int{1, 2, 3, 4, 5})) // Output: [5 4 3 2 1]
```

### ReverseInPlace

Reverses a slice in-place.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 4, 5}
ms.ReverseInPlace(arr)
fmt.Println(arr) // Output: [5 4 3 2 1]
```

### ReplaceFirst

Replaces the first occurrence of `needle` with `replacement` in `haystack`.

```go
import ms "github.com/xyproto/missingstdlib"

fmt.Println(ms.ReplaceFirst("hello world", "l", "x")) // Output: hexlo world
```

### ReplaceLast

Replaces the last occurrence of `needle` with `replacement` in `haystack`.

```go
import ms "github.com/xyproto/missingstdlib"

fmt.Println(ms.ReplaceLast("hello world", "l", "x")) // Output: hello worxd
```

### ReplaceFirstInSlice

Replaces the first occurrence of `needle` with `replacement` in a slice.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 1, 4}
fmt.Println(ms.ReplaceFirstInSlice(arr, 1, 9)) // Output: [9 2 3 1 4]
```

### ReplaceLastInSlice

Replaces the last occurrence of `needle` with `replacement` in a slice.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 1, 4}
fmt.Println(ms.ReplaceLastInSlice(arr, 1, 9)) // Output: [1 2 3 9 4]
```

### Partition

Partitions a slice into two slices based on a predicate function.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 4, 5}
even, odd := ms.Partition(arr, func(n int) bool { return n%2 == 0 })
fmt.Println(even, odd) // Output: [2 4] [1 3 5]
```

### Intersection

Returns the intersection of two slices.

```go
import ms "github.com/xyproto/missingstdlib"

a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}
fmt.Println(ms.Intersection(a, b)) // Output: [3 4]
```

### Difference

Returns the difference between two slices.

```go
import ms "github.com/xyproto/missingstdlib"

a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}
fmt.Println(ms.Difference(a, b)) // Output: [1 2]
```

### Flatten

Flattens a nested slice into a single-level slice.

```go
import ms "github.com/xyproto/missingstdlib"

nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
fmt.Println(ms.Flatten(nested)) // Output: [1 2 3 4 5 6]
```

### Unique

Removes duplicate elements from a slice.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 2, 3, 4, 4, 5}
fmt.Println(ms.Unique(arr)) // Output: [1 2 3 4 5]
```

### Shuffle

Randomly shuffles the elements of a slice.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 4, 5}
ms.Shuffle(arr)
fmt.Println(arr) // Output: [3 5 4 1 2] (example, actual output may vary)
```

### Zip

Combines multiple slices into a slice of tuples.

```go
import ms "github.com/xyproto/missingstdlib"

a := []int{1, 2, 3}
b := []string{"a", "b", "c"}
fmt.Println(ms.Zip(a, b)) // Output: [{1 a} {2 b} {3 c}]
```

### Contains

Checks if a slice contains a specific element.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 4, 5}
fmt.Println(ms.Contains(arr, 3)) // Output: true
```

### Filter

Filters a slice based on a predicate function.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 4, 5}
even := ms.Filter(arr, func(n int) bool { return n%2 == 0 })
fmt.Println(even) // Output: [2 4]
```

### Map

Applies a function to each element in a slice and returns a new slice.

```go
import ms "github.com/xyproto/missingstdlib"

strArr := []string{"hello", "world"}
upperStrArr := ms.Map(strArr, strings.ToUpper)
fmt.Println(upperStrArr) // Output: [HELLO WORLD]
```

### Reduce

Reduces a slice to a single value using a reduction function.

```go
import ms "github.com/xyproto/missingstdlib"

arr := []int{1, 2, 3, 4, 5}
sum := ms.Reduce(arr, 0, func(acc, n int) int { return acc + n })
fmt.Println(sum) // Output: 15
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the BSD-3 License - see the [LICENSE](LICENSE) file for details.

## Version

1.0.0
