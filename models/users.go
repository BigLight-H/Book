package models

type Users struct {
	Id      int
	Name    string
	Pwd     string
	mobile  string
	Email   string
}

func (m *Users) TableName() string {
	return TableName("users")
}