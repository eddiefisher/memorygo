package store_test

import (
	"testing"

	"github.com/eddiefisher/memorygo/src/store"
)

//Set means "store this data"
func TestSet(t *testing.T) {
	store.Messages = map[string]string{}
	if len(store.Messages) > 0 {
		t.Error("map Messages is not empty", len(store.Messages))
	}
	store.Set("1", "a")
	if len(store.Messages) == 0 {
		t.Error("map Messages is empty, value is not set")
	}
}

//Add means "store this data, but only if the server *doesn't* already hold data for this key".
func TestAdd(t *testing.T) {
	store.Messages = map[string]string{}
	store.Set("1", "a")
	if store.Add("1", "b") == nil {
		t.Error("key is exist")
	}
	if a, _ := store.Get("1"); a != "a" {
		t.Error("key was changed")
	}
	if store.Add("2", "c") != nil {
		t.Error("key not exist, but have error")
	}
}

//Replace means "store this data, but only if the server *does* already hold data for this key".
func TestReplace(t *testing.T) {
	store.Messages = map[string]string{}
	store.Set("1", "a")
	if err := store.Replace("1", "b"); err != nil {
		t.Error("key exist but was not replaced", err)
	}
	if b, _ := store.Get("1"); b != "b" {
		t.Error("key was not replaced")
	}
	if err := store.Replace("2", "b"); err == nil {
		t.Error("key is not exist", err)
	}
}

//Append means "add this data to an existing key after existing data".
func TestAppend(t *testing.T) {
	store.Messages = map[string]string{}
	store.Set("1", "a")
	if err := store.Append("1", "b"); err != nil {
		t.Error("key exist but was not append", err)
	}
	if b, _ := store.Get("1"); b != "ab" {
		t.Error("key was not append")
	}
	if err := store.Append("2", "b"); err == nil {
		t.Error("key is not exist", err)
	}
}

//Prepend means "add this data to an existing key before existing data".
func TestPrepend(t *testing.T) {
	store.Messages = map[string]string{}
	store.Set("1", "a")
	if err := store.Prepend("1", "b"); err != nil {
		t.Error("key exist but was not prepend", err)
	}
	if b, _ := store.Get("1"); b != "ba" {
		t.Error("key was not prepend")
	}
	if err := store.Prepend("2", "b"); err == nil {
		t.Error("key is not exist", err)
	}
}
