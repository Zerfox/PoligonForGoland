package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (list *List) AddTask(task Task) error {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	if _, ok := list.tasks[task.Title]; ok {
		list.mtx.Unlock()
		return ErrTaskAlreadyExistst
	}

	list.tasks[task.Title] = task

	return nil
}

func (list *List) GetTask(title string) (Task, error) {
	list.mtx.RLock()
	defer list.mtx.RUnlock()

	task, ok := list.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (list *List) ListTask() map[string]Task {
	list.mtx.RLock()
	defer list.mtx.RUnlock()

	tmp := make(map[string]Task, len(list.tasks))

	for k, v := range list.tasks {
		tmp[k] = v
	}
	return tmp
}

func (list *List) ListUncompletedTask() map[string]Task {
	list.mtx.RLock()
	defer list.mtx.RUnlock()

	UncompletedTask := make(map[string]Task)

	for tatle, task := range list.tasks {
		if !task.Completed {
			UncompletedTask[tatle] = task
		}
	}
	return nil
}

func (list *List) CompletedTask(title string) (Task, error) {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	task, ok := list.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Complete()
	list.tasks[title] = task

	return list.tasks[title], nil
}
func (list *List) UncompletedTask(title string) (Task, error) {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	task, ok := list.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Uncompleted()
	list.tasks[title] = task
	return list.tasks[title], nil
}

func (list *List) DeleteTask(title string) error {
	list.mtx.Lock()
	defer list.mtx.Unlock()

	_, ok := list.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	delete(list.tasks, title)
	return nil
}
