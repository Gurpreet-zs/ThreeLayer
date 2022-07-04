package book

import (
	"ThreeLayer/entities"
	"database/sql"
	"errors"
	"log"
)

type BookStorer struct {
	db *sql.DB
}

func New(db *sql.DB) BookStorer {
	return BookStorer{db: db}
}

func (a BookStorer) GetAllBook(title, includeAuthor string) ([]entities.Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if err != nil {
		log.Print(err)
	}
	if title == "" {
		rows, err = a.db.Query("select * from Books;")
	} else {
		rows, err = a.db.Query("select * from Books where title=?;", title)
	}

	books := []entities.Book{}

	for rows.Next() {
		book := entities.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.ID)

		if err != nil {
			log.Print(err)
		}

		if includeAuthor == "true" {
			row := a.db.QueryRow("select * from Authors where id=?", book.Author.ID)
			err = row.Scan(&book.Author.ID, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)

			if err != nil {
				log.Print(err)
			}
		}

		books = append(books, book)
	}

	return books, nil

}

func (a BookStorer) GetBookByID(id int) (entities.Book, error) {
	row := a.db.QueryRow("SELECT * FROM Books WHERE id=?;", id)
	var book entities.Book
	err := row.Scan(&book.ID, &book.Title, &book.Publication, &book.PublishedDate, &book.Author.ID)
	if err != nil {
		return entities.Book{}, err
	}

	authorRow := a.db.QueryRow("SELECT * FROM Authors WHERE id=?", book.Author.ID)
	err = authorRow.Scan(&book.Author.ID, &book.Author.FirstName, &book.Author.LastName, &book.Author.Dob, &book.Author.PenName)
	if err != nil {
		return entities.Book{}, err
	}
	return book, nil
}

func (a BookStorer) CreateBook(book entities.Book) (entities.Book, error) {

	res, err := a.db.Exec("INSERT INTO Books (title,publication,publication_date,author_id) VALUES(?,?,?,?)", book.Title, book.Publication,
		book.PublishedDate, book.Author.ID)
	if err != nil {
		return entities.Book{}, err
	}

	id, _ := res.LastInsertId()
	book.ID = int(id)

	return book, nil
}

func (a BookStorer) UpdateBook(id int, book entities.Book) (entities.Book, error) {
	_, err := a.db.Exec("UPDATE Books SET title = ? ,publication = ? ,publication_date = ?,author_id=?  WHERE id =?",
		&book.Title, &book.Publication, &book.PublishedDate, &book.Author.ID, id)
	if err != nil {
		return entities.Book{}, err
	}

	return book, nil
}

func (a BookStorer) DeleteBook(id int) error {
	_, err := a.db.Exec("delete from Books where id=?;", id)
	if err != nil {
		return errors.New("InValid ID")
	}

	return nil
}
