package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const ntpServer = "0.beevik-ntp.pool.ntp.org"

func main() {
	timeCurrent := time.Now()
	timeExact, err := ntp.Time(ntpServer)
	if err != nil {
		log.Fatalf("Ошибка запроса: %e", err)
	}
	fmt.Printf("current time: %s\n", timeCurrent.Round(time.Second).String())
	fmt.Printf("exact time: %s\n", timeExact.Round(time.Second).String())
}
