package iptablenforcer

import (
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
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

func ReadConfig() {
	config.SetConfigType("yaml")
	config.SetConfigFile("./iptables.config")
	//config.SetConfigName("iptables.config") // name of config file (without extension)
	//config.AddConfigPath(".")               // optionally look for config in the working directory
	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
		log.Errorf("Fatal error config file: %s \n", err)
	}
	//onfig.Get("GUO_OPENSHIFT_INPUT")
	log.Infoln(config.Get("GUO_OPENSHIFT_INPUT"))
}
