package game

type Game struct {
	ID      string
	Players map[string]struct{}
}
