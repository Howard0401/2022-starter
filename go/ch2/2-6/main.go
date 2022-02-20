package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// io read write reader writer
func main() {
	reader := strings.NewReader("Test String.")
	p := make([]byte, 8)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
	meats := []string{
		"chicken",
		",beef",
		",pork",
		",mutton",
	}
	var writer bytes.Buffer
	for _, p := range meats {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Printf("err=%v", err)
		}
		if n != len(p) {
			fmt.Printf("n != len(p), failed\n")
		}
	}

	fmt.Println(writer.String())
}
