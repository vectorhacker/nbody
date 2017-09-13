GO = go

buildflags = -a -ldflags '-extldflags "-static"'

BINARY_NAME ?= bodies
ADVANCEMENTS ?= 50000000

build:
	@echo "building..."
	@export CGO_ENABLED=0
	@$(GO) build $(buildflags) -o ./bin/$(BINARY_NAME) main.go

bench: build 
	@echo "benchmarking..."
	@time ./bin/$(BINARY_NAME) $(ADVANCEMENTS) >> /dev/null


run: build
	@echo "running..."
	@./bin/$(BINARY_NAME) $(ADVANCEMENTS)

clean:
	@echo "cleaning..."
	@if [ -a bin ]; then rm -rf bin; fi;
	@echo "done"