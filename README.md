# Learn Go

## chapter 1:
### installation
Download msi/pkg file from official go website based on the architecture.
Check the go environment by running this command in bash/cmd.

```go
>> go version
go version go1.25.3 windows/amd64

>>where go
C:\Program Files\Go\bin\go.exe
```

### Go Tooling
All of the Go development tools are accessed via the go command. 

- compiler (**go build**)
- code formatter (**go fmt**)
- dependency manager (**go mod**)
- test runner (**go test**)
- a tool that scans for common coding mistakes (**go vet**), and more

### Making a Go Module
The first thing need to do is create a directory to hold program.
```go
>> mkdir chapter1
>> cd chapter1
```
Inside the directory, run the go mod init command to mark this directory as a Go
module.Go project is called a module.:
```go
>> go mod init module_name
go: creating new go.mod: module module_name
```

A module is not just source code. It is
also an exact specification of the dependencies of the code within the module. Every
module has a `go.mod` file in its root directory. Running `go mod init` creates this file. 

*`go.mod`* file look like this:
```go
module module_name

go 1.25.3
```






