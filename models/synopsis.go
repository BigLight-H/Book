package models

type Synopsis struct {
	Id      int
	Root    string
	Writer  string
	Img     string
	Name    string
	Synopsis   string
	RenewTime  string
}

func (m *Synopsis) TableName() string {
	return TableName("synopsis")
}
