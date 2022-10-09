package cli

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCanOpenDB(t *testing.T) {
	db, err := sql.Open("sqlite3", "python/gpa_dataset.db")

	if err != nil {
		t.Errorf("could not open database")
	}

	db.Close()
}

func TestInvalidGETThrowsError(t *testing.T) {

}
