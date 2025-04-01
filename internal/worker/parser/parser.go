package parser

import (
	"bufio"
	"strconv"
	"strings"
)

// ParseNFSIoStat parses the NFS IO stats
func ParseNFSIoStat(input string) []*NFSIoStatType {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var (
		currentMount string
		currentStats NFSIoStatType
		statsList    []*NFSIoStatType
	)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "mounted on") {
			if currentMount != "" {
				statsList = append(statsList, &currentStats)
			}

			parts := strings.Split(line, " mounted on ")
			path := parts[0]
			currentMount = parts[1]
			currentStats = NFSIoStatType{Path: path, MountPoint: currentMount}
		} else if strings.HasPrefix(line, "ops/s") {
			// skip header
		} else if strings.HasPrefix(line, "read:") || strings.HasPrefix(line, "write:") {
			var mode string
			if strings.HasPrefix(line, "read:") {
				mode = "read"
			} else {
				mode = "write"
			}

			scanner.Scan()

			values := strings.Fields(scanner.Text())
			metrics := &NFSIoStatMetricsType{
				OpsPerSec: ParseFloat(values[0]),
				KBPerSec:  ParseFloat(values[1]),
				KBPerOp:   ParseFloat(values[2]),
				Retrans:   ParseFloat(strings.Split(values[3], " ")[0]),
				RTT:       ParseFloat(values[4]),
				Exec:      ParseFloat(values[5]),
				Queue:     ParseFloat(values[6]),
				Errors:    ParseFloat(strings.Split(values[7], " ")[0]),
			}

			if mode == "read" {
				currentStats.Read = metrics
			} else if mode == "write" {
				currentStats.Write = metrics
			}
		} else if line != "" {
			values := strings.Fields(line)
			if len(values) == 2 {
				currentStats.OpsPerSec = ParseFloat(values[0])
				currentStats.RPCBklog = ParseFloat(values[1])
			}
		}
	}

	if currentMount != "" {
		statsList = append(statsList, &currentStats)
	}

	return statsList
}

// ParseFloat parses a float64 from a string
func ParseFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}
