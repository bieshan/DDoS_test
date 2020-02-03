package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	num := 500
	for cnt := 0; cnt < num; cnt++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			request()
		}()
	}
	wg.Wait()
}

func request() {
	url := "http://raspberrypi.local:8080/ping"
	for {
		resp, _ := http.Get(url)
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}
