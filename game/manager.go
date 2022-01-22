package game

import (
	"context"
	"sync"
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

func (m *Manager) LoadOrStore(game *Game) (actual *Game, loaded bool) {
	actual, loaded = m.Load(game.ID())
	if loaded {
		return actual, loaded
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.games[game.ID()] = game
	return game, false
}

func (m *Manager) Load(id string) (*Game, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	game, ok := m.games[id]
	return game, ok
}

func (m *Manager) Store(game *Game) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.games[game.ID()] = game
}

func (m *Manager) List() []*Game {
	result := []*Game{}
	for _, game := range m.games {
		result = append(result, game)
	}
	return result
}

func (m *Manager) Delete(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.games[id]
	if !ok {
		return
	}

	for _, p := range m.games[id].players {
		p.Leave()
	}
	delete(m.games, id)
}
