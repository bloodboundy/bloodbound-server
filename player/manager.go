package player

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (m *Manager) Register() (string, error) {
	var id string
	for {
		uu, err := uuid.NewUUID()
		if err != nil {
			return "", errors.Wrap(err, "new game id")
		}
		id = uu.String()
		if m.tryRegister(id) {
			return id, nil
		}
	}
}

func (m *Manager) tryRegister(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.players[id]; ok {
		return false
	}
	m.players[id] = NewPlayer(id, "")
	return true
}
