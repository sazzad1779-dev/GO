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

## Literals

A Go literal is an explicitly specified number, character, or string. 

Basic rules for all literals:
-In Go single quotes and double quotes are not interchangeable.
- different prefixes are used to indicate other bases
[`0b for binary (base 2), 0o for octal
(base 8), or 0x for hexadecimal (base 16)`]
- can use either upper- or lowercase letters for the prefix
- To easier to read put underscores in the middle of literal. `1234`-> `1_2_3_4/1_234`
- Literals are considered `untyped`.

**integer literal**  
- is a sequence of numbers.
- base 10 by default

**floating-point literal**
- can define an exponent number with letter `e`. (such as `6.03e23`)
- for 0x prefix,letter p for indicating any exponent (`0x12.34p5`)

**rune literal**
- a character and is surrounded by `single` quotes.
- rune type is an alias for the int32 type
- can be written as,
  - single Unicode characters ('a')
  - 8-bit octal numbers ('\141')
  - 8-bit hexadecimal numbers ('\x61')
  - 16-bit hexadecimal numbers('\u0061')
  - 32-bit Unicode numbers ('\U00000061').
- `Avoid using any of the `**numeric escapes**` for rune literals, unless the context makes your code clearer.`
- There are also several
backslash-escaped rune literals, with the most useful ones being newline ('\n'), tab('\t'), single quote ('\''), and backslash ('\\').

**string literals**

- double quotes `""`
- The only characters that cannot appear are unescaped backslashes `(\)`, unescaped newlines`(\n)`, and unescaped double quotes`("")`.
- The zero value for a string is the empty string.
- concatenated by using the + operator.

**raw string literal**

- delimited with backquotes ``(`)``
- can contain any character except a backquote.
- can write a multiline, all characters are included as is.

- ``Greetings and``\
  ``"Salutations"``

## Booleans

- Variables of bool type can have one of two values: true or false.
- The zero value for a bool is false.

## Numeric Types

### 1. Integer types
Go provides both signed and unsigned integers in a variety of sizes, from one to eight bytes.


*signed numbers*: **`-(n/2) to +(n/2)`**\
*unsigned numbers*: **`0 to n`**(the values are always 0 or positive).

Integer numbers size variation:
- 8bits: int8 /uint8
- 16bits: int8 /uint8
- 32bits: int32 /uint32
- 64bits: int64 /uint64

Note:
- byte is an alias for uint8

### 2. Floating-point types
Floating-point numbers size variation: `float32` / `float64`

Note:
- floats aren’t exact,A floating-point number cannot represent a decimal value exactly.
- Do not use them to represent money or any other value that must have an exact decimal representation!
- a `nonzero` floating-point variable by `0` returns `+Inf or -Inf`.
- Dividing  0 by 0 returns NaN (Not a Number).
- don’t do == and != to compare floats, because of the inexact
nature.

### Explicit Type Conversion

- Go doesn’t allow automatic type promotion between variables.




