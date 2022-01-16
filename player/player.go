package player

import (
	"fmt"
	"sync"
)

type Player struct {
	mu   *sync.Mutex
	ID   string `json:"id"`
	game string
}

func NewPlayer(id string, game string) *Player {
	return &Player{
		mu:   &sync.Mutex{},
		ID:   id,
		game: game,
	}
}

func (p *Player) Join(game string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.game != "" {
		return fmt.Errorf("already in game")
	}
	p.game = game
	return nil
}

func (p *Player) Leave() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.game == "" {
		return
	}
	p.game = ""
}

func (p *Player) Game() string {
	return p.game
}
