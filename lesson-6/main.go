package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	//1. Написать программу, которая использует мьютекс для безопасного доступа к данным
	//из нескольких потоков. Выполните трассировку программы

	mutex()

	//2. Написать многопоточную программу, в которой будет использоваться явный вызов
	//планировщика. Выполните трассировку программы

	sheduler()

	//3. Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”

	race()

	// go -race main.go - to check
}

func mutex() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var (
		counter int
		lock    sync.Mutex
		wg      sync.WaitGroup
	)

	wg.Add(1000)
	for i := 0; i < 1000; i += 1 {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func sheduler() {
	f, err := os.Create("trace_sheduler.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	go log.Println("I'm working!")
	for i := 0; i < 1000; i += 1 {
		for j := 1; j < 1000; j += 1 {
			if i+j > 50 {
				runtime.Gosched()
			}
		}
	}
}

func race() {
	var (
		counter int
		wg      sync.WaitGroup
	)

	wg.Add(1000)
	for i := 0; i < 1000; i += 1 {
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
