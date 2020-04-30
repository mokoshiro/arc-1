
.PHONY: release
release:
	GOOS=linux GOARCH=amd64 go build -o bin/txm_linux_amd64
.PHONY: build
build:
	go build -o txm .
.PHONY: run-executor
run-executor: build
	./txm executor --config ./executor/executor.yaml
