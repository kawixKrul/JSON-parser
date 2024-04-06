parser: src/main.go src/verifier.go
	go build -o parser ./src

test: parser
	go test ./src

clean:
	rm parser