package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("do before tests")
	val := m.Run()
	fmt.Println("do after tests")
	os.Exit(val)
}
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
func TestLog(t *testing.T) {
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
