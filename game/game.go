package game

import (
	"fmt"
	"sync"
	"time"

	"github.com/bloodboundy/bloodbound-server/config"
	"github.com/bloodboundy/bloodbound-server/game/fsm"
	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Status int

const (
	WAITING Status = iota
	STARTED
	ENDED
)

type Game struct {
	// meta data
	id        string
	createdAt uint64
	owner     string

	// settings
	maxPlayers uint32
	password   string

	// contained data
	players map[string]*player.Player

	mus    *sync.Mutex // protect state, status and actions
	state  *fsm.State
	status Status
}

func NewGame(createdBy string) *Game {
	return &Game{
		id:        uuid.NewString(),
		createdAt: uint64(time.Now().Unix()),
		owner:     createdBy,
		players:   make(map[string]*player.Player),
		mus:       &sync.Mutex{},
	}
}

func (g *Game) ID() string { return g.id }

// SetMaxPlayers when max==nil, 12 is used
// when max not in (6,12], error returned
func (g *Game) SetMaxPlayers(max *uint32) error {
	if max == nil {
		g.maxPlayers = config.GameMaxPlayers
		return nil
	}

	if *max < config.GameMinPlayers || *max > config.GameMaxPlayers {
		return fmt.Errorf("unacceptable max_players: %v, expected [%d, %d]", *max, config.GameMinPlayers, config.GameMaxPlayers)
	}
	g.maxPlayers = *max
	return nil
}

func (g *Game) GetMaxPlayers() uint32 {
	return g.maxPlayers
}

func (g *Game) IsPrivate() bool {
	return g.password != ""
}

func (g *Game) Password() string { return g.password }

func (g *Game) AddPlayer(p *player.Player) error {
	if len(g.players) >= int(g.GetMaxPlayers()) {
		return fmt.Errorf("the game is full")
	}
	if p.Game() == g.id {
		return nil
	}
	if err := p.Join(g.id); err != nil {
		return err
	}
	g.players[p.ID()] = p
	return nil
}

func (g *Game) RemovePlayer(p *player.Player) {
	if _, ok := g.players[p.ID()]; !ok {
		return
	}
	p.Leave()
	delete(g.players, p.ID())
}

func (g *Game) ListPlayers() []*player.Player {
	result := make([]*player.Player, 0, len(g.players))
	for _, p := range g.players {
		result = append(result, p)
	}
	return result
}

func (g *Game) Owner() string { return g.owner }

func (g *Game) Start() error {
	g.mus.Lock()
	defer g.mus.Unlock()

	if g.status != WAITING {
		return errors.Errorf("game is not in waiting status, started or ended")
	}
	state, err := fsm.NewState(g.ID(), g.ListPlayers())
	if err != nil {
		return errors.Wrap(err, "Start")
	}
	g.state = state
	g.status = STARTED
	return nil
}

func (g *Game) Status() Status {
	return g.status
}

func (g *Game) ApplyAction() error {
	g.mus.Lock()
	defer g.mus.Unlock()

	return nil
}
