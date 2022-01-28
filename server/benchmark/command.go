package benchmark

type command struct {
	workDir string
	command string
}

var commands = map[string]command{
	"isucon-test":      {workDir: ".", command: "sleep 1m && echo %s"},
	"isucon11-qualify": {workDir: "/home/isucon/bench", command: "./bench -all-addresses 127.0.0.11 -target %s -tls -jia-service-url http://127.0.0.1:4999"},
}
