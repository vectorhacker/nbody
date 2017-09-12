GO = go

buildflags = -a -ldflags '-extldflags "-static"'

build: main.go
	export CGO_ENABLED=0
	$(GO) build $(buildflags) -o bodies main.go