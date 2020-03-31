package models

type Users struct {
	Id      int
	Name    string
	Pwd     string
	Mobile  string
	Email   string
}

func (m *Users) TableName() string {
	return TableName("users")
}