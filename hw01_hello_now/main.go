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
	fmt.Printf("current time: %s\n", timeFormat(timeCurrent))
	fmt.Printf("exact time: %s\n", timeFormat(timeExact))
}

func timeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05 -0700 UTC")
}
