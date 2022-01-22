package player

type PlayerJSON struct {
	ID       string `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

func (p *Player) Dump() *PlayerJSON {
	return &PlayerJSON{
		ID:       p.id,
		Nickname: p.nickname,
	}
}
