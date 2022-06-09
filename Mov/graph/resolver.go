package graph

import (
	"github.com/arukasar/golang-graphql/repository"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MovieRepository repository.MovieRepository
	Db              *gorm.DB
}
