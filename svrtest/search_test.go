package main

import (
	"fmt"
	"net/http"
	"testing"
)

//压力测试程序
func BenchmarkSend(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Send()
	}
}

//mockFile 代表测试样本数据
func Send() {
	newRequest, err := http.NewRequest("GET", "127.0.0.1:9000/api/v1/search/go", nil)
	if err != nil {
		return
	}
	response, err := http.DefaultClient.Do(newRequest)
	fmt.Println(response)
	if err != nil {
		fmt.Println(err)
		return
	}
}
