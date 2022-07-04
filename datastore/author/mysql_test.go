package author

import (
	"ThreeLayer/driver"
	"ThreeLayer/entities"
	"database/sql"
	"log"
	"reflect"
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
		log.Print(err)
	}
	return db
}

func TestDatastore(t *testing.T) {
	db := startMySql(t)
	a := New(db)
	testAuthorStorer_Create(t, a)
	testAuthorStorer_Put(t, a)
	testAuthorStorer_Delete(t, a)
}

func testAuthorStorer_Create(t *testing.T, db AuthorStorer) {

	testcases := []struct {
		desc     string
		req      entities.Author
		response entities.Author
	}{
		{"Valid details", entities.Author{FirstName: "JP", LastName: "Garg", Dob: "2/11/1999", PenName: "Sharma"},
			entities.Author{1, "JP", "Garg", "2/11/1999", "Sharma"}},
		{"Valid details", entities.Author{FirstName: "hc", LastName: "Verma", Dob: "2/11/1999",
			PenName: "Sharma"},
			entities.Author{2, "hc", "Verma", "2/11/1999", "Sharma"}},
	}

	for i, v := range testcases {
		resp, _ := db.CreateAuthor(v.req)
		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}

func testAuthorStorer_Put(t *testing.T, db AuthorStorer) {
	testcases := []struct {
		desc    string
		reqID   int
		reqData entities.Author
	}{
		{"Valid case UPDATE firstname.", 1, entities.Author{1, "Gurpreet", "Saini", "22/08/2000", "GP"}},
		//	{"Valid case UPDATE dob", 1, entities.Author{1, "Gurpreet", "Saini", "01/07/2001", "GP"}},
	}
	for i, v := range testcases {
		resp, err := db.PutAuthor(v.reqID, v.reqData)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(resp, v.reqData) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.reqData)
		}
	}
}
func testAuthorStorer_Delete(t *testing.T, db AuthorStorer) {
	testcases := []struct {
		desc     string
		reqID    int
		expError error
	}{
		{"success test case with id 1", 1, nil},
	}
	for i, v := range testcases {
		err := db.DeleteAuthor(v.reqID)
		if err != v.expError {
			t.Errorf("[TEST%d]Failed. Expected %v\n", i+1, v.expError)

		}

	}
}
