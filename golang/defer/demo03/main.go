package main

import (
	"fmt"
	"time"
)

func main() {
	signChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	defer ticker.Stop()

	signChan <- true
	time.Sleep(1)

	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker")
			signChan <- true
		case <-signChan:
			fmt.Println("chan")
		}
	}
}
