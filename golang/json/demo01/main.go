package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	name := `{"hello":"world"}`
	addr := "\"shanghai\""

	type A struct {
		Name interface{}
		Addr interface{}
		Age  int64
	}

	var a A
	a.Age = 10
	if err := json.Unmarshal([]byte(name), &a.Name); err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(addr), &a.Addr)
	fmt.Println(a)
}
