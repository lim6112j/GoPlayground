package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("tests start")
	val := m.Run()
	fmt.Println("do after tests")
	os.Exit(val)
}
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File context

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
func TestLog(t *testing.T) {
	// 로그를 저장할 파일 설정.
	fpLog, err := os.OpenFile("logfile.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("logfile open failed => %v", err)
	}
	// 파일 닫기를  프로그램 종료시로 지연시킴.
	defer fpLog.Close()
	mw := io.MultiWriter(os.Stdout, fpLog)
	log.SetOutput(mw)
	fmt.Println("Testing logrus")

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log.WithField("source", "thecheat_complete.go").Warn("error while decoding")
}
func TestB(t *testing.T) {
	fmt.Println("Testing B")
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			fmt.Println(ip)
		}
	}
}
func TestC(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Println(localAddr.IP)
}
func TestContext(t *testing.T) {

}
