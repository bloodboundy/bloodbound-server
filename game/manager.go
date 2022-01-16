package game

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Manager struct {
	mu *sync.Mutex

	games map[string]*Game
}

func NewManager() *Manager {
	return &Manager{
		mu:    &sync.Mutex{},
		games: make(map[string]*Game),
	}
}

type ctxValueKey string

const gameManagerCtxKey ctxValueKey = "gm"

// PickManager pick manager from ctx
func PickManager(ctx context.Context) *Manager {
	return ctx.Value(gameManagerCtxKey).(*Manager)
}

// MixManager mix manager into ctx and return the child ctx
func MixManager(ctx context.Context, m *Manager) context.Context {
	return context.WithValue(ctx, gameManagerCtxKey, m)
}

func (m *Manager) NewGame() (string, error) {
	var id string
	for {
		uu, err := uuid.NewUUID()
		if err != nil {
			return "", errors.Wrap(err, "new game id")
		}
		id = uu.String()
		if m.tryNewGame(id) {
			return id, nil
		}
	}
}

func (m *Manager) tryNewGame(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.games[id]; ok { // dup
		return false
	}
	m.games[id] = &Game{ID: id, State: WAITING, Players: make(map[string]struct{})}
	return true
}

func (m *Manager) GetGame(id string) (*Game, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	game, ok := m.games[id]
	return game, ok
}

func (m *Manager) RemoveGame(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.games[id]
	if !ok {
		return fmt.Errorf("game not exist: %v", id)
	}
	delete(m.games, id)
	return nil
}
