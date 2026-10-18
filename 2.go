package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Стоп")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Работа")
			}
		}
	}()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()
	fmt.Println("Нажмите 'Enter' для остановки программы")
	<-done
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Программа остановлена")
}
