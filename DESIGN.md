# Design

## What is the purpose of this project?

The primary motivation for creating this project was to build a set of reliable tools that feel like Game Maker but with a focus on making local/online multiplayer games easily, such as executing the logic of multiple rooms at once for cases where two players might be in two seperate areas.

I chose Go as my language of choice for a few reasons.

- It has incredibly fast compile times which is important for game design iteration.
- Its compiler and libraries are open-source, which means in 10+ years I’ll be able to actually compile my projects without a VM, unlike my old Game Maker 5.X and 6.X projects.
- Code is easy to read
- Testing and benchmarking tools are fantastic
- Cross-compilation allows me to build headless Linux builds and easily spin up cheap servers for my networked projects.

## What is this project NOT trying to be?

This project is not trying to be a direct replacement for Game Maker or to even be an open-source alternative to Game Maker. After getting feedback from a few friends, I’ve come to the conclusion that Game Maker’s visual IDE and object event system is a core part of what makes Game Maker so fantastic and easy to use.

Game Maker hides a lot of what it does for you under the hood, which I think is fantastic for beginners as it stops someone from being overwhelmed with a lot of code that they wouldn’t yet be able to read. However, I don’t want any magic or hidden behaviour in my framework. I want the code to be explicit and predictable so that someone who understands the fundamentals of software engineering can easily understand what the engine is doing.

The fuzzy rule of thumb is, if you peek under the hood and see what the code is doing, will you feel good about it? Or will you think it’s overly complex and hard to reason about? I want to lean towards the former as much as possible.

## What is the history of GML-Go?

- GML-Go was created with the intention of creating a game framework that:
  - Strikes a balance between feeling like Game Maker and Go, but leans more on the latter. The aim is to simplify the API as much as honestly as possible without costing the developer large performance problems. This means code might not “look pretty” and terseness will be sacrificed in some cases.
- Allows for multiple rooms to be simulated at the same time. 
  - Game Maker only has the concept of there being 1 room running at a time which makes sense for single player experiences.  However the purpose of this feature is to enable the ability to create networked games, where two or more rooms might need to be run by a server. It also opens up room for split-screen games that takes place over multiple rooms.
- Runs in the browser, and does it well. 
  - When exporting HTML5 games with Game Maker, I’ve found that performance can be pretty poor and that playing looping music hitches the browser. I’m hoping this is simply a problem with Game Maker’s HTML5 code output and not just with Chrome and Firefox.
- Is free, not a blackbox and can be preserved. 
  - When recovering old Game Maker projects from the 5.0 to 8.1 era, I’ve found that those files are impossible to playback on my Windows 10 machine. I want my games to be buildable into the distant future.
- Is strongly-typed. 
  - Typos cannot cause crashes, the game will perform incredibly fast as it’s natively compiled and in theory, I should be able to build tools similar to Go’s fix tool to update my old games from old versions of the framework to new versions.
- Allows for easy serialization of object instances, possibly through code generation. 
  - By having native mechanisms for serialization, programmers can easily implement time-rewind mechanisms, lag-compensation for networked games, write out the state of the game world to a file and more.

## What are the differences with Game Maker?