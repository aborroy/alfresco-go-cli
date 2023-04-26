EXECUTABLE=alfresco
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_arm64
VERSION=$(shell git describe --tags --always --long --dirty)

all: build

build: windows linux darwin
	@echo version: $(VERSION)

windows: $(WINDOWS) 

linux: $(LINUX) 

darwin: $(DARWIN) 

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./alfresco.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./alfresco.go

$(DARWIN):
	env GOOS=darwin GOARCH=arm64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./alfresco.go

clean: 
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)
