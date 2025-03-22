package services

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestGenerateRandomBook(t *testing.T) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("Skipping test: OPENAI_API_KEY not set")
	}
	fmt.Println("Running test: TestGenerateRandomBook")

	ctx := context.Background()
	book, err := GenerateRandomBook(ctx)
	if err != nil {
		t.Fatalf("GenerateRandomBook failed: %v", err)
	}

	if book.Title == "" {
		t.Error("Book title should not be empty")
	}
	if book.Author == "" {
		t.Error("Book author should not be empty")
	}
	if book.Description == "" {
		t.Error("Book description should not be empty")
	}
	if book.Genre == "" {
		t.Error("Book genre should not be empty")
	}
	if book.Year == 0 {
		t.Error("Book year should not be zero")
	}
}
