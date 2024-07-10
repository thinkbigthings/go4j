# go4j
Golang for Java

# Building

cd hello; go build; ./hello

cd greetings; go test

# Cleaning

This command will recursively remove all installed object files and binaries in the current directory and all its subdirectories.

go clean -i ./...

# To update the go version

- update the go version in each go.mod file
- run go mod tidy
- update the golang docker container in my-workflow.yaml