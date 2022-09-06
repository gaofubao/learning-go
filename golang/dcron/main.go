package main

import (
	"fmt"
	"github.com/libi/dcron"
	"github.com/libi/dcron/driver/redis"
)

func main() {
	drv, _ := redis.NewDriver(&redis.Conf{
		Host:     "",
		Port:     0,
		Password: "",
	})
	dc := dcron.NewDcron("server1", drv)
	dc.AddFunc("", "", func() {
		fmt.Println("")
	})
	dc.Start()
}
