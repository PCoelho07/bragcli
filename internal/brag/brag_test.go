package brag_test

import (
	"brag/internal/brag"
	"errors"
	"testing"
)

type stubStore struct {
	initialized   bool
	isInitErr     error
	initializeErr error
	saveErr       error
	savedItem     *brag.BragItem
}

func (s *stubStore) IsInitialized() (bool, error) {
	return s.initialized, s.isInitErr
}

func (s *stubStore) Initialize() error {
	return s.initializeErr
}

func (s *stubStore) Save(item *brag.BragItem) error {
	s.savedItem = item
	return s.saveErr
}

func (s *stubStore) ReadAll() ([]brag.BragItem, error) {
	return nil, nil
}

func TestBragInit(t *testing.T) {
	t.Run("initializes when not yet initialized", func(t *testing.T) {
		stub := &stubStore{initialized: false}
		if err := brag.New(stub).Init(); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("no-op when already initialized", func(t *testing.T) {
		stub := &stubStore{initialized: true, initializeErr: errors.New("should not be called")}
		if err := brag.New(stub).Init(); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("propagates IsInitialized error", func(t *testing.T) {
		storeErr := errors.New("disk error")
		stub := &stubStore{isInitErr: storeErr}
		if err := brag.New(stub).Init(); !errors.Is(err, storeErr) {
			t.Fatalf("expected %v, got %v", storeErr, err)
		}
	})

	t.Run("propagates Initialize error", func(t *testing.T) {
		storeErr := errors.New("disk full")
		stub := &stubStore{initialized: false, initializeErr: storeErr}
		if err := brag.New(stub).Init(); !errors.Is(err, storeErr) {
			t.Fatalf("expected %v, got %v", storeErr, err)
		}
	})
}

func TestBragCreate(t *testing.T) {
	t.Run("saves item on valid input", func(t *testing.T) {
		stub := &stubStore{initialized: true}
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

	t.Run("returns ErrNotInitialized when not initialized", func(t *testing.T) {
		stub := &stubStore{initialized: false}
		err := brag.New(stub).Create("valid title", "valid description")
		if !errors.Is(err, brag.ErrNotInitialized) {
			t.Fatalf("expected ErrNotInitialized, got %v", err)
		}
	})

	t.Run("returns error without calling store on blank title", func(t *testing.T) {
		stub := &stubStore{initialized: true}
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
		stub := &stubStore{initialized: true, saveErr: storeErr}
		err := brag.New(stub).Create("valid title", "valid description")
		if !errors.Is(err, storeErr) {
			t.Fatalf("expected %v, got %v", storeErr, err)
		}
	})
}
