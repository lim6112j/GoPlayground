package main

import (
	"../util"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type ApiWriter struct {
}

func (aw *ApiWriter) Write(p []byte) (n int, err error) {
	logUrl := "https://thecheat.co.kr/thecheat/app/___PutError.php"
	b64data := base64.StdEncoding.EncodeToString(p)
	fmt.Println(b64data)
	dataString := fmt.Sprintf("contents=%s", b64data)
	req, err := http.NewRequest("POST", logUrl, strings.NewReader(dataString))
	if err != nil {
		fmt.Errorf("error while requesting logifile to server : %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	timeout := time.Duration(20 * time.Second)
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("file log to server error  => %v", err)
		return
	}
	util.PrintJsonResponse(resp)
	resp.Body.Close()
	return len(p), nil
}
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
	aw := &ApiWriter{}
	mw := io.MultiWriter(os.Stdout, fpLog, aw)
	log.SetOutput(mw)
	entry := log.WithFields(log.Fields{
		"host":    "127.0.0.1",
		"source":  "taillog.go",
		"service": "naver_crawl",
	})
	entry.Warn("A walrus appears")
	entry.Warn("error while decoding")
	c := make(chan int)
	go util.AFunc(c)
	entry.Warnf("channel returns %v", <-c)
}
