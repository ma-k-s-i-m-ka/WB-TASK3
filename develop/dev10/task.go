package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		host string
		port string
	)
	if len(os.Args) < 3 {
		fmt.Println("Команда для запуска: develop/dev09/task.go [--timeout=10s] host port")
		os.Exit(1)
	}

	timeout := 10 * time.Second
	if os.Args[1] == "--timeout=" {
		i := 1
		var err error
		timeout, err = time.ParseDuration(os.Args[i+1])
		if err != nil {
			fmt.Println("Ошибка передачи timeout:", err)
			return
		}
		host = os.Args[i+2]
		port = os.Args[i+3]
	} else {
		host = os.Args[1]
		port = os.Args[2]
	}
	fmt.Printf("Время на подключение: %s\n", timeout)

	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	fmt.Printf("Подключение по адресу: %s выполнено\n", addr)
	defer conn.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("Ошибка чтения из STDIN:", err)
				break
			}
			if n == 1 && buf[0] == 4 {
				break
			}
			conn.Write(buf[:n])
		}
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Соединение закрыто со стороны сервера")
				os.Exit(0)
				return
			}
			os.Stdout.Write(buf[:n])
		}
	}()

	<-sigCh
	fmt.Println("Соединение закрыто")
}
