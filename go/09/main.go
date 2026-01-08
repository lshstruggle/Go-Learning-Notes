package main

import (
	"fmt"
)

func main() {
	var ch = make(chan int, 10)
	// 必须先有数据写入，或者在另一个 goroutine 中写入，否则 range 会一直阻塞等待
	// 这里演示先写入数据，再关闭通道，然后 range 才能正常读取并退出
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 关键：写入完毕后关闭通道

	// 使用 range 自动读取，直到通道关闭且为空
	for v := range ch {
		fmt.Println(v)
	}

	// 下面是导致死锁的错误写法，已删除
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }
}
