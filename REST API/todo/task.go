package todo

import "time"

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`

	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"complited_time"`
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	completedAt := time.Now()

	t.Completed = true
	t.CompletedAt = &completedAt
}
func (t *Task) Uncompleted() {
	t.Completed = false
	t.CompletedAt = nil
}
