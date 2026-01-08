package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func num(n int) {
	defer wg.Done()
	for i := (n-1)*3000000 + 1; i <= n*3000000; i++ {
		if i > 1 {
			flag := true

			for j := 2; j < i; j++ {
				if i%j == 0 {
					flag = false
					break
				}

			}
			if flag {
				// fmt.Println(i, " is a prime number")
			}
		}
	}
}

func main() {
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go num(i)
	}
	wg.Wait()
	fmt.Println("执行完毕")
	end := time.Now().Unix()
	fmt.Println(end - start)
}
