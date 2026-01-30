# Poker Server Backend

A poker game logic web server built with Go, providing deck management, card dealing, and game state management for poker applications.

## Features

- **Deck Management**: Standard 52-card deck initialization and shuffling
- **Card Dealing**: Deal cards from the deck with quantity control
- **Stack Data Structure**: LIFO (Last In First Out) stack implementation for deck management
- **Event System**: Observer pattern-based event management for game events
- **REST API**: HTTP endpoints for deck operations
- **Comprehensive Tests**: Full test coverage for core game logic

## Project Structure

```
poker-server-backend/
├── core/                    # Core game logic
│   ├── gamelogic.go        # Deck, Card, Player, and Game types
│   ├── gamelogic_test.go   # Tests for game logic
│   ├── stack.go            # Stack data structure implementation
│   ├── stack_test.go       # Tests for stack
│   ├── eventmanager.go     # Event system (Observer pattern)
│   └── rest.go             # REST API handlers
├── infra/                   # Infrastructure layer
│   ├── webserver.go        # HTTP server setup
│   └── websocket.go        # WebSocket support (TODO)
├── main.go                  # Application entry point
└── go.mod                   # Go module dependencies
```

## Installation

### Prerequisites

- Go 1.25.5 or later

### Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd poker-server-backend
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build
```

## Usage

### Running the Server

Start the server:
```bash
go run main.go
```

The server will start on `localhost:8080`.

### API Endpoints

#### Deal Cards
**GET** `/api/poker/deck/dealer`

Deal a specified quantity of cards from a shuffled deck.

**Query Parameters:**
- `qty` (required): Number of cards to deal (integer)

**Example Request:**
```bash
curl "http://localhost:8080/api/poker/deck/dealer?qty=2"
```

**Example Response:**
```json
[
  {
    "naipe": "H",
    "code": "K",
    "value": 13
  },
  {
    "naipe": "D",
    "code": "A",
    "value": 14
  }
]
```

**Response Codes:**
- `200 OK`: Success
- `400 Bad Request`: Invalid `qty` parameter
- `422 Unprocessable Entity`: JSON marshaling error

#### Poker Endpoint
**GET** `/api/poker`

Returns a simple counter response.

**Example Request:**
```bash
curl "http://localhost:8080/api/poker"
```

**Example Response:**
```
Counter
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test ./core -v
```

Run specific test:
```bash
go test ./core -v -run TestDeck_Init
```

### Test Coverage

The project includes comprehensive tests for:
- Deck initialization and shuffling
- Card dealing (including edge cases)
- Stack operations (push, pop, empty, size, clear)
- LIFO behavior verification

## Card Structure

Cards are represented with the following structure:
- **Naipe**: Suit (C=Clubs, D=Diamonds, H=Hearts, S=Spades)
- **Code**: Card code (2-10 for number cards, J/Q/K/A for face cards)
- **Value**: Numeric value (2-10 for numbers, 11-14 for J/Q/K/A)

## Development

### Current Status

The project is in active development. Current features include:
- ✅ Deck initialization and shuffling
- ✅ Card dealing functionality
- ✅ Stack data structure
- ✅ Basic REST API
- ✅ Comprehensive test suite
- ✅ Event management system (Observer pattern)

### TODO

- [ ] Complete game coordination object (multiple poker hands)
- [ ] Player management (minimum 2 players)
- [ ] Button position management
- [ ] Small Blind and Big Blind assignment
- [ ] Betting system (initial bet values: 50/100, 100/200, etc.)
- [ ] WebSocket support for real-time gameplay
- [ ] Route mapping system (map routes to handler methods)
- [ ] Separate Observer Pattern into dedicated package
- [ ] Add more comprehensive game logic tests

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Add your license here]
