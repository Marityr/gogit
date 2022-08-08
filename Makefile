PHONY: build
build:
	rm -rf build && mkdir build && go build -o build/app -v ./
  
.PHONY: run
run:
	go run main.go
