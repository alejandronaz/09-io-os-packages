package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	// generate a reader from a string
	text := "lorem ipsum dolor sit amet"
	reader := strings.NewReader(text)

	// read by 8 bytes per time (iteration)
	buffer := make([]byte, 8)
	for {
		n, err := reader.Read(buffer) // Read puts data into buffer
		if err == io.EOF {
			break
		}
		// conver buffer (slice of bytes) to string and print it
		fmt.Print(string(buffer[:n]), " - ")
	}

}
