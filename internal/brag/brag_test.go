package brag_test

import (
	"brag/internal/brag"
	"errors"
	"testing"
)

type stubStore struct {
	initErr   error
	saveErr   error
	savedItem *brag.BragItem
}

func (s *stubStore) Init() error {
	return s.initErr
}

func (s *stubStore) Save(item *brag.BragItem) error {
	s.savedItem = item
	return s.saveErr
}

func TestBragInit(t *testing.T) {
	t.Run("delegates to store", func(t *testing.T) {
		stub := &stubStore{}
		if err := brag.New(stub).Init(); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("propagates store error", func(t *testing.T) {
		storeErr := errors.New("disk full")
		stub := &stubStore{initErr: storeErr}
		if err := brag.New(stub).Init(); !errors.Is(err, storeErr) {
			t.Fatalf("expected %v, got %v", storeErr, err)
		}
	})
}

func TestBragCreate(t *testing.T) {
	t.Run("saves item on valid input", func(t *testing.T) {
		stub := &stubStore{}
		err := brag.New(stub).Create("led incident response", "resolved p0 outage in 23 minutes")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if stub.savedItem == nil {
			t.Fatal("expected store.Save to be called, but savedItem is nil")
		}
		if stub.savedItem.Title != "led incident response" {
			t.Errorf("expected title %q, got %q", "led incident response", stub.savedItem.Title)
		}
	})

	t.Run("returns error without calling store on blank title", func(t *testing.T) {
		stub := &stubStore{}
		err := brag.New(stub).Create("", "some description")
		if err == nil {
			t.Fatal("expected error for blank title, got nil")
		}
		if stub.savedItem != nil {
			t.Error("expected store.Save not to be called")
		}
	})

	t.Run("propagates store error", func(t *testing.T) {
		storeErr := errors.New("disk full")
		stub := &stubStore{saveErr: storeErr}
		err := brag.New(stub).Create("valid title", "valid description")
		if !errors.Is(err, storeErr) {
			t.Fatalf("expected %v, got %v", storeErr, err)
		}
	})
}
