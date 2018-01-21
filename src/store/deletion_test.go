package store_test

import (
	"testing"

	"github.com/eddiefisher/memorygo/src/store"
)

func TestDeleteExist(t *testing.T) {
	store.Set("a", "1")
	if err := store.Delete("a"); err != nil {
		t.Error(err)
	}
	s, err := store.Get("a")
	if s == "1" {
		t.Error(err)
	}
}

func TestDeleteNotExist(t *testing.T) {
	if err := store.Delete("a"); err == nil {
		t.Error(err)
	}
}
