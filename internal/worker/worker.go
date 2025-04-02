package worker

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/metrics"
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker/parser"
	"github.com/amirhnajafiz/nfs-metrics-exporter/pkg/execute"

	"go.uber.org/zap"
)

// Start starts the NFS I/O statistics collection at the specified interval
func Start(interval time.Duration, me *metrics.Metrics, logr *zap.Logger) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		// run the command to collect NFS I/O statistics
		output, err := execute.Command("cat", "./example.out")
		if err != nil {
			return err
		}

		// parse the output
		stats := parser.ParseNFSIoStat(output)

		// Print the parsed output
		for _, stat := range stats {
			fmt.Println("Path:", stat.Path)
			fmt.Println("MountPoint:", stat.MountPoint)
			fmt.Printf("OpsPerSec: %.2f\n", stat.OpsPerSec)
			fmt.Println("RPCBklog:", stat.RPCBklog)
			fmt.Printf("Read OpsPerSec: %.2f\n", stat.Read.OpsPerSec)
			fmt.Printf("Read KBPerSec: %.2f\n", stat.Read.KBPerSec)
			fmt.Printf("Read KBPerOp: %.2f\n", stat.Read.KBPerOp)
			fmt.Println("Read Retrans:", stat.Read.Retrans)
			fmt.Printf("Read RTT: %.2f\n", stat.Read.RTT)
			fmt.Printf("Read Exec: %.2f\n", stat.Read.Exec)
			fmt.Println("Read Queue:", stat.Read.Queue)
			fmt.Println("Read Errors:", stat.Read.Errors)
			fmt.Printf("Write OpsPerSec: %.2f\n", stat.Write.OpsPerSec)
			fmt.Printf("Write KBPerSec: %.2f\n", stat.Write.KBPerSec)
			fmt.Printf("Write KBPerOp: %.2f\n", stat.Write.KBPerOp)
			fmt.Println("Write Retrans:", stat.Write.Retrans)
			fmt.Printf("Write RTT: %.2f\n", stat.Write.RTT)
			fmt.Printf("Write Exec: %.2f\n", stat.Write.Exec)
			fmt.Println("Write Queue:", stat.Write.Queue)
			fmt.Println("Write Errors:", stat.Write.Errors)
		}
	}

	return nil
}
