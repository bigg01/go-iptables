package main

import (
	"github.com/bigg01/ocp-iptables/pkg/iptablenforcer"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	//&log.TextFormatter{
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Infof("Start Iptables enforcer done...")
	iptablenforcer.ApplRules()

}
