package main

import (
	"fmt"
)

func funcRecover() error {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("panic recover, v=%v\n", v)
		}
	}()
	return funcCook()
}

func funcCook() error {
	panic("funcCook panic")
	// return errors.New("funcCook return err")
}

func main() {
	err := funcRecover()
	if err != nil {
		fmt.Printf("err is %v\n", err)
	} else {
		fmt.Printf("err is nil\n")
	}

}
