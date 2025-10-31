# Learn Go

## *CHAPTER 1*

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

- The `go.mod` file declares the name of the module
- the minimum supported version of
  Go for the module,
- and any other modules that your module depends on.
- the `go get` and `go mod tidy`
  commands to manage changes to the file.

It as being similar to the `requirements.txt` file used by Python.

## *CHAPTER 2*

**The Predeclared Types:** Go has many types built into the language, called predeclared types.They
are similar to types that are found in other languages: `booleans, integers, floats, and strings`.

**The Zero Value:** Go assigns a default zero value to any variable that is declared but not assigned a value.

**Literals**

A Go literal is an explicitly specified number, character, or string. 

Basic rules for all literals:
-In Go single quotes and double quotes are not interchangeable.
- different prefixes are used to indicate other bases
[`0b for binary (base 2), 0o for octal
(base 8), or 0x for hexadecimal (base 16)`]
- can use either upper- or lowercase letters for the prefix
- To easier to read put underscores in the middle of literal. `1234`-> `1_2_3_4/1_234`

**integer literal**  
- is a sequence of numbers.
- base 10 by default

**floating-point literal**
- can define an exponent number with letter `e`. (such as `6.03e23`)
- for 0x prefix,letter p for indicating any exponent (`0x12.34p5`)

**rune literal**
- a character and is surrounded by `single` quotes.
- can be written as,
  - single Unicode characters ('a')
  - 8-bit octal numbers ('\141')
  - 8-bit hexadecimal numbers ('\x61')
  - 16-bit hexadecimal numbers('\u0061')
  - 32-bit Unicode numbers ('\U00000061').
- `Avoid using any of the `**numeric escapes**` for rune literals, unless the context makes your code clearer.`
