package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	time1 := nowTime.Add(-1 * time.Hour)
	//endTime := nowTime.Format("2006-01-02 15:04:05")
	//startTime := nowTime.Add(-30 * time.Minute).Format("2006-01-02 15:04:05")
	//fmt.Println(startTime, endTime)
	//
	////
	//h := nowTime.Add(12 * time.Hour).Hour()
	fmt.Println(nowTime.Add(-1 * time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Println(nowTime.Sub(time1).String())
}
