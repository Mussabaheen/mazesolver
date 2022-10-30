test:
	$(info 🧪 Testing...)
	go test ./...

format:
	$(info 🖊️ formatting...)
	@go fmt ./...


build:
	$(info 📦 Building...)
	go build -o build/ ./...

clean:
	rm -rf ./build/* 
