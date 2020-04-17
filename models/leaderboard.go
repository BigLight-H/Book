package models

type Leaderboard struct {
	Id      int
	Root    string
	BookName      string
	Domain        string
	DomainPc      string
	BookNameLink  string
	BookAuthor    string
	BookType      string
	BookNextPage  string
	BookLastPage  string
	BookFirstPage string
	BookPrevious  string
	BookPageRoot  string
}

func (m *Leaderboard) TableName() string {
	return TableName("leaderboard_list")
}
