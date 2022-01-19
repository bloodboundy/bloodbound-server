package ws

import (
	"context"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	mu *sync.Mutex

	conns map[string]*websocket.Conn
}

func NewManager() *Manager {
	return &Manager{
		mu:    &sync.Mutex{},
		conns: make(map[string]*websocket.Conn),
	}
}

type ctxValueKey string

const wsManagerCtxKey ctxValueKey = "wm"

func PickManager(ctx context.Context) *Manager {
	return ctx.Value(wsManagerCtxKey).(*Manager)
}

func MixManager(ctx context.Context, m *Manager) context.Context {
	return context.WithValue(ctx, wsManagerCtxKey, m)
}

// Store when id already exists, an error will be raised
func (m *Manager) Store(id string, conn *websocket.Conn) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.conns[id]; ok {
		return fmt.Errorf("the id already exists")
	}
	m.conns[id] = conn
	return nil
}

func (m *Manager) Load(id string) (*websocket.Conn, bool) {
	m.mu.Lock()
	defer m.mu.Lock()

	conn, ok := m.conns[id]
	return conn, ok
}

func (m *Manager) Delete(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.conns, id)
	return
}
