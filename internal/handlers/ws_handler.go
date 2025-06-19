package handlers

import (
	//"encoding/json"
	"net/http"
	"nls-go-messaging/internal/utils"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ConnectionWrapper struct {
	*websocket.Conn
	mu sync.Mutex
}

func (c *ConnectionWrapper) SafeWriteJSON(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Conn.WriteJSON(v)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// HandleWebSocket godoc
// @Summary WebSocket protégé par JWT
// @Description Ouvre une connexion WebSocket sur un thread
// @Tags ws
// @Produce  json
// @Param threadId path string true "ID du thread"
// @Security BearerAuth
// @Success 101 {string} string "Switching Protocols"
// @Failure 401 {object} map[string]string
// @Router /ws/{threadId} [get]
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["threadId"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "WebSocket upgrade failed")
		return
	}

	wrappedConn := &ConnectionWrapper{Conn: conn}
	client := &Client{
		Conn:     wrappedConn,
		ThreadID: threadID,
		Send:     make(chan Message),
	}

	thread := GetOrCreateThread(threadID)
	thread.Clients[client] = true

	// Envoie des messages déjà existants
	for _, msg := range thread.Messages {
		client.Conn.SafeWriteJSON(msg)
	}

	go client.readPump(thread)
	go client.writePump()
	// Add gestion errors and close connections

}

func (c *Client) readPump(t *Thread) {
	defer func() {
		c.Conn.Close()
		delete(t.Clients, c)
	}()

	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		t.Broadcast(msg)
	}
}

func (c *Client) writePump() {
	for msg := range c.Send {
		err := c.Conn.SafeWriteJSON(msg)
		if err != nil {
			break
		}
	}
}
