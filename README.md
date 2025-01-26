# WebSocket Server in Go with Gorilla

This repository contains a simple WebSocket server built with Go and the Gorilla WebSocket library. It demonstrates real-time communication between a server and a client, making it a great starting point for building chat applications, live notifications, and more.

---

## Features

- WebSocket server using Go and Gorilla WebSocket
- Simple echo server functionality
- Client-side interface built with HTML and JavaScript
- Handles real-time bi-directional communication

---

## Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/websocket-golang-gorilla.git
   cd websocket-golang-gorilla
   ```

2. Install dependencies:

   ```bash
   go get -u github.com/gorilla/websocket
   ```

---

## Folder Structure

```
websocket-server/
├── main.go           # Entry point of the application
├── handlers/         # Contains WebSocket handlers
│   └── websocket.go  # Handles WebSocket connections and messages
├── static/           # Contains client-side HTML/JS files
│   └── index.html    # A simple client interface
└── go.mod            # Go module file
```

---

## Usage

### Run the Server

1. Start the server:

   ```bash
   go run main.go
   ```

2. Open your browser and navigate to:

   [http://localhost:8080](http://localhost:8080)

### Client Interaction

- Type a message in the input box and click "Send."
- Messages will be echoed back by the server and displayed in real-time.

---

## Example Code

### Server (main.go)

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"websocket-server/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.HandleWebSocket)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
```

### Client (index.html)

```html
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Client</title>
</head>
<body>
    <h1>WebSocket Client</h1>
    <input type="text" id="messageInput" placeholder="Type a message...">
    <button id="sendButton">Send</button>
    <ul id="messages"></ul>

    <script>
        const ws = new WebSocket("ws://localhost:8080/ws");
        const messageInput = document.getElementById("messageInput");
        const sendButton = document.getElementById("sendButton");
        const messages = document.getElementById("messages");

        ws.onmessage = (event) => {
            const li = document.createElement("li");
            li.textContent = `Server: ${event.data}`;
            messages.appendChild(li);
        };

        sendButton.onclick = () => {
            const message = messageInput.value;
            ws.send(message);
            const li = document.createElement("li");
            li.textContent = `You: ${message}`;
            messages.appendChild(li);
            messageInput.value = "";
        };
    </script>
</body>
</html>
```

---

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue to suggest improvements or report bugs.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Gorilla WebSocket](https://github.com/gorilla/websocket) for providing an easy-to-use WebSocket library for Go.
