package main

import (
	"database/sql"
)

type Store interface {
	CreateBook(book *Book) error
	GetBooks() ([]Book, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBook(book *Book) error {
	_, err := store.db.Query("INSERT INTO books(isbn, title) VALUES ($1,$2)", book.Isbn, book.Title)
	return err
}

func (store *dbStore) GetBooks() ([]*Book, error) {
	rows, err := store.db.Query("SELECT isbn, title from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*Book{}
	for rows.Next() {
		book := &Book{}
		if err := rows.Scan(&book.Isbn, &book.Title); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
