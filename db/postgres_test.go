package db

import (
	"../schema"
	"../testdb"
	"reflect"
	"testing"
	"time"
)

func TestPostgres_Insert(t *testing.T) {
	postgres := &Postgres{testdb.Setup()}
	defer postgres.Close()

	todo := &schema.Todo{
		Title:   "title1",
		Note:    "note1",
		DueDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}

	got, err := postgres.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	want := 1

	if got != want {
		t.Fatal(err)
	}
}




func equal(got interface{}, want interface{}) bool {
	return reflect.DeepEqual(got, want)
}