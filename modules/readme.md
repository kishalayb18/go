# Go Modules

- One module consists of multiple packages

## Commands to create modules
- run this command `go mod init sample.github.com/hello-module`
- this will create `go.mod` 
- run `go build` command
- this will create the binary file based on the module name you have provided, here `hello-module`
- from terminal run it with `./<binary name>`, here `./hello-module`