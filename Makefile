export

.PHONY: build

build:
	@echo "Building plugin..."
	@go build -buildmode=plugin -o build/hello-plugin.so -ldflags="-s -w" -gcflags="all=-N -l" src/main.go