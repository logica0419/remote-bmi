package benchmark

var commands = map[string]string{
	"isucon-test":      "sleep 1m; echo %s",
	"isucon11-qualify": "cd bench && ./bench -all-addresses 127.0.0.11 -target %s -tls -jia-service-url http://127.0.0.1:4999",
}
