# Game Maker Language Go

[![Build Status](https://travis-ci.org/silbinarywolf/gml-go.svg?branch=master)](https://travis-ci.org/silbinarywolf/gml-go)
[![Documentation](https://godoc.org/github.com/silbinarywolf/gml-go?status.svg)](https://github.com/silbinarywolf/gml-go)
[![Report Card](https://goreportcard.com/badge/github.com/silbinarywolf/gml-go)](https://godoc.org/github.com/silbinarywolf/gml-go)

**NOTE: This project is currently undergoing a large refactoring effort to help ease workflow and serialization. I'm also aiming to improve the documentation, add examples and improve test coverage. This is still just a hobby project for now!**

This is a library / framework that aims to create workflow like Game Maker, but utilizing the Go programming language.

## Install

```
go get github.com/silbinarywolf/gml-go
```

## Requirements

* Golang 1.11+

## Documentation

* TODO when this library has been refactored
* [License](LICENSE.md)

# Rough Roadmap

Below are some rough ideas on where I want this project to go. 
This project is mostly for fun and I have no intentions to get anything done unless public interest is shown.

* Implement various frequently used Game Maker functions, ideally with the same names / parameter count.
* Have tutorials for using VSCode and/or Sublime Text.
	- Getting started, Hello world
	 	- This should be easy enough for someone with no knowledge of programming to follow.
		- This should cover installing "gofmt" / "goimports"
	- Transitioning from GML to Golang, major / minor differences
* Add build tools to help with:
	- [x] Auto generating entity IDs and the like.
	- Packing assets into texture atlases

## Credits

* [Hajime Hoshi](https://github.com/hajimehoshi/ebiten) for their fantastically simple 2D game library, [https://github.com/hajimehoshi/ebiten](Ebiten).
* [Yann Le Coroller ](www.yannlecoroller.com) for their free to use Helvetica style font. 