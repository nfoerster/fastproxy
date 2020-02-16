package main

import (
	"flag"
	"os"

	"github.com/nfoerster/fastproxy"
)

func main() {

	var configPath string

	flag.StringVar(&configPath, "configPath", "./config.yml", "the path to config.yml")
	flag.Parse()

	var c chan os.Signal
	c = make(chan os.Signal, 1)

	config := fastproxy.ReadConfig(configPath)
	fastproxy.StartProxies(config)

	<-c
}
