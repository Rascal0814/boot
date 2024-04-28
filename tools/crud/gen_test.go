package orm

import (
	"testing"
)

func TestCrud(t *testing.T) {
	curd, err := NewGenCurd("mysql://root:123456@tcp(43.143.80.123:3306)/order-dishes?charset=utf8&parseTime=True&loc=Local", "test/model", "dao")
	if err != nil {
		t.Fatal(err)
	}

	err = curd.Gen()
	if err != nil {
		t.Fatal(err)
	}
}
