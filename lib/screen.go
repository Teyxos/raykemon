package lib

type Screens int32

const (
	WorldScreen Screens = iota
	BattleScreen
	MenuScreen
)

var currentScreen Screens

func SetScreen(screen Screens) {
	currentScreen = Screens(screen)
}

func GetScreen() Screens {
	return Screens(currentScreen)
}
