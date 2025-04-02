package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Namespace is the namespace for the application metrics
	Namespace = "nfs-metrics-exporter"
	// Subsystem is the subsystem for the application metrics
	Subsystem = "exporter"
	// Labels are the labels for the application metrics
	Labels = []string{"path", "node", "mounted"}
)

// Metrics holds the Prometheus metrics for the application
type Metrics struct {
	AverageOperationsPerSecond *prometheus.GaugeVec
	RPCBklogSize               *prometheus.GaugeVec
	ReadOperationsRatio        *prometheus.GaugeVec
	ReadLatency                *prometheus.GaugeVec
	ReadThroughput             *prometheus.GaugeVec
	ReadRetransmits            *prometheus.GaugeVec
	ReadAverageRTT             *prometheus.GaugeVec
	ReadAverageExecutionTime   *prometheus.GaugeVec
	ReadAverageQueueTime       *prometheus.GaugeVec
	ReadErrors                 *prometheus.GaugeVec
	WriteOperationsRatio       *prometheus.GaugeVec
	WriteLatency               *prometheus.GaugeVec
	WriteThroughput            *prometheus.GaugeVec
	WriteRetransmits           *prometheus.GaugeVec
	WriteAverageRTT            *prometheus.GaugeVec
	WriteAverageExecutionTime  *prometheus.GaugeVec
	WriteAverageQueueTime      *prometheus.GaugeVec
	WriteErrors                *prometheus.GaugeVec
}

// NewMetrics initializes and registers a GaugeVec with the given name and help text
func newGaugeVec(name, help string) *prometheus.GaugeVec {
	// create a new GaugeVec with the given name, labels, and help text
	gaugeVec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      name,
			Help:      help,
		},
		Labels,
	)

	// register the GaugeVec with Prometheus
	if err := prometheus.Register(gaugeVec); err != nil && errors.Is(err, prometheus.AlreadyRegisteredError{}) {
		panic(err)
	}

	return gaugeVec
}

// NewMetrics initializes and registers the Prometheus metrics for the application
func NewMetrics() *Metrics {
	return &Metrics{
		AverageOperationsPerSecond: newGaugeVec("NFS_ME_average_operations_per_second", "Average operations per second"),
		RPCBklogSize:               newGaugeVec("NFS_ME_rpc_bklog_size", "RPC backlog size"),
		ReadOperationsRatio:        newGaugeVec("NFS_ME_read_operations_ratio", "Read operations ratio"),
		ReadLatency:                newGaugeVec("NFS_ME_read_latency", "Read latency"),
		ReadThroughput:             newGaugeVec("NFS_ME_read_throughput", "Read throughput"),
		ReadRetransmits:            newGaugeVec("NFS_ME_read_retransmits", "Read retransmits"),
		ReadAverageRTT:             newGaugeVec("NFS_ME_read_average_rtt", "Read average RTT"),
		ReadAverageExecutionTime:   newGaugeVec("NFS_ME_read_average_execution_time", "Read average execution time"),
		ReadAverageQueueTime:       newGaugeVec("NFS_ME_read_average_queue_time", "Read average queue time"),
		ReadErrors:                 newGaugeVec("NFS_ME_read_errors", "Read errors"),
		WriteOperationsRatio:       newGaugeVec("NFS_ME_write_operations_ratio", "Write operations ratio"),
		WriteLatency:               newGaugeVec("NFS_ME_write_latency", "Write latency"),
		WriteThroughput:            newGaugeVec("NFS_ME_write_throughput", "Write throughput"),
		WriteRetransmits:           newGaugeVec("NFS_ME_write_retransmits", "Write retransmits"),
		WriteAverageRTT:            newGaugeVec("NFS_ME_write_average_rtt", "Write average RTT"),
		WriteAverageExecutionTime:  newGaugeVec("NFS_ME_write_average_execution_time", "Write average execution time"),
		WriteAverageQueueTime:      newGaugeVec("NFS_ME_write_average_queue_time", "Write average queue time"),
		WriteErrors:                newGaugeVec("NFS_ME_write_errors", "Write errors"),
	}
}
