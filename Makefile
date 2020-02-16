# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=fastproxy

linux: 
		GOOS=linux $(GOBUILD) -o build/$(BINARY_NAME) -v cmd/main.go 
windows64: 
		GOOS=windows GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME).exe -v cmd/main.go 
windows32: 
		GOOS=windows GOARCH=386 $(GOBUILD) -o build/$(BINARY_NAME).exe -v cmd/main.go 
clean: 
		$(GOCLEAN)
		rm -f build