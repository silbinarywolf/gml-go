# Game Maker Language Go

[![Build Status](https://travis-ci.org/silbinarywolf/gml-go.svg?branch=master)](https://travis-ci.org/symbiote/silverstripe-gridfieldextensions)

**NOTE: This is currently a hobby project and not meant for any use other than my own. If you are interested in this or use this, please let me know so I can improve documentation, tagging, etc**

This is a library / framework that aims to create workflow like Game Maker, but utilizing the Go programming language.

## Install

```
go get github.com/silbinarywolf/gml-go
```

## Requirements

* Golang 1.10+

## Documentation

* TODO when the library has progressed
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
	- Packing assets into texture atlases
	- Converting Tiled maps into an internal engine format
	- (maybe) Auto generating entity IDs and the like.

## Credits

* [Hajime Hoshi](https://github.com/hajimehoshi/ebiten) for his fantastically simple 2D game library, [https://github.com/hajimehoshi/ebiten](Ebiten).
