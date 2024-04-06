parser: src/main.go src/verifier.go
	go build -o parser ./src

test: src/verifier_test.go
	go test ./src

clean:
	rm parser