# Trivia Game

This project is a simple trivia game implemented using Go for the server and JavaScript for the client. It utilizes WebSocket communication to provide real-time updates to connected clients.

## Features

- **Real-Time Updates:** Players receive real-time updates on the game state, including questions, options, and scores.
- **User Interaction:** Players can select answers to questions by clicking on clickable buttons.
- **Scalable:** Designed to handle multiple players connecting to the game simultaneously.

## Technologies Used

- **Go:** Used for the server-side implementation.
- **JavaScript:** Used for the client-side implementation.
- **Gorilla WebSocket:** A WebSocket implementation for the Go programming language.
  
## Getting Started

### Prerequisites

- Go installed on your machine.
- A modern web browser.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/trivia-game.git
    cd trivia-game
    ```

2. Run the server:

    ```bash
    go run server.go
    ```

3. Open `index.html` in a web browser.

## Usage

1. Open the application in your web browser.
2. Enter your username when prompted.
3. Answer trivia questions by clicking on the provided options.
4. See real-time updates on the main screen.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue.

## License

This project is licensed under the [MIT License](LICENSE).
