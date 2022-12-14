test:
	$(info ๐งช Testing...)
	go test ./...

format:
	$(info ๐๏ธ formatting...)
	@go fmt ./...


build:
	$(info ๐ฆ Building...)
	go build -o build/ ./...

clean:
	$(info ๐งน Cleaning...)
	rm -rf ./build/* 

run:
	$(info ๐ป Running...)
	go run main.go