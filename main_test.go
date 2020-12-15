package main

import (
	"fmt"
	"os"
	"testing"
)
func TestMain(m *testing.M){
	fmt.Println("do before tests")
	val := m.Run()
	fmt.Println("do after tests")
	os.Exit(val)
}
func TestA(t *testing.T) {
	fmt.Println("testing A")
}
func TestB(t *testing.T) {
	fmt.Println("Testing B")
}