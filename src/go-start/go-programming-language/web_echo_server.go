package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "URL.Path = %q \n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "count %d\n", count)
	mu.Unlock()
}

/*
Methods and interfaces（方法和接口）: 方法是关联到一个命名类型的函数。
接口是一种把不同类型同等对待的抽象类型
*/
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counnt", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
