package main

import (
	"fmt"
	"github.com/shenhailuanma/goutils/cmd"
)

func main() {

	cmdOptions := cmd.Options{
		Buffered:    true,
		Streaming:   false,
		Interlacing: true,
	}

	// Create Cmd with options
	command := cmd.NewCmdOptions(cmdOptions, "ffmpeg", "-v")

	// Run and wait for Cmd to return Status
	status := <-command.Start()

	// Print STDOUT from Cmd
	fmt.Println(status.Stdout)

}
