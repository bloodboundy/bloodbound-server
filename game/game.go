package game

import (
	"fmt"

	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type Game struct {
	// meta data
	ID        string
	CreatedAt uint64
	CreatedBy string

	// settings
	MaxPlayers *uint32
	Password   string

	// contained data
	players map[string]*player.Player
}

// GameJSON represents the JSON format of a game, used to communicate
type GameJSON struct {
	// meta data
	ID        string `json:"id,omitempty"`
	CreatedAt uint64 `json:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`

	// settings
	MaxPlayers *uint32 `json:"max_players,omitempty"`
	IsPrivate  bool    `json:"is_private,omitempty"`
	Password   string  `json:"password,omitempty"`

	// contained data
	Players []*player.Player `json:"players,omitempty"`
}

// Load settings from `src`
func (g *Game) Load(src *GameJSON) error {
	if err := g.SetMaxPlayers(src.MaxPlayers); err != nil {
		return errors.Wrap(err, "setMaxPlayers")
	}
	g.Password = src.Password
	return nil
}

// Dump game to GameJSON
//
// normally, the "password", "players" were filter out from ret-val
// include it in addition to add it into ret-val
func (g *Game) Dump(addition ...string) *GameJSON {
	var isPrivate bool
	if g.Password != "" {
		isPrivate = true
	}

	gj := &GameJSON{
		ID:         g.ID,
		MaxPlayers: g.MaxPlayers,
		IsPrivate:  isPrivate,
		CreatedAt:  g.CreatedAt,
		CreatedBy:  g.CreatedBy,
	}

	for _, field := range addition {
		switch field {
		case "password":
			gj.Password = g.Password
		case "players":
			gj.Players = g.ListPlayers()
		}
	}

	return gj
}

const (
	MIN_PLAYERS = 6
	MAX_PLAYERS = 12
)

// SetMaxPlayers when max==nil, 12 is used
// when max not in (6,12], error returned
func (g *Game) SetMaxPlayers(max *uint32) error {
	if max == nil {
		g.MaxPlayers = proto.Uint32(MAX_PLAYERS)
		return nil
	} else {
		if *max > MIN_PLAYERS && *max <= MAX_PLAYERS {
			g.MaxPlayers = max
		} else {
			return fmt.Errorf("unacceptable max_players: %v, expected (%d, %d]", *max, MIN_PLAYERS, MAX_PLAYERS)
		}
	}
	return nil
}

func (g *Game) GetMaxPlayers() uint32 {
	if g.MaxPlayers == nil {
		return 12
	}
	return *g.MaxPlayers
}

func (g *Game) IsPrivate() bool {
	return g.Password != ""
}

func (g *Game) AddPlayer(p *player.Player) {
	if err := p.Join(g.ID); err != nil {
		return
	}
	g.players[p.ID] = p
}

func (g *Game) RemovePlayer(p *player.Player) {
	if _, ok := g.players[p.ID]; !ok {
		return
	}
	p.Leave()
	delete(g.players, p.ID)
}

func (g *Game) ListPlayers() []*player.Player {
	result := make([]*player.Player, 0, len(g.players))
	for _, p := range g.players {
		result = append(result, p)
	}
	return result
}
