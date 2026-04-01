package todo

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExistst = errors.New("task already exists")
