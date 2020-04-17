package models

type Type struct {
	Id      int
	Root    string
	BookName      string
	Domain        string
	DomainPc      string
	BookNameLink  string
	BookAuthor    string
	BookNextPage  string
	BookLastPage  string
	BookFirstPage string
	BookPrevious  string
	BookPageRoot  string
}

func (m *Type) TableName() string {
	return TableName("type")
}