win.simple:
	go build -o ./bin/winsimple ./examples/winsimple/main.go
	cd ./bin && ./winsimple
fmt:
	go fmt ./...
