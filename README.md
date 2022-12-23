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

The rules of tichu are a bit complex but can be found [here](https://en.wikipedia.org/wiki/Tichu).

## Installation

### Quick Installation
You can install this program via the go package manager:

```go
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
$ go build
```

you can find configuration for tichu in the .tichu file of the directory you installed it in.

e.g if you installed it in your home directory you can find the configuration file in ~/.tichu.conf/...

## Usage