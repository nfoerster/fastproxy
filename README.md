# What is Fastproxy?

Fastproxy is an extreme lightweight proxy application for fast routing. 
It supports multiple parallel routing through a straightforward configuration file.

# Why Fastproxy?

Fastproxy is completely written in a few Go lines and is therefore extreme lightweight. 
The executable takes only a few megabytes, the application consumes very low memory and the routing is fast.

# Install

```go
go get github.com/nfoerster/fastproxy
```

# How to use Fastproxy?

Go these steps to use Fastproxy:
* compile the program for linux or windows with the supported Makefile or with go directly: 
```go
go build cmd/main.go
```
* define your routing options in the config.yml
* start the program with the associated configuration.
```go
go run cmd/main.go -configPath="./config.yml" // fastproxy -configPath="./config.yml"
```

# GO usage

## Import 
```go
import (
    "github.com/nfoerster/fastproxy
)
```
## Usage example 
```go
fastproxy.GoProxy("127.0.0.1", 3128, 192.168.2.12, 6666)
```
