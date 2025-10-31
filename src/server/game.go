package server

import (
	"net/http"
)

type checkForAWinnerStruct struct {
	Player        int
	IsThereWinner bool
	IsDraw        bool
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	if IsGameStarted {
		loadRows()
		IsGameStarted = true
	}
	LoadPage(w, r, "./src/client/index.html")
}

func NewParty() {
	ServerData.Rows = make([][]RowStruct, 6)
	ServerData.IsLineFull = make([]bool, 7)
	ServerData.Win = WinStruct{}
	loadRows()
	IsGameStarted = true
	}
	
func loadRows() {
	for i := 0; i <= 6; i++ {
		row := make([]RowStruct, 7)
		for y := 0; y <= 7; y++ {
			row[y] = RowStruct{Player: 0, IsPlaced: false}
		}
		ServerData.Rows[i] = row
	}
}

func PlaceCoinLine(col int) {
	if ServerData.Win.IsWin || ServerData.Win.IsDraw {
		return
	}

	if !appendCoinInsideRow(col) {
		win := checkForAWinner()

		if win.IsDraw {
			ServerData.Win.IsDraw = true
			ServerData.Leaderboard = append(ServerData.Leaderboard, LeaderboardScores{Player: 0, IsWinner: false})
			return
		}

		if win.IsThereWinner {
			ServerData.Win.IsWin = true
			ServerData.Win.Winner = win.Player
			ServerData.Leaderboard = append(ServerData.Leaderboard, LeaderboardScores{Player: win.Player, IsWinner: true})
			return
		}

		if ServerData.PlayerToPlay == 1 {
			ServerData.PlayerToPlay = 2
		}else{
			ServerData.PlayerToPlay = 1
		}
	}
}


func appendCoinInsideRow(col int) bool {
	for row := 5; row >= 0; row-- {
		if !ServerData.Rows[row][col].IsPlaced {
			ServerData.Rows[row][col] = RowStruct{
				Player:   ServerData.PlayerToPlay,
				IsPlaced: true,
			}

			if row == 0 {
				ServerData.IsLineFull[col] = true
			}
			return false
		}
	}

	return true
}

func checkForAWinner() checkForAWinnerStruct {
	full := true
	for _, l := range ServerData.IsLineFull {
		if !l {
			full = false
		}
	}
	if full {
		return checkForAWinnerStruct{IsDraw: true}
	}

	for _, row := range ServerData.Rows {
		count := 1
		for c := 1; c < 7; c++ {
			if row[c].Player != 0 && row[c].Player == row[c-1].Player {
				count ++
				if count == 4 {
					return checkForAWinnerStruct{Player: row[c].Player, IsThereWinner: true}
			}
			} else {
				count = 1
			}
		}
	}

	for c := 0; c < 7; c++ {
		count := 1
		for r := 1; r < 6; r++ {
			if ServerData.Rows[r][c].Player != 0 && ServerData.Rows[r][c].Player == ServerData.Rows[r-1][c].Player {
				count ++
				if count == 4 {
					return checkForAWinnerStruct{Player: ServerData.Rows[r][c].Player, IsThereWinner: true}
			}
			} else {
				count = 1
			}
		}
	}

	for r := 0; r < 3; r++ {
		for c := 0; c < 4; c++ {
			p := ServerData.Rows[r][c].Player
			if p != 0 &&
				p == ServerData.Rows[r+1][c+1].Player &&
				p == ServerData.Rows[r+2][c+2].Player &&
				p == ServerData.Rows[r+3][c+3].Player {
				return checkForAWinnerStruct{Player: p, IsThereWinner: true}
			}
		}
	}

	for r := 3; r < 6; r++ {
		for c := 0; c < 4; c++ {
			p := ServerData.Rows[r][c].Player
			if p != 0 &&
				p == ServerData.Rows[r-1][c+1].Player &&
				p == ServerData.Rows[r-2][c+2].Player &&
				p == ServerData.Rows[r-3][c+3].Player {
				return checkForAWinnerStruct{Player: p, IsThereWinner: true}
			}
		}
	}

	return checkForAWinnerStruct{}
}