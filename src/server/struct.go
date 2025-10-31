package server

type PlayerData struct {
	Slot string
}

type LeaderboardScores struct {
	Player   int
	IsWinner bool
}

type WinStruct struct {
	Winner int
	IsWin  bool
	IsDraw bool
}

type RowStruct struct {
	Player   int
	IsPlaced bool
}

type ServerStruct struct {
	Title               string
	PlayerSelectedIndex int
	Players             []PlayerData
	PlayerToPlay        int
	Leaderboard         []LeaderboardScores
	Win                 WinStruct
	Rows                [][]RowStruct
	IsLineFull          []bool
	AvailableSlotColors []string
}
