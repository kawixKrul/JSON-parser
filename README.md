## Parser

### Build

To build the `parser` executable, run the following command:

```
make parser
```

This command compiles the source files `src/main.go` and `src/verifier.go` into an executable named `parser`.

### Test

To run the tests for the `parser`, execute:

```
make test
```

This command runs the tests defined in the file `src/verifier_test.go`.

### Clean

To clean up generated files, including the `parser` executable, run:

```
make clean
```

### Usage

The `parser` program accepts two flags:

- `--file`: Use this flag if you want to read JSON from a file. Example usage: `./parser --file filename.json`
- `--json`: Use this flag if you want to directly process JSON. Example usage: `./parser --json '{"key": "value"}'`

If both `--file` and `--json` flags are provided or none of them are provided, the program will display a usage message and exit with an error.

### Example

Here's an example usage of the `parser` program:

```
./parser --file test.json
```

This command reads JSON from the file `test.json`, validates it, and prints the validation result.

```
./parser --json '{"PolicyName": "root", "PolicyDocument": {"Version": "2012-10-17", "Statement": [{"Sid": "IamListAccess", "Effect": "Allow", "Action": ["iam:ListRoles", "iam:ListUsers"], "Resource": "*"}]}}'
```

This command directly processes the provided JSON, validates it, and prints the validation result.

### Note

Input data must be in AWS::IAM::Role Policy format otherwise the program will give you the corresponding error.
