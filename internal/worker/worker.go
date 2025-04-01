package worker

import (
	"github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker/parser"
	"github.com/amirhnajafiz/nfs-metrics-exporter/pkg/execute"
)

func Start() error {
	// run nfsiostat command
	output, err := execute.Command("cat", "./example.out")
	if err != nil {
		return err
	}

	// parse the output
	stats := parser.ParseNFSIoStat(output)

	// print the parsed output
	for _, stat := range stats {
		println(stat.MountPoint)
		println(stat.OpsPerSec)
		println(stat.RPCBklog)
		println(stat.Read.OpsPerSec)
		println(stat.Read.KBPerSec)
		println(stat.Read.KBPerOp)
		println(stat.Read.Retrans)
		println(stat.Read.RTT)
		println(stat.Read.Exec)
		println(stat.Read.Queue)
		println(stat.Read.Errors)
		println(stat.Write.OpsPerSec)
		println(stat.Write.KBPerSec)
		println(stat.Write.KBPerOp)
		println(stat.Write.Retrans)
		println(stat.Write.RTT)
		println(stat.Write.Exec)
		println(stat.Write.Queue)
		println(stat.Write.Errors)
	}

	return nil
}
