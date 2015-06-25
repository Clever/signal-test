package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"gopkg.in/clever/kayvee-go.v2"
)

type m map[string]interface{}

func main() {
	go heartbeat(time.Second)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	for {
		select {
		case sig := <-sigChan:
			log.Println(kayvee.Format(m{
				"source": "signal-test",
				"title":  "signal-received",
				"level":  kayvee.Info,
				"signal": sig,
			}))
		}
	}
}

func heartbeat(interval time.Duration) {
	ticker := time.Tick(interval)
	for _ = range ticker {
		log.Println(kayvee.Format(m{
			"source": "signal-test",
			"title":  "heartbeat",
			"level":  kayvee.Info,
		}))
	}
}
