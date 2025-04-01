package parser

// NFSIoStatMetricsType represents the metrics of NFS IO stats
type NFSIoStatMetricsType struct {
	OpsPerSec float64
	KBPerSec  float64
	KBPerOp   float64
	Retrans   float64
	RTT       float64
	Exec      float64
	Queue     float64
	Errors    float64
}

// NFSIoStatType represents the NFS IO stats
type NFSIoStatType struct {
	Path       string
	MountPoint string
	OpsPerSec  float64
	RPCBklog   float64
	Read       *NFSIoStatMetricsType
	Write      *NFSIoStatMetricsType
}
