package handlers

import (
	"strings"
	"sync"
)

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

type Thread struct {
	ID       string
	Clients  map[*Client]bool
	Messages []Message
	mu       sync.Mutex
}

type Client struct {
	Conn     *ConnectionWrapper
	ThreadID string
	Send     chan Message
}

var (
	threads     = make(map[string]*Thread)
	threadsLock = sync.Mutex{}
	alertWords  = []string{"urgent", "alerte", "immédiat"}
)

func GetOrCreateThread(id string) *Thread {
	threadsLock.Lock()
	defer threadsLock.Unlock()

	if t, ok := threads[id]; ok {
		return t
	}

	thread := &Thread{
		ID:      id,
		Clients: make(map[*Client]bool),
	}
	threads[id] = thread
	return thread
}

func (t *Thread) Broadcast(msg Message) {
	t.mu.Lock()
	t.Messages = append(t.Messages, msg)
	t.mu.Unlock()

	for c := range t.Clients {
		select {
		case c.Send <- msg:
		default:
			close(c.Send)
			delete(t.Clients, c)
		}
	}

	for _, word := range alertWords {
		if strings.Contains(strings.ToLower(msg.Content), word) {
			logAlert(t.ID, msg)
			break
		}
	}
}

func logAlert(threadID string, msg Message) {
	log := "[ALERTE] Dans thread " + threadID + " : " + msg.User + " a dit : " + msg.Content
	println(log)
}
