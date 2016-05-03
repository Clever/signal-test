package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/clever/kayvee-go.v2"
)

type m map[string]interface{}

var signals = []string{}

func main() {
	// go heartbeat(time.Second)
	go listenForSignals()

	http.HandleFunc("/signal", SignalHandler)
	http.HandleFunc("/", HealthCheckHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

func listenForSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case sig := <-sigChan:
			log.Println(kayvee.Format(m{
				"source": "signal-test",
				"title":  "signal-received",
				"level":  kayvee.Info,
				"signal": sig,
			}))
			signals = append(signals, sig.String())
		}
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func SignalHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(signals); err != nil {
		log.Println(kayvee.Format(m{
			"source": "signal-test",
			"title":  "signal-handler",
			"level":  kayvee.Error,
			"error":  err.Error(),
		}))
	}
	log.Println(kayvee.Format(m{
		"source": "signal-test",
		"title":  "signal-handler",
		"level":  kayvee.Info,
		"msg":    "Served /signal request",
	}))
	return
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
