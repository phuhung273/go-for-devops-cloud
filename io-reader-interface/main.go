package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	Content string
	pos     int
}

// Read implements io.Reader.
func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos < len(m.Content) {
		copy(p, m.Content[m.pos:m.pos+1])
		m.pos++
		return 1, nil
	}
	return 0, io.EOF
}

func main() {

	mySlowReaderInstance := &MySlowReader{
		Content: "hello go interface",
	}

	data, err := io.ReadAll(mySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data: %s", data)
}
