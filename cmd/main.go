package main

import (
	"fmt"
	"service-account/config"
)

func main() {
	config.LoadConfig("./.env")
	fmt.Println("hello world")
}
