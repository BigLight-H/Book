package models

type Source struct {
	Id                     int
	Domain                 string
	PcDomain               string
	Root                   string
	TypeTitle              string
	MoreList               string
	Img					   string
	BookName               string
	BookAuthor             string
	BookMark               string
	BookNameLink           string
	BookAuthorLink         string
	ListTypeRoot           string
	ListTypeName           string
	ListBookName           string
	ListBookAuthor         string
	ListTypeLink           string
	ListBookNameLink       string
	ListBookAuthorLink     string
}

func (m *Source) TableName() string {
	return TableName("source")
}
