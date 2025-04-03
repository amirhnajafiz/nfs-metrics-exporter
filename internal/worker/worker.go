package worker

import (
	"time"

	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/metrics"
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker/parser"
	"github.com/amirhnajafiz/nfs-metrics-exporter/pkg/execute"
	"github.com/prometheus/client_golang/prometheus"

	"go.uber.org/zap"
)

// Worker is a struct that contains necessary components to collect and export metrics
type Worker struct {
	Hostname string
	Logr     *zap.Logger
	Metrics  *metrics.Metrics
}

// Start starts the NFS I/O statistics collection at the specified interval
func (w *Worker) Start(interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		w.Logr.Info("collecting NFS I/O statistics", zap.Duration("interval", interval))

		// run the command to collect NFS I/O statistics using nsenter
		output, err := execute.Command("nsenter", "--mount=/proc/1/ns/mnt", "nfsiostat", "1", "1")
		if err != nil {
			w.Logr.Error("failed to execute nfsiostat", zap.Error(err))
			continue
		}

		// parse the output
		stats := parser.ParseNFSIoStat(output)
		w.Logr.Debug("parsed NFS I/O statistics", zap.Int("stats", len(stats)))

		// convert the stats to prometheus metrics
		for _, stat := range stats {
			w.exportToProms(stat)
			w.Logr.Debug("exported NFS I/O statistics to Prometheus",
				zap.String("path", stat.Path),
				zap.String("mounted", stat.MountPoint),
			)
		}
	}

	return nil
}

// export the NFS IO stats to prometheus metrics
func (w *Worker) exportToProms(stat *parser.NFSIoStatType) {
	// create the labels for each observation
	labels := prometheus.Labels{"node": w.Hostname, "path": stat.Path, "mounted": stat.MountPoint}

	// write the metrics
	w.Metrics.AverageOperationsPerSecond.With(labels).Add(stat.OpsPerSec)
	w.Metrics.RPCBklogSize.With(labels).Add(stat.RPCBklog)

	w.Metrics.ReadOperationsRatio.With(labels).Add(stat.Read.OpsPerSec)
	w.Metrics.ReadLatency.With(labels).Add(stat.Read.KBPerSec)
	w.Metrics.ReadThroughput.With(labels).Add(stat.Read.KBPerOp)
	w.Metrics.ReadRetransmits.With(labels).Add(stat.Read.Retrans)
	w.Metrics.ReadAverageRTT.With(labels).Add(stat.Read.RTT)
	w.Metrics.ReadAverageExecutionTime.With(labels).Add(stat.Read.Exec)
	w.Metrics.ReadAverageQueueTime.With(labels).Add(stat.Read.Queue)
	w.Metrics.ReadErrors.With(labels).Add(stat.Read.Errors)

	w.Metrics.WriteOperationsRatio.With(labels).Add(stat.Write.OpsPerSec)
	w.Metrics.WriteLatency.With(labels).Add(stat.Write.KBPerSec)
	w.Metrics.WriteThroughput.With(labels).Add(stat.Write.KBPerOp)
	w.Metrics.WriteRetransmits.With(labels).Add(stat.Write.Retrans)
	w.Metrics.WriteAverageRTT.With(labels).Add(stat.Write.RTT)
	w.Metrics.WriteAverageExecutionTime.With(labels).Add(stat.Write.Exec)
	w.Metrics.WriteAverageQueueTime.With(labels).Add(stat.Write.Queue)
	w.Metrics.WriteErrors.With(labels).Add(stat.Write.Errors)
}
