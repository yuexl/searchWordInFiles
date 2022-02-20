package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	fmt.Println(time.Now().String())

	//wg := sync.WaitGroup{}
	//wg.Add(5000)
	//for i := 0; i < 5000; i++ {
	//	go Send(&wg)
	//}
	//wg.Wait()

	appName := os.Args[0]
	fmt.Println(appName)
	fmt.Println(filepath.Base(appName))
	fmt.Println(filepath.Abs(appName))

	fmt.Println(time.Now().String())
}

func Send(wg *sync.WaitGroup) {
	defer wg.Done()
	newRequest, err := http.NewRequest("GET", "http://127.0.0.1:9000/api/v1/search/redis", nil)
	if err != nil {
		return
	}
	http.DefaultClient.Do(newRequest)
	//data := make([]byte, 1000)
	//bufio.NewReader(response.Body).Read(data)
	//fmt.Println(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
}
