package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	urls := strings.Split(os.Getenv("PING_URLS"), ",")

	for _, url := range urls {
		go pingUrl(url)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting Down")
}

func pingUrl(url string) {
	for {
		_, err := http.Get(url)

		log.Println("Pinging " + url)

		if err != nil {
			log.Println("There was on error pingin" + url)
		}

		time.Sleep(time.Second * 5)
	}
}
