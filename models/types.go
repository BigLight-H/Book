package models

type Types struct {
	Id        int
	Name      string
	Domain    string
}

func (m *Types) TableName() string {
	return TableName("types")
}