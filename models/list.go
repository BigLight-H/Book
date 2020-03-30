package models

type Lists struct {
	Id      int
	Root    string
	Link    string
	Name    string
}

func (m *Lists) TableName() string {
	return TableName("list")
}