package player

import (
	"fmt"
	"sync"
)

type Player struct {
	mu       *sync.Mutex
	ID       string `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	game     string // which game this player joined now
}

// NewPlayer creates a new Player
// `game` can be ommited when created without a certain game
func NewPlayer(id string, nickname string, game string) *Player {
	return &Player{
		mu:       &sync.Mutex{},
		ID:       id,
		Nickname: nickname,
		game:     game,
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
