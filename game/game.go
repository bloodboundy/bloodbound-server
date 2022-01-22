package game

import (
	"fmt"
	"time"

	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

const (
	MIN_PLAYERS = 6
	MAX_PLAYERS = 12
)

type Game struct {
	// meta data
	id        string
	createdAt uint64
	createdBy string

	// settings
	maxPlayers uint32
	password   string

	// contained data
	players map[string]*player.Player
}

func NewGame(createdBy string) *Game {
	return &Game{
		id:        uuid.NewString(),
		createdAt: uint64(time.Now().Unix()),
		createdBy: createdBy,
		players:   make(map[string]*player.Player),
	}
}

// Load settings from `src`
func (g *Game) Load(src *GameJSON) error {
	if err := g.SetMaxPlayers(src.MaxPlayers); err != nil {
		return errors.Wrap(err, "setMaxPlayers")
	}
	g.password = src.Password
	return nil
}

// Dump game to GameJSON
//
// normally, the "password", "players" were filter out from ret-val
// include it in addition to add it into ret-val
func (g *Game) Dump(addition ...string) *GameJSON {
	gj := &GameJSON{
		ID:         g.id,
		MaxPlayers: proto.Uint32(g.maxPlayers),
		IsPrivate:  g.IsPrivate(),
		CreatedAt:  g.createdAt,
		CreatedBy:  g.createdBy,
	}

	for _, field := range addition {
		switch field {
		case "password":
			gj.Password = g.password
		case "players":
			gj.Players = g.ListPlayers()
		}
	}

	return gj
}

func (g *Game) ID() string { return g.id }

// SetMaxPlayers when max==nil, 12 is used
// when max not in (6,12], error returned
func (g *Game) SetMaxPlayers(max *uint32) error {
	if max == nil {
		g.maxPlayers = MAX_PLAYERS
		return nil
	}

	if *max < MIN_PLAYERS || *max > MAX_PLAYERS {
		return fmt.Errorf("unacceptable max_players: %v, expected [%d, %d]", *max, MIN_PLAYERS, MAX_PLAYERS)
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
