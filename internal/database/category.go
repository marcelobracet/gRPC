package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/marcelobracet/grpc/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Category struct {
	DB        *sql.DB
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(db *sql.DB) *Category {
	return &Category{DB: db}
}

func (c *Category) CreateCategory(name string, description *string) error {
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (c *Category) ListCategories() ([]*pb.Category, error) {
	rows, err := c.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*pb.Category
	for rows.Next() {
		var category pb.Category
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (c *Category) UpdateCategory(id string, name string, description *string) error {
	_, err := c.DB.Exec("UPDATE categories SET name = $1, description = $2 WHERE id = $3", name, description, id)
	return err
}
