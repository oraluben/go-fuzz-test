all: build

deps:
	go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

build: deps
	cd fuzz && go-fuzz-build

ast-demo:
	go run ast/main.go