package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех
	fmt.Println(waitG(10))

	//  2. Реализуйте функцию для разблокировки мьютекса с помощью defer
	fmt.Println(mutex(10))

}

func waitG(n int) int {
	var sum int64
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			atomic.AddInt64(&sum, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	return int(sum)
}

func mutex(n int) int {
	sum := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			sum++
			wg.Done()
		}()
	}
	wg.Wait()
	return sum
}
