---
title: 'The "new" function changes in Go 1.26'
date: 2025-12-26
draft: false
tags: ["golang"]
categories: []
---

Here is how you can create a pointer to a boolean in
[Go 1.26](https://go.dev/doc/go1.26):

```go
// New way in Go 1.26
b := new(true)
```

This is equivalent to the older, more verbose method:

```go
// Old way
val := true
b := &val
```

It works for all native types because the operand can be any expression. In
addition to booleans, you can use it for:

- Strings: `s := new("Hello, Go 1.26")`
- Integers: `i := new(42)`
- Floating-point numbers: `f := new(3.14159)`
- Complex numbers: `c := new(1 + 2i)`
- Composite types:
  `l := new([]int{1, 2, 3}) or m := new(map[string]int{"a": 1})`
- Function results: `p := new(time.Now())`

Essentially, if you can write an expression that returns a value of type T, you
can now pass that expression to new() to get a `*T` initialized with that value.

> [!TIP]
>
> You can try it out today using
> [`gotip`](https://pkg.go.dev/golang.org/dl/gotip):
>
> ```bash
> go install golang.org/dl/gotip@latest
> gotip download
> ```
