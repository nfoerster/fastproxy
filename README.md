## Install

```console
go get github.com/nfoerster/fastproxy
```

## What is Fastproxy?

Fastproxy is an extreme lightweight proxy application for fast routing. 
It supports multiple parallel routing through a straightforward configuration file.

## Why Fastproxy?

Fastproxy is completely written in a few Go lines and is therefore extreme lightweight. 
The executable takes only a few megabytes, the application consumes very low memory and the routing is fast.

## How to use Fastproxy?

Go these steps to use Fastproxy:
* compile the program for linux or windows with the supported Makefile or with go directly: go build cmd/main.go
* define your routing options in the config.yml
* start the program with the associated configuration. E.g. go run cmd/main.go -configPath="./config.yml" // fastproxy -configPath="./config.yml"

## License & Contribution

Please feel free to fork or use it as you like.