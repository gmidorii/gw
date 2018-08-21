BIN=gw
CMD=./cmd

build:
	cd $(CMD); go build -o $(BIN)

run: build
	$(CMD)/$(BIN) ls -l

run-err: build
	$(CMD)/$(BIN) sh error.sh
