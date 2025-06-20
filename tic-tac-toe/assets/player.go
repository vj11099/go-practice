package game

type Player struct {
	name   string
	marker string
}

func MakePlayer(name, marker string) Player {
	return Player{name, marker}
}
