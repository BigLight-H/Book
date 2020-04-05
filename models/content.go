package models


type Content struct {
	Id       int
	Name     string
	Root     string
	Content  string `html:"Html"`
	SPage    string
	XPage    string
	List     string
	Domain   string
}

func (m *Content) TableName() string {
	return TableName("content")
}