package storage_test

import (
	"sql02/storage"
	"testing"
)

func TestDummyStorage(t *testing.T) {
	// Create new storage
	s := storage.NewDummyStorage()

	s.CreateUser(&storage.User{ID: 34, Name: "Extra", Age: 21})

	u := storage.User{
		ID:   1,
		Name: "Dani",
		Age:  22,
	}
	// try to get not existing user
	if _, err := s.GetUser(u.ID); err != storage.ErrNotFound {
		t.Fatalf("Expected %s got %s", storage.ErrNotFound, err)
	}
	// try to create user
	if err := s.CreateUser(&u); err != nil {
		t.Fatal(err)
	}
	// try to create duplicate user
	if err := s.CreateUser(&u); err != storage.ErrDuplicate {
		t.Fatalf("Expected %s got %s", storage.ErrNotFound, err)
	}

	t.Cleanup(func() { s.DeleteUser(u.ID) })

	u2, err := s.GetUser(u.ID)
	if err != nil {
		t.Fatal(err)
	}

	if u != u2 {
		t.Fatalf("Expected %v got %v", u, u2)
	}

	if err := s.UpdateUser(&storage.User{ID: 34232}); err != storage.ErrNotFound {
		t.Fatalf("Expected %s got %s", storage.ErrNotFound, err)
	}

	u2.Age = 23

	if err := s.UpdateUser(&u2); err != nil {
		t.Fatal(err)
	}

	u, err = s.GetUser(u.ID)
	if err != nil {
		t.Fatal(err)
	}

	if u != u2 {
		t.Fatalf("Expected %v got %v", u2, u)
	}

	// try to delete user
	if err := s.DeleteUser(u.ID); err != nil {
		t.Fatal(err)
	}

	// try to delete user a second time
	if err := s.DeleteUser(u.ID); err != storage.ErrNotFound {
		t.Fatalf("Expected %s got %s", storage.ErrNotFound, err)
	}
}
