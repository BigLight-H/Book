package models

type Bookshelf struct {
	Id        int
	UserId    int
	HubId     int
	Link      string
	Img       string
	Domain    string
	BookName  string
	Author    string
}

func (m *Bookshelf) TableName() string {
	return TableName("bookshelf")
}