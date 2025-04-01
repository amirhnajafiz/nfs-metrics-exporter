package cmd

// CMD interface defines the methods that any command type must implement.
type CMD interface {
	Command() string // Returns the name of the command
	Run() error      // Executes the command logic
}

// CommandList function returns a slice of all available commands.
func CommandList() []CMD {
	// This function returns a list of all available commands.
	// add more commands to this list as needed.
	return []CMD{
		&CMDExporter{},
		&CMDHealth{},
	}
}
