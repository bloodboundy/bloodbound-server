package game

import (
	"context"
	"sync"
	"time"

	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/google/uuid"
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

func (m *Manager) NewGame(creatorID string) *Game {
	var id string
	for {
		uu := uuid.New()
		id = uu.String()
		if g := m.tryNewGame(id, creatorID); g != nil {
			return g
		}
	}
}

func (m *Manager) tryNewGame(id string, creatorID string) *Game {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.games[id]; ok { // dup
		return nil
	}
	g := &Game{
		ID:        id,
		CreatedAt: uint64(time.Now().Unix()),
		CreatedBy: creatorID,
		players:   make(map[string]*player.Player),
	}
	m.games[id] = g
	return g
}

func (m *Manager) Load(id string) (*Game, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	game, ok := m.games[id]
	return game, ok
}

// List all public games
func (m *Manager) List() []*Game {
	result := []*Game{}
	for _, game := range m.games {
		if game.IsPrivate() {
			continue
		}
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
	delete(m.games, id)
}
