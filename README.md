# bypass
[![Build Status](https://travis-ci.org/gsquire/bypass.svg?branch=master)](https://travis-ci.org/gsquire/bypass)
[![GoDoc](https://godoc.org/github.com/gsquire/bypass?status.svg)](https://godoc.org/github.com/gsquire/bypass)

This is trivial middleware for Go in which it allows you omit middleware chains for
certain routes that don't require all of the handlers you are using.

### Install
You can always run `go get` or use some package manager for Go like `gb` or `glide` to use it.
Bypass plays nicely with [Alice](https://github.com/justinas/alice) as well.

### License
MIT
