# Slice

<!--<img alt="Slice logo" src="img/logo.png" width=192 align=right>-->

`slice` is a Go package that provides a collection of utility functions that are often needed but missing from the Go standard library. This package includes functions for manipulating slices, such as reversing, shuffling, finding unique elements, and more.

## Installation

To install the package, run:

```sh
go get github.com/xyproto/slice
```

## Usage

Here are some examples of how to use the functions provided by `slice`.

### Reverse

Creates and returns a new slice that is the reverse of the input slice.

```go
import "github.com/xyproto/slice"

fmt.Println(slice.Reverse([]int{1, 2, 3, 4, 5})) // Output: [5 4 3 2 1]
```

### ReverseInPlace

Reverses a slice in-place.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 4, 5}
slice.ReverseInPlace(arr)
fmt.Println(arr) // Output: [5 4 3 2 1]
```

### ReplaceFirst

Replaces the first occurrence of `needle` with `replacement` in a slice.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 1, 4}
fmt.Println(slice.ReplaceFirst(arr, 1, 9)) // Output: [9 2 3 1 4]
```

### ReplaceLast

Replaces the last occurrence of `needle` with `replacement` in a slice.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 1, 4}
fmt.Println(slice.ReplaceLast(arr, 1, 9)) // Output: [1 2 3 9 4]
```

### Partition

Partitions a slice into two slices based on a predicate function.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 4, 5}
even, odd := slice.Partition(arr, func(n int) bool { return n%2 == 0 })
fmt.Println(even, odd) // Output: [2 4] [1 3 5]
```

### Intersection

Returns the intersection of two slices.

```go
import "github.com/xyproto/slice"

a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}
fmt.Println(slice.Intersection(a, b)) // Output: [3 4]
```

### Difference

Returns the difference between two slices.

```go
import "github.com/xyproto/slice"

a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}
fmt.Println(slice.Difference(a, b)) // Output: [1 2]
```

### Flatten

Flattens a nested slice into a single-level slice.

```go
import "github.com/xyproto/slice"

nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
fmt.Println(slice.Flatten(nested)) // Output: [1 2 3 4 5 6]
```

### Unique

Removes duplicate elements from a slice.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 2, 3, 4, 4, 5}
fmt.Println(slice.Unique(arr)) // Output: [1 2 3 4 5]
```

### Shuffle

Randomly shuffles the elements of a slice.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 4, 5}
slice.Shuffle(arr)
fmt.Println(arr) // Output: [3 5 4 1 2] (example, actual output may vary)
```

### Zip

Combines multiple slices into a slice of tuples.

```go
import "github.com/xyproto/slice"

a := []int{1, 2, 3}
b := []string{"a", "b", "c"}
fmt.Println(slice.Zip(a, b)) // Output: [{1 a} {2 b} {3 c}]
```

### Contains

Checks if a slice contains a specific element.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 4, 5}
fmt.Println(slice.Contains(arr, 3)) // Output: true
```

### Filter

Filters a slice based on a predicate function.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 4, 5}
even := slice.Filter(arr, func(n int) bool { return n%2 == 0 })
fmt.Println(even) // Output: [2 4]
```

### Map

Applies a function to each element in a slice and returns a new slice.

```go
import "github.com/xyproto/slice"

strArr := []string{"hello", "world"}
upperStrArr := slice.Map(strArr, strings.ToUpper)
fmt.Println(upperStrArr) // Output: [HELLO WORLD]
```

### Reduce

Reduces a slice to a single value using a reduction function.

```go
import "github.com/xyproto/slice"

arr := []int{1, 2, 3, 4, 5}
sum := slice.Reduce(arr, 0, func(acc, n int) int { return acc + n })
fmt.Println(sum) // Output: 15
```

### Capitalize

Capitalizes the first letter of a slice of runes and returns a new slice.

```go
import "github.com/xyproto/slice"

fmt.Println(string(slice.Capitalize([]rune("hello world")))) // Output: Hello world
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the BSD-3 License - see the [LICENSE](LICENSE) file for details.

## Version

1.0.0
