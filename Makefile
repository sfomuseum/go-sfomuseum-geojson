fmt:
	go fmt *.go
	go fmt feature/*.go
	go fmt properties/*.go
	go fmt properties/*/*.go

tools:
	go build -o bin/depicts cmd/depicts/main.go
