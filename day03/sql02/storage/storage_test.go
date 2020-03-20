package storage_test

import (
	"sql02/storage"
	"testing"
)

func TestDummyStorage(t *testing.T) {
	// Create storage
	s := storage.NewDummyStorage()
	u := storage.User{
		ID: 1,
		Name: "Dani",
		Age: 22,
	}
	if _, err := s.GetUser(u.ID); err != storage.ErrNotFound {
		t.Fatalf("Expected %s got %s", storage.ErrNotFound, err)
	}
	if err := s.CreateUser(&u); err != nil {
		t.Fatal(err)
	}

	u2, err := s.GetUser(u.ID)
	if err != nil {
		t.Fatal(err)
	}
	if u != u2 {
		t.Fatalf("Expected %v got %v", u, u2)
	}
}
