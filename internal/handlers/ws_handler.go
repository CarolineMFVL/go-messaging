package handlers

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
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

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	threadID := vars["threadId"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
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
/* tokenStr := r.URL.Query().Get("token")
claims := &Claims{}
token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
	return jwtKey, nil
})

if err != nil || !token.Valid {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
} username := claims.Username*/

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
