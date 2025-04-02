package cmd

import (
	"time"

	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/config"
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/logr"
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/metrics"
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker"

	"go.uber.org/zap"
)

// CMDExporter is a command that starts the exporter
type CMDExporter struct{}

func (c *CMDExporter) Command() string {
	return "exporter"
}

func (c *CMDExporter) Run() error {
	// load configs
	cfg := config.Load()

	// create a new logger instance
	logger := logr.NewZapLogger(cfg.DebugMode)

	// create a new metrics instance
	me := metrics.NewMetrics()

	// start the metrics server
	logger.Info("starting metrics server", zap.String("port", cfg.ServicePort))
	go func() {
		if err := metrics.NewServer(cfg.ServicePort, cfg.SecretKey).Start(); err != nil {
			panic(err)
		}
	}()

	// build and start a worker
	wo := worker.Worker{
		Logr:    logger.Named("worker"),
		Metrics: me,
	}
	wo.Start(time.Duration(cfg.ExportInterval) * time.Second)

	return nil
}
