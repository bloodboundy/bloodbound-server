package player

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Player struct {
	mu       *sync.Mutex
	id       string
	nickname string
	game     string // which game this player joined now
}

// NewPlayer creates a new Player
// `game` can be omitted when created without a certain game
func NewPlayer(nickname string) *Player {
	return &Player{
		mu:       &sync.Mutex{},
		id:       uuid.NewString(),
		nickname: nickname,
	}
}

func (p *Player) ID() string { return p.id }

func (p *Player) Nickname() string { return p.nickname }

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
