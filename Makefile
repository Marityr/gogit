PHONY: build
build:
	rm -rf build && mkdir build && go build -o build/gogit -v ./


.PHONY: run
run:
	go run main.go
