package benchmark

import (
	"fmt"
	"strings"

	"github.com/logica0419/remote-bmi/server/repository"
)

type command struct {
	workDir   string
	createCmd func(servers []*repository.Server, serverNumber int) (string, error)
}

var commands = map[string]command{
	"isucon-test": {
		workDir: ".",
		createCmd: func(servers []*repository.Server, serverNumber int) (string, error) {
			return "sleep 2s", nil
		}},

	"isucon11-qualify": {
		workDir: "/home/isucon/bench",
		createCmd: func(servers []*repository.Server, serverNumber int) (string, error) {
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

			return fmt.Sprintf("./bench -all-addresses %s -target %s -tls -jia-service-url http://127.0.0.1:4999", allAddresses, target), nil
		},
	},
}
