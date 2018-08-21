BIN=gwrap

build:
	go build -o $(BIN)

run: build
	./$(BIN) ls -l
