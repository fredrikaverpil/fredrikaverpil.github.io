---
title: "Go's secret weapon: the standard library interfaces"
date: 2025-12-28
draft: false
tags: ["golang"]
categories: []
---

Go is peculiar in the sense that there's a very tight social contract between
the community and the [standard library](https://pkg.go.dev/std). Seemingly
universally, the community agrees to use these specific, small interfaces, as
glue between unrelated libraries. In fact, I think there's even an expectation
that implementations should, whenever possible, snap nicely together with
standard library interfaces, like LEGO bricks.

I think that at least one enabler to this community culture stems from Go's
implicit interface design and the distinct level of composability that comes out
of it. In languages with explicit implementation (like Java or C#), you have to
plan to implement an interface. In Go, you can implement `io.Writer` by accident
just by having a `Write([]byte) (int, error)` method.

> [!HINT] Implicit interfaces
>
> Because you don't have to import "io" to implement `io.Reader`, libraries
> don't become coupled. That specific lack of coupling is what makes the
> standard library interfaces so powerful in Go compared to similar interfaces
> in other languages. This behavior is technically referred to as
> [structural typing](https://en.wikipedia.org/wiki/Structural_type_system) and
> can be observed in other languages, like TypeScript. In contrast, Java, C#,
> Rust and Swift use (explicit)
> [nominal typing](https://en.wikipedia.org/wiki/Nominal_type_system) and for
> Python, Ruby and JavaScript you have
> [duck typing](https://en.wikipedia.org/wiki/Duck_typing).
>
> Go feels unique here because it is one of the few mainstream, compiled
> languages where structural typing is the default and primary way to model
> abstraction.

What I find especially appealing by all this is that by leveraging the standard
library interfaces you don't just solve your specific problem; you might
actually create tools that fit into the wider Go ecosystem, allowing others to
use your code in ways you might not have anticipated.

However, this implies you know about said interfaces. When I was completely new
to Go, I came across just a handful of such interfaces naturally. In this post
I've gathered some interfaces from the Go standard library that I think are nice
to know about, especially for newcomers to the language. I'm also focusing on
examples around _accepting_ said interfaces, as this is where I think
composability shines. Where helpful, implementation examples are also included
to show the other side of the coin.

> [!EXAMPLE] Example on "accepting interfaces"
>
> Consider a function that parses data. If you write it to accept a `string`
> filepath and open the file internally, users can e.g. only use it with files
> on disk. But if you write it to accept an `io.Reader`, your function can work
> with:
>
> - Files on disk
> - HTTP response bodies
> - In-memory string buffers
> - Network connections
> - Standard input (stdin)
> - Anything else that implements `io.Reader`

---

## Interfaces

### Core Built-ins

- **[`error`](https://pkg.go.dev/builtin#error)**: The most fundamental
  interface. Functions that accept `error` can work with any error type.
- **[`any`](https://pkg.go.dev/builtin#any)**: Alias for the empty interface
  `interface{}`. Represents a value of any type.
- **[`comparable`](https://pkg.go.dev/builtin#comparable)**: A built-in
  constraint for all types that can be compared with `==` and `!=`. Essential
  when writing generic functions. Note: this is a _constraint_, not a
  traditional interface, and can only be used with generics.

> [!EXAMPLE-] Accepting Error, Any, Comparable and Implementing Custom Errors
>
> **Code Examples**
>
> {{< code file="examples/builtin.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/builtin_test.go" >}}

### Formatting

- **[`fmt.Stringer`](https://pkg.go.dev/fmt#Stringer)**: Implement this to
  customize how your type is printed. Functions in the fmt package (like
  `Println` or `Printf`) accepts any type that implements `Stringer` when using
  `%s` or `%v`.
- **[`fmt.GoStringer`](https://pkg.go.dev/fmt#GoStringer)**: Implement this to
  customize the `%#v` output.

> [!EXAMPLE-] Accepting and Implementing fmt.Stringer and fmt.GoStringer
>
> **Code Examples**
>
> {{< code file="examples/formatting.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/formatting_test.go" >}}

### Testing

- **[`testing.TB`](https://pkg.go.dev/testing#TB)**: The common interface for
  `*testing.T` and `*testing.B`. By accepting `testing.TB`, your helper
  functions work in both tests and benchmarks.

> [!EXAMPLE-] Accepting testing.TB for Reusable Test Helpers
>
> **Code Examples**
>
> {{< code file="examples/testing.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/testing_test.go" >}}

### Context

- **[`context.Context`](https://pkg.go.dev/context#Context)**: Defines the
  lifecycle of requests, handling cancellation, deadlines, and request-scoped
  values.

> [!EXAMPLE-] Accepting context.Context for Cancellation and Timeout Control
>
> **Code Examples**
>
> {{< code file="examples/context.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/context_test.go" >}}

### Logging

- **[`slog.Handler`](https://pkg.go.dev/log/slog#Handler)**: The core interface
  for structured logging. By accepting a `slog.Handler`, you can wrap or
  decorate any logging backend, such as the built-in
  [`slog.TextHandler`](https://pkg.go.dev/log/slog#TextHandler) or
  [`slog.JSONHandler`](https://pkg.go.dev/log/slog#JSONHandler).
- **[`slog.LogValuer`](https://pkg.go.dev/log/slog#LogValuer)**: Implement this
  to control how your type appears in structured logs. The `slog` package
  accepts any type that implements this interface.

> [!EXAMPLE-] Accepting and Implementing slog Interfaces
>
> **Code Examples**
>
> {{< code file="examples/logging.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/logging_test.go" >}}

### Generics

- **[`cmp.Ordered`](https://pkg.go.dev/cmp#Ordered)**: A constraint that accepts
  any orderable type (`int`, `float64`, `string`, etc.). Use this in generic
  function signatures to enable comparison operators.

> [!EXAMPLE-] Accepting cmp.Ordered Constraint in Generic Functions
>
> **Code Examples**
>
> {{< code file="examples/generics.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/generics_test.go" >}}

### Iteration

The `iter` package introduces standard signatures for iterators. By accepting
these types, your functions can work with any iterable source: slices, maps,
channels, custom data structures, or generated sequences.

- **[`iter.Seq[V]`](https://pkg.go.dev/iter#Seq)**: The standard iterator for a
  sequence of values.
- **[`iter.Seq2[K, V]`](https://pkg.go.dev/iter#Seq2)**: The standard iterator
  for key-value pairs.

> [!EXAMPLE-] Accepting and Implementing iter.Seq Iterators
>
> **Code Examples**
>
> {{< code file="examples/iter.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/iter_test.go" >}}

### Command Line Flags

The `flag` package accepts any type implementing this interface, allowing custom
flag parsing.

- **[`flag.Value`](https://pkg.go.dev/flag#Value)**: Implement this to create
  custom flag types.

> [!EXAMPLE-] Implementing flag.Value for Custom Flag Parsing
>
> **Code Examples**
>
> {{< code file="examples/flags.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/flags_test.go" >}}

### Concurrency

- **[`sync.Locker`](https://pkg.go.dev/sync#Locker)**: Represents an object that
  can be locked and unlocked. By accepting `sync.Locker`, you allow callers to
  provide any lock implementation (`sync.Mutex`, `sync.RWMutex`, or a no-op for
  testing).

> [!EXAMPLE-] Accepting sync.Locker for Flexible Locking Strategies
>
> **Code Examples**
>
> {{< code file="examples/concurrency.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/concurrency_test.go" >}}

### Sorting & Containers

- **[`sort.Interface`](https://pkg.go.dev/sort#Interface)**: By accepting
  `sort.Interface`, functions like `sort.Sort` work with any sortable
  collection.
  - Note that you should probably use
    [`slices.Sort`](https://pkg.go.dev/slices#Sort) or
    [`slices.SortFunc`](https://pkg.go.dev/slices#SortFunc) for slices, as they
    don't require implementing an interface.
- **[`heap.Interface`](https://pkg.go.dev/container/heap#Interface)**: Extends
  `sort.Interface`. The `heap` package accepts this to provide heap operations.

> [!EXAMPLE-] Accepting and Implementing sort.Interface and heap.Interface
>
> **Code Examples**
>
> {{< code file="examples/sorting.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/sorting_test.go" >}}

### Encoding & Serialization

"Marshalling" (or serialization) is the process of converting a Go value (like a
struct) into a format suitable for storage or transmission (like JSON text or
binary data). Conversely, "unmarshalling" is parsing that data back into a Go
value.

These interfaces are primarily designed to be _implemented_ by your types to
customize this process. The standard library functions (like `json.Marshal`)
check if your type implements these interfaces before falling back to default
behaviors.

- **[`encoding.TextMarshaler`](https://pkg.go.dev/encoding#TextMarshaler)** /
  **[`encoding.TextUnmarshaler`](https://pkg.go.dev/encoding#TextUnmarshaler)**:
  Text serialization. Used by `json`, `xml`, `yaml`, and other encoders.
- **[`encoding.BinaryMarshaler`](https://pkg.go.dev/encoding#BinaryMarshaler)**
  /
  **[`encoding.BinaryUnmarshaler`](https://pkg.go.dev/encoding#BinaryUnmarshaler)**:
  Binary serialization. Used by `gob` and other binary encoders.
- **[`json.Marshaler`](https://pkg.go.dev/encoding/json#Marshaler)** /
  **[`json.Unmarshaler`](https://pkg.go.dev/encoding/json#Unmarshaler)**:
  JSON-specific serialization.
- **[`xml.Marshaler`](https://pkg.go.dev/encoding/xml#Marshaler)** /
  **[`xml.Unmarshaler`](https://pkg.go.dev/encoding/xml#Unmarshaler)**:
  XML-specific serialization.

> [!EXAMPLE-] Accepting and Implementing Encoding Interfaces
>
> **Code Examples**
>
> {{< code file="examples/encoding.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/encoding_test.go" >}}

### Input / Output

The `io` package is about **moving bytes**. It abstracts the concept of a stream
of data, allowing you to process inputs and outputs without caring about the
source (file, network, memory buffer). Because these interfaces are so
universal, the standard library provides powerful combinators like
[`io.TeeReader`](https://pkg.go.dev/io#TeeReader) (read and write
simultaneously, e.g., hash a file while uploading it) and
[`io.MultiWriter`](https://pkg.go.dev/io#MultiWriter) (write to multiple
destinations at once).

- **[`io.Reader`](https://pkg.go.dev/io#Reader)**: Reads a stream of bytes. Used
  for files, network connections, and more.
- **[`io.Writer`](https://pkg.go.dev/io#Writer)**: Writes a stream of bytes.
- **[`io.StringWriter`](https://pkg.go.dev/io#StringWriter)**: Writes strings
  efficiently. Used by `io.WriteString()`.
- **[`io.Closer`](https://pkg.go.dev/io#Closer)**: Closes a resource (file,
  connection). Often used with `defer`.
- **[`io.Seeker`](https://pkg.go.dev/io#Seeker)**: Moves the current offset in a
  stream.
- **[`io.ReadWriter`](https://pkg.go.dev/io#ReadWriter)**: Combines `Reader` and
  `Writer`.
- **[`io.ReadCloser`](https://pkg.go.dev/io#ReadCloser)**: Combines `Reader` and
  `Closer` (e.g., `http.Response.Body`).
- **[`io.WriteCloser`](https://pkg.go.dev/io#WriteCloser)**: Combines `Writer`
  and `Closer` (e.g., `gzip.Writer`).
- **[`io.ReadWriteCloser`](https://pkg.go.dev/io#ReadWriteCloser)**: Combines
  all three (e.g., `net.Conn`).
- **[`io.ReaderAt`](https://pkg.go.dev/io#ReaderAt)**: Reads from a specific
  offset without changing the underlying state.
- **[`io.WriterAt`](https://pkg.go.dev/io#WriterAt)**: Writes to a specific
  offset.
- **[`io.ReaderFrom`](https://pkg.go.dev/io#ReaderFrom)**: Reads data _from_ a
  generic reader (optimizing copy operations).
- **[`io.WriterTo`](https://pkg.go.dev/io#WriterTo)**: Writes data _to_ a
  generic writer (optimizing copy operations).

> [!EXAMPLE-] Accepting and Implementing io Interfaces
>
> **Code Examples**
>
> {{< code file="examples/io.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/io_test.go" >}}

### File Systems

While `io` deals with open streams, `io/fs` deals with the **structure** of
files (names, directories, hierarchy).

Crucially, `io/fs` interfaces are **read-only**. This constraint is intentional:
it allows the same code to safely navigate files stored on a local disk,
embedded inside the binary, wrapped in a ZIP archive, or hosted in the cloud,
without needing to handle complex (and perhaps impossible) write operations for
those different backends.

- **[`fs.FS`](https://pkg.go.dev/io/fs#FS)**: The base interface for a file
  system. Helper functions like `fs.ReadFile`, `fs.Glob`, and `fs.WalkDir`
  accept this interface.
- **[`fs.File`](https://pkg.go.dev/io/fs#File)**: Represents an open file.
- **[`fs.DirEntry`](https://pkg.go.dev/io/fs#DirEntry)**: An item in a directory
  (faster than `FileInfo` for listings).
- **[`fs.FileInfo`](https://pkg.go.dev/io/fs#FileInfo)**: Metadata about a file.

> [!EXAMPLE-] Accepting fs.FS for File System Abstraction
>
> **Code Examples**
>
> {{< code file="examples/filesystem.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/filesystem_test.go" >}}

### HTTP & Networking

- **[`http.Handler`](https://pkg.go.dev/net/http#Handler)**: The core interface
  for responding to HTTP requests (`ServeHTTP`).
- **[`http.ResponseWriter`](https://pkg.go.dev/net/http#ResponseWriter)**: Used
  to construct an HTTP response.
- **[`http.RoundTripper`](https://pkg.go.dev/net/http#RoundTripper)**:
  Represents a single HTTP transaction. Used for client middleware.
- **[`http.Flusher`](https://pkg.go.dev/net/http#Flusher)**: Flushes buffered
  data to the client. Essential for streaming responses (SSE, chunked transfer).
- **[`http.Hijacker`](https://pkg.go.dev/net/http#Hijacker)**: Takes over the
  underlying connection. Used for WebSockets and protocol upgrades.
- **[`net.Conn`](https://pkg.go.dev/net#Conn)**: A generic network connection.
- **[`net.Listener`](https://pkg.go.dev/net#Listener)**: Listens for incoming
  connections.
- **[`net.Addr`](https://pkg.go.dev/net#Addr)**: A network address.

> [!EXAMPLE-] Accepting and Implementing HTTP and Networking Interfaces
>
> **Code Examples**
>
> {{< code file="examples/http.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/http_test.go" >}}

### Database

Like the encoding interfaces, these are designed to be _implemented_ by your
types. The `database/sql` package accepts any type and checks if it implements
these interfaces when reading from or writing to the database.

- **[`sql.Scanner`](https://pkg.go.dev/database/sql#Scanner)**: Implement this
  to control how database values are read into your type.
- **[`driver.Valuer`](https://pkg.go.dev/database/sql/driver#Valuer)**:
  Implement this to control how your type is written to the database.

> [!EXAMPLE-] Accepting and Implementing Database Interfaces
>
> **Code Examples**
>
> {{< code file="examples/database.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/database_test.go" >}}

### Cryptography & Hashing

- **[`hash.Hash`](https://pkg.go.dev/hash#Hash)**: By accepting `hash.Hash`,
  your functions work with any hash algorithm (MD5, SHA256, SHA512, etc.).
- **[`crypto.Signer`](https://pkg.go.dev/crypto#Signer)**: Accept this to sign
  with any private key type (RSA, ECDSA, Ed25519).

> [!EXAMPLE-] Accepting hash.Hash and crypto.Signer
>
> **Code Examples**
>
> {{< code file="examples/crypto.go" >}}
>
> **Usage Examples**
>
> {{< code file="examples/crypto_test.go" >}}

---

## Workflow

This post only scratches the surface. The standard library is full of interfaces
waiting to be discovered, and the best way to find them is to actively explore
the [standard library documentation](https://pkg.go.dev/std). Don't stop there
though; third-party libraries often define useful interfaces too. When you're
about to define a new type or write a function signature, ask yourself: _is
there an existing interface that fits here?_

You can quickly look up interfaces from your terminal using `go doc` (e.g.,
`go doc io.Reader`). Leverage [stdsym](https://github.com/lotusirous/gostdsym)
to search third-party packages. If you're a Neovim user, the
[godoc.nvim](https://github.com/fredrikaverpil/godoc.nvim) plugin lets you
search, preview and inspect documentation directly in your editor.

A really great Neovim plugin is
[goplements.nvim](https://github.com/maxandron/goplements.nvim), which makes
"implemented by" and "implements" information pop up inline, as you satisfy an
interface or implement one. Super helpful!

## Conclusion

The pattern is simple:
[accept interfaces, return concrete types](https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8).
Returning concrete types (like structs or custom types) allows you to add new
methods to them later without breaking backward compatibility, making it easier
to extend functionality. When your functions accept interfaces like `io.Reader`,
`context.Context`, or `hash.Hash`, they become building blocks that others can
combine in ways you never anticipated. When your types implement interfaces like
`error`, `fmt.Stringer`, or `json.Marshaler`, they plug seamlessly into the Go
ecosystem.

> [!NOTE]
>
> This post is about leveraging _existing_ standard library interfaces. You
> generally shouldn't
> [create your own interfaces preemptively](https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a).
> Go's [implicit interfaces](https://go.dev/tour/methods/10) mean the caller can
> define an interface when they actually need one.

This is Go's secret weapon for composability. The standard library provides the
shared vocabulary; you just need to speak it.

## Further Reading

If you want to dive deeper into Go's interface philosophy and design:

- [Go Proverbs](https://go-proverbs.github.io/) - Rob Pike's wisdom, including
  "The bigger the interface, the weaker the abstraction"
- [Effective Go: Interfaces](https://go.dev/doc/effective_go#interfaces) - The
  official guide on interface conventions
- [Interface Segregation Principle](https://en.wikipedia.org/wiki/Interface_segregation_principle) -
  The SOLID principle that Go's standard library exemplifies
- [Robustness Principle](https://en.wikipedia.org/wiki/Robustness_principle) -
  Also known as Postel's Law: "Be conservative in what you send, and liberal in
  what you accept"
- [Errors are values](https://go.dev/blog/errors-are-values) by Rob Pike -
  Creative patterns with the `error` interface
- [The Laws of Reflection](https://go.dev/blog/laws-of-reflection) -
  Understanding how interfaces work at runtime
