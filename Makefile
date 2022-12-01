build: 
	go build -o .bin/main cmd/web/main.go

run: build
	.bin/main