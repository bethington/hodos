package hud

type inputCallbackListener interface {
	OnPlayerMove(x, y float64)
	OnPlayerCast(skillID int, x, y float64)
}
