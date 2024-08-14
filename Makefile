.PHONY: build
build:
	mkdir -p bin
	go build -trimpath -o bin/single-dev-env

.PHONY: run
run:
	go run main.go

serve:
	air --build.cmd "go build -o bin/api main.go" --build.bin "./bin/api"