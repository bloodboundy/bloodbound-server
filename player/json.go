package player

type JSON struct {
	ID       string `json:"ID,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

func (p *Player) Dump() *JSON {
	return &JSON{
		ID:       p.id,
		Nickname: p.nickname,
	}
}
