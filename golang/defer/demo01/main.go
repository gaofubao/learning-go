package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	hello()
}

func hello() (err error) {
	rollback := func(rollback func()) {
		if err != nil {
			rollback()
		}
	}

	if err = initA(); err != nil {
		return err
	}
	defer rollback(rollbackInitA)

	if err = initB(); err != nil {
		return err
	}
	defer rollback(rollbackInitB)

	return nil
}

func initA() (err error) {
	fmt.Println("init A")
	return nil
}

func rollbackInitA() {
	fmt.Println("rollback init A")
}

func initB() (err error) {
	fmt.Println("init B")
	return errors.New("init B error")
}

func rollbackInitB() {
	fmt.Println("rollback init B")
}
