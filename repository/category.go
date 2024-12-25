package repository

import (
	"database/sql"
	"fmt"
	"quiz-sanbercode/structs"
	"time"
)

func GetAllCategories(db *sql.DB) ([]structs.Category, error) {
	sql := `SELECT id, name, created_at, created_by, modified_at, modified_by 
            FROM categories`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		err = rows.Scan(
			&category.ID, &category.Name, &category.CreatedAt,
			&category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategory(db *sql.DB, id int) (structs.Category, error) {
	var category structs.Category
	sql := `SELECT id, name, created_at, created_by, modified_at, modified_by 
            FROM categories WHERE id = $1`

	err := db.QueryRow(sql, id).Scan(
		&category.ID, &category.Name, &category.CreatedAt,
		&category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy,
	)
	return category, err
}

func CreateCategory(db *sql.DB, category structs.Category) error {
	sql := `INSERT INTO categories (name, created_at, created_by) 
            VALUES ($1, $2, $3)`

	_, err := db.Exec(sql, category.Name, time.Now(), category.CreatedBy)
	return err
}

func UpdateCategory(db *sql.DB, category structs.Category) error {
	sql := `UPDATE categories SET name = $1, modified_at = $2, modified_by = $3 
            WHERE id = $4`

	_, err := db.Exec(sql, category.Name, time.Now(),
		category.ModifiedBy, category.ID)
	return err
}

func DeleteCategory(db *sql.DB, id int) error {
	// First check if category has any books
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM books WHERE category_id = $1", id).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("cannot delete category with existing books")
	}

	sql := "DELETE FROM categories WHERE id = $1"
	_, err = db.Exec(sql, id)
	return err
}
