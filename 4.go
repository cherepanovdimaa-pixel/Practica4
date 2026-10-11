package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{"https://ktk-45.ru","https://github.com/cherepanovdimaa-pixel","https://ya.ru","https://playerok.com", }
	semafor := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func() {
			semafor <- struct{}{}
			defer func () {<- semafor; wg.Done()}()
			resp, error := http.Get(url)
			if error != nil {
				fmt.Printf("Ошибка: %v\n", url, error)
				return
			}
			fmt.Printf("%-20s -> %s\n", url, resp.Status)
		}()
	}
	wg.Wait()
	fmt.Println("Завершено")
}