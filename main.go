package main

import (
	"ag/cmd"
	"ag/config"
	"fmt"
)

func main() {
	if err := config.Load(); err != nil {
		fmt.Println(err)
		return
	}

	cmd.Execute()
}
