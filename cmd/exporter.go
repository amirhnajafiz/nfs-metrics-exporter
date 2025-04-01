package cmd

import (
	"fmt"

	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker"
)

type CMDExporter struct{}

func (c *CMDExporter) Command() string {
	return "exporter"
}

func (c *CMDExporter) Run() error {
	m := worker.Start()
	fmt.Println(m)

	return nil
}
