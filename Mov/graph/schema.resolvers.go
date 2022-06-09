package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/arukasar/golang-graphql/graph/generated"
	"github.com/arukasar/golang-graphql/graph/model"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.MovieInput) (*model.Movie, error) {
	movie, err := r.MovieRepository.CreateMovie(&input)
	movieCreated := &model.Movie{
		URL:         movie.URL,
		ReleaseDate: movie.ReleaseDate,
		Title:       movie.Title,
		ID:          movie.ID,
	}
	if err != nil {
		return nil, err
	}
	return movieCreated, nil
}

func (r *mutationResolver) DeleteMovie(ctx context.Context, id int) (string, error) {
	err := r.MovieRepository.DeleteMovie(id)
	if err != nil {
		return "", err
	}
	successMessage := "successfully deleted"
	return successMessage, nil
}

func (r *mutationResolver) UpdateMovie(ctx context.Context, id int, input model.MovieInput) (string, error) {
	err := r.MovieRepository.UpdateMovie(&input, id)
	if err != nil {
		return "nil", err
	}
	successMessage := "successfully updated"

	return successMessage, nil
}

func (r *queryResolver) GetAllMovies(ctx context.Context) ([]*model.Movie, error) {
	movies, err := r.MovieRepository.GetAllMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *queryResolver) GetOneMovie(ctx context.Context, id int) (*model.Movie, error) {
	movie, err := r.MovieRepository.GetOneMovie(id)
	selectedMovie := &model.Movie{
		ID:          movie.ID,
		URL:         movie.URL,
		ReleaseDate: movie.ReleaseDate,
		Title:       movie.Title,
	}
	if err != nil {
		return nil, err
	}
	return selectedMovie, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
