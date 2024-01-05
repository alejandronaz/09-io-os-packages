package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Open a file as reader mode
	// file es *os.File, e implementa io.Reader (por eso se puede usar como reader)
	file, err := os.Open("customers.txt")

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	// // read by 8 bytes per time (iteration)
	buffer := make([]byte, 8)
	for {
		n, err := file.Read(buffer) // Read puts data into buffer
		if err == io.EOF {
			break
		}
		// conver buffer (slice of bytes) to string and print it
		fmt.Print(string(buffer[:n]), " - ")
	}

}
