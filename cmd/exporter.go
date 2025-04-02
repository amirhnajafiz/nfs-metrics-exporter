package cmd

import (
	"time"

	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/config"
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/metrics"
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker"
)

// CMDExporter is a command that starts the exporter
type CMDExporter struct{}

func (c *CMDExporter) Command() string {
	return "exporter"
}

func (c *CMDExporter) Run() error {
	// load configs
	cfg := config.Load()

	// create a new metrics instance
	me := metrics.NewMetrics()

	// start the metrics server
	go func() {
		if err := metrics.NewServer(cfg.ServicePort).Start(); err != nil {
			panic(err)
		}
	}()

	// start the worker
	worker.Start(time.Duration(cfg.ExportInterval)*time.Second, me)

	return nil
}
