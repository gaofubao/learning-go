package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {
	err := test3()
	log.Printf("%+v", err)
}

func test3() error {
	return test2()
}

func test2() error {
	err := test1()
	return errors.Errorf("%v", err)
}

func test1() error {
	err := fmt.Errorf("hello")
	//return errors.Errorf("test1: %v", err)
	return err
}
