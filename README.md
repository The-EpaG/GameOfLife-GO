# Game of Life in Go

This repository contains a reimplementation of Conway's Game of Life using the Go programming language. The primary goal of this project is to learn and explore the Go language by building a well-known cellular automaton simulation.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Conway's Game of Life is a zero-player game, meaning its evolution is determined by its initial state, requiring no further input. The game consists of a grid of cells that can live, die, or multiply based on a few mathematical rules. This project reimplements the Game of Life in Go to serve as a practical learning exercise for the language.

## Getting Started

### Prerequisites

To run this project, you will need to have Go installed on your system. You can download it from the [official website](https://golang.org/dl/).

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/The-EpaG/GameOfLife-GO.git
    cd GameOfLife-GO
    ```

2. Build the project:
    ```sh
    go build
    ```

## Usage
After building the project, you can run the Game of Life simulation with the following command

### Generate the first image
You can generate the first image in the right position using the -i (init) parameters

```sh
./bin/GameOfLife-GO -i
```

You can also customize the grid size:

```sh
./bin/GameOfLife-GO -i -width 50 -height 30
```

### Start a simulation
You can start the simulation using the -s (start) parameters

```sh
./bin/GameOfLife-GO -s
```

You can also customize the number of generations through command-line arguments:


```sh
./bin/GameOfLife-GO -s -generations 100
```

### Command-line Arguments

- `-i`: Create the first image
- `-s`: Start the simulation
- `-width`: Width of the grid (default: 100)
- `-height`: Height of the grid (default: 100)
- `-generations`: Number of generations to simulate 
(default: 1)

## Features

- Implementation of Conway's Game of Life rules.
- Customizable grid size and number of generations.
- Simple and clean Go code to facilitate learning.

## Contributing

Contributions are welcome! If you have any improvements or suggestions, please open an issue or submit a pull request.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch: `git checkout -b my-feature-branch`
3. Make your changes and commit them: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin my-feature-branch`
5. Open a pull request.

## License

This project is licensed under the GNU3 License - see the [LICENSE](LICENSE) file for details.
