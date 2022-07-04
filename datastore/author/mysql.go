package author

import (
	"ThreeLayer/entities"
	"database/sql"
	"errors"
	"log"
)

type AuthorStorer struct {
	db *sql.DB
}

func New(db *sql.DB) AuthorStorer {
	return AuthorStorer{db: db}
}

func (a AuthorStorer) CreateAuthor(author entities.Author) (entities.Author, error) {
	res, err := a.db.Exec("INSERT INTO Authors (first_name, last_name, dob, pen_name)\nVALUES (?,?,?,?);",
		author.FirstName, author.LastName, author.Dob, author.PenName)
	if err != nil {
		log.Print(err)
	}
	id, _ := res.LastInsertId()
	author.ID = int(id)
	return author, nil
}

func (a AuthorStorer) PutAuthor(id int, author entities.Author) (entities.Author, error) {
	_, err := a.db.Exec("UPDATE Authors SET first_name=? ,last_name=? ,dob=? ,pen_name=? WHERE id=?",
		author.FirstName, author.LastName, author.Dob, author.PenName, id)
	if err != nil {
		return entities.Author{}, err
	}
	return author, nil
}
func (a AuthorStorer) DeleteAuthor(id int) error {
	_, err := a.db.Exec("DELETE FROM Authors WHERE id=?;", id)
	if err != nil {
		return errors.New("INVALID ID FOUND")
	}
	return nil
}
