package pkgutil

import "github.com/dimasbayuseno/cisdi-go-test/config"

func GetPort(ports ...string) string {
	if len(ports) > 0 {
		return ":" + ports[0]
	}
	port := config.Get().HttpPort
	if port != "" {
		return ":" + port
	}
	return ":8888"
}
