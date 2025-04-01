package cmd

import (
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker"
)

type CMDExporter struct{}

func (c *CMDExporter) Command() string {
	return "exporter"
}

func (c *CMDExporter) Run() error {
	return worker.Start()
}
