package benchmark

import (
	"fmt"
	"strings"

	"github.com/logica0419/remote-bmi/server/repository"
)

type command struct {
	workDir   string
	createCmd func(benchServerIP string, servers []*repository.Server, serverNumber int) (string, error)
}

var commands = map[string]command{
	"isucon-test": {
		workDir: ".",
		createCmd: func(benchServerIP string, servers []*repository.Server, serverNumber int) (string, error) {
			return "sleep 2s", nil
		}},

	"isucon11-qualify": {
		workDir: "/home/isucon/bench",
		createCmd: func(benchServerIP string, servers []*repository.Server, serverNumber int) (string, error) {
			allAddressesArr := []string{}
			target := ""

			for _, server := range servers {
				allAddressesArr = append(allAddressesArr, server.Address)
				if server.ServerNumber == serverNumber {
					target = server.Address
				}
			}

			if target == "" {
				return "", fmt.Errorf("target server not found")
			}
			allAddresses := strings.Join(allAddressesArr, ",")

			return fmt.Sprintf("./bench -all-addresses %s -target %s -tls -jia-service-url http://%s:4999", allAddresses, target, benchServerIP), nil
		},
	},

	"isucon11-final": {
		workDir: "/home/isucon/bench",
		createCmd: func(benchServerIP string, servers []*repository.Server, serverNumber int) (string, error) {
			target := ""

			for _, server := range servers {
				if server.ServerNumber == serverNumber {
					target = server.Address
				}
			}

			if target == "" {
				return "", fmt.Errorf("target server not found")
			}

			return fmt.Sprintf("./bin/benchmarker -target %s:443 -tls", target), nil
		},
	},
}
