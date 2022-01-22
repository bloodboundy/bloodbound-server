package player

import (
	"context"
	"sync"
)

type Manager struct {
	mu *sync.Mutex

	players map[string]*Player
}

func NewManager() *Manager {
	return &Manager{
		mu:      &sync.Mutex{},
		players: make(map[string]*Player),
	}
}

type ctxValueKey string

const playerManagerCtxKey ctxValueKey = "pm"

// PickManager pick manager from ctx
func PickManager(ctx context.Context) *Manager {
	return ctx.Value(playerManagerCtxKey).(*Manager)
}

// MixManager mix manager into ctx and return the child ctx
func MixManager(ctx context.Context, m *Manager) context.Context {
	return context.WithValue(ctx, playerManagerCtxKey, m)
}

func (m *Manager) Register(nickname string) (*Player, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for {
		p := NewPlayer(nickname)
		if _, ok := m.players[p.ID()]; ok {
			continue
		}
		m.players[p.ID()] = p
		return p, nil
	}
}

func (m *Manager) Load(id string) (*Player, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	p, ok := m.players[id]
	return p, ok
}
