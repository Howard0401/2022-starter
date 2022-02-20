package main

import (
	"bufio"
	"fmt"
	"os"
)

func Cook() {
	defer fmt.Println("3")
	fmt.Println("1")
	defer fmt.Println("2")
	panic("panic")
}

func WriteMenu(fileName string, foods []string) {
	curDir, _ := os.Getwd()
	path := curDir + fileName
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, item := range foods {
		fmt.Fprintln(w, item)
	}
}

func main() {
	// Cook()
	s := []string{"str1", "str2", "str3", "str4"}
	WriteMenu("/output", s)
}
