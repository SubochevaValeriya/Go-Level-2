package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	/*	1. С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из
		которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться, что при
		каждом запуске программы итоговое число равно 1000.*/
	workers()
	/*	2. Написать программу, которая при получении в канал сигнала SIGTERM останавливается не
		позднее, чем за одну секунду (установить таймаут).*/
	sigterm()
}

func workers() {
	var sum int
	wg := sync.WaitGroup{}
	wg.Add(1000)
	var workers = make(chan struct{}, 1)
	defer close(workers)
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}

		go func() {
			defer func() {
				<-workers
			}()
			sum++
			wg.Done()
		}()
	}
	wg.Wait()

	if sum == 1000 {
		fmt.Println(true)
	}
}

func sigterm() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	time.Sleep(time.Second * 1)
	fmt.Println("exiting")
}
