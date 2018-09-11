
/*
读锁定（RLock），对读操作进行锁定
读解锁（RUnlock），对读锁定进行解锁
写锁定（Lock），对写操作进行锁定
写解锁（Unlock），对写锁定进行解锁

如何理解读写锁呢？

同时只能有一个 goroutine 能够获得写锁定。
同时可以有任意多个 gorouinte 获得读锁定。
同时只能存在写锁定或读锁定（读和写互斥）。
也就是说，当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；
当有一个 goroutine 获得读锁定，其它读锁定任然可以继续；当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。
所以说这里的读锁定（RLock）目的其实是告诉写锁定：有很多人正在读取数据，你给我站一边去，等它们读（读解锁）完你再来写（写锁定）。
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var rw sync.RWMutex

func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
    //time.Sleep(time.Second*10)
	for i := 0; i < 10; i++ {
		<-ch
	}
}

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}