package main

import (
	"fmt"
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
	var workers = make(chan struct{}, 1)
	defer close(workers)
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}

		go func(job int) {
			defer func() {
				<-workers
			}()
			fmt.Println(job)
		}(i)
	}

	time.Sleep(2 * time.Second)
}

func sigterm() {
	ch := make(chan string)
	defer close(ch)
	go func() {
		ch <- "SIGTERM"
	}()

	select {
	case <-ch:
		return

	case <-time.After(1 * time.Second):
		fmt.Println("Second passed")
	}
}
