package repository

import (
	"github.com/RidhoFahrizal/Golang-Personal-Blog/internal/model"
)

type PostRepository interface {
	Create(post *model.Post) error 
	GetByID(id string) (*model.Post, error)
	List() ([]model.Post, error)
	Delete(id string) error
	Update(post *model.Post) error
}


