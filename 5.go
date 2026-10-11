package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func hesh(jobs <-chan string) {
	for NameF := range jobs {
		data, errorr := os.ReadFile(NameF)
		if errorr != nil {
			fmt.Printf("Ошибка открытия: ", NameF)
			continue
		}
		hesh := md5.Sum(data)
		fmt.Printf("MD5(%s): %x\n", NameF, hesh)
	}
}
func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}
	jobs := make(chan string)
	go hesh(jobs)
	for _, NameF := range files {
		jobs <- NameF
	}
	close(jobs)
	fmt.Println("Выполнено")
}
