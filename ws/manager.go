package ws

import (
	"context"
	"github.com/pkg/errors"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	mu *sync.Mutex

	conns map[string]*websocket.Conn // key: player id
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
		return errors.Errorf("the id already exists")
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
}

func (m *Manager) BroadCast(jso interface{}, ids ...string) error {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	errs := &strings.Builder{}
	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			if err := m.Send(jso, id); err != nil {
				mu.Lock()
				errs.WriteString("#" + id + ": " + err.Error())
				mu.Unlock()
			}
			wg.Done()
		}(id)
	}
	if err := errs.String(); err != "" {
		return errors.New(err)
	}
	return nil
}

func (m *Manager) Send(jso interface{}, id string) error {
	ws, ok := m.Load(id)
	if !ok {
		return errors.Errorf("ws conn not found for: %v", id)
	}
	if err := ws.WriteJSON(jso); err != nil {
		return errors.Wrap(err, "ws.Write")
	}
	return nil
}
