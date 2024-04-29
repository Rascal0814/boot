package crud

import (
	"testing"
)

func TestCrud(t *testing.T) {
	curd, err := NewGenCurd("mysql://user:password@tcp(xxx:3306)/xxx?charset=utf8&parseTime=True&loc=Local", "test/model", "dao")
	if err != nil {
		t.Fatal(err)
	}

	err = curd.Gen()
	if err != nil {
		t.Fatal(err)
	}
}
