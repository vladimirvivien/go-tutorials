# Getting Started With Go 
This document is a collection of topics designed to help newcomers to get started with the Go proramming language.  All code samples, discussed in the sections below, can be found in this repository.

## About Go
Go was created as system language at Google in 2007 primarily by Robert Griesemer, Rob Pike, Ken Thomson.  The language was an answer to handle the needs of application development at Google-scale. The designers of Go wanted to mitigate issues such as slow build cycle, language complexity, magical syntax while creating a new language that is simple, safe, consistent, and predictable. 

> “Go is an attempt to combine the safety and performance of a statically-typed language with the expressiveness and convenience of a dynamically-typed interpreted language.” -- Rob Pike

- BSD-style licensed open source language
- It borrows ideas from many languages before it
- Go has a simple and concise syntax (that slightly resemble C)
- Short mortal-decipherable spec document
- Strongly and statically typed for safety
- Type declaration with dynamic language feel
- Support for objects (not class-based)
- Near-zero compilation time with script language-like performance
- Compiled to native binaries for fast runtime execution
- Easy cross compilation for popular OS and architectures
- Uses channels with send-receive idiom
- Simple concurrency idioms

## Where to start
There are several ways to start with the Go programming language

- Download Go tools from [golang.org/dl](https://golang.org/dl/) [download and install the Go toolchain.  Configure your local environment and use your own editor to craft sweet Go code]
- Try it out at the Playground ([play.golang.org](https://play.golang.org)) [ for instant gratification, you can try this browser-based IDE that let you run your own code to test and share ideas]
- Take the Go Tour at [tour.golang.org](https://tour.golang.org) [ a step-by-step exploration of the features of the language with browser-based editor]

### Supported OSs / Architectures
The Go compiler can target 8 architectures: 

- arm
- arm64
- mips
- mips64
- ppc64
- s390
- x86
- x86_64

The Go compiler can target 9 operating systems:

- DragonFly BSD
- FreeBSD
- Linux
- NetBSD
- OpenBSD
- OS X
- Plan 9
- Solaris
- Windows

This means you can create your program from the comfort your Mac running OS X on a x86_64 and compile that code to run on Linux running on an ARM6.

### Installing Go
Go can be installed on your local machine in either of the following ways:
- From archive files - for Linux, Mac OS, FreeBSD, and Windows
- Installer - available for Mac OS and Windows
- Build from source 
Use URL [https://golang.org/doc/install](https://golang.org/doc/install) to follow instruction on your preferred method of installation for your target environment.

## Setup your workspace
Before you get introduced to your first Go program, let us take a moment to discuss your environment for local development.  After you install your Go toolchain, it is crucial that you properly setup your workspace.  It is an arbitrary directory which stores your Go source files, organized as packages, and built artifacts (such as object files and executable binaries) as shown in the following sample workspace:
```sh
$HOME/go
 +- src/
 |  +- github.com/vladimirvivien/getting-started-with-go
 |  |  +- ex01/hello.go 
 |  |  +- ex02/mettaloids.go
 |  +- github.com/vladimirvivien/automi
 |  |  +- operators/batch/funcs.go
 |  |  +- stream/stream.go
 |  + ... (may contain many more source packages) ...
 |
 +- pkg/
 |  +- darwin_amd64
 |     +- github.com/vladimirvivien/getting-started-with-go/ex01.a
 |     +- github.com/vladimirvivien/getting-started-with-go/ex02.a
 |     +- github.com/vladimirvivien/automi/operators/batch.a
 |     +- github.com/vladimirvivien/automi/stream.a
 |     +- ... (each package compiled as an object file) ...
 |
 +- bin/
    +- hello
    +- ... (compiled programs stored here) ...
```
### Setting GOPATH
By convention, the Go toolchain assumes the workspace is located at `$HOME/go`. If different, you will need to set env variable `$GOPATH` pointing to the workspace location.  You can use the `go env` command to verify the location of your `GOPATH` as shown below.
```sh
go env GOPATH
/Users/vvivien/DEV/go
```

## Editors and IDE
Because of its minimal syntax, working with Go can be done using your favorite text editor.  Some editors include features such as syntax higlights while others have deep integrations with the Go toolchain to provide an IDE-like experience.  Some of the more well-known editors with Go support include:
- Atom
- Emacs
- Sublime Text
- TextWrangler
- Vim/Neovim
- ... and more

If you feel more comfortable with working with a full IDE, you will be well-served there as well.  There are several IDE's providing full support for Go including the followings:
- IntelliJ Gogland - Go IDE
- Visual Studio Code - Support with vscode-go plugin
- GoClipse - Eclipse based Go IDE

## Your first Go program
To get started with Go, let us write the obligatory *Hello World* program.  For now, create directory  `$HOME/go/src/hello` .  Later we will have a full discussion on code organization and Go packages.
```
mkdir -p $HOME/go/src/hello
```
Next create and save the following code in a file called  `hello_world.go`.
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
After the source code is saved, we can compile and run it with the `go run` command as shown in the following example:
```sh
$> go run hello_world.go
Hello, World!
```
The `go run` command, shown above, is a convenience tool that compiles of the source code and immediately runs the resulting program.   This command can be handy when prototyping ideas or running a simple Go program.  As we will see later, there are other, more suitable, ways to compile your Go code for distribution.

## The Go source file
It is important to understand the make up of a Go source file.  The following figure highlights the major attributes of a Go source file.

![Go Source](go-source.png)

#### `package`
All Go source files must start with the he `package` directive.  It declares the name of a package to which the source in the file belongs.  In the source snippet above, the source belongs to package `main` which is a reserved name for packages that get compiled to runnable binaries.  

#### `import`
The `import` statement specifies an  `import path` for packages that are used in the source code.  Your code can import from the standard library or other user-provided packages that are in the workspace.  The code above imports package `fmt` so that it can invoke function `fmt.Println`.

#### `func main()`
Go functions are declared using the `func` keyword.  Function `main` is a special function that is used as the entry point of an executable program.  Function `main` must be defined in the main package.

#### File `hello_world.go`
Go source files can have arbitrary names followed by the `.go` extension.  There is no relations between the code elements in the source file and its name.  However, by convention, the file is usually named something meaningful that is usually kept short and use the underscore (`_`) as word separator.

#### // Comment
Go uses C-style comments which are used by Go tools to generate documentation automatically.

#### Optional semi-colon
One more thing that is notable in Go sources is the lack of semi-colon.  While the compiler understands and can parse source code with semi-column, they are almost always omitted in idiomatic Go.

## Packages

### Building your code
When you want to build your Go project for distribution,  you will use command `go build` which compiles your Go source files into a binary.
```
> go build hello_world.go
```
The previous command will compile the specified source 

Command `go install` which compiles and installs (in `$GOPATH/bin`) your Go programs.  For instance the following command will compile and install 

```
go install hello_world.go
```
However, for program releases, you would use the Go compiler to properly compile your binaries for distribution.
### The Go program source
 While the program 

 ## Go Packages

##The tools
