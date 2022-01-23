package component

type Token struct {
	t     string // type, color
	gotAt uint32 // got at which round
}

type TokenJSON struct {
	Type  string `json:"type"`
	GotAt uint32 `json:"got_at"`
}

func NewBToken(gotAt uint32) *Token { return &Token{t: BLUE_CLAN, gotAt: gotAt} }

func NewRToken(gotAt uint32) *Token { return &Token{t: RED_CLAN, gotAt: gotAt} }

func NewSToken(gotAt uint32) *Token { return &Token{t: SEC_CLAN, gotAt: gotAt} }

func (t *Token) Type() string { return t.t }

func (t *Token) GotAt() uint32 { return t.gotAt }

func (t *Token) Dump() *TokenJSON { return &TokenJSON{Type: t.t, GotAt: t.gotAt} }
