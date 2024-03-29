package game

import (
	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// GameJSON represents the JSON format of a game, used to communicate
type GameJSON struct {
	// meta data
	ID        string `json:"ID,omitempty"`
	CreatedAt uint64 `json:"created_at,omitempty"`
	Owner     string `json:"owner,omitempty"`

	// settings
	MaxPlayers *uint32 `json:"max_players,omitempty"`
	IsPrivate  bool    `json:"is_private,omitempty"`
	Password   string  `json:"password,omitempty"`

	// contained data
	Players []*player.PlayerJSON `json:"players,omitempty"`
}

// Load settings from `src`
func (g *Game) Load(src *GameJSON) error {
	var e error
	if err := g.SetMaxPlayers(src.MaxPlayers); err != nil {
		e = errors.Wrap(err, "setMaxPlayers")
	}
	g.password = src.Password
	return e
}

// Dump game to GameJSON
//
// normally, the "password" were filter out from ret-val
// include it in addition to add it into ret-val
func (g *Game) Dump(addition ...string) *GameJSON {
	var players []*player.PlayerJSON
	for _, p := range g.ListPlayers() {
		players = append(players, p.Dump())
	}

	gj := &GameJSON{
		ID:         g.id,
		MaxPlayers: proto.Uint32(g.maxPlayers),
		IsPrivate:  g.IsPrivate(),
		CreatedAt:  g.createdAt,
		Owner:      g.owner,
		Players:    players,
	}

	for _, field := range addition {
		switch field {
		case "password":
			gj.Password = g.password
		default:
		}
	}
	return gj
}
