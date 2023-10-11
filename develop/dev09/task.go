package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Команда для запуска: develop/dev09/task.go <URL>")
		return
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка запроса к сайту %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: %s вернул статус %d\n", url, resp.StatusCode)
		return
	}

	fileName := "C:/Users/Maks/GolandProjects/WBL2/develop/dev09/test"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при копировании данных: %v\n", err)
		return
	}

	fmt.Printf("Сайт успешно скачан и сохранен в %s\n", fileName)
}
