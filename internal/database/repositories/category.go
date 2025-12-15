package repositories

import "github.com/marcelobracet/grpc/internal/database"

type CategoryRepository interface {
	Create(name string) error
	GetAll() ([]database.Category, error)
	GetById(id uint) (database.Category, error)
	Update(id uint, name string) error
	Delete(id uint) error
}
