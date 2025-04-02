package cmd

import "github.com/amirhnajafiz/nfs-metrics-exporter/pkg/execute"

// CMDHealth is a command that checks the health of the application
type CMDHealth struct{}

func (c *CMDHealth) Command() string {
	return "health"
}

func (c *CMDHealth) Run() error {
	if _, err := execute.Command("nfsiostat", "-h"); err != nil {
		return err
	}

	return nil
}
