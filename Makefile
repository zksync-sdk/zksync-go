all: tests

download:
	mkdir -p libs
ifeq ($(shell uname -s),Darwin)
	test -f ./libs/libzks-crypto.a || curl -L https://github.com/zksync-sdk/zksync-crypto-c/releases/download/v0.1.1/zks-crypto-macos-x64.a --output ./libs/libzks-crypto.a
else
	test -f ./libs/libzks-crypto.so || curl -L https://github.com/zksync-sdk/zksync-crypto-c/releases/download/v0.1.1/zks-crypto-linux-x64.so --output ./libs/libzks-crypto.so
endif

tests: download
	CGO_LDFLAGS="-L./libs" LD_LIBRARY_PATH="./libs" go test -race -v -count=1 .

integration-test: download
	CGO_LDFLAGS="-L./libs" LD_LIBRARY_PATH="../libs" go test -race -v -count=1 ./IntegrationTests

generate:
	go install github.com/vektra/mockery/cmd/mockery
	go install github.com/golang/mock/mockgen
	go generate .
