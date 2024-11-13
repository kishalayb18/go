# Go Modules

- One module consists of multiple packages

## Commands to create modules
- run this command `go mod init sample.github.com/hello-module`
- this will create `go.mod` 
- run `go build` command
- this will create the binary file based on the module name you have provided, here `hello-module`
- from terminal run it with `./<binary name>`, here `./hello-module`

## Why "main" package name is important
The `main` package name tells Go that this is the entry point of the program, if Go does not find a package name called **_main_** then during `go build` it will not generate any binary file

## Why "main" function is important
- Go will call and run the **_main_** function when the application starts
- There should only one _main_ function, not more than that

