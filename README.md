![main](https://github.com/Diogenesoftoronto/tichu/actions/workflows/main.yml/badge.svg)
![run](https://github.com/Diogenesoftoronto/tichu/actions/workflows/run.yml/badge.svg)
![test](https://github.com/Diogenesoftoronto/tichu/actions/workflows/test.yml/badge.svg)



# Tichu
Tichu is a card game. This is a game written in go for playing tichu on the command line! 🎇 🀄
## Description

It is played with a standard 52-card deck, plus 4 jokers.
Each player is dealt 14 cards, and the remaining cards are placed face down in the middle of the table.
The object of the game is to win tricks containing cards of higher value than those played by the other players.
Each trick is won by the player who played the highest-ranking card of the suit led.
Players may also play special cards called Tichus or Grand Tichus to win tricks without playing the highest card.
At the end of the game, the player with the most points wins.
Points are awarded for winning tricks and for holding certain cards in hand at the end of the game.

## Rules

The rules of tichu are a bit complex but can be found [here](TICHU_RULES.md).

## Motivation

I wanted to create a game with Go and the Bubbletea library for my friends to play if they wanted.
I also wanted to try seeing if I could use this commandline tool in a larger project I am working on with GDscript and [Godot](https://www.godotengine.org).
This is also my first time using the Bubbletea library so I wanted to see how it worked in a smaller project.

## Features

- [x] Play tichu on the command line
- [x] Play with 4 players

## Dependencies

- [task*](https://taskfile.dev/#/installation)
- [go](https://golang.org/doc/install)
- [docker*](https://www.docker.com)

Note: * -> denotes a dev or optional dependency.
## Installation

### Quick Installation
You can install this program via the go package manager:

```bash
$ go install github.com/Diogenesoftoronto/tichu
```

### Installation from source

Install this package by running: 
1. First clone the repository
```bash
$ git clone https://github.com/Diogenesoftoronto/tichu
```
2. Then cd into the directory it was installed in

```bash
$ cd tichu
```
3. Then compile the program
```bash
$ go build -o ./bin/
```
you can also now compile the program with the following command:
```bash
$ task
```
More compilation options can be found in the [taskfile](Taskfile.yml)

Keep in mind that you can only do this if you have the task package installed.

### Installation with Docker

Docker is the best way to install this program cross-platform.

1. Build the docker image

``` bash
docker build --tag tichu <path/to/project/root>
```

2. Run the image!

``` bash
docker run tichu
```
That's all you need to do! Obviously this requires having docker already installed.
## Configuration

You can find configuration for tichu in the .tichu file of the directory you installed it in.

e.g if you installed it in your home directory you can find the configuration file in ~/.tichu.conf/...

## Usage

To begin the game you can run the following command:

```bash
$ tichu play
```

## Contributing

If you would like to contribute to this project please feel free to do so. I am open to any suggestions or improvements. I am also open to any issues you may find.

## Credits

I would like to thank the following people for their help in creating this project:

- Sinthrill

## License

This project is licensed under the MPL v2 License - see the [LICENSE](LICENSE) file for details.
