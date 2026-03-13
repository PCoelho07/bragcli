package brag_test

import (
	"brag/internal/brag"
	"testing"
)

func TestNewBragItem(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		item, err := brag.NewBragItem("shipped auth refactor", "replaced JWT with opaque tokens")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if item.Title != "shipped auth refactor" {
			t.Errorf("expected title %q, got %q", "shipped auth refactor", item.Title)
		}
		if item.Description != "replaced JWT with opaque tokens" {
			t.Errorf("expected description %q, got %q", "replaced JWT with opaque tokens", item.Description)
		}
		if item.CreatedAt.IsZero() {
			t.Error("expected CreatedAt to be set")
		}
	})

	t.Run("blank title", func(t *testing.T) {
		_, err := brag.NewBragItem("", "some description")
		if err == nil {
			t.Fatal("expected error for blank title, got nil")
		}
	})

	t.Run("whitespace-only title", func(t *testing.T) {
		_, err := brag.NewBragItem("   ", "some description")
		if err == nil {
			t.Fatal("expected error for whitespace-only title, got nil")
		}
	})

	t.Run("blank description", func(t *testing.T) {
		_, err := brag.NewBragItem("some title", "")
		if err == nil {
			t.Fatal("expected error for blank description, got nil")
		}
	})

	t.Run("whitespace-only description", func(t *testing.T) {
		_, err := brag.NewBragItem("some title", "   ")
		if err == nil {
			t.Fatal("expected error for whitespace-only description, got nil")
		}
	})
}
