package main

import "fmt"

func main() {
	echo("hello")
}

func echo(format string, args ...any) {
	res := fmt.Sprintf(format, args...)
	fmt.Println(res)
}
