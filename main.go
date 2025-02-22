package main

import (
	"fmt"
	"ag/config"
	"ag/cmd"
)

func main() {
	if err := config.Load(); err != nil {
        fmt.Println(err)
        return
    }

	cmd.Execute()
}
