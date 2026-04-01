package library

import "time"

type Book struct {
	Name                 string     `json:"Name"`
	Autor                string     `json:"Autor"`
	NumberOfPages        int        `json:"NumberOfPages"`
	ReadingCompleted     bool       `json:"ReadingCompleted"`
	TimeAddBook          time.Time  `json:"TimeAddBook"`
	TimeReadingCompleted *time.Time `json:"TimeReadingCompleted"`
}

func NewBook(name string, autor string, numberOfPages int) Book {
	return Book{
		Name:                 name,
		Autor:                autor,
		NumberOfPages:        numberOfPages,
		ReadingCompleted:     false,
		TimeAddBook:          time.Now(),
		TimeReadingCompleted: nil,
	}
}
func (b *Book) Completed() {
	b.ReadingCompleted = true

	timeCompleted := time.Now()
	b.TimeReadingCompleted = &timeCompleted
}
func (b *Book) Uncompleted() {
	b.ReadingCompleted = false
	b.TimeReadingCompleted = nil
}
