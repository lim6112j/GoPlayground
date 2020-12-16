package myutil

import (
	log "github.com/sirupsen/logrus"
)

func AFunc(c chan int) {
	entry := log.WithFields(log.Fields{
		"host":    "127.0.0.1",
		"source":  "a.go",
		"service": "naver_crawl",
	})
	entry.Warn("log in A")
	c <- 1
}
