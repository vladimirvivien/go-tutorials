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

## Supported OSs / Architectures
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

This means you can create an application from the comfort your Mac running OS X on a x86_64 and compile that application to run on Linux running on an ARM6. 

## Setup your workspace
Before you get introduced to your first Go program, let us take a moment to discuss your environment for local development.  After you install your Go toolchain, it is crucial that you properly setup your workspace.  It is an arbitrary directory which stores your Go source files, organized as packages, and built artifacts (such as object files and executable binaries) as shown in the following sample workspace:
```shell
$HOME/go
 +- src/
 |  +- github.com/vladimirvivien/getting-started-with-go
 |  |  +- ex01/hello.go 
 |  |  +- ex02/mettaloids.go
 |  +- github.com/vladimirvivien/automi
 |  |  +- operators/batch/funcs.go
 |  |  +- stream/stream.go
 |  |
 |  + ... (may contain many more source packages) ...
 +- pkg/
 |  +- darwin_amd64
 |     +- github.com/vladimirvivien/getting-started-with-go/ex01.a
 |     +- github.com/vladimirvivien/getting-started-with-go/ex02.a
 |     +- github.com/vladimirvivien/automi/operators/batch.a
 |     +- github.com/vladimirvivien/automi/stream.a
 |     +- ... (each package compiled as an object file) ...
 +- bin/
    +- hello
    +- ... (compiled programs stored here) ...
```
By convention, the Go toolchain assumes the workspace is located at `$HOME/go` (if different, you will need to set env variable `$GOPATH` pointing to the workspace location).

## Your first Go program
