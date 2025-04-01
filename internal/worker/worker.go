package worker

import (
	"bufio"
	"strconv"
	"strings"
)

type Metrics struct {
	OpsPerSec float64
	KBPerSec  float64
	KBPerOp   float64
	Retrans   float64
	RTT       float64
	Exec      float64
	Queue     float64
	Errors    float64
}

type NFSStats struct {
	MountPoint string
	OpsPerSec  float64
	RPCBklog   float64
	Read       Metrics
	Write      Metrics
}

func Start() []NFSStats {
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

	scanner := bufio.NewScanner(strings.NewReader(input))
	var statsList []NFSStats
	var currentMount string
	var currentStats NFSStats

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "mounted on") {
			if currentMount != "" {
				statsList = append(statsList, currentStats)
			}
			parts := strings.Split(line, " mounted on ")
			currentMount = parts[1]
			currentStats = NFSStats{MountPoint: currentMount}
		} else if strings.HasPrefix(line, "ops/s") {
			// Skip header
		} else if strings.HasPrefix(line, "read:") || strings.HasPrefix(line, "write:") {
			mode := strings.TrimSuffix(line, ":")
			scanner.Scan()
			values := strings.Fields(scanner.Text())
			metrics := Metrics{
				OpsPerSec: parseFloat(values[0]),
				KBPerSec:  parseFloat(values[1]),
				KBPerOp:   parseFloat(values[2]),
				Retrans:   parseFloat(strings.Split(values[3], " ")[0]),
				RTT:       parseFloat(values[4]),
				Exec:      parseFloat(values[5]),
				Queue:     parseFloat(values[6]),
				Errors:    parseFloat(strings.Split(values[7], " ")[0]),
			}
			if mode == "read" {
				currentStats.Read = metrics
			} else if mode == "write" {
				currentStats.Write = metrics
			}
		} else if line != "" {
			values := strings.Fields(line)
			if len(values) == 2 {
				currentStats.OpsPerSec = parseFloat(values[0])
				currentStats.RPCBklog = parseFloat(values[1])
			}
		}
	}

	if currentMount != "" {
		statsList = append(statsList, currentStats)
	}

	return statsList
}

func parseFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}
