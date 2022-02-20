package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

func Show() {
	fmt.Println("Echo Show func")
}

type CheckOut func(int, int) int

// props? We could measure business logic in anonymous function,
// without change original func GetTotal()
// GetTotal()
func GetTotal(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

type GenerateRadom func() int

func RandomSum() GenerateRadom {
	fmt.Println("RandomSum()")
	// https://stackoverflow.com/questions/8288679/difficulty-with-go-rand-package
	// https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())
	a, b := rand.Intn(10), rand.Intn(20)
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
	return func() int {
		a, b = b, a+b
		return a
	}
}

// overwrite Read func for GenerateRadom,
// so we could pass anonymous g = func() int { a, b = b, a+b return a }
func (g GenerateRadom) Read(p []byte) (n int, err error) {
	next := g()
	if next > 21 {
		fmt.Println(">21")
		fmt.Println(next)
		fmt.Println(">21 end...")
		return 0, io.EOF
	}
	s := fmt.Sprintf("Read %d\n", next)
	return strings.NewReader(s).Read(p)
}

func PrintResult(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("scanner.Text()=%v\n", scanner.Text())
	}
}

func main() {
	Show()
	show := Show
	show()
	show2 := func() {
		fmt.Println("Anonymous Function")
	}
	show2()

	var checkOut CheckOut = func(a, b int) int {
		return a + b
	}

	fmt.Printf("checkOut(68, 98)=%v\n", checkOut(68, 98))

	total := GetTotal(68)
	sum := total(100)
	fmt.Println(sum) // 168

	total = GetTotal(sum)
	sum = total(50) // 218

	total = GetTotal(sum)
	sum = total(200)

	fmt.Println(sum) // 418

	r := RandomSum()
	PrintResult(r)
}
