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
	$(info 🧹 Cleaning...)
	rm -rf ./build/* 

run:
	$(info 💻 Running...)
	go run main.go