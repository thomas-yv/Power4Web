package server

var ServerData = ServerStruct{
	Title:               "Power4Web",
	PlayerToPlay:        1,
	PlayerSelectedIndex: 0,
	Players: []PlayerData{
		{Slot: "red"},
		{Slot: "yellow"},
	},
	Leaderboard: []LeaderboardScores{},
	Win: WinStruct{
		Winner: 0,
		IsWin:  false,
		IsDraw: false,
	},
	Rows:       make([][]RowStruct, 6),
	IsLineFull: make([]bool, 7),
	AvailableSlotColors: []string{
		"red", "yellow",
	},
}

var IsGameStarted = false
var NumOfPlayers = 2

func IsValidColor(color string) bool {
	for _, c := range ServerData.AvailableSlotColors {
		if color == c {
			return true
		}
	}

	return false
}