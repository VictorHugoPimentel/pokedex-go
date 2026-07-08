package main

import (
	"fmt"
	"os"
)

func callbackExit() error{
	fmt.Println("Shutting down!")
	os.Exit(0)
	return nil
}