# go4j
Golang for Java

# Building

cd hello; go build; ./hello

cd greetings; go test

## Running the tests

We can see the test coverage by running the following command:

    go test -cover

The output is very basic and is just a summary


For more detailed information, use these commands to save the coverage into a file

    go test -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out

In the first command, we use -coverprofile to save coverage results to the file. And then, we print detailed results by using Go's cover tool.

By using the same cover tool, we can also view coverage result as an HTML page:

    go test -coverprofile=coverage.out ./...; go tool cover -html=coverage.out


# Cleaning

This command will recursively remove all installed object files and binaries in the current directory and all its subdirectories.

go clean -i ./...

# To update the go version

- update the go version in each go.mod file
- run go mod tidy
- update the golang docker container in my-workflow.yaml