BIN=gw
CMD=./cmd

init:
	dep init

build:
	cd $(CMD); go build -o $(BIN)

build-multi:
	cd $(CMD); gox -output="$(BIN)_{{.OS}}_{{.Arch}}" --osarch "linux/amd64 darwin/amd64 windows/amd64"

run: build
	$(CMD)/$(BIN) ls -l

run-err: build
	$(CMD)/$(BIN) sh error.sh

