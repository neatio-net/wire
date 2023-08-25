.PHONY: all test get_deps

all: test install

install: get_deps
	go install github.com/nio-net/wire/cmd/...

test:
	go test --race github.com/nio-net/wire/...

get_deps:
	go get -d github.com/nio-net/wire/...

pigeon:
	pigeon -o expr/expr.go expr/expr.peg
