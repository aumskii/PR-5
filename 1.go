package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Ответ от сервера получен"
	}()
	select {
	case r := <-ch:
		fmt.Println(r)
	case <-time.After(2 * time.Second):
		fmt.Println("Ошибка, превышено время соединения")
	}
}
