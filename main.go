package main

import (
	"flag"
	"log"

	"github.com/amirhnajafiz/nfs-metrics-exporter/cmd"
)

func main() {
	// main flags
	var (
		flagCmd = flag.String("cmd", "exporter", "Specify the command to run (e.g., exporter, requirements). This is required.")
	)

	flag.Parse()

	// get the list of commands from the cmd package
	commands := cmd.CommandList()

	// find and execute the command based on the provided flag
	for _, command := range commands {
		if command.Command() == *flagCmd {
			// execute the command
			log.Printf("Running command: %s\n", command.Command())
			if err := command.Run(); err != nil {
				panic(err)
			}
			return // exit after running the command
		}
	}
}
