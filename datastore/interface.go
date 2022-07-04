package datastore

import "ThreeLayer/entities"

type Author interface {
	CreateAuthor(author entities.Author) (entities.Author, error) //post
	PutAuthor(author entities.Author) (entities.Author, error)
	DeleteAuthor(id int) (entities.Author, error)
}

type Book interface {
	GetAllBook(book entities.Book) (entities.Book, error)
	GetBookByID(int entities.Book) (entities.Book, error)
	CreateBook(book entities.Book) (entities.Book, error)
	UpdateBook(book entities.Book) (entities.Book, error)
	DeleteBook(book entities.Book) (entities.Book, error)
}
