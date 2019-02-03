# Game Maker Language Go

[![Build Status](https://travis-ci.org/silbinarywolf/gml-go.svg?branch=master)](https://travis-ci.org/silbinarywolf/gml-go)
[![Documentation](https://godoc.org/github.com/silbinarywolf/gml-go?status.svg)](https://godoc.org/github.com/silbinarywolf/gml-go)
[![Report Card](https://goreportcard.com/badge/github.com/silbinarywolf/gml-go)](https://goreportcard.com/report/github.com/silbinarywolf/gml-go)

**NOTE: This project is currently undergoing a large refactoring effort to help ease workflow and serialization. I'm also aiming to improve the documentation, add examples and improve test coverage. This is still just a hobby project for now!**

This is an engine that aims to strike a balance between capturing the simplicity of the Game Maker API whilst not losing any performance given to you by Go. It has been designed with multiplayer games in mind and differs itself from Game Maker by allowing you to "run" multiple rooms at once so that every player does not have to be in the same room.

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

* [Hajime Hoshi](https://github.com/hajimehoshi/ebiten) for their fantastically simple 2D game library, [Ebiten](https://github.com/hajimehoshi/ebiten).
* [Yann Le Coroller](http://www.yannlecoroller.com) for their free to use Helvetica style font.
* [milkroscope](https://www.artstation.com/milkroscope) for their artwork on Worm In The Pipes (example/worm)
