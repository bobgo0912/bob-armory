package constant

type GameEvent string

const (
	Begin        GameEvent = "begin"
	WaitPrize    GameEvent = "wait_prize"
	PrizeOnGoing GameEvent = "prize_on_going"
	Finish       GameEvent = "finish"
	WaitStart    GameEvent = "wait_start"
	Cancel       GameEvent = "cancel"
)
