package main

import (
	"fmt"
	"github.com/shenhailuanma/goutils/cmd"
)

func main()  {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("env")

	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()


	// Print STDOUT from Cmd
	fmt.Println(status.Stdout)
}
