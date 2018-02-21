package main

import(
	"fmt"
	"sync"
	"math/rand"
	"time"
	"sync/atomic"
)
//原子操作 goroutine的计数器
var lock sync.Mutex //互斥锁  读写锁sync.RWMutex()
var rwlock sync.RWMutex
func testMap2(){
	var a map[int]int
	a = make(map[int]int,5)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[19] = 10

	for i :=1 ; i<2 ; i++{

		go func(b map[int]int){
			lock.Lock()
			b[1] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second)
}

func testRWLock(){

	var a map[int]int
	a = make(map[int]int,5)
	var count int32

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[19] = 10

	for i :=1 ; i<2 ; i++{
		go func(b map[int]int){
			rwlock.Lock()
			b[1] = rand.Intn(100)
			time.Sleep(10*time.Millisecond)
			rwlock.Unlock()
		}(a)
	}

	for i := 0;i<20;i++{
		go func(b map[int]int){
			for{
				rwlock.RLock()
				//fmt.Println(a)
				time.Sleep(time.Millisecond)
				rwlock.RUnlock()
				atomic.AddInt32(&count,1)
			}
		}(a)
	}

	time.Sleep(time.Second*3)
	
	fmt.Println(atomic.LoadInt32(&count))
}

func main(){

	testRWLock()

}