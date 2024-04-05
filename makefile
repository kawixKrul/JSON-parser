parser: src/main.go src/verifier.go
	go build -o parser ./src

test: parser
	go test ./src

	./parser --file 'tests/test.json'
	./parser --json '{"PolicyName": "root", "PolicyDocument": {"Version": "2012-10-17","Statement": [{"Sid": "IamListAccess","Effect": "Allow","Action": ["iam:ListRoles","iam:ListUsers"],"Resource": "*"}]}}'

clean:
	rm parser