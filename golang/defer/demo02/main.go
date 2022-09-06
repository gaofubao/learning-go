package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	greet()
}

func greet() (err error) {
	var rollbacks []func()
	//rollback := func(rollback func()) {
	//	if err != nil {
	//		rollback()
	//	}
	//}
	defer func() {
		fmt.Println(">>>>> rollbacking")
		if err == nil {
			return
		}

		for _, f := range rollbacks {
			f()
		}
	}()

	inputs := []string{"hi", "tom", "hello", "world"}
	for _, input := range inputs {
		//if err = hello(input); err != nil {
		//	return err
		//}
		rollbacks = append(rollbacks, rollbackHello(input))
	}

	return nil
}

func hello(in string) (err error) {
	if in == "hello" {
		return errors.Errorf("error %s", in)
	}
	return nil
}

func rollbackHello(in string) func() {
	return func() {
		fmt.Printf("rollback %s...\n", in)
	}
}
