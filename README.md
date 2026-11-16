# Pokedex CLI

A command-line interactive Pokedex (REPL) written in Go. It connects to the [PokéAPI](https://pokeapi.co/) to let trainers browse the Pokémon world, explore wild Pokémon habitats, and catch Pokémon in their terminal.

## Features

- **Interactive REPL**: Read-Eval-Print Loop designed with dynamic command dispatching and input sanitization.
- **Thread-Safe In-Memory Cache (`pokecache`)**: Custom cache layer with a background cleanup goroutine (TTL reaping) to reduce network requests and comply with PokéAPI fair use guidelines.
- **Location Pagination**: Paginate through 20 location areas at a time in both forward and backward directions.
- **Area Exploration**: Look up any location area to discover all wild Pokémon species that can be encountered there.
- **Catching Mechanics**: Attempt to catch Pokémon with catch rates dynamically scaled based on the target's `base_experience` using randomized rolls. Caught Pokémon are stored in your active session Pokedex.

---

## Available Commands

| Command | Usage | Description |
|---|---|---|
| `help` | `help` | Displays a list of all available commands and their descriptions. |
| `map` | `map` | Fetches and displays the next 20 location areas in the Pokémon world. |
| `mapb` | `mapb` | Fetches and displays the previous 20 location areas (or alerts if on the first page). |
| `explore` | `explore <location_area>` | Lists all wild Pokémon that can be encountered in the specified area. |
| `catch` | `catch <pokemon_name>` | Throws a Pokéball at a Pokémon to attempt capture. |
| `exit` | `exit` | Exits the Pokedex application. |

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (version 1.22 or higher)

### Installation & Running

Clone the repository and run the application:

```bash
git clone git@github.com:Clinton-Okumu/Pokedox.git
cd Pokedox
go run .
```

---

## Example Usage

```text
Pokedex > help

Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Get the next page of locations
mapb: Get the previous page of locations
explore: Explore a location
catch: Attempt to catch a pokemon
exit: Exit the Pokedex

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados
 - remoraid
 - octillery

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > exit
Closing the Pokedex... Goodbye!
```

---

## Running Tests

Run the test suite across all packages:

```bash
go test -v ./...
```
