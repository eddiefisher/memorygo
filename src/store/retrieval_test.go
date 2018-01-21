package store_test

import (
	"testing"

	"github.com/eddiefisher/memorygo/src/store"
)

func TestGet(t *testing.T) {
	store.Messages = map[string]string{}
	store.Set("1", "a")
	a, _ := store.Get("1")
	if a != "a" {
		t.Error("can't get existed record")
	}

	_, err := store.Get("2")
	if err == nil {
		t.Error("got existed record", err)
	}
}

func TestGets(t *testing.T) {
	store.Messages = map[string]string{}
	store.Set("1", "a")
	store.Set("2", "b")
	store.Set("3", "c")
	a, _ := store.Gets("1", "3")
	if a[0] != "a" && a[1] != "c" && len(a) != 2 {
		t.Error("can't get existed records")
	}

	_, err := store.Gets("1", "4")
	if err == nil {
		t.Error("got existed records")
	}
}
