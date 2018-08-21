BIN=gw

build:
	go build -o $(BIN)

run: build
	./$(BIN) ls -l

run-err: build
	./$(BIN) sh error.sh
