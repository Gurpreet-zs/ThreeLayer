package book

import (
	"ThreeLayer/driver"
	"ThreeLayer/entities"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func startMySql(t *testing.T) *sql.DB {
	conf := driver.MySQLConfig{
		Host:     "localhost",
		User:     "root",
		Password: "Gurpreet@0848",
		Port:     "3306",
		Db:       "test",
	}

	var err error
	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		t.Errorf("could not connect to sql, err:%v", err)
	}

	return db
}

func TestDatastore(t *testing.T) {
	db := startMySql(t)
	a := New(db)
	testCreateBook(t, a)
	testGetAllBook(t, a)
	testGetBookByID(t, a)
	testUpdateBook(t, a)
	testDeleteBook(t, a)
}

func testGetAllBook(t *testing.T, db BookStorer) {
	testCases := []struct {
		desc          string
		title         string
		includeAuthor string
		expRes        []entities.Book
		expErr        error
	}{
		{"get all books", "", "", []entities.Book{{1, "the city", entities.Author{ID: 2}, "Scholastic", "26/5/2017"}}, nil},
		//	{"get all books with query param", "the wall", "", []entities.Book{{1, "the wall", entities.Author{}, "Penguin", "25/03/2001"}}, nil},
		//	{"get all books with query param", "", "true", []entities.Book{{1, "the wall", entities.Author{}, "Penguin", "10/02/2008"}}, nil},
	}
	for i, v := range testCases {
		resp, err := db.GetAllBook(v.title, v.includeAuthor)

		assert.Equalf(t, err, v.expErr, "%v test case failed %v", i, v.desc)
		assert.Equalf(t, resp, v.expRes, "%v test case failed %v", i, v.desc)
	}
}

func testGetBookByID(t *testing.T, db BookStorer) {
	testCases := []struct {
		desc     string
		id       int
		response entities.Book
		expErr   error
	}{
		{"valid case", 1, entities.Book{ID: 1, Title: "the city", Author: entities.Author{ID: 2, FirstName: "hc", LastName: "Verma", Dob: "2/11/1999", PenName: "Sharma"}, Publication: "Scholastic", PublishedDate: "26/5/2017"}, nil},
		//	{"valid case", 2, entities.Book{ID: 1, Title: "the city", Author: entities.Author{ID: 2, FirstName: "RD", LastName: "Sharma", Dob: "09-05-2004", PenName: "RD"}, Publication: "Scholastic", PublishedDate: "26/5/2017"}, nil},
		//	{"invalid case: id not found", 100, entities.Book{}, errors.New("missing ID")},
	}
	for i, v := range testCases {
		resp, err := db.GetBookByID(v.id)

		assert.Equalf(t, err, v.expErr, "%v test case failed %v", i, v.desc)
		assert.Equalf(t, resp, v.response, "%v test case failed %v", i, v.desc)
	}
}

func testCreateBook(t *testing.T, db BookStorer) {
	testCases := []struct {
		desc     string
		req      entities.Book
		response entities.Book
		expErr   error
	}{
		//{"valid case: successfully posted", entities.Book{Title: "the wall", Author: entities.Author{ID: 1, FirstName: "JP", LastName: "Garg", Dob: "2/11/1999", PenName: "Sharma"}, Publication: "ABC", PublishedDate: "6/7/2017"}, entities.Book{Title: "the wall", Author: entities.Author{1, "JP", "Garg", "2/11/1999", "Sharma"}, Publication: "ABC", PublishedDate: "6/7/2017"}, nil},
		//{"Valid Case  successfullt posted", entities.Book{Title: "the wall", Author: entities.Author{1, "JP", "Garg", "2/11/1999", "Sharma"}, Publication: "ABC", PublishedDate: "6/7/2017"}, entities.Book{ID: 1, Title: "the wall", Author: entities.Author{1, "JP", "Garg", "2/11/1999", "Sharma"}, Publication: "ABC", PublishedDate: "6/7/2017"}, nil},
		{"valid case: successfully posted", entities.Book{Title: "the city", Author: entities.Author{ID: 2, FirstName: "hc", LastName: "Verma", Dob: "2/11/1999", PenName: "Sharma"}, Publication: "Scholastic", PublishedDate: "26/5/2017"}, entities.Book{ID: 1, Title: "the city", Author: entities.Author{2, "hc", "Verma", "2/11/1999", "Sharma"}, Publication: "Scholastic", PublishedDate: "26/5/2017"}, nil},
	}
	for i, v := range testCases {
		resp, err := db.CreateBook(v.req)

		assert.Equalf(t, err, v.expErr, "%v test case failed %v", i, v.desc)
		assert.Equalf(t, resp, v.response, "%v test case failed %v", i, v.desc)
	}
}

func testUpdateBook(t *testing.T, db BookStorer) {
	testCases := []struct {
		desc     string
		id       int
		req      entities.Book
		response entities.Book
		expErr   error
	}{
		{"Valid Details", 1, entities.Book{Title: "the wall", Author: entities.Author{ID: 2}, Publication: "Arihanth", PublishedDate: "04/05/2017"}, entities.Book{Title: "the wall", Author: entities.Author{ID: 2}, Publication: "Arihanth", PublishedDate: "04/05/2017"}, nil},
		//	{"Invalid Publication", 1, entities.Book{Title: "city", Author: entities.Author{ID: 1}, Publication: "Arihant", PublishedDate: "17/8/2001"}, entities.Book{}, errors.New("Invalid Publication")},
		//	{"Invalid Date", 1, entities.Book{Title: "the wall", Author: entities.Author{ID: 1}, Publication: "the sun", PublishedDate: "13/8/2004"}, entities.Book{}, errors.New("Invalid Date")},
		//	{"Missing details", 1, entities.Book{}, entities.Book{}, errors.New("Missing details")},
	}

	for i, v := range testCases {
		resp, err := db.UpdateBook(v.id, v.req)

		assert.Equalf(t, err, v.expErr, "%v test case failed %v", i, v.desc)
		assert.Equalf(t, resp, v.response, "%v test case failed %v", i, v.desc)
	}
}

func testDeleteBook(t *testing.T, db BookStorer) {
	testCases := []struct {
		desc     string
		id       int
		response error
	}{
		{"Valid details", 1, nil},
		//	{"Book does not exist", 100, errors.New("Missing Book")},
	}

	for i, v := range testCases {
		resp := db.DeleteBook(v.id)

		assert.Equalf(t, resp, v.response, "%v test case failed %v", i, v.desc)
	}

}
