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
By convention (and for the remainder of this discussion), the Go toolchain assumes the workspace is located at `$HOME/go`. If different, you will need to set env variable `$GOPATH` pointing to the workspace location.  You can use the `go env` command to verify the location of your `GOPATH` as shown below.
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
The `go run` command, shown above, is a convenience tool that compiles of the source code and immediately runs the resulting program.   This command can be handy when prototyping ideas or running a simple Go program.  As we will see later though, there are other and more suitable ways to compile your Go code and packages for distribution.

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

#### `// Comment`
Go uses C-style comments which are used by Go tools to generate documentation automatically.

#### Optional semi-colon
One more thing that is notable in Go sources is the lack of semi-colon.  In idiomatic Go, semi-colons are optional and are always omitted.  However, the Go compiler inserts them during compilation as they are required by Go's formal grammar.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 

## Packages
All Go source files must belong to a package which is simply a directory where the files are saved in the workspace under `$HOME/go/src`.  It should be made clear that all files in a directory must declare the same package or the compiler will not be happy.  



There are two types of Go packages: `programs` which compile to executables and `libraries` that are reusable assets in the workspace.

### A program package
As mentioned earlier, a program is a package where all source files declare `package main` and at most one file include the special function `main()`.  For instance, the following shows the layout for a program in directory  `greetings/`  as its package.  
```sh 
$HOME/go
 +-src/
   +-greetings/
     +-greet.go
``` 

Source file `greet.go` is shown below declare `package main` and includes declaration `func main()`.  This implies that package `greetings` will be compiled as an executable program which, when compiled, prints `"Hello World"` in a language specified at the command-line.  
> If you don't understand this code, that is OK, we will cover more as we go.
```go
package main

import (
	"fmt"
	"os"
)

var greetings = map[string]string{
	"English": "Hello, World!",
	"French":  "Salut Monde",
	"Chinese": "世界您好",
	"Klingon": "qo' vIvan",
	"Hindi":   "हैलो वर्ल्ड",
	"Korean":  "안녕하세요",
	"Russian": "привет мир",
	"Swahili": "Wapendwa Dunia",
	"Spanish": "Hola Mundo",
	"Turkish": "Merhaba Dünya",
}

func main() {
	lang := "English"
	if len(os.Args) >= 2 {
		lang = os.Args[1]
	}
	if greeting, ok := greetings[lang]; ok {
		fmt.Println(greeting)
	} else {
		fmt.Println(greetings["English"])
	}
}
```
### Compile the package
The Go compiler builds `packages` and uses the command `go build <package dir | import path>`.   One way to easily build a package is to `cd` into it and run the `go build .` command as shown below:

```sh
> cd greetings
> go build .
```
In the previous command, the `.` specifies the current directory as the package.  By default, the `go build` command creates a binary with the same name as the package directory.

```sh
> ls -l
total 1520
-rw-rw-r-- 1 vladimir vladimir     582 Jul  8 08:30 greet.go
-rwxrwxr-x 1 vladimir vladimir 1551885 Jul  8 09:05 greetings
```
We can run the greetings program as follows:

```sh
> ./greetings Korean
안녕하세요
```
The name of the binary output can be controlled with the `-o` flag.  The following will build the program and output an executable binary file named `worldgreet`:

```sh
> go build -o worldgreet .
```
We can also provide a relative directory for the package name when building it:

```sh
> go build -o worldgreet ./greetings
```
### Package installation
Your package, along with its dependencies, can also be compiled and *installed* in your workspace.  If the package is program, the resulting binaries are copied to path `$HOME/go/bin`.  For instance, the following installs program `worldgreet`:

```sh
> go install -o worldgreet ./greetings
```
It is recommended practice to add `$HOME/go/bin` to your local system $PATH to make your compiled binary available.

### The import path
One crucial concept to understand in Go is the `import path`.  It is the relative path of each package in the workspace path `$HOME/go/src` which constitutes an identifier that uniquely identify a package.	To drive the concept home, the following shows workspace paths and their and their *import path* values:

| Workspace Path | Import Path |
|---|---|
|`$HOME/go/src/foo`|`foo`|
|`$HOME/go/src/foo/bar`|`foo/bar`|
|`$HOME/go/src/foo/bar/bazz`|`foo/bar/bazz`|

In the previous example, for instance, the import path forth for workspace directory `$HOME/go/srg/greetings` is, as you would guess,  `greetings`.

A common, and accepted, practice in Go is to include the path of a source code repository and the user name of the repository as part of the path.  The following shows this with an example of an import path that uses GitHub repository path `github.com/vladimirvivien/automi`:
```
$HOME/go/src
 +-github.com/vladimirvivien/automi
```
> This is a completely arbitrary practice.  The import path for your Go code does not have to be stored in a source code repository (local or remote) for the code to compile properly.

### A library package
Libraries, on the other hand, by convention uses the name of the directory as their package name.  For instance, in the example earlier, package `hello` compiles to a program because file `hello_world.go` starts with package declaration `package main`.
