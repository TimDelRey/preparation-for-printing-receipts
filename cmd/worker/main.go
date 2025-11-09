package main

import (
	"fmt"
	"script_for_receipts/pkg/service"
)

func main() {
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
