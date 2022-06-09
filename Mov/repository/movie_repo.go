package repository

import (
	"github.com/arukasar/golang-graphql/graph/model"
	"github.com/arukasar/golang-graphql/models"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movieInput *model.MovieInput) (*models.Movie, error)
	UpdateMovie(movieInput *model.MovieInput, id int) error
	DeleteMovie(id int) error
	GetOneMovie(id int) (*models.Movie, error)
	GetAllMovies() ([]*model.Movie, error)
}

type MovieService struct {
	Db *gorm.DB
}

var _ MovieRepository = &MovieService{}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{
		Db: db,
	}
}

func (b *MovieService) CreateMovie(movieInput *model.MovieInput) (*models.Movie, error) {
	movie := &models.Movie{
		Title:       movieInput.Title,
		URL:         movieInput.URL,
		ReleaseDate: movieInput.ReleaseDate,
	}
	err := b.Db.Create(&movie).Error

	return movie, err
}
func (b *MovieService) UpdateMovie(movieInput *model.MovieInput, id int) error {
	movie := models.Movie{
		ID:          id,
		Title:       movieInput.Title,
		URL:         movieInput.URL,
		ReleaseDate: movieInput.ReleaseDate,
	}
	err := b.Db.Model(&movie).Where("id = ?", id).Updates(movie).Error
	return err
}

func (b *MovieService) DeleteMovie(id int) error {
	movie := &models.Movie{}
	err := b.Db.Delete(movie, id).Error
	return err
}

func (b *MovieService) GetOneMovie(id int) (*models.Movie, error) {
	movie := &models.Movie{}
	err := b.Db.Where("id = ?", id).First(movie).Error
	return movie, err
}

func (b *MovieService) GetAllMovies() ([]*model.Movie, error) {
	movies := []*model.Movie{}
	err := b.Db.Find(&movies).Error
	return movies, err

}
