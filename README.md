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
 |  +- github.com/vladimirvivien/go-tutorial
 |  |  +- hello/hello_world.go 
 |  |  +- greetings2/greet.go
 |  +- github.com/vladimirvivien/automi
 |  |  +- operators/batch/funcs.go
 |  |  +- stream/stream.go
 |  + ... (may contain many more source packages) ...
 |
 +- pkg/
 |  +- darwin_amd64
 |     +- github.com/vladimirvivien/go-tutorial/greetings2.a
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
In Go, a package is the unit of code that can be compiled either as an executable `program` or a reusable code `library`.  Packages are directories in the workspace under `$HOME/go/src`.  A package can consist of one or more source files which are compiled into one logical unit.  The membership of a source to a package is declared in the with the `package` directive (discussed later).  All files in a package must declare the same package name or the compiler will not be happy.  

### Package import path and the default name
The `import path` of a package is the unique directory path of the package in the workspace, relative to path `$HOME/go/src` as shown in the following table:

| Workspace Path | Import Path | Default Name
|---|---|---|
|`$HOME/go/src/foo`|`foo`|`foo`|
|`$HOME/go/src/foo/bar`|`foo/bar`|`bar`|
|`$HOME/go/src/foo/bar/bazz`|`foo/bar/bazz`|`bazz`|

The previous table shows examples of workspace paths, their resolved import paths, and their default package name.  These concepts are crucial to Go and its tools.  The import path is used in Go source files when specifying packages to `import` for instance.  The default package name is resolved as the last element of the import path and is used in source files as an identifier for the package.  For instance, the following example uses import path  `foo/bar/bazz` which is assigned package identifier `bazz` to access package elements using  a dot-notation:
```go
package main
import "foo/bar/bazz"
func main() {
    bazz.Blat()
}
...
```
It should be noted that the default package name in a source file can be explicitly specified by preceding the import path with an identifier as shown below:
```go
package main
import ba "foo/bar/bazz"
func main() {
    ba.Blat()
}
...
```
### Naming your packages
An accepted practice in Go is to give the package path a unique name to avoid name collisions.  This is specially important if you plan to distribute your code for others to consume.  The most common approach is to include a unique identifier such as a source code repository and username as part of the path.  Others also use a company name or a project name, when naming the package directory.  

|Import path|Qualifier|
|---|---|
|github.com/vladimirvivien/go-tutorial|github.com/vladimirvivien|
|github.com/stretchr/testify|github.com/stretchr|
|k8s.io/client-go/pkg/api/v1|k8s.io/client-go|
|gopkg.in/yaml.v1/|gopkg.in|

### A Go program
As mentioned earlier,  a program is a package that can be compiled into executable code.  All source files of a program must declare `package main` and at most one source file in the directory must include the special function `main()`.  

For instance, the following shows the layout for a program in directory  `greetings/` :  
```sh 
$HOME/go/src/
 +-github.com/vladimirvivien/go-tutorial/
   +-greetings/
     +-greet.go
``` 

The program package directory is `/greetings` with one source file,  `greet.go`, shown below:  
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
#### Compiling and running the program
We can compile the program package, along with its dependencies, using the `go build` command-line tool by specifying the relative path of the package or its *import path*.  For instance, the following will compile the program in package `greetings`:
```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial/greetings
$ go build .
```
Or,
```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial
$ go build ./greetings
```
In the previous commands are equivalent. They `.` specify the relative directory path of package `greetings` to compile.  By default, the `go build` command creates a binary with the same name as the package.

```sh
> ls -l
total 1520
-rw-rw-r-- 1 vvivien vvivien     582 Jul  8 08:30 greet.go
-rwxrwxr-x 1 vvivien vvivien 1551885 Jul  8 09:05 greetings
```
We can run the greetings program as follows:

```sh
> ./greetings Korean
안녕하세요
```
>Note that we can use the `-o` flag to specify the name of the binary compiled. 

We can also compile the program by specifying its full import path:

```sh
$ go build github.com/vladimirvivien/go-tutorial/greetings
```
We can also use `go install` which is a tool that compiles the package and its dependencies and installs the resulting binary in `$HOME/go/bin`.  This command also caches any dependent packages in the workspace to avoid future unnecessary compilation.

> See `go help build` and `go help install` for detail.
 
### A Go library
Libraries are packages that use the same `go` command tools and are compiled into archive files (instead of executable code) that can be reused by other packages.  To demonstrate a library, we will rewrite the previous greeting program.  In this version, we will extract the greeting functionality and place it into library `greetlib` so that it can be imported by other packages:  
```sh 
$HOME/go/src/
 +-github.com/vladimirvivien/go-tutorial
   +-greetlib
     +-lib.go
   +-greetings2/
     +-greet.go
```
The source code for the library package is stored in source file `greetlib/lib.go` as shown below:
```
package greetlib

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

// GreetIn returns a greeting in specified lang
func GreetIn(lang string) string {
	if greeting, ok := greetings[lang]; ok {
		return greeting
	}
	return greetings["English"]
}
```
As a convention, and to make things easy, the source files of a package declare a package name that matches the directory where they are located.  The previous source snippet, for instance, declares `package greetlib` since the directory where the file is located is called `greetlib`.

The program which uses the library is in package `greetings2`.  The source code in that package imports the library by specifying its import path as `github.com/vladimirvivien/go-tutorial/greetlib` to access its exported code elements.
```go
package main

import (
	"fmt"
	"os"
	
	"github.com/vladimirvivien/go-tutorial/greetlib"
)

func main() {
	lang := "English"
	if len(os.Args) >= 2 {
		lang = os.Args[1]
	}
	fmt.Println(greetlib.GreetIn(lang))
}
```
#### Compiling a library
Compiling a library is done using the `go build` tool by specifying the package's import path or its relative directory path as was done before.  For instance, the following would compile the `greetlib` package:
```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial/greetlib
$ go build . 
```
Or we can ensure that the compiled artifacts is cached in the workspace and is available for other packages by using the `go install` command:
```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial/
$ go install ./greetlib
```
It should be noted that when we compile the program, the `go` tool will properly resolve the dependency to the `greetlib` library and compile that as well.  So, the following will compile and install both the program and its dependent library together:
```
$ cd HOME/go/src/github.com/vladimirvivien/go-tutorial/
$ go install ./greetings2
```

> See `go help build` and `go help install` for detail.

### Package element visibility
Go has a simple rule for element visibility when a package is imported from another:

>*Capitalized identifiers are visible to other packages*

For instance, we can see this in action from the `greetlib` package example above. Variable `greetings` is an identifier with lowercase, therefore it cannot be accessed from other packages.
```
var greetings = map[string]string{
	"English": "Hello, World!",
	"French":  "Salut Monde",
	"Chinese": "世界您好",
	...
```
On the other hand, function `GreetIn()` , shown below is capitalized which means it can be accessed by other packages.
```
func GreetIn(lang string) string {
	if greeting, ok := greetings[lang]; ok {
		return greeting
	}
	return greetings["English"]
}
```
### Remote packages
Go includes the `go get` tool which can retrieve and install packages stored on a remote source control repository server such as Git or Mercurial.  The tool uses the import path to figure out where the file is located on the server.  For instance, the following command will pull and install package `greetings2` from this repository:
```
$ go get github.com/vladimirvivien/go-tutorial/greetings2
```
Since package `greetings2` imports package `github.com/vladimirvivien/go-tutorial/greetlib`,  `go get` will transitively attempt to resolve, download, and install package `greetlib` if it is not found in the workspace.

> see `go help get` for more detail.
