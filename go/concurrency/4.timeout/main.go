package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type Result struct {
	Value string
}

func PipeLine(servers ...*httptest.Server) (Result, error) {
	c := make(chan Result)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// call apis
	callAPIFunc := func(i int, server *httptest.Server) {
		url := server.URL
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("[GET] %s, req err:%v \n", url, err)
			return
		}
		req = req.WithContext(ctx)
		log.Printf("[GET] %s sent... \n", url)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("[GET] %s, res err:%v \n", url, err)
			return
		}
		log.Printf("[GET] %s, res=%v \n", url, res)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		c <- Result{
			Value: string(body),
		}
		return
	}

	// goroutine req
	for k, v := range servers {
		go callAPIFunc(k, v)
	}

	// select res staus
	select {
	case r := <-c:
		return r, nil
	case <-time.After(700 * time.Millisecond):
		return Result{}, fmt.Errorf("timeout")
	}
}

func mockServer(name string, interval int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("recv http req:%v \n", name)
		time.Sleep(time.Duration(interval) * time.Millisecond)
		w.Write([]byte(name + ":ok"))
	}))
}

// Is it similar to Promise.race?
func main() {
	res, err := PipeLine(mockServer("server1", 200), mockServer("server2", 300), mockServer("server3", 600))
	if err != nil {
		log.Printf("PipeLine err=%v \n", err)
	}
	fmt.Printf("res=%v\n", res)
	time.Sleep(10 * time.Second)
}
