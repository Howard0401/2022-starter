package main

import (
	"fmt"
	"time"
)

func main() {
	// Select類似switch case default
	// select {
	// case <-ch1:
	// case data:=<-ch2:
	// case ch3<-123:
	// default:
	//
	//
	// }
	start := time.Now()
	ch1 := make(chan interface{})
	ch2 := make(chan string)
	ch3 := make(chan string)
	go func() {
		time.Sleep(4 * time.Second)
		close(ch1)
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "ch2 str"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "ch3 str"
	}()
	fmt.Println("loading...")

	// listen first msg from chan
	select {
	case <-ch1:
		fmt.Printf("not block= %v\n", time.Since(start))
	case data := <-ch2:
		fmt.Printf("ch2=%v\n", data)
	case data := <-ch3:
		fmt.Printf("ch3=%v\n", data)
		// default:
		// fmt.Println("default")
	}

	ch4 := make(chan string, 8)
	ch5 := make(chan string, 8)
	var chs = []chan string{ch4, ch5}
	names := []string{"names1", "names2", "names3"}
	//從上到下從左到右
	select {
	case getChan(0, chs) <- getName(2, names):
		fmt.Printf("getChan(0, chs) <- getName(2, names)\n")
	case getChan(1, chs) <- getName(1, names):
		fmt.Printf("getChan(1, chs) <- getName(1, names)\n")
	default:
		fmt.Printf("default\n")
	}

	ch := make(chan string)
	go func() {
		for {
			ch <- "str test anonymous"
		}
	}()

	for {
		select {
		case data := <-ch:
			fmt.Println(data)
			goto exit
		default:
			fmt.Printf("default\n")
		}
		time.Sleep(3 * time.Second)
		fmt.Printf("end of main\n")
	}

exit:
	fmt.Println("退出循環")
}

func getChan(i int, chs []chan string) chan string {
	fmt.Printf("getChan chs=%d\n", i)
	return chs[i]
}

func getName(i int, names []string) string {
	fmt.Printf("getChan chs=%d\n", i)
	return names[i]
}

// 參考當筆記用，非答案 https://www.jb51.net/article/202231.htm
// select中break是只跳脱了select体，而不是结束for循环
// 跳出 select

// method 1

// for {
// 	select{
// 	case <-tick.C:
// 	 //d o someting
// 	case <- stop:
// 	 break //break的不是for循环, 而是跳脱select，执行doNext()
// 	}
// 	doNext()
//  }

// method 2

// for {
// 	select{
// 	case <-tick.C:
// 	 //do someting
// 	case <- stop:
// 	 return //干净利落，适合退出goroutin的场景
// 	}
// 	doNext()
//  }
//  doOther()

// method 3

// LOOP: for {
// 	select{
// 	case <-tick.C:
// 	 //do someting
// 	case <- stop:
// 	 break LOOP//break的for循环,跳转执行doOther()
// 	}
// 	doNext()
//  }
//  doOther()
