package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to upgrade connection:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Client connected")

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Read error:", err)
            break
        }

        fmt.Printf("Received: %s\n", message)

        if err := conn.WriteMessage(messageType, message); err != nil {
            fmt.Println("Write error:", err)
            break
        }
    }
}