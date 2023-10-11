package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	time2 "time"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		err = fmt.Errorf("[Ошибка: %s] при подключения к серверу NTP", err)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("\nТочное время с NTP-сервера: %s\n", time.Format(time2.DateTime))
}
