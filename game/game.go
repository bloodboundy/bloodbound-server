package game

type Game struct {
	ID      string
	State   GameState
	Players map[string]struct{}
}
