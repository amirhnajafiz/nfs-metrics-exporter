package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Namespace is the namespace for the application metrics
	Namespace = "NFS"
	// Subsystem is the subsystem for the application metrics
	Subsystem = "ME"
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
		AverageOperationsPerSecond: newGaugeVec("average_operations_per_second", "Average operations per second"),
		RPCBklogSize:               newGaugeVec("rpc_bklog_size", "RPC backlog size"),
		ReadOperationsRatio:        newGaugeVec("read_operations_ratio", "Read operations ratio"),
		ReadLatency:                newGaugeVec("read_latency", "Read latency"),
		ReadThroughput:             newGaugeVec("read_throughput", "Read throughput"),
		ReadRetransmits:            newGaugeVec("read_retransmits", "Read retransmits"),
		ReadAverageRTT:             newGaugeVec("read_average_rtt", "Read average RTT"),
		ReadAverageExecutionTime:   newGaugeVec("read_average_execution_time", "Read average execution time"),
		ReadAverageQueueTime:       newGaugeVec("read_average_queue_time", "Read average queue time"),
		ReadErrors:                 newGaugeVec("read_errors", "Read errors"),
		WriteOperationsRatio:       newGaugeVec("write_operations_ratio", "Write operations ratio"),
		WriteLatency:               newGaugeVec("write_latency", "Write latency"),
		WriteThroughput:            newGaugeVec("write_throughput", "Write throughput"),
		WriteRetransmits:           newGaugeVec("write_retransmits", "Write retransmits"),
		WriteAverageRTT:            newGaugeVec("write_average_rtt", "Write average RTT"),
		WriteAverageExecutionTime:  newGaugeVec("write_average_execution_time", "Write average execution time"),
		WriteAverageQueueTime:      newGaugeVec("write_average_queue_time", "Write average queue time"),
		WriteErrors:                newGaugeVec("write_errors", "Write errors"),
	}
}
