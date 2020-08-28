package main

import (
	"flag"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	var timeout time.Duration
	var wg sync.WaitGroup
	flag.DurationVar(&timeout, "timeout", time.Second*10, "Set connection timeout. Default = 10s")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("incorrect host/port")
	}
	addr := net.JoinHostPort(args[0], args[1])
	client := NewTelnetClient(addr, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatal("Can't connect: ", err.Error())
	}
	log.Println("...connected to", addr)
	defer client.Close()

	wg.Add(2)
	go readRoutine(&wg, client)
	go writeRoutine(&wg, client)
	wg.Wait()
}

func readRoutine(wg *sync.WaitGroup, client TelnetClient) {
	defer wg.Done()
	for {
		if err := client.Receive(); err != nil {
			log.Fatal(err)
			return
		}
	}
}

func writeRoutine(wg *sync.WaitGroup, client TelnetClient) {
	defer wg.Done()
	for {
		if err := client.Send(); err != nil {
			log.Fatal(err)
			return
		}
	}
}
