package examples

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ExampleProcessLogs demonstrates accepting io.Reader and io.Writer
func ExampleProcessLogs() {
	input := strings.NewReader("INFO: ok\nERROR: fail\n")
	var output bytes.Buffer
	ProcessLogs(&output, input)
	fmt.Print(output.String())
	// Output: ERROR: fail
}

// ExampleWriteLabel demonstrates accepting io.StringWriter
func ExampleWriteLabel() {
	var b strings.Builder
	WriteLabel(&b, "INFO")
	fmt.Println(b.String())
	// Output: [INFO]
}

// ExampleStreamSize demonstrates accepting io.Seeker
func ExampleStreamSize() {
	data := bytes.NewReader([]byte("hello world"))
	size, _ := StreamSize(data)
	fmt.Println("Stream size:", size)
	// Output: Stream size: 11
}

// ExampleReadChunk demonstrates accepting io.ReaderAt
func ExampleReadChunk() {
	data := bytes.NewReader([]byte("hello world"))
	chunk, _ := ReadChunk(data, 0, 5)
	fmt.Println(string(chunk))
	// Output: hello
}

// ExampleTransfer demonstrates accepting io.WriterTo
func ExampleTransfer() {
	var src bytes.Buffer
	src.WriteString("payload")
	var dst bytes.Buffer
	Transfer(&dst, &src)
	fmt.Println(dst.String())
	// Output: payload
}

// ExampleAReader demonstrates implementing io.Reader
func ExampleAReader() {
	data, _ := io.ReadAll(io.LimitReader(AReader{}, 5))
	fmt.Println(string(data))
	// Output: AAAAA
}
