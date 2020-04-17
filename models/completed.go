package models

type Completed struct {
	Id      int
	Root    string
	BookName      string
	Domain        string
	DomainPc      string
	BookNameLink  string
	BookAuthor    string
	BookType      string
	BookNextPage  string
	BookLastPage  string
	BookFirstPage string
	BookPrevious  string
	BookPageRoot  string
}

func (m *Completed) TableName() string {
	return TableName("completed_list")
}