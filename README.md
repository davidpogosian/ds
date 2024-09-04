# Generic Concurrency-Safe Data Structures Library for Go

This library provides robust, thread-safe implementations of commonly used data structures, optimized for performance and ease of use.

## Features

- **Concurrency-Safe:** All data structures are designed to be safe for concurrent use, leveraging Go's synchronization primitives.
- **Generic Support:** Utilizes Go's generics for flexibility, allowing you to work with any data type.
- **Comprehensive Test Coverage:** Each data structure is thoroughly tested to ensure correctness and performance.

## What's in the Box?

- *Stack*
- *Queue*
- *List*
- *Set*
- *Binary Search Tree*
- *Priority Queue*

## Documentation

The full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/davidpogosian/ds)

## Installation

Install via `go get`:

```bash
go get github.com/davidpogosian/ds
```

## Example of using the Stack

```go
package main

import (
	"fmt"
	"log"

	"github.com/davidpogosian/ds/comparators"
	"github.com/davidpogosian/ds/stack"
)

func main() {
	s := stack.NewEmpty(comparators.ComparatorInt)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // Output: 3
}
```

## Licence

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
