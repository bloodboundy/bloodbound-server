package game

import "github.com/bloodboundy/bloodbound-server/player"

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
