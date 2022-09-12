# Go redirector for Arcaflow

This repository hosts the `go.flow.arcalot.io` redirector, which handles the Golang package paths. To create a new package, simply amend the [packages.json](packages.json) file.

## Developing the redirector

The redirector is located in [generate.go](generate.go) and you can run the test suite in [generate_test.go](generate_test.go) by simply running `go test` in the project directory.