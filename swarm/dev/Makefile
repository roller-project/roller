.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    /Volumes/DATA/node_project/coins/cmd/swarm
	go build -o bin/geth     /Volumes/DATA/node_project/coins/cmd/geth
	go build -o bin/bootnode /Volumes/DATA/node_project/coins/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v /Volumes/DATA/node_project/coins/swarm/...
