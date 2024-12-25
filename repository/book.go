package repository

import (
	"database/sql"
	"quiz-sanbercode/structs"
	"time"
)

func GetAllBooks(db *sql.DB) ([]structs.Book, error) {
	sql := `SELECT id, title, description, image_url, release_year, price, 
            total_page, thickness, category_id, created_at, created_by, 
            modified_at, modified_by FROM books`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var book structs.Book
		err = rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.ImageURL,
			&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
			&book.CategoryID, &book.CreatedAt, &book.CreatedBy,
			&book.ModifiedAt, &book.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBook(db *sql.DB, id int) (structs.Book, error) {
	var book structs.Book
	sql := `SELECT id, title, description, image_url, release_year, price, 
            total_page, thickness, category_id, created_at, created_by, 
            modified_at, modified_by FROM books WHERE id = $1`

	err := db.QueryRow(sql, id).Scan(
		&book.ID, &book.Title, &book.Description, &book.ImageURL,
		&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
		&book.CategoryID, &book.CreatedAt, &book.CreatedBy,
		&book.ModifiedAt, &book.ModifiedBy,
	)
	return book, err
}

func CreateBook(db *sql.DB, book structs.Book) error {
	sql := `INSERT INTO books (title, description, image_url, release_year, 
            price, total_page, thickness, category_id, created_at, created_by) 
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// Calculate thickness based on total_page
	thickness := "tipis"
	if book.TotalPage > 100 {
		thickness = "tebal"
	}

	_, err := db.Exec(sql,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear,
		book.Price, book.TotalPage, thickness, book.CategoryID,
		time.Now(), book.CreatedBy,
	)
	return err
}

func UpdateBook(db *sql.DB, book structs.Book) error {
	sql := `UPDATE books SET title = $1, description = $2, image_url = $3, 
            release_year = $4, price = $5, total_page = $6, thickness = $7, 
            category_id = $8, modified_at = $9, modified_by = $10 
            WHERE id = $11`

	// Calculate thickness based on total_page
	thickness := "tipis"
	if book.TotalPage > 100 {
		thickness = "tebal"
	}

	_, err := db.Exec(sql,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear,
		book.Price, book.TotalPage, thickness, book.CategoryID,
		time.Now(), book.ModifiedBy, book.ID,
	)
	return err
}

func DeleteBook(db *sql.DB, id int) error {
	sql := "DELETE FROM books WHERE id = $1"
	_, err := db.Exec(sql, id)
	return err
}

func GetBooksByCategory(db *sql.DB, categoryID int) ([]structs.Book, error) {
	sql := `SELECT id, title, description, image_url, release_year, price, 
            total_page, thickness, category_id, created_at, created_by, 
            modified_at, modified_by FROM books WHERE category_id = $1`

	rows, err := db.Query(sql, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var book structs.Book
		err = rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.ImageURL,
			&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
			&book.CategoryID, &book.CreatedAt, &book.CreatedBy,
			&book.ModifiedAt, &book.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
