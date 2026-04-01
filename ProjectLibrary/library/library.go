package library

import "sync"

type Library struct {
	books map[string]Book
	mtx   sync.Mutex
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}

func (lib *Library) AddBook(book Book) error {
	lib.mtx.Lock(
	defer lib.mtx.Unlock()

	if _, ok := lib.books[book.Name]; ok {
		return ErrBookAlreadyExist
	}

	lib.books[book.Name] = book
	return nil
}

func (lib *Library) GetBook(name string) (Book, error) {
	lib.mtx.Lock()
	defer lib.mtx.Unlock()

	book, ok := lib.books[name]
	if !ok {
		return Book{}, ErrBookNotFound
	}
	return book, nil
}

func (lib *Library) ListBook() map[string]Book {
	lib.mtx.Lock()
	defer lib.mtx.Unlock()

	copyMap := make(map[string]Book, len(lib.books))

	for k, v := range lib.books {
		copyMap[k] = v
	}
	return copyMap
}

func (lib *Library) ListUncompletedBook() map[string]Book {
	lib.mtx.Lock()
	defer lib.mtx.Unlock()

	uncompeledCopyMap := make(map[string]Book)

	for name, book := range lib.books {
		if !book.ReadingCompleted {
			uncompeledCopyMap[name] = book
		}
	}
	return uncompeledCopyMap
}

func (lib *Library) CompletedBook(name string) (Book, error) {
	lib.mtx.Lock()
	defer lib.mtx.Unlock()

	book, ok := lib.books[name]
	if !ok {
		return Book{}, ErrBookNotFound
	}
	book.Completed()
	lib.books[name] = book
	return lib.books[name], nil
}

func (lib *Library) UncompletedBook(name string) (Book, error) {
	lib.mtx.Lock()
	defer lib.mtx.Unlock()

	book, ok := lib.books[name]
	if !ok {
		return Book{}, ErrBookNotFound
	}
	book.Uncompleted()
	lib.books[name] = book
	return lib.books[name], nil
}

func (lib *Library) DeleteBook(name string) error {
	lib.mtx.Lock()
	defer lib.mtx.Unlock()

	_, ok := lib.books[name]
	if !ok {
		return ErrBookNotFound
	}
	delete(lib.books, name)
	return nil
}
