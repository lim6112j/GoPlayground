package main

import (
	"../myutil"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File context

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
func main() {
	// 로그를 저장할 파일 설정.
	fpLog, err := os.OpenFile("../logfile.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("logfile open failed => %v", err)
	}
	// 파일 닫기를  프로그램 종료시로 지연시킴.
	defer fpLog.Close()
	mw := io.MultiWriter(os.Stdout, fpLog)
	log.SetOutput(mw)
	entry := log.WithFields(log.Fields{
		"host":    "127.0.0.1",
		"source":  "taillog.go",
		"service": "naver_crawl",
	})
	entry.Warn("A walrus appears")
	entry.Warn("error while decoding")
	c := make(chan int)
	go myutil.AFunc(c)
	entry.Warnf("channel returns %v", <-c)
}
