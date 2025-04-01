package worker

import "github.com/amirhnajafiz/nfs-metrics-exporter/internal/worker/parser"

func Start() []parser.NFSIoStatType {
	input := `nap:/home/anajafizadeh mounted on /home/anajafizadeh:

           ops/s       rpc bklog
          73.187           0.000

read:              ops/s            kB/s           kB/op         retrans    avg RTT (ms)    avg exe (ms)  avg queue (ms)          errors
                   0.000           0.000           9.096        0 (0.0%)           2.158           2.225           0.042        0 (0.0%)
write:             ops/s            kB/s           kB/op         retrans    avg RTT (ms)    avg exe (ms)  avg queue (ms)          errors
                   0.000           0.000          11.367        0 (0.0%)           0.342           0.421           0.061        0 (0.0%)

nap:/workloads/sunyibm/ibm2 mounted on /var/lib/kubelet-nfs:

           ops/s       rpc bklog
          81.835           0.000

read:              ops/s            kB/s           kB/op         retrans    avg RTT (ms)    avg exe (ms)  avg queue (ms)          errors
                  13.069         207.339          15.866        0 (0.0%)           0.536           0.595           0.032        0 (0.0%)
write:             ops/s            kB/s           kB/op         retrans    avg RTT (ms)    avg exe (ms)  avg queue (ms)          errors
                   0.772          31.224          40.453        0 (0.0%)           1.083           4.515           3.397        0 (0.0%)`

	return parser.ParseNFSIoStat(input)
}
