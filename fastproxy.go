package fastproxy

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/spf13/viper"
)

//ReadConfig reads the config from given path
func ReadConfig(configPath string) *Config {
	var config Config

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

//StartProxies starts all proxies for the given configuration
func StartProxies(config *Config) {
	for _, proxy := range config.Proxies {
		go GoProxy(proxy.FromHost,
			proxy.FromPort,
			proxy.ToHost,
			proxy.ToPort)
	}
}

func handleConnection(sourceConn net.Conn, destConn net.Conn) {
	go func() {
		defer destConn.Close()
		defer sourceConn.Close()
		io.Copy(destConn, sourceConn)
	}()
	go func() {
		defer destConn.Close()
		defer sourceConn.Close()
		io.Copy(sourceConn, destConn)
	}()
}

func GoProxy(fromhost string, fromport int, tohost string, toport int) {
	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", fromhost, fromport))
	if err != nil {
		log.Fatal(err)
	}

	for {
		sourceConn, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}

		destConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", tohost, toport))
		if err != nil {
			log.Fatal(err)
		}

		handleConnection(sourceConn, destConn)
	}
}
