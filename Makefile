GO_CMD = go

all: build test

build:
	$(GO_CMD) build -o out/$(service) ./$(service)/cmd/main.go

test:
	$(GO_CMD) test ./...

clean:
	rm -rf out/
