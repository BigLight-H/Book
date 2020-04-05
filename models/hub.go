package models

type Hub struct {
	Id        int
	BookHub   string
	Link      string
	Root      string
	Name      string
	Author    string
	Suffix    string
	NewList   string
	NewListLink   string
	RenewTime     string
	Status    string
	Image     string
	Mark      string
}

func (m *Hub) TableName() string {
	return TableName("hub")
}
