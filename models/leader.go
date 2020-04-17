package models

type Leader struct {
	Id        int
	LeaderboardName      string
	Domain    string
}

func (m *Leader) TableName() string {
	return TableName("leaderboard")
}