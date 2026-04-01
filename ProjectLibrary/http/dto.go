package http

import (
	"encoding/json"
	"errors"
	"time"
)

type BookDTO struct {
	Name          string `json:"name"`
	Autor         string
	NumberOfPages int
}

func (b BookDTO) ValidationForCreate() error {
	if b.Name == "" {
		return errors.New("book title is empty")
	}
	if b.Autor == "" {
		return errors.New("autor is empty")
	}
	if b.NumberOfPages == 0 {
		return errors.New("numberOfPages is empty")
	}
	return nil
}

type CompletedBookDTO struct {
	CompleteBook bool `json:"completeBook"`
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func NewErrorDTO(err error) *ErrorDTO {
	return &ErrorDTO{Message: err.Error(), Time: time.Now()}
}
