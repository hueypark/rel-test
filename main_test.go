package main

import (
	"context"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	// create a mocked repository.
	var (
		repo = reltest.New()
		book = Book{
			ID:       1,
			Title:    "Go for dummies",
			Category: "learning",
			AuthorID: 1,
		}
		author = Author{ID: 1, Name: "CZ2I28 Delta"}
	)

	// mock find and return result
	repo.ExpectFind(where.Eq("id", 1)).Result(book)

	// mock find and return result using query builder.
	repo.ExpectFind(rel.Select().Where(where.Eq("id", 1)).Limit(1)).Result(book)

	// mock preload and return result
	repo.ExpectPreload("author").ForType("main.Book").Result(author)

	// mocks transaction
	repo.ExpectTransaction(func(repo *reltest.Repository) {
		// mock updates
		repo.ExpectUpdate().ForType("main.Book")
		repo.ExpectUpdate(rel.Set("discount", false)).ForType("main.Book")
		repo.ExpectUpdate(rel.Dec("stock")).ForType("main.Book")
	})

	// run and asserts
	assert.Nil(t, Example(context.Background(), repo))
	repo.AssertExpectations(t)
}
